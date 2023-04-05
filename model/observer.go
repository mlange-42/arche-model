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
	Initialize(w *ecs.World)       // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)           // Update the observer.
	Header() []string              // Header/column names in the same order as data values.
	Values(w *ecs.World) []float64 // Values for the current model tick.
}

// TableObserver interface. Reports column headers, and multiple data rows per call.
//
// Reporters like [system.SnapshotCSV] require a TableObserver instance to extract data from the ECS world.
//
// See also [RowObserver].
type TableObserver interface {
	Initialize(w *ecs.World)         // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)             // Update the observer.
	Header() []string                // Header/column names in the same order as data values.
	Values(w *ecs.World) [][]float64 // Values for the current model tick.
}

// MatrixObserver interface. Reports dimensionality, and a matrix of values per call.
//
// See also [GridObserver].
type MatrixObserver interface {
	Initialize(w *ecs.World)       // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)           // Update the observer.
	Dims() (int, int)              // Matrix dimensions.
	Values(w *ecs.World) []float64 // Values for the current model tick, in row-major order (i.e. idx = row*ncols + col).
}

// GridObserver interface. Reports dimensionality, axis information, and a matrix of values per call.
//
// See also [MatrixObserver].
type GridObserver interface {
	MatrixObserver // Methods from MatrixObserver.
	X() []float64  // X axis values.
	Y() []float64  // Y axis values.
}
