package reporter_test

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
)

func ExampleCallback() {
	// Create a new model.
	m := model.New()

	data := [][]float64{}

	// Add a Print reporter with an Observer.
	m.AddSystem(&reporter.Callback{
		Observer: &ExampleObserver{},
		Callback: func(step int, row []float64) {
			data = append(data, row)
		},
		HeaderCallback: func(header []string) {},
	})

	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 3})

	// Run the simulation.
	m.Run()

	fmt.Println(data)
	// Output:
	// [[1 2 3] [1 2 3] [1 2 3]]
}
