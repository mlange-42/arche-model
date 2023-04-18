package observer

import (
	"github.com/mlange-42/arche/ecs"
)

// MatrixToGrid creates an observer that serves as adapter from a [Matrix] observer to a [Grid] observer.
func MatrixToGrid(obs Matrix, origin *[2]float64, cellsize *[2]float64) Grid {
	m := matrixToGrid{
		Observer: obs,
	}
	if origin != nil {
		m.Origin = *origin
	}
	if cellsize != nil {
		m.CellSize = *cellsize
	}
	return &m
}

// matrixToGrid is an observer that serves as adapter from a [Matrix] observer to a [Grid] observer.
type matrixToGrid struct {
	Observer Matrix     // The wrapped Matrix observer.
	Origin   [2]float64 // Origin. Optional, defaults to (0, 0)
	CellSize [2]float64 // CellSize. Optional, defaults to (1, 1).
}

// Initialize the child observer.
func (o *matrixToGrid) Initialize(w *ecs.World) {
	o.Observer.Initialize(w)

	if o.CellSize[0] == 0 {
		o.CellSize[0] = 1
	}
	if o.CellSize[1] == 0 {
		o.CellSize[1] = 1
	}
}

// Update the child observer.
func (o *matrixToGrid) Update(w *ecs.World) {
	o.Observer.Update(w)
}

// Dims returns the matrix dimensions.
func (o *matrixToGrid) Dims() (int, int) {
	return o.Observer.Dims()
}

// Values for the current model tick.
func (o *matrixToGrid) Values(w *ecs.World) []float64 {
	return o.Observer.Values(w)
}

// X axis coordinates.
func (o *matrixToGrid) X(c int) float64 {
	return o.Origin[0] + o.CellSize[0]*float64(c)
}

// Y axis coordinates.
func (o *matrixToGrid) Y(r int) float64 {
	return o.Origin[1] + o.CellSize[1]*float64(r)
}
