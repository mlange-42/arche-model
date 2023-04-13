package resource

import (
	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

// Rand is a PRNG resource to be used in [System] implementations.
type Rand struct {
	rand.Source // Source to use for PRNGs in [System] implementations.
}

// Tick is a resource holding the model's time step.
type Tick struct {
	Tick int64 // The current model tick.
}

// Termination is a resource holding whether the simulation should terminate after the current step.
type Termination struct {
	Terminate bool // Whether the simulation run is finished. Can be set by systems.
}

// SelectedEntity is a resource holding the currently selected entity.
// Primarily for the use in UI systems for inspecting entities.
type SelectedEntity struct {
	Selected ecs.Entity
}
