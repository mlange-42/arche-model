package reporter_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
)

func ExamplePrint() {
	m := model.New()

	m.AddSystem(&reporter.Print{
		Observer:       &ExampleObserver{},
		UpdateInterval: 10,
	})

	m.AddSystem(&system.FixedTermination{Steps: 20})

	m.Run()
	// Output:
	// [A B C]
	// [1 2 3]
	// [A B C]
	// [1 2 3]
}
