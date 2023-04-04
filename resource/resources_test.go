package resource_test

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

func ExampleRand() {
	m := model.New()

	src := ecs.GetResource[resource.Rand](&m.World)
	rng := rand.New(src.Source)
	_ = rng.NormFloat64()
	// Output:
}

func ExampleTick() {
	m := model.New()

	tick := ecs.GetResource[resource.Tick](&m.World)

	fmt.Println(tick.Tick)
	// Output: 0
}

func ExampleTermination() {
	m := model.New()

	term := ecs.GetResource[resource.Termination](&m.World)

	fmt.Println(term.Terminate)
	// Output: false
}
