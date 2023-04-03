package model

import (
	"github.com/mlange-42/arche/ecs"
)

// Observer interface.
//
// Reporters like [systems.CSV] require an Observer instance to extract data from the ECS world.
type Observer interface {
	Initialize(w *ecs.World)       // Initialize the observer.
	Update(w *ecs.World)           // Update the observer.
	Header(w *ecs.World) []string  // Header/column names in the same order as data values.
	Values(w *ecs.World) []float64 // Values for the current model tick.
}

// MatrixObserver interface
//
// Reporters like [systems.SnapshotCSV] require a MatrixObserver instance to extract data from the ECS world.
type MatrixObserver interface {
	Initialize(w *ecs.World)         // Initialize the observer.
	Update(w *ecs.World)             // Update the observer.
	Header(w *ecs.World) []string    // Header/column names in the same order as data values.
	Values(w *ecs.World) [][]float64 // Values for the current model tick.
}
