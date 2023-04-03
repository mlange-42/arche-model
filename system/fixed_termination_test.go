package system_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func ExampleFixedTermination() {
	m := model.New()

	m.AddSystem(&system.FixedTermination{Steps: 100})

	m.Run()
	// Output:
}
