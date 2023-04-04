package model_test

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// TestUISystem is an example for implementing [UISystem].
type TestUISystem struct {
	timeRes generic.Resource[model.Tick]
}

// Initialize the system.
func (s *TestUISystem) InitializeUI(w *ecs.World) {
	s.timeRes = generic.NewResource[model.Tick](w)
}

// Update the system.
func (s *TestUISystem) UpdateUI(w *ecs.World) {
	time := s.timeRes.Get()
	fmt.Println(time.Tick)
}

// PostUpdate the system.
func (s *TestUISystem) PostUpdateUI(w *ecs.World) {}

// Finalize the system.
func (s *TestUISystem) FinalizeUI(w *ecs.World) {}

func ExampleUISystem() {
	// Create a new model.
	m := model.New()

	// Add the test ui system.
	m.AddUISystem(&TestUISystem{})

	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 30})

	// Run the simulation.
	m.Run()
}
