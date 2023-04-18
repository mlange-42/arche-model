package observer_test

import (
	"github.com/mlange-42/arche-model/observer"
)

func ExampleLayersToGrid() {
	// Multiple Matrix observers
	var matrix1 observer.Matrix = &MatrixObserver{}
	var matrix2 observer.Matrix = &MatrixObserver{}
	var matrix3 observer.Matrix = &MatrixObserver{}

	// A MatrixToGrid observer, wrapping the Matrix observers
	var layers observer.MatrixLayers = &observer.MatrixToLayers{
		Observers: []observer.Matrix{matrix1, matrix2, matrix3},
	}

	// A GridLayers observer, wrapping the MatrixLayers observer
	var _ observer.GridLayers = &observer.LayersToGrid{
		Observer: layers,
	}
	// Output:
}
