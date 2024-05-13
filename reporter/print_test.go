package reporter_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
)

func ExamplePrint() {
	// Create a new model.
	m := model.New()

	// Add a Print reporter with an Observer.
	m.AddSystem(&reporter.Print{
		Observer: &ExampleObserver{},
	})

	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 3})

	// Run the simulation.
	m.Run()
	// Output:
	// [A B C]
	// [1 2 3]
	// [A B C]
	// [1 2 3]
	// [A B C]
	// [1 2 3]
}
