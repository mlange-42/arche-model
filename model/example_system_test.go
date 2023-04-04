package model_test

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// TestSystem is an example for implementing [System].
type TestSystem struct {
	timeRes generic.Resource[model.Time]
}

// Initialize the system.
func (s *TestSystem) Initialize(w *ecs.World) {
	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system.
func (s *TestSystem) Update(w *ecs.World) {
	time := s.timeRes.Get()
	fmt.Println(time.Tick)
}

// Finalize the system.
func (s *TestSystem) Finalize(w *ecs.World) {}

func ExampleSystem() {
	// Create a new model.
	m := model.New()

	// Add the test system.
	m.AddSystem(&TestSystem{})

	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 30})

	// Run the simulation.
	m.Run()
}