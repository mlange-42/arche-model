// The example from the README.
package main

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Position component
type Position struct {
	X float64
	Y float64
}

// Velocity component
type Velocity struct {
	X float64
	Y float64
}

func main() {
	// Create a new, seeded model.
	m := model.New().Seed(123)
	// Limit the simulation speed.
	m.TPS = 30

	// Add systems to the model.
	m.AddSystem(&VelocitySystem{EntityCount: 1000})
	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 100})

	// Run the model.
	m.Run()
}

// VelocitySystem is an example system adding velocity to position.
// For simplicity, it also creates entities during initialization.
type VelocitySystem struct {
	EntityCount int
	filter      generic.Filter2[Position, Velocity]
}

// Initialize the system
func (s *VelocitySystem) Initialize(w *ecs.World) {
	s.filter = *generic.NewFilter2[Position, Velocity]()

	mapper := generic.NewMap2[Position, Velocity](w)
	mapper.NewBatch(s.EntityCount)
}

// Update the system
func (s *VelocitySystem) Update(w *ecs.World) {
	query := s.filter.Query(w)

	for query.Next() {
		pos, vel := query.Get()
		pos.X += vel.X
		pos.Y += vel.Y
	}
}

// Finalize the system
func (s *VelocitySystem) Finalize(w *ecs.World) {}
