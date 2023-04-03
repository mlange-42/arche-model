package model

import "golang.org/x/exp/rand"

// Rand is a PRNG resource to be used in [System] implementations.
type Rand struct {
	rand.Source // Source to use for PRNGs in [System] implementations.
}

// Time is a resource holding the model's time step.
type Time struct {
	Tick     int64 // The current model tick.
	Finished bool  // Whether the model run is finished. Can be set by systems.
}
