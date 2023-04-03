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
		m.Tps = 120
		m.Fps = 60

		termSys := system.FixedTermination{
			Steps: 1000,
		}
		uiSys := uiSystem{}

		m.AddSystem(&termSys)
		m.AddSystem(&system.FixedTermination{
			Steps: 100,
		})
		m.AddUISystem(&uiSys)

		m.AddSystem(&removerSystem{
			Remove:   []model.System{&termSys},
			RemoveUI: []model.UISystem{&uiSys},
		})

		m.Run()

		assert.Panics(t, func() { m.AddSystem(&termSys) })
		assert.Panics(t, func() { m.AddUISystem(&uiSys) })
	}
}

type uiSystem struct{}

func (s *uiSystem) InitializeUI(w *ecs.World) {}
func (s *uiSystem) UpdateUI(w *ecs.World)     {}
func (s *uiSystem) PostUpdateUI(w *ecs.World) {}
func (s *uiSystem) FinalizeUI(w *ecs.World)   {}

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
