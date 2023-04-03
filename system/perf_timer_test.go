package system_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func ExamplePerfTimer() {
	m := model.New()

	m.AddSystem(&system.PerfTimer{UpdateInterval: 10})
	m.AddSystem(&system.FixedTermination{Steps: 30})

	// m.Run()

	// Output:
}
