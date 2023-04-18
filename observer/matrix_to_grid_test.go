package observer_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

type matObs struct {
	values []float64
	Cols   int
	Rows   int
}

func (o *matObs) Initialize(w *ecs.World) {
	if o.Cols == 0 {
		o.Cols = 30
	}
	if o.Rows == 0 {
		o.Rows = 20
	}
	o.values = make([]float64, o.Cols*o.Rows)
}

func (o *matObs) Update(w *ecs.World) {}

func (o *matObs) Dims() (int, int) {
	return o.Cols, o.Rows
}

func (o *matObs) Values(w *ecs.World) []float64 {
	return o.values
}

func TestMatrixToGrid(t *testing.T) {
	m := model.New()

	var mat observer.Matrix = &matObs{}
	var grid observer.Grid = observer.MatrixToGrid(
		mat,
		&[...]float64{1, 2},
		&[...]float64{5, 10},
	)

	grid.Initialize(&m.World)
	grid.Update(&m.World)

	v := grid.Values(&m.World)
	assert.Equal(t, make([]float64, 20*30), v)

	w, h := grid.Dims()
	assert.Equal(t, 30, w)
	assert.Equal(t, 20, h)

	assert.Equal(t, 1.0, grid.X(0))
	assert.Equal(t, 2.0, grid.Y(0))

	assert.Equal(t, 6.0, grid.X(1))
	assert.Equal(t, 12.0, grid.Y(1))

	grid = observer.MatrixToGrid(mat, nil, nil)
	grid.Initialize(&m.World)
	assert.Equal(t, 1.0, grid.X(1))
	assert.Equal(t, 1.0, grid.Y(1))

	data := grid.Values(&m.World)
	assert.Equal(t, 20*30, len(data))
}
