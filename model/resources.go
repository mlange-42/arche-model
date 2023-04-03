package model

import "golang.org/x/exp/rand"

// Rand is a PRNG resource to be used in systems.
type Rand struct {
	rand.Source
}

// Time is a resource holding the model's time step.
type Time struct {
	Tick     int64
	Finished bool
}
