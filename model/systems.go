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
// [System] instances are updated with a frequency given by Tps.
// [UISystem] instances are updated independently of normal systems, with a frequency given by Fps.
//
// [Systems] is an embed in [Model] and it's methods are usually only used through a [Model] instance.
// By also being a resource of each [Model], however, systems can access it and e.g. remove themselves from a model.
type Systems struct {
	Fps float64 // Frames per second for UI systems. Values <= 0 (the default) sync FPS with TPS.
	Tps float64 // Ticks per second for normal systems. Values <= 0 (the default) mean as fast as possible.

	world      *ecs.World
	systems    []System
	uiSystems  []UISystem
	toRemove   []System
	uiToRemove []UISystem

	lastDraw   time.Time
	lastUpdate time.Time

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
	for _, sys := range s.uiToRemove {
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
	s.toRemove = s.toRemove[:0]
	s.uiToRemove = s.uiToRemove[:0]
}

// Initialize all systems.
func (s *Systems) initialize() {
	s.tickRes = generic.NewResource[resource.Tick](s.world)
	s.termRes = generic.NewResource[resource.Termination](s.world)

	if s.initialized {
		panic("model is already initialized")
	}
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
	t := time.Now()
	nextUpdate := t
	if s.Tps > 0 {
		nextUpdate = s.lastUpdate.Add(time.Second / time.Duration(s.Tps))
	}
	if s.Fps > 0 {
		nextUpdate2 := s.lastDraw.Add(time.Second / time.Duration(s.Fps))
		if nextUpdate2.Before(nextUpdate) {
			nextUpdate = nextUpdate2
		}
	}
	wait := nextUpdate.Sub(t)
	if wait > 0 {
		time.Sleep(wait)
	}
}

// Update normal systems.
func (s *Systems) updateSystems() bool {
	update := false
	if s.Tps <= 0 {
		update = true
		for _, sys := range s.systems {
			sys.Update(s.world)
		}
	} else {
		t := time.Now()
		frameDur := t.Sub(s.lastUpdate)
		update = 1.0/frameDur.Seconds() <= s.Tps
		if update {
			s.lastUpdate = t
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
		if s.Fps <= 0 {
			if updated {
				for _, sys := range s.uiSystems {
					sys.UpdateUI(s.world)
				}
				for _, sys := range s.uiSystems {
					sys.PostUpdateUI(s.world)
				}
			}
		} else {
			t := time.Now()
			frameDur := t.Sub(s.lastDraw)
			if 1.0/frameDur.Seconds() <= s.Fps {
				s.lastDraw = t
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

	s.lastDraw = time.Time{}
	s.lastUpdate = time.Time{}

	s.initialized = false
	s.tickRes = generic.Resource[resource.Tick]{}
}
