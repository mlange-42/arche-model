package observer

import (
	"github.com/mlange-42/arche/ecs"
)

// Row observer interface. Reports column headers, and a single data row per call.
//
// Reporters like [system.CSV] require an Row instance to extract data from the ECS world.
//
// See also [Table].
type Row interface {
	Initialize(w *ecs.World)       // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)           // Update the observer.
	Header() []string              // Header/column names in the same order as data values.
	Values(w *ecs.World) []float64 // Values for the current model tick.
}

// Table observer interface. Reports column headers, and multiple data rows per call.
//
// Reporters like [system.SnapshotCSV] require a Table instance to extract data from the ECS world.
//
// See also [Row].
type Table interface {
	Initialize(w *ecs.World)         // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)             // Update the observer.
	Header() []string                // Header/column names in the same order as data values.
	Values(w *ecs.World) [][]float64 // Values for the current model tick.
}

// Matrix observer interface. Reports dimensionality, and a matrix of values per call.
//
// See also [Grid].
type Matrix interface {
	Initialize(w *ecs.World)       // Initialize the observer. No other methods are called before this.
	Update(w *ecs.World)           // Update the observer.
	Dims() (int, int)              // Matrix dimensions.
	Values(w *ecs.World) []float64 // Values for the current model tick, in row-major order (i.e. idx = row*ncols + col).
}

// Grid observer interface. Reports dimensionality, axis information, and a matrix of values per call.
//
// See also [Matrix].
type Grid interface {
	Matrix        // Methods from MatrixObserver.
	X() []float64 // X axis values.
	Y() []float64 // Y axis values.
}
