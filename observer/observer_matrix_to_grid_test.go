package observer_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
)

func ExampleMatrixToGrid() {
	m := model.New()

	// A Matrix observer
	var matrix observer.Matrix = &MatrixObserver{}
	var _ []float64 = matrix.Values(&m.World)

	// A MatrixToGrid observer, wrapping the Matrix observer
	var grid observer.Grid = &observer.MatrixToGrid{
		Observer: matrix,
		Origin:   [...]float64{100, 200},
		CellSize: [...]float64{1000, 1000},
	}
	var _ []float64 = grid.Values(&m.World)
	// Output:
}
