package reporter_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
)

func ExampleCSV() {
	// Create a new model.
	m := model.New()

	// Add a CSV reporter with an Observer.
	m.AddSystem(&reporter.CSV{
		Observer: &ExampleObserver{},
		File:     "../out/test.csv",
		Sep:      ";",
	})

	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 100})

	// Run the simulation.
	m.Run()
}

// ExampleObserver to generate some simple time series.
type ExampleObserver struct{}

func (o *ExampleObserver) Initialize(w *ecs.World) {}
func (o *ExampleObserver) Update(w *ecs.World)     {}
func (o *ExampleObserver) Header() []string {
	return []string{"A", "B", "C"}
}
func (o *ExampleObserver) Values(w *ecs.World) []float64 {
	return []float64{1, 2, 3}
}
