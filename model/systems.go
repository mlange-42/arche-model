package model

import (
	"fmt"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

// System is the interface for ECS systems.
type System interface {
	// Initialize the system
	Initialize(m *Model)
	// Update the system
	Update(m *Model)
	// Finalize the system
	Finalize(m *Model)
}

// UISystem is the interface for ECS systems that display UI.
type UISystem interface {
	// InitializeUI the system
	InitializeUI(m *Model)
	// UpdateUI/update the system
	UpdateUI(m *Model)
	// FinalizeUI the system
	FinalizeUI(m *Model)
	// Returns the system's window if there is any.
	Window() *pixelgl.Window
}

// systems manages ECS systems.
type systems struct {
	model *Model
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

	Step        int64
	Finished    bool
	initialized bool
}

// AddSystem adds a system to the model
func (s *systems) AddSystem(sys System) {
	if sys, ok := sys.(UISystem); ok {
		panic(fmt.Sprintf("System %T is also an UI system. Must be added via AddSystem.", sys))
	}
	s.systems = append(s.systems, sys)
}

// AddUISystem adds an UI system to the model
func (s *systems) AddUISystem(sys UISystem) {
	s.uiSystems = append(s.uiSystems, sys)
	if sys, ok := sys.(System); ok {
		s.systems = append(s.systems, sys)
	}
}

// RemoveSystem removes a system from the model
func (s *systems) RemoveSystem(sys System) {
	s.toRemove = append(s.toRemove, sys)
}

// RemoveUISystem removes an UI system from the model
func (s *systems) RemoveUISystem(sys UISystem) {
	s.uiToRemove = append(s.uiToRemove, sys)
}

func (s *systems) removeSystems() {
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
		s.systems[idx].Finalize(s.model)
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
		s.uiSystems[idx].FinalizeUI(s.model)
		s.uiSystems = append(s.uiSystems[:idx], s.uiSystems[idx+1:]...)
	}
	s.toRemove = s.toRemove[:0]
	s.uiToRemove = s.uiToRemove[:0]
}

// initialize all systems
func (s *systems) initialize() {
	if s.initialized {
		panic("model is already initialized")
	}
	for _, sys := range s.systems {
		sys.Initialize(s.model)
	}
	for _, sys := range s.uiSystems {
		sys.InitializeUI(s.model)
	}
	s.removeSystems()
	s.initialized = true
}

// update all systems
func (s *systems) update() {
	update := s.updateSystems()
	s.updateUISystems(update)

	s.removeSystems()

	if update {
		s.Step++
	}
}

func (s *systems) updateSystems() bool {
	update := false
	if s.Tps <= 0 {
		update = true
		for _, sys := range s.systems {
			sys.Update(s.model)
		}
	} else {
		t := time.Now()
		frameDur := t.Sub(s.lastUpdate)
		update = 1.0/frameDur.Seconds() <= s.Tps
		if update {
			s.lastUpdate = t
			for _, sys := range s.systems {
				sys.Update(s.model)
			}
		}
	}
	return update
}

func (s *systems) updateUISystems(updated bool) {
	if len(s.uiSystems) > 0 {
		if s.Fps <= 0 {
			if updated {
				for _, sys := range s.uiSystems {
					sys.UpdateUI(s.model)
				}
				for _, sys := range s.uiSystems {
					win := sys.Window()
					if win != nil {
						win.Update()
					}
				}
			}
		} else {
			t := time.Now()
			frameDur := t.Sub(s.lastDraw)
			if 1.0/frameDur.Seconds() <= s.Fps {
				s.lastDraw = t
				for _, sys := range s.uiSystems {
					sys.UpdateUI(s.model)
				}
				for _, sys := range s.uiSystems {
					win := sys.Window()
					if win != nil {
						win.Update()
					}
				}
			}
		}
	}
}

// finalize all systems
func (s *systems) finalize() {
	for _, sys := range s.systems {
		sys.Finalize(s.model)
	}
	for _, sys := range s.uiSystems {
		sys.FinalizeUI(s.model)
	}
	s.removeSystems()
}

func (s *systems) run() {
	if !s.initialized {
		s.initialize()
	}

	s.Step = 0
	for !s.Finished {
		s.update()
	}

	s.finalize()
}
