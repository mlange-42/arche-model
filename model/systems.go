package model

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// System is the interface for ECS systems.
//
// See also [UISystem] for systems with an independent graphics step.
type System interface {
	Initialize(w *ecs.World) // Initialize the system.
	Update(w *ecs.World)     // Update the system.
	Finalize(w *ecs.World)   // Finalize the system.
}

// UISystem is the interface for ECS systems that display UI in an independent graphics step.
//
// See also [System] for normal systems.
type UISystem interface {
	InitializeUI(w *ecs.World) // InitializeUI the system.
	UpdateUI(w *ecs.World)     // UpdateUI/update the system.
	PostUpdateUI(w *ecs.World) // PostUpdateUI does the final part of updating, e.g. update the GL window.
	FinalizeUI(w *ecs.World)   // FinalizeUI the system.
}

// Systems manages and schedules ECS [System] and [UISystem] instances.
//
// [System] instances are updated with a frequency given by TPS (ticks per second).
// [UISystem] instances are updated independently of normal systems, with a frequency given by FPS (frames per second).
//
// [Systems] is an embed in [Model] and it's methods are usually only used through a [Model] instance.
// By also being a resource of each [Model], however, systems can access it and e.g. remove themselves from a model.
type Systems struct {
	TPS    float64 // Ticks per second for normal systems. Values <= 0 (the default) mean as fast as possible.
	FPS    float64 // Frames per second for UI systems. A zero/unset value defaults to 30 FPS. Values < 0 sync FPS with TPS.
	Paused bool    // Whether the simulation is currently paused. When paused, only UI updates but no normal updates are performed.

	world      *ecs.World
	systems    []System
	uiSystems  []UISystem
	toRemove   []System
	uiToRemove []UISystem

	nextDraw   time.Time
	nextUpdate time.Time

	initialized bool
	locked      bool

	tickRes generic.Resource[resource.Tick]
	termRes generic.Resource[resource.Termination]
}

// AddSystem adds a [System] to the model.
//
// Panics if the system is also a [UISystem].
// To add systems that implement both [System] and [UISystem], use [Systems.AddUISystem]
func (s *Systems) AddSystem(sys System) {
	if s.initialized {
		panic("adding systems after model initialization is not implemented yet")
	}
	if sys, ok := sys.(UISystem); ok {
		panic(fmt.Sprintf("System %T is also an UI system. Must be added via AddSystem.", sys))
	}
	s.systems = append(s.systems, sys)
}

// AddUISystem adds an [UISystem] to the model.
//
// Adds the [UISystem] also as a normal [System] if it implements the interface.
func (s *Systems) AddUISystem(sys UISystem) {
	if s.initialized {
		panic("adding systems after model initialization is not implemented yet")
	}
	s.uiSystems = append(s.uiSystems, sys)
	if sys, ok := sys.(System); ok {
		s.systems = append(s.systems, sys)
	}
}

// RemoveSystem removes a system from the model.
//
// Systems can also be removed during a model run.
// However, this will take effect only after the end of the full model step.
func (s *Systems) RemoveSystem(sys System) {
	if sys, ok := sys.(UISystem); ok {
		panic(fmt.Sprintf("System %T is also an UI system. Must be removed via RemoveUISystem.", sys))
	}
	s.toRemove = append(s.toRemove, sys)
	if !s.locked {
		s.removeSystems()
	}
}

// RemoveUISystem removes an UI system from the model.
//
// Systems can also be removed during a model run.
// However, this will take effect only after the end of the full model step.
func (s *Systems) RemoveUISystem(sys UISystem) {
	s.uiToRemove = append(s.uiToRemove, sys)
	if !s.locked {
		s.removeSystems()
	}
}

// Removes systems that were removed during the model step.
func (s *Systems) removeSystems() {
	for _, sys := range s.toRemove {
		s.removeSystem(sys)
	}
	for _, sys := range s.uiToRemove {
		if sys, ok := sys.(System); ok {
			s.removeSystem(sys)
		}
		s.removeUISystem(sys)
	}
	s.toRemove = s.toRemove[:0]
	s.uiToRemove = s.uiToRemove[:0]
}

func (s *Systems) removeSystem(sys System) {
	if s.locked {
		panic("can't remove a system in locked state")
	}
	idx := -1
	for i := 0; i < len(s.systems); i++ {
		if sys == s.systems[i] {
			idx = i
			break
		}
	}
	if idx < 0 {
		panic(fmt.Sprintf("can't remove system %T: not in the model", sys))
	}
	s.systems[idx].Finalize(s.world)
	s.systems = append(s.systems[:idx], s.systems[idx+1:]...)
}

func (s *Systems) removeUISystem(sys UISystem) {
	if s.locked {
		panic("can't remove a system in locked state")
	}
	idx := -1
	for i := 0; i < len(s.uiSystems); i++ {
		if sys == s.uiSystems[i] {
			idx = i
			break
		}
	}
	if idx < 0 {
		panic(fmt.Sprintf("can't remove UI system %T: not in the model", sys))
	}
	s.uiSystems[idx].FinalizeUI(s.world)
	s.uiSystems = append(s.uiSystems[:idx], s.uiSystems[idx+1:]...)
}

// Initialize all systems.
func (s *Systems) initialize() {
	if s.initialized {
		panic("model is already initialized")
	}

	if s.FPS == 0 {
		s.FPS = 30
	}

	s.tickRes = generic.NewResource[resource.Tick](s.world)
	s.termRes = generic.NewResource[resource.Termination](s.world)

	s.locked = true
	for _, sys := range s.systems {
		sys.Initialize(s.world)
	}
	for _, sys := range s.uiSystems {
		sys.InitializeUI(s.world)
	}
	s.locked = false
	s.removeSystems()
	s.initialized = true

	s.nextDraw = time.Time{}
	s.nextUpdate = time.Time{}
}

// Update all systems.
func (s *Systems) update() {
	s.locked = true
	update := s.updateSystems()
	s.updateUISystems(update)
	s.locked = false

	s.removeSystems()

	if update {
		time := s.tickRes.Get()
		time.Tick++
	} else {
		s.wait()
	}
}

// Calculates and waits the time until the next update of UI update.
func (s *Systems) wait() {
	var nextUpdate time.Time
	if s.TPS > 0 {
		nextUpdate = s.nextUpdate
	}
	if (s.Paused || s.FPS > 0) && s.nextDraw.Before(nextUpdate) {
		nextUpdate = s.nextDraw
	}
	if nextUpdate.IsZero() {
		return
	}

	t := time.Now()
	wait := nextUpdate.Sub(t)
	// Wait only if time is sufficiently long, as time.Sleep only guaranties minimum waiting time.
	if wait > time.Millisecond {
		time.Sleep(wait)
	}
}

// Update normal systems.
func (s *Systems) updateSystems() bool {
	if s.Paused {
		return false
	}
	update := false
	if s.TPS <= 0 {
		update = true
		for _, sys := range s.systems {
			sys.Update(s.world)
		}
	} else {
		update = !time.Now().Before(s.nextUpdate)
		if update {
			s.nextUpdate = nextTime(s.nextUpdate, s.TPS)
			for _, sys := range s.systems {
				sys.Update(s.world)
			}
		}
	}
	return update
}

// Update UI systems.
func (s *Systems) updateUISystems(updated bool) {
	if len(s.uiSystems) > 0 {
		if !s.Paused && s.FPS <= 0 {
			if updated {
				for _, sys := range s.uiSystems {
					sys.UpdateUI(s.world)
				}
				for _, sys := range s.uiSystems {
					sys.PostUpdateUI(s.world)
				}
			}
		} else {
			if !time.Now().Before(s.nextDraw) {
				fps := s.FPS
				if s.Paused {
					fps = 30
				}
				s.nextDraw = nextTime(s.nextDraw, fps)
				for _, sys := range s.uiSystems {
					sys.UpdateUI(s.world)
				}
				for _, sys := range s.uiSystems {
					sys.PostUpdateUI(s.world)
				}
			}
		}
	}
}

// Finalize all systems.
func (s *Systems) finalize() {
	s.locked = true
	for _, sys := range s.systems {
		sys.Finalize(s.world)
	}
	for _, sys := range s.uiSystems {
		sys.FinalizeUI(s.world)
	}
	s.locked = false
	s.removeSystems()
}

// Run the model.
func (s *Systems) run() {
	if !s.initialized {
		s.initialize()
	}

	time := s.tickRes.Get()
	time.Tick = 0
	terminate := s.termRes.Get()

	for !terminate.Terminate {
		s.update()
	}

	s.finalize()
}

// Removes all systems.
func (s *Systems) reset() {
	s.systems = []System{}
	s.uiSystems = []UISystem{}
	s.toRemove = []System{}
	s.uiToRemove = []UISystem{}

	s.nextDraw = time.Time{}
	s.nextUpdate = time.Time{}

	s.initialized = false
	s.tickRes = generic.Resource[resource.Tick]{}
}
