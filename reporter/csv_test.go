package reporter_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
)

func ExampleCSV() {
	m := model.New()

	m.AddSystem(&reporter.CSV{
		Observer:       &ExampleObserver{},
		File:           "../out/test.csv",
		Sep:            ";",
		UpdateInterval: 10,
	})

	m.AddSystem(&system.FixedTermination{Steps: 100})

	m.Run()
	// Output:
}

type ExampleObserver struct{}

func (o *ExampleObserver) Initialize(w *ecs.World) {}
func (o *ExampleObserver) Update(w *ecs.World)     {}
func (o *ExampleObserver) Header(w *ecs.World) []string {
	return []string{"A", "B", "C"}
}
func (o *ExampleObserver) Values(w *ecs.World) []float64 {
	return []float64{1, 2, 3}
}
