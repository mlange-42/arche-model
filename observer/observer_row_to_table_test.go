package observer_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
)

func ExampleRowToTable() {
	m := model.New()

	// A Row observer
	var row observer.Row = &RowObserver{}
	var _ []float64 = row.Values(&m.World)

	// A RowToTable observer, wrapping the Row observer
	var table observer.Table = &observer.RowToTable{Observer: row}
	var _ [][]float64 = table.Values(&m.World)
	// Output:
}
