package reporters

import "github.com/mlange-42/arche-model/model"

// Observer interface
type Observer interface {
	Initialize(m *model.Model)
	Update(m *model.Model)
	Header(m *model.Model) []string
	Values(m *model.Model) []float64
}

// MatrixObserver interface
type MatrixObserver interface {
	Initialize(m *model.Model)
	Update(m *model.Model)
	Header(m *model.Model) []string
	Values(m *model.Model) [][]float64
}
