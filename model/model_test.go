package model_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {
	m := model.New()

	for i := 0; i < 3; i++ {
		m.Reset()
		m.Seed(123)

		m.AddSystem(&system.FixedTermination{
			Steps: 10,
		})

		m.Run()
	}
}

func TestModelStep(t *testing.T) {
	m := model.New()

	for i := 0; i < 3; i++ {
		m.Reset()
		m.Seed(123)

		m.AddSystem(&system.FixedTermination{
			Steps: 10,
		})

		assert.Panics(t, func() { m.Update() })
		assert.Panics(t, func() { m.UpdateUI() })

		m.Initialize()

		m.Paused = true
		m.Update()
		m.Paused = false

		for m.Update() {
			m.UpdateUI()
		}
		m.Finalize()
	}
}

func TestModelSeed(t *testing.T) {
	m := model.New()
	m.Seed(123)

	rand := ecs.GetResource[resource.Rand](&m.World)
	r1 := rand.Uint64()

	m.Seed(123)
	assert.Equal(t, r1, rand.Uint64())

	m.Seed()
	assert.NotEqual(t, r1, rand.Uint64())

	assert.Panics(t, func() { m.Seed(1, 2, 3) })
}

func ExampleModel() {
	// Create a new, seeded model.
	m := model.New().Seed(123)

	// Add systems.
	m.AddSystem(&system.FixedTermination{
		Steps: 100,
	})

	// Run the simulation.
	m.Run()
	// Output:
}

func ExampleModel_Reset() {
	// Create a new model.
	m := model.New()

	// Do many model runs.
	for i := 0; i < 10; i++ {
		// Reset the model to clear entities, systems etc. before the run.
		m.Reset()

		// Seed the model for the run.
		m.Seed(uint64(i))

		// Add systems.
		m.AddSystem(&system.FixedTermination{
			Steps: 100,
		})

		// Run the simulation.
		m.Run()

	}
	// Output:
}
