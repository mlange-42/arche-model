package observer_test

import (
	"github.com/mlange-42/arche-model/observer"
)

func ExampleGridToLayers() {
	// Multiple Grid observers
	var grid1 observer.Grid = &GridObserver{}
	var grid2 observer.Grid = &GridObserver{}
	var grid3 observer.Grid = &GridObserver{}

	// A MatrixToGrid observer, wrapping the Grid observers
	var _ observer.GridLayers = observer.GridToLayers(grid1, grid2, grid3)
	// Output:
}
