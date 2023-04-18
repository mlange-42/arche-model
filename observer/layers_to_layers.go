package observer

import (
	"github.com/mlange-42/arche/ecs"
)

// LayersToLayers creates an observer that serves as adapter from a [MatrixLayers] observer to a [GridLayers] observer.
func LayersToLayers(obs MatrixLayers, origin *[2]float64, cellsize *[2]float64) GridLayers {
	m := layersToLayers{
		Observer: obs,
	}
	if origin != nil {
		m.Origin = *origin
	}
	if cellsize == nil {
		m.CellSize = [2]float64{1, 1}
	} else {
		m.CellSize = *cellsize
	}
	return &m
}

// layersToLayers is an observer that serves as adapter from a [MatrixLayers] observer to a [GridLayers] observer.
type layersToLayers struct {
	Observer MatrixLayers // The wrapped MatrixLayers observers.
	Origin   [2]float64   // Origin. Optional, defaults to (0, 0)
	CellSize [2]float64   // CellSize. Optional, defaults to (1, 1).
}

// Initialize the child observer.
func (o *layersToLayers) Initialize(w *ecs.World) {
	o.Observer.Initialize(w)
}

// Update the child observer.
func (o *layersToLayers) Update(w *ecs.World) {
	o.Observer.Update(w)
}

// Dims returns the matrix dimensions.
func (o *layersToLayers) Dims() (int, int) {
	return o.Observer.Dims()
}

// Layers returns the number of layers.
func (o *layersToLayers) Layers() int {
	return o.Observer.Layers()
}

// Values for the current model tick.
func (o *layersToLayers) Values(w *ecs.World) [][]float64 {
	return o.Observer.Values(w)
}

// X axis coordinates.
func (o *layersToLayers) X(c int) float64 {
	return o.Origin[0] + o.CellSize[0]*float64(c)
}

// Y axis coordinates.
func (o *layersToLayers) Y(r int) float64 {
	return o.Origin[1] + o.CellSize[1]*float64(r)
}
