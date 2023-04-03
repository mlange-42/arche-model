package reporters

import (
	"github.com/mlange-42/arche/ecs"
)

// Observer interface
type Observer interface {
	Initialize(w *ecs.World)
	Update(w *ecs.World)
	Header(w *ecs.World) []string
	Values(w *ecs.World) []float64
}

// MatrixObserver interface
type MatrixObserver interface {
	Initialize(w *ecs.World)
	Update(w *ecs.World)
	Header(w *ecs.World) []string
	Values(w *ecs.World) [][]float64
}
