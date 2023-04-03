package model

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// System is the interface for ECS systems.
type System interface {
	// Initialize the system
	Initialize(w *ecs.World)
	// Update the system
	Update(w *ecs.World)
	// Finalize the system
	Finalize(w *ecs.World)
}

// UISystem is the interface for ECS systems that display UI.
type UISystem interface {
	// InitializeUI the system.
	InitializeUI(w *ecs.World)
	// UpdateUI/update the system.
	UpdateUI(w *ecs.World)
	// PostUpdateUI does the final part of updating, e.g. update the GL window.
	PostUpdateUI(w *ecs.World)
	// FinalizeUI the system.
	FinalizeUI(w *ecs.World)
}

// Systems manages ECS Systems.
type Systems struct {
	world *ecs.World
	// Frames per second for UI systems.
	// Values <= 0 sync FPS with TPS.
	Fps float64
	// Ticks per second for normal systems.
	// Values <= 0 mean as fast as possible.
	Tps float64

	systems    []System
	uiSystems  []UISystem
	toRemove   []System
	uiToRemove []UISystem

	lastDraw   time.Time
	lastUpdate time.Time

	initialized bool

	timeRes generic.Resource[Time]
}

// AddSystem adds a system to the model
func (s *Systems) AddSystem(sys System) {
	if sys, ok := sys.(UISystem); ok {
		panic(fmt.Sprintf("System %T is also an UI system. Must be added via AddSystem.", sys))
	}
	if s.initialized {
		panic("adding systems after model initialization is not implemented yet")
	}
	s.systems = append(s.systems, sys)
}

// AddUISystem adds an UI system to the model
func (s *Systems) AddUISystem(sys UISystem) {
	if s.initialized {
		panic("adding systems after model initialization is not implemented yet")
	}
	s.uiSystems = append(s.uiSystems, sys)
	if sys, ok := sys.(System); ok {
		s.systems = append(s.systems, sys)
	}
}

// RemoveSystem removes a system from the model
func (s *Systems) RemoveSystem(sys System) {
	s.toRemove = append(s.toRemove, sys)
}

// RemoveUISystem removes an UI system from the model
func (s *Systems) RemoveUISystem(sys UISystem) {
	s.uiToRemove = append(s.uiToRemove, sys)
}

func (s *Systems) removeSystems() {
	for _, sys := range s.toRemove {
		idx := -1
		for idx = 0; idx < len(s.systems); idx++ {
			if sys == s.systems[idx] {
				break
			}
		}
		if idx < 0 {
			panic("System not in the model")
		}
		s.systems[idx].Finalize(s.world)
		s.systems = append(s.systems[:idx], s.systems[idx+1:]...)
	}
	for _, sys := range s.uiToRemove {
		idx := -1
		for idx = 0; idx < len(s.uiSystems); idx++ {
			if sys == s.uiSystems[idx] {
				break
			}
		}
		if idx < 0 {
			panic("System not in the model")
		}
		s.uiSystems[idx].FinalizeUI(s.world)
		s.uiSystems = append(s.uiSystems[:idx], s.uiSystems[idx+1:]...)
	}
	s.toRemove = s.toRemove[:0]
	s.uiToRemove = s.uiToRemove[:0]
}

// initialize all systems
func (s *Systems) initialize() {
	s.timeRes = generic.NewResource[Time](s.world)

	if s.initialized {
		panic("model is already initialized")
	}
	for _, sys := range s.systems {
		sys.Initialize(s.world)
	}
	for _, sys := range s.uiSystems {
		sys.InitializeUI(s.world)
	}
	s.removeSystems()
	s.initialized = true
}

// update all systems
func (s *Systems) update() {
	update := s.updateSystems()
	s.updateUISystems(update)

	s.removeSystems()

	if update {
		time := s.timeRes.Get()
		time.Tick++
	}
}

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

// finalize all systems
func (s *Systems) finalize() {
	for _, sys := range s.systems {
		sys.Finalize(s.world)
	}
	for _, sys := range s.uiSystems {
		sys.FinalizeUI(s.world)
	}
	s.removeSystems()
}

func (s *Systems) run() {
	if !s.initialized {
		s.initialize()
	}

	time := s.timeRes.Get()
	time.Tick = 0

	for !time.Finished {
		s.update()
	}

	s.finalize()
}
