# Arche Model

[![Test status](https://img.shields.io/github/actions/workflow/status/mlange-42/arche-model/tests.yml?branch=main&label=Tests&logo=github)](https://github.com/mlange-42/arche-model/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/mlange-42/arche-model/badge.svg?branch=main)](https://coveralls.io/github/mlange-42/arche-model?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mlange-42/arche-model)](https://goreportcard.com/report/github.com/mlange-42/arche-model)
[![Go Reference](https://pkg.go.dev/badge/github.com/mlange-42/arche-model.svg)](https://pkg.go.dev/github.com/mlange-42/arche-model)
[![GitHub](https://img.shields.io/badge/github-repo-blue?logo=github)](https://github.com/mlange-42/arche-model)
[![MIT license](https://img.shields.io/github/license/mlange-42/arche-model)](https://github.com/mlange-42/arche-model/blob/main/LICENSE)

*Arche Model* provides a wrapper around the [Arche](https://github.com/mlange-42/arche) Entity Component System (ECS), and some common systems and resources.
It's purpose is to get started with prototyping and developing simulation models immediately, focussing on the model logic.

## Features

* Scheduler for running logic and UI systems with independent update rates.
* Interfaces for ECS systems and observers.
* Ready-to-use systems for common tasks like writing CSV files or terminating a simulation.
* Common ECS resources, like central PRNG source or the current model tick.

## Installation

```
go get github.com/mlange-42/arche-model
```

## Usage

See the [API docs](https://pkg.go.dev/github.com/mlange-42/arche-model) for more details and examples.  
[![Go Reference](https://pkg.go.dev/badge/github.com/mlange-42/arche-model.svg)](https://pkg.go.dev/github.com/mlange-42/arche-model)

```go
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
	// Limit simulation speed
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
	mapper.NewEntities(s.EntityCount)
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
```

## License

This project is distributed under the [MIT licence](./LICENSE).
