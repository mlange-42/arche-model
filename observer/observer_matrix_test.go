package observer_test

import (
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// Example observer, reporting a matrix with z = i^2 + j.
type MatrixObserver struct {
	cols   int
	rows   int
	values []float64
}

func (o *MatrixObserver) Initialize(w *ecs.World) {
	o.cols = 24
	o.rows = 16
	o.values = make([]float64, o.cols*o.rows)
}

func (o *MatrixObserver) Update(w *ecs.World) {}

func (o *MatrixObserver) Dims() (int, int) {
	return o.cols, o.rows
}

func (o *MatrixObserver) Values(w *ecs.World) []float64 {
	for idx := 0; idx < len(o.values); idx++ {
		i := idx % o.cols
		j := idx / o.cols
		o.values[idx] = float64(i*i + j)
	}
	return o.values
}

func ExampleMatrix() {
	var _ observer.Matrix = &MatrixObserver{}
	// Output:
}
