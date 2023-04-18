package observer

import (
	"github.com/mlange-42/arche/ecs"
)

// LayersToGrid is an observer that serves as adapter from a [MatrixLayers] observer to a [GridLayers] observer.
type LayersToGrid struct {
	Observer MatrixLayers // The wrapped MatrixLayers observers.
	Origin   [2]float64   // Origin. Optional, defaults to (0, 0)
	CellSize [2]float64   // CellSize. Optional, defaults to (1, 1).
}

// Initialize the child observer.
func (o *LayersToGrid) Initialize(w *ecs.World) {
	o.Observer.Initialize(w)

	if o.CellSize[0] == 0 {
		o.CellSize[0] = 1
	}
	if o.CellSize[1] == 0 {
		o.CellSize[1] = 1
	}
}

// Update the child observer.
func (o *LayersToGrid) Update(w *ecs.World) {
	o.Observer.Update(w)
}

// Dims returns the matrix dimensions.
func (o *LayersToGrid) Dims() (int, int) {
	return o.Observer.Dims()
}

// Layers returns the number of layers.
func (o *LayersToGrid) Layers() int {
	return o.Observer.Layers()
}

// Values for the current model tick.
func (o *LayersToGrid) Values(w *ecs.World) [][]float64 {
	return o.Observer.Values(w)
}

// X axis coordinates.
func (o *LayersToGrid) X(c int) float64 {
	return o.Origin[0] + o.CellSize[0]*float64(c)
}

// Y axis coordinates.
func (o *LayersToGrid) Y(r int) float64 {
	return o.Origin[1] + o.CellSize[1]*float64(r)
}
