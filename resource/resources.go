package resource

import (
	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

// Rand is a PRNG resource to be used in [System] implementations.
//
// This resource is provided by [github.com/mlange-42/arche-model/model.Model] per default.
type Rand struct {
	rand.Source // Source to use for PRNGs in [System] implementations.
}

// Tick is a resource holding the model's time step.
//
// This resource is provided by [github.com/mlange-42/arche-model/model.Model] per default.
type Tick struct {
	Tick int64 // The current model tick.
}

// Termination is a resource holding whether the simulation should terminate after the current step.
//
// This resource is provided by [github.com/mlange-42/arche-model/model.Model] per default.
type Termination struct {
	Terminate bool // Whether the simulation run is finished. Can be set by systems.
}

// SelectedEntity is a resource holding the currently selected entity.
//
// The primarily purpose is communication between UI systems, e.g. for entity inspection or manipulation by the user.
type SelectedEntity struct {
	Selected ecs.Entity
}
