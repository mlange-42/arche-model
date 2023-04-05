package observer

import (
	"github.com/mlange-42/arche/ecs"
)

// Row observer interface. Provides column headers, and a single data row per call.
//
// See also [Table]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type Row interface {
	Initialize(w *ecs.World)       // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)           // Update the observer.
	Header() []string              // Header/column names in the same order as data values.
	Values(w *ecs.World) []float64 // Values for the current model tick.
}

// Table observer interface. Provides column headers, and multiple data rows per call.
//
// See also [Row]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type Table interface {
	Initialize(w *ecs.World)         // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)             // Update the observer.
	Header() []string                // Header/column names in the same order as data values.
	Values(w *ecs.World) [][]float64 // Values for the current model tick.
}

// Matrix observer interface. Provides dimensionality, and a matrix of values per call.
//
// See also [Grid]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type Matrix interface {
	Initialize(w *ecs.World)       // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)           // Update the observer.
	Dims() (int, int)              // Matrix dimensions.
	Values(w *ecs.World) []float64 // Values for the current model tick, in row-major order (i.e. idx = row*ncols + col).
}

// Grid observer interface. Provides dimensionality, axis information, and a matrix of values per call.
//
// See also [Matrix]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type Grid interface {
	Matrix        // Methods from Matrix observer.
	X() []float64 // X axis values.
	Y() []float64 // Y axis values.
}
