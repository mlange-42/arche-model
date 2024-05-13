package reporter_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
)

func ExampleSnapshotCSV() {
	// Create a new model.
	m := model.New()

	// Add a SnapshotCSV reporter with an Observer.
	m.AddSystem(&reporter.SnapshotCSV{
		Observer:    &ExampleSnapshotObserver{},
		FilePattern: "../out/test-%06d.csv",
		Sep:         ";",
	})

	// Add a termination system that ends the simulation.
	m.AddSystem(&system.FixedTermination{Steps: 100})

	// Run the simulation.
	m.Run()
}

// ExampleSnapshotObserver to generate some simple tables.
type ExampleSnapshotObserver struct{}

func (o *ExampleSnapshotObserver) Initialize(w *ecs.World) {}
func (o *ExampleSnapshotObserver) Update(w *ecs.World)     {}
func (o *ExampleSnapshotObserver) Header() []string {
	return []string{"A", "B", "C"}
}
func (o *ExampleSnapshotObserver) Values(w *ecs.World) [][]float64 {
	return [][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
}
