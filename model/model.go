package model

import (
	"time"

	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

// Model is the top-level ECS entrypoint.
//
// Model provides access to the ECS world, and manages the scheduling
// of [System] and [UISystem] instances via [Systems].
// [System] instances are updated with a frequency given by TPS.
// [UISystem] instances are updated independently of normal systems,
// with a frequency given by FPS.
//
// The [Systems] scheduler, the model's [resource.Tick], [resource.Termination]
// and a central [resource.Rand] PRNG source can be accessed by systems as resources.
type Model struct {
	Systems             // Systems manager and scheduler
	World     ecs.World // The ECS world
	rand      resource.Rand
	time      resource.Tick
	terminate resource.Termination
}

// New creates a new model.
//
// See [ecs.NewWorld] for the arguments.
func New(initialCapacity ...int) *Model {
	var mod = Model{
		World: ecs.NewWorld(initialCapacity...),
	}
	mod.FPS = 30
	mod.TPS = 0
	mod.Systems.world = &mod.World

	mod.rand = resource.Rand{
		Source: rand.NewSource(uint64(time.Now().UnixNano())),
	}
	ecs.AddResource(&mod.World, &mod.rand)
	mod.time = resource.Tick{}
	ecs.AddResource(&mod.World, &mod.time)
	mod.terminate = resource.Termination{}
	ecs.AddResource(&mod.World, &mod.terminate)

	ecs.AddResource(&mod.World, &mod.Systems)

	return &mod
}

// Seed sets the random seed of the model's [resource.Rand].
// Call without an argument to seed from the current time.
//
// Systems should always use the Rand resource for PRNGs.
func (m *Model) Seed(seed ...uint64) *Model {
	switch len(seed) {
	case 0:
		m.rand.Seed(uint64(time.Now().UnixNano()))
	case 1:
		m.rand.Seed(seed[0])
	default:
		panic("can only use a single random seed")
	}
	return m
}

// Run the model, updating systems and ui systems according to Model.TPS and Model.FPS, respectively.
// Initializes the model if it is not already initialized.
// Finalizes the model after the run.
//
// Runs until Terminate in the resource resource.Termination is set to true
// (see [resource.Termination]).
//
// To perform updates manually, see [Model.Update] and [Model.UpdateUI],
// as well as [Model.Initialize] and [Model.Finalize].
func (m *Model) Run() {
	m.Systems.run()
}

// Initialize the model.
func (m *Model) Initialize() {
	m.Systems.initialize()
}

// Update the model's systems.
// Return whether the run should continue.
//
// Ignores Model.TPS.
//
// Panics if [Model.Initialize] was not called.
func (m *Model) Update() bool {
	return m.Systems.updateSystems()
}

// UpdateUI the model's UI systems.
//
// Ignores Model.FPS.
//
// Panics if [Model.Initialize] was not called.
func (m *Model) UpdateUI() {
	m.Systems.updateUISystems()
}

// Finalize the model.
func (m *Model) Finalize() {
	m.Systems.finalize()
}

// Reset resets the world and removes all systems.
//
// Can be used to run systematic simulations without the need to re-allocate memory for each run.
// Accelerates re-populating the world by a factor of 2-3.
func (m *Model) Reset() {
	m.World.Reset()
	m.Systems.reset()

	m.rand = resource.Rand{
		Source: rand.NewSource(uint64(time.Now().UnixNano())),
	}
	ecs.AddResource(&m.World, &m.rand)
	m.time = resource.Tick{}
	ecs.AddResource(&m.World, &m.time)
	m.terminate = resource.Termination{}
	ecs.AddResource(&m.World, &m.terminate)

	ecs.AddResource(&m.World, &m.Systems)
}
