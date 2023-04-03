package model

import (
	"time"

	"golang.org/x/exp/rand"

	"github.com/faiface/pixel/pixelgl"
	"github.com/mlange-42/arche/ecs"
)

// Model is the top-level ecs entrypoint.
type Model struct {
	ecs.World
	systems
	rand Rand
}

// New creates a new model.
func New(config ...ecs.Config) *Model {
	var mod = Model{
		World: ecs.NewWorld(config...),
	}
	mod.Fps = 30
	mod.Tps = 0
	mod.systems.model = &mod

	mod.rand = Rand{rand.NewSource(uint64(time.Now().UnixNano()))}
	ecs.AddResource(&mod.World, &mod.rand)

	return &mod
}

// Seed sets the random seed of the model.
// Call without an argument to seed from the current time.
func (m *Model) Seed(seed ...uint64) {
	switch len(seed) {
	case 0:
		m.rand.Seed(uint64(time.Now().UnixNano()))
	case 1:
		m.rand.Seed(seed[0])
	default:
		panic("can only use a single random seed")
	}
}

// Run runs a model
func (m *Model) Run() {
	pixelgl.Run(m.systems.run)
}
