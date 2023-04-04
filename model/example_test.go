package model_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func Example() {
	m := model.New()
	m.Seed(123)

	m.AddSystem(&system.FixedTermination{
		Steps: 100,
	})

	m.Run()
	// Output:
}
