package observer_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
	"github.com/stretchr/testify/assert"
)

func TestGridToLayers(t *testing.T) {
	m := model.New()

	var mat1 observer.Matrix = &matObs{}
	var mat2 observer.Matrix = &matObs{}
	var mat3 observer.Matrix = &matObs{}

	var grid1 observer.Grid = observer.MatrixToGrid(mat1, nil, nil)
	var grid2 observer.Grid = observer.MatrixToGrid(mat2, nil, nil)
	var grid3 observer.Grid = observer.MatrixToGrid(mat3, nil, nil)

	var layers observer.GridLayers = observer.GridToLayers(grid1, grid2, grid3)

	layers.Initialize(&m.World)
	layers.Update(&m.World)

	assert.Equal(t, 3, layers.Layers())

	w, h := layers.Dims()

	assert.Equal(t, 30, w)
	assert.Equal(t, 20, h)

	assert.Equal(t, 1.0, layers.X(1))
	assert.Equal(t, 1.0, layers.Y(1))

	data := layers.Values(&m.World)
	assert.Equal(t, 3, len(data))
	assert.Equal(t, 20*30, len(data[0]))
}

func TestGridToLayersFail(t *testing.T) {
	m := model.New()

	var mat1 observer.Matrix = &matObs{}
	var mat2 *matObs = &matObs{}
	mat2.Rows = 15

	var grid1 observer.Grid = observer.MatrixToGrid(mat1, nil, nil)
	var grid2 observer.Grid = observer.MatrixToGrid(mat2, nil, nil)

	var layers observer.GridLayers = observer.GridToLayers(grid1, grid2)
	assert.Panics(t, func() { layers.Initialize(&m.World) })

	assert.Panics(t, func() { observer.GridToLayers() })
}
