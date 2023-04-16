package observer

import "github.com/mlange-42/arche/ecs"

// RowToTable is an observer that serves as adapter from a [Row] observer to a [Table] observer.
type RowToTable struct {
	Observer Row
	values   [1][]float64
}

// Initialize the child observer.
func (o *RowToTable) Initialize(w *ecs.World) {
	o.Observer.Initialize(w)
}

// Update the child observer.
func (o *RowToTable) Update(w *ecs.World) {
	o.Observer.Update(w)
}

// Header / column names of the child observer in the same order as data values.
func (o *RowToTable) Header() []string {
	return o.Observer.Header()
}

// Values for the current model tick in table format.
func (o *RowToTable) Values(w *ecs.World) [][]float64 {
	o.values[0] = o.Observer.Values(w)
	return o.values[:]
}
