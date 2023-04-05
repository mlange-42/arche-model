package model

import (
	"github.com/mlange-42/arche/ecs"
)

// RowObserver interface. Reports column headers, and a single data row per call.
//
// Reporters like [system.CSV] require an RowObserver instance to extract data from the ECS world.
//
// See also [TableObserver].
type RowObserver interface {
	Initialize(w *ecs.World)       // Initialize the observer.
	Update(w *ecs.World)           // Update the observer.
	Header(w *ecs.World) []string  // Header/column names in the same order as data values.
	Values(w *ecs.World) []float64 // Values for the current model tick.
}

// TableObserver interface. Reports column headers, and multiple data rows per call.
//
// Reporters like [system.SnapshotCSV] require a TableObserver instance to extract data from the ECS world.
//
// See also [RowObserver].
type TableObserver interface {
	Initialize(w *ecs.World)         // Initialize the observer.
	Update(w *ecs.World)             // Update the observer.
	Header(w *ecs.World) []string    // Header/column names in the same order as data values.
	Values(w *ecs.World) [][]float64 // Values for the current model tick.
}

// MatrixObserver interface. Reports axis information, and a matrix of values per call.
//
// See also [Observer].
type MatrixObserver interface {
	Initialize(w *ecs.World)       // Initialize the observer.
	Update(w *ecs.World)           // Update the observer.
	X(w *ecs.World) []float64      // X axis values.
	Y(w *ecs.World) []float64      // Y axis values.
	Values(w *ecs.World) []float64 // Values for the current model tick, in row-major order (i.e. idx = row*ncols + col).
}
