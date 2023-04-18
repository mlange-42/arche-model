package observer_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
	"github.com/stretchr/testify/assert"
)

func TestMatrixToLayers(t *testing.T) {
	m := model.New()

	var mat1 observer.Matrix = &matObs{}
	var mat2 observer.Matrix = &matObs{}
	var mat3 observer.Matrix = &matObs{}

	var layers observer.MatrixLayers = observer.MatrixToLayers(mat1, mat2, mat3)

	layers.Initialize(&m.World)
	layers.Update(&m.World)

	assert.Equal(t, 3, layers.Layers())

	w, h := layers.Dims()

	assert.Equal(t, 30, w)
	assert.Equal(t, 20, h)

	data := layers.Values(&m.World)
	assert.Equal(t, 3, len(data))
	assert.Equal(t, 20*30, len(data[0]))
}

func TestMatrixToLayersFail(t *testing.T) {
	m := model.New()

	var mat1 observer.Matrix = &matObs{}
	var mat2 *matObs = &matObs{}
	mat2.Rows = 15

	var layers observer.MatrixLayers = observer.MatrixToLayers(mat1, mat2)
	assert.Panics(t, func() { layers.Initialize(&m.World) })

	assert.Panics(t, func() { observer.MatrixToLayers() })
}
