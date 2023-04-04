package model_test

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"golang.org/x/exp/rand"
)

func ExampleRand() {
	m := model.New()

	src := ecs.GetResource[model.Rand](&m.World)
	rng := rand.New(src.Source)
	_ = rng.NormFloat64()
	// Output:
}

func ExampleTick() {
	m := model.New()

	tick := ecs.GetResource[model.Tick](&m.World)

	fmt.Println(tick.Tick)
	// Output: 0
}

func ExampleTermination() {
	m := model.New()

	term := ecs.GetResource[model.Termination](&m.World)

	fmt.Println(term.Terminate)
	// Output: false
}
