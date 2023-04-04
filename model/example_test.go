package model_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func Example() {
	// Create a new model.
	m := model.New()
	// Seed the model's PRNG.
	m.Seed(123)

	// Add systems.
	m.AddSystem(&system.FixedTermination{
		Steps: 100,
	})

	// Run the simulation.
	m.Run()
	// Output:
}
