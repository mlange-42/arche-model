package system_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func TestPerfTimer(t *testing.T) {
	m := model.New()

	m.AddSystem(&system.PerfTimer{UpdateInterval: 10, Stats: true})
	m.AddSystem(&system.FixedTermination{Steps: 30})

	m.Run()
}

func ExamplePerfTimer() {
	m := model.New()

	m.AddSystem(&system.PerfTimer{UpdateInterval: 10})
	m.AddSystem(&system.FixedTermination{Steps: 30})

	// Uncomment the next line.

	// m.Run()
	// Output:
}
