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

func ExampleTime() {
	m := model.New()

	time := ecs.GetResource[model.Time](&m.World)

	fmt.Println(time.Tick)
	// Output: 0
}
