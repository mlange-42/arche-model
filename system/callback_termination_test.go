package system_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func ExampleCallbackTermination() {
	m := model.New()

	m.AddSystem(&system.CallbackTermination{
		Callback: func(t int64) bool {
			return t >= 100
		}},
	)

	m.Run()
	// Output:
}
