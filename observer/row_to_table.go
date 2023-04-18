package observer

import "github.com/mlange-42/arche/ecs"

// RowToTable creates an observer that serves as adapter from a [Row] observer to a [Table] observer.
func RowToTable(row Row) Table {
	return &rowToTable{
		Observer: row,
	}
}

// rowToTable is an observer that serves as adapter from a [Row] observer to a [Table] observer.
type rowToTable struct {
	Observer Row
	values   [1][]float64
}

// Initialize the child observer.
func (o *rowToTable) Initialize(w *ecs.World) {
	o.Observer.Initialize(w)
}

// Update the child observer.
func (o *rowToTable) Update(w *ecs.World) {
	o.Observer.Update(w)
}

// Header / column names of the child observer in the same order as data values.
func (o *rowToTable) Header() []string {
	return o.Observer.Header()
}

// Values for the current model tick in table format.
func (o *rowToTable) Values(w *ecs.World) [][]float64 {
	o.values[0] = o.Observer.Values(w)
	return o.values[:]
}
