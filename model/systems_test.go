package model_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestSystems(t *testing.T) {
	m := model.New()
	for i := 0; i < 3; i++ {
		m.Reset()

		m.Seed()
		m.Seed(123)
		m.TPS = 120
		m.FPS = 60

		termSys := system.FixedTermination{
			Steps: 1000,
		}
		uiSys := uiSystem{}
		dualSys := dualSystem{}

		m.AddSystem(&termSys)
		m.AddSystem(&system.FixedTermination{
			Steps: 10,
		})
		m.AddUISystem(&uiSys)

		assert.Panics(t, func() { m.AddSystem(&dualSys) })
		m.AddUISystem(&dualSys)

		m.AddSystem(&removerSystem{
			Remove:   []model.System{&termSys},
			RemoveUI: []model.UISystem{&uiSys},
		})

		m.Run()

		assert.Panics(t, func() { m.RemoveSystem(&termSys) })
		assert.Panics(t, func() { m.RemoveUISystem(&uiSys) })

		assert.Panics(t, func() { m.AddSystem(&termSys) })
		assert.Panics(t, func() { m.AddUISystem(&uiSys) })
	}
}

func TestSystemsInit(t *testing.T) {
	m := model.New()
	m.AddSystem(&system.FixedTermination{Steps: 1})
	m.Run()

	assert.Equal(t, 30.0, m.FPS)
	assert.Equal(t, 0.0, m.TPS)
}

type uiSystem struct{}

func (s *uiSystem) InitializeUI(w *ecs.World) {}
func (s *uiSystem) UpdateUI(w *ecs.World)     {}
func (s *uiSystem) PostUpdateUI(w *ecs.World) {}
func (s *uiSystem) FinalizeUI(w *ecs.World)   {}

type dualSystem struct{}

func (s *dualSystem) Initialize(w *ecs.World)   {}
func (s *dualSystem) InitializeUI(w *ecs.World) {}
func (s *dualSystem) Update(w *ecs.World)       {}
func (s *dualSystem) UpdateUI(w *ecs.World)     {}
func (s *dualSystem) PostUpdateUI(w *ecs.World) {}
func (s *dualSystem) Finalize(w *ecs.World)     {}
func (s *dualSystem) FinalizeUI(w *ecs.World)   {}

type removerSystem struct {
	Remove   []model.System
	RemoveUI []model.UISystem
	step     int
}

func (s *removerSystem) Initialize(w *ecs.World) {}
func (s *removerSystem) Update(w *ecs.World) {
	if s.step == 3 {
		systems := ecs.GetResource[model.Systems](w)
		for _, sys := range s.Remove {
			systems.RemoveSystem(sys)
		}
		for _, sys := range s.RemoveUI {
			systems.RemoveUISystem(sys)
		}
	}
	s.step++
}
func (s *removerSystem) Finalize(w *ecs.World) {}

func ExampleSystems() {
	// Create a new model.
	m := model.New()

	// The model contains Systems as an embed, TPS and FPS are accessible through the model directly.
	m.TPS = 1000
	m.FPS = 60

	// Create a system
	sys := system.FixedTermination{
		Steps: 10,
	}

	// Add the system the usual way, through the model.
	// The model contains Systems as an embed, so actually [Systems.AddSystem] is called.
	m.AddSystem(&sys)

	// Inside systems, [Systems] can be accessed as a resource.
	systems := ecs.GetResource[model.Systems](&m.World)

	// Pause the simulation, e.g. based on user input.
	systems.Paused = true

	// Remove the system using the resource.
	systems.RemoveSystem(&sys)
	// Output:
}
