package observer_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
	"github.com/stretchr/testify/assert"
)

func TestLayersToGrid(t *testing.T) {
	m := model.New()

	var mat1 observer.Matrix = &matObs{}
	var mat2 observer.Matrix = &matObs{}
	var mat3 observer.Matrix = &matObs{}

	var layers observer.MatrixLayers = observer.MatrixToLayers(mat1, mat2, mat3)

	var grid observer.GridLayers = observer.LayersToLayers(layers, &[2]float64{0, 0}, &[2]float64{1, 1})

	grid.Initialize(&m.World)
	grid.Update(&m.World)

	assert.Equal(t, 3, grid.Layers())

	w, h := grid.Dims()

	assert.Equal(t, 30, w)
	assert.Equal(t, 20, h)

	assert.Equal(t, 1.0, grid.X(1))
	assert.Equal(t, 1.0, grid.Y(1))

	data := grid.Values(&m.World)
	assert.Equal(t, 3, len(data))
	assert.Equal(t, 20*30, len(data[0]))
}
