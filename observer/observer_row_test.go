package observer_test

import (
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// Example observer, reporting the number of entities.
type RowObserver struct{}

func (o *RowObserver) Initialize(w *ecs.World) {}

func (o *RowObserver) Update(w *ecs.World) {}

func (o *RowObserver) Header() []string {
	return []string{"TotalEntities"}
}

func (o *RowObserver) Values(w *ecs.World) []float64 {
	query := w.Query(ecs.All())
	cnt := query.Count()
	query.Close()

	return []float64{float64(cnt)}
}

func ExampleRow() {
	var _ observer.Row = &RowObserver{}
	// Output:
}
