package observer_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// Example observer, reporting a grid with z = x^2 + y.
type GridObserver struct {
	cols     int
	rows     int
	cellsize float64
	xOrigin  float64
	yOrigin  float64
	values   []float64
}

func (o *GridObserver) Initialize(w *ecs.World) {
	// In a real example, we would get these values from the world,
	// e.g. from a resource representing a global grid.
	o.cols = 24
	o.rows = 16
	o.cellsize = 1000
	o.xOrigin = 123_000
	o.yOrigin = 234_000

	o.values = make([]float64, o.cols*o.rows)
}

func (o *GridObserver) Update(w *ecs.World) {}

func (o *GridObserver) Dims() (int, int) {
	return o.cols, o.rows
}

func (o *GridObserver) X(c int) float64 {
	return o.xOrigin + o.cellsize*float64(c)
}

func (o *GridObserver) Y(r int) float64 {
	return o.yOrigin + o.cellsize*float64(r)
}

func (o *GridObserver) Values(w *ecs.World) []float64 {
	for idx := 0; idx < len(o.values); idx++ {
		x := o.X(idx % o.cols)
		y := o.Y(idx / o.cols)
		o.values[idx] = float64(x*x + y)
	}
	return o.values
}

func ExampleGrid() {
	m := model.New()

	var obs observer.Grid = &GridObserver{}
	_ = obs.Values(&m.World)
	// Output:
}
