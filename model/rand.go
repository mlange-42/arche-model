package model

import "golang.org/x/exp/rand"

// Rand is a PRNG resource to be used in systems.
type Rand struct {
	rand.Source
}
