package model_test

import (
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/system"
)

func TestModel(t *testing.T) {
	m := model.New()

	for i := 0; i < 3; i++ {
		m.Reset()
		m.Seed(123)

		m.AddSystem(&system.FixedTermination{
			Steps: 10,
		})

		m.Run()
	}
}
