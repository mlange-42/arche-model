package model

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/mlange-42/arche/ecs"
)

// Model is the top-level ecs entrypoint.
type Model struct {
	ecs.World
	systems
}

// New creates a new model.
func New(config ...ecs.Config) *Model {
	world := ecs.NewWorld(config...)
	var mod = Model{
		World: world,
	}
	mod.Fps = 30
	mod.Tps = 0
	mod.systems.model = &mod
	return &mod
}

// Seed sets the random seed of the model.
// Call without an argument to seed from the current time.
func (m *Model) Seed(seed ...int64) {
	switch len(seed) {
	case 0:
		rand.Seed(time.Now().UnixNano())
	case 1:
		rand.Seed(seed[0])
	default:
		panic("can only use a single random seed")
	}
}

// Run runs a model
func (m *Model) Run() {
	m.systems.model = m
	pixelgl.Run(m.systems.run)
}
