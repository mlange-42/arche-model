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
	x        []float64
	y        []float64
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

	o.x = make([]float64, o.cols)
	o.y = make([]float64, o.rows)
	o.values = make([]float64, o.cols*o.rows)

	for i := 0; i < o.cols; i++ {
		o.x[i] = o.xOrigin + o.cellsize*float64(i)
	}
	for i := 0; i < o.cols; i++ {
		o.y[i] = o.yOrigin + o.cellsize*float64(i)
	}
}

func (o *GridObserver) Update(w *ecs.World) {}

func (o *GridObserver) Dims() (int, int) {
	return o.cols, o.rows
}

func (o *GridObserver) X() []float64 {
	return o.x
}

func (o *GridObserver) Y() []float64 {
	return o.y
}

func (o *GridObserver) Values(w *ecs.World) []float64 {
	for idx := 0; idx < len(o.values); idx++ {
		x := o.x[idx%o.cols]
		y := o.y[idx/o.cols]
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
