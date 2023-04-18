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

// MatrixLayers observer interface. Provides dimensionality, and multiple matrices of values per call.
//
// See also [Matrix] and [GridLayers]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type MatrixLayers interface {
	Initialize(w *ecs.World)         // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)             // Update the observer.
	Layers() int                     // Number of layers.
	Dims() (int, int)                // Matrix dimensions.
	Values(w *ecs.World) [][]float64 // Values for the current model tick, in row-major order (i.e. idx = row*ncols + col). First index is the layer.
}

// Grid observer interface. Provides dimensionality, axis information, and a matrix of values per call.
//
// See also [Matrix] and [GridLayers]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type Grid interface {
	Matrix           // Methods from Matrix observer.
	X(c int) float64 // X axis coordinates.
	Y(r int) float64 // Y axis coordinates.
}

// GridLayers observer interface. Provides dimensionality, axis information, and multiple matrices of values per call.
//
// See also [Grid], [Matrix] and [GridLayers]. See package [github.com/mlange-42/arche-model/reporter] for usage examples.
type GridLayers interface {
	MatrixLayers     // Methods from Matrix observer.
	X(c int) float64 // X axis coordinates.
	Y(r int) float64 // Y axis coordinates.
}
