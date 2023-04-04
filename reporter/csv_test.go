package reporter_test

import (
	"os"
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/stretchr/testify/assert"
)

func TestCSV(t *testing.T) {
	m := model.New()

	m.AddSystem(&reporter.CSV{
		Observer:       &ExampleObserver{},
		File:           "../out/test.csv",
		Sep:            ";",
		UpdateInterval: 10,
	})
	m.AddSystem(&system.FixedTermination{Steps: 100})

	m.Run()

	_, err := os.Stat("../out/test.csv")
	assert.Nil(t, err)
}
