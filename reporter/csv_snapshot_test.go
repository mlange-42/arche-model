package reporter_test

import (
	"os"
	"testing"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/reporter"
	"github.com/mlange-42/arche-model/system"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotCSV(t *testing.T) {
	m := model.New()

	m.AddSystem(&reporter.SnapshotCSV{
		Observer:       &ExampleSnapshotObserver{},
		FilePattern:    "../out/test-%06d.csv",
		UpdateInterval: 10,
	})
	m.AddSystem(&system.FixedTermination{Steps: 100})

	m.Run()

	_, err := os.Stat("../out/test-000000.csv")
	assert.Nil(t, err)
	_, err = os.Stat("../out/test-000090.csv")
	assert.Nil(t, err)
}
