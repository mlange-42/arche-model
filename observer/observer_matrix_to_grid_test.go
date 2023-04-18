package observer_test

import (
	"github.com/mlange-42/arche-model/observer"
)

func ExampleMatrixToGrid() {
	// A Matrix observer
	var matrix observer.Matrix = &MatrixObserver{}

	// A MatrixToGrid observer, wrapping the Matrix observer
	var _ observer.Grid = observer.MatrixToGrid(
		matrix,
		&[...]float64{100, 200},
		&[...]float64{1000, 1000},
	)
	// Output:
}
