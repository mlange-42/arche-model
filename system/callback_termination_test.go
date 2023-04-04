package system_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestCallbackTermination(t *testing.T) {
	m := model.New()

	m.AddSystem(&system.CallbackTermination{
		Callback: func(t int64) bool {
			return t >= 99
		}},
	)

	m.Run()

	time := ecs.GetResource[resource.Tick](&m.World)
	assert.Equal(t, 100, int(time.Tick))
}

func ExampleCallbackTermination() {
	m := model.New()

	m.AddSystem(&system.CallbackTermination{
		Callback: func(t int64) bool {
			return t >= 99
		}},
	)

	m.Run()
	// Output:
}
