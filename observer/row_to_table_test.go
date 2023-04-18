package observer_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

type rowObs struct {
	header []string
	values []float64
}

func (o *rowObs) Initialize(w *ecs.World) {}

func (o *rowObs) Update(w *ecs.World) {}

func (o *rowObs) Header() []string {
	return o.header
}

func (o *rowObs) Values(w *ecs.World) []float64 {
	return o.values
}

func TestRowToTable(t *testing.T) {
	m := model.New()

	var row observer.Row = &rowObs{
		header: []string{"A", "B"},
		values: []float64{1, 2},
	}

	var table observer.Table = &observer.RowToTable{Observer: row}

	table.Initialize(&m.World)
	table.Update(&m.World)

	h := table.Header()
	assert.Equal(t, []string{"A", "B"}, h)

	v := table.Values(&m.World)
	assert.Equal(t, [][]float64{{1, 2}}, v)
}
