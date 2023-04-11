package model_test

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
	"github.com/mlange-42/arche/ecs"
)

func ExampleSystems() {
	// Create a new model.
	m := model.New()

	// The model contains Systems as an embed, TPS and FPS are accessible through the model directly.
	m.TPS = 1000
	m.FPS = 60

	// Create a system
	sys := system.FixedTermination{
		Steps: 10,
	}

	// Add the system the usual way, through the model.
	// The model contains Systems as an embed, so actually [Systems.AddSystem] is called.
	m.AddSystem(&sys)

	// Inside systems, [Systems] can be accessed as a resource.
	systems := ecs.GetResource[model.Systems](&m.World)

	// Pause the simulation, e.g. based on user input.
	systems.Paused = true

	// Remove the system using the resource.
	systems.RemoveSystem(&sys)
	// Output:
}
