package system_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestFixedTermination(t *testing.T) {
	m := model.New()

	m.AddSystem(&system.FixedTermination{Steps: 100})

	m.Run()

	time := ecs.GetResource[model.Tick](&m.World)
	assert.Equal(t, 100, int(time.Tick))
}

func ExampleFixedTermination() {
	m := model.New()

	m.AddSystem(&system.FixedTermination{Steps: 100})

	m.Run()
	// Output:
}
