package model

import (
	"time"

	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

// Model is the top-level ECS entrypoint.
//
// Model provides access to the ECS world, and manages the scheduling of [System] and [UISystem] instances.
// The [Systems] scheduler, model [Time] and a central [Rand] PRNG source can be accessed by systems as resources.
type Model struct {
	Systems           // Systems manager and scheduler
	World   ecs.World // The ECS world
	rand    Rand
	time    Time
}

// New creates a new model.
func New(config ...ecs.Config) *Model {
	var mod = Model{
		World: ecs.NewWorld(config...),
	}
	mod.Fps = 30
	mod.Tps = 0
	mod.Systems.world = &mod.World

	mod.rand = Rand{rand.NewSource(uint64(time.Now().UnixNano()))}
	ecs.AddResource(&mod.World, &mod.rand)
	mod.time = Time{}
	ecs.AddResource(&mod.World, &mod.time)
	ecs.AddResource(&mod.World, &mod.Systems)

	return &mod
}

// Seed sets the random seed of the model.
// Call without an argument to seed from the current time.
//
// Systems should always use the [Rand] resource for PRNGs.
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

// Run runs the model.
func (m *Model) Run() {
	m.Systems.run()
}

// Reset resets the world and removes all systems.
func (m *Model) Reset() {
	m.World.Reset()
	m.Systems.reset()

	m.rand = Rand{rand.NewSource(uint64(time.Now().UnixNano()))}
	ecs.AddResource(&m.World, &m.rand)
	m.time = Time{}
	ecs.AddResource(&m.World, &m.time)
	ecs.AddResource(&m.World, &m.Systems)

	m.time.Tick = 0
	m.time.Finished = false
}
