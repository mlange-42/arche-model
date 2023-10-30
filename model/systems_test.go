package model

import (
	"testing"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"github.com/stretchr/testify/assert"
)

func TestSystems(t *testing.T) {
	m := New()
	for i := 0; i < 3; i++ {
		m.Reset()

		m.Seed()
		m.Seed(123)

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

		m.locked = true
		assert.Panics(t, func() { m.removeSystem(&termSys) })
		assert.Panics(t, func() { m.removeUISystem(&uiSys) })
		m.locked = false

		assert.Panics(t, func() { m.AddSystem(&dualSys) })
		m.AddUISystem(&dualSys)

		m.AddSystem(&removerSystem{
			Remove:   []System{&termSys},
			RemoveUI: []UISystem{&uiSys},
		})

		assert.Equal(t, m.systems, m.Systems.Systems())
		assert.Equal(t, m.uiSystems, m.Systems.UISystems())

		assert.Panics(t, func() { m.RemoveSystem(&dualSys) })

		assert.Equal(t, 4, len(m.systems))
		assert.Equal(t, 2, len(m.uiSystems))
		assert.Equal(t, 0, len(m.toRemove))
		assert.Equal(t, 0, len(m.uiToRemove))

		m.Run()

		assert.Equal(t, 3, len(m.systems))
		assert.Equal(t, 1, len(m.uiSystems))
		assert.Equal(t, 0, len(m.toRemove))
		assert.Equal(t, 0, len(m.uiToRemove))

		assert.Panics(t, func() { m.initialize() })

		m.RemoveUISystem(&dualSys)

		assert.Equal(t, 2, len(m.systems))
		assert.Equal(t, 0, len(m.uiSystems))
		assert.Equal(t, 0, len(m.toRemove))
		assert.Equal(t, 0, len(m.uiToRemove))

		assert.Panics(t, func() { m.RemoveUISystem(&dualSys) })

		assert.Panics(t, func() { m.RemoveSystem(&termSys) })

		assert.Panics(t, func() { m.RemoveUISystem(&uiSys) })

		assert.Panics(t, func() { m.AddSystem(&termSys) })
		assert.Panics(t, func() { m.AddUISystem(&uiSys) })
	}
}

func TestSystemsInit(t *testing.T) {
	m := New()
	m.TPS = 0
	m.FPS = 0

	m.AddSystem(&system.FixedTermination{Steps: 5})
	m.AddUISystem(&uiSystem{})
	m.Run()

	assert.Equal(t, 30.0, m.FPS)
	assert.Equal(t, 0.0, m.TPS)
	assert.Equal(t, 5, int(m.time.Tick))

	m = New()
	m.TPS = 10

	m.AddSystem(&system.FixedTermination{Steps: 5})
	m.AddUISystem(&uiSystem{})

	m.Run()

	m = New()
	m.TPS = 10
	m.FPS = 30

	m.AddSystem(&system.FixedTermination{Steps: 5})
	m.AddUISystem(&uiSystem{})

	m.Run()

	m = New()
	m.TPS = 10
	m.FPS = -1

	m.AddSystem(&system.FixedTermination{Steps: 5})
	m.AddUISystem(&uiSystem{})

	m.Run()
}

func TestSystemsPaused(t *testing.T) {
	m := New()
	m.TPS = 0
	m.FPS = 0

	m.AddSystem(&system.FixedTermination{Steps: 5})
	m.AddUISystem(&uiTerminationSystem{Steps: 100})

	m.Paused = true
	m.Run()

	assert.Equal(t, 0, int(m.time.Tick))
}

type uiSystem struct{}

func (s *uiSystem) InitializeUI(w *ecs.World) {}
func (s *uiSystem) UpdateUI(w *ecs.World)     {}
func (s *uiSystem) PostUpdateUI(w *ecs.World) {}
func (s *uiSystem) FinalizeUI(w *ecs.World)   {}

type uiTerminationSystem struct {
	Steps   int
	step    int
	termRes generic.Resource[resource.Termination]
}

func (s *uiTerminationSystem) InitializeUI(w *ecs.World) {
	s.termRes = generic.NewResource[resource.Termination](w)
	s.step = 0
}

func (s *uiTerminationSystem) UpdateUI(w *ecs.World) {
	if s.step >= s.Steps {
		term := s.termRes.Get()
		term.Terminate = true
	}
	s.step++
}
func (s *uiTerminationSystem) PostUpdateUI(w *ecs.World) {}
func (s *uiTerminationSystem) FinalizeUI(w *ecs.World)   {}

type dualSystem struct{}

func (s *dualSystem) Initialize(w *ecs.World)   {}
func (s *dualSystem) InitializeUI(w *ecs.World) {}
func (s *dualSystem) Update(w *ecs.World)       {}
func (s *dualSystem) UpdateUI(w *ecs.World)     {}
func (s *dualSystem) PostUpdateUI(w *ecs.World) {}
func (s *dualSystem) Finalize(w *ecs.World)     {}
func (s *dualSystem) FinalizeUI(w *ecs.World)   {}

type removerSystem struct {
	Remove   []System
	RemoveUI []UISystem
	step     int
}

func (s *removerSystem) Initialize(w *ecs.World) {}
func (s *removerSystem) Update(w *ecs.World) {
	if s.step == 3 {
		systems := ecs.GetResource[Systems](w)
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
