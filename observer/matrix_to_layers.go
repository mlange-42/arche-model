package observer

import "github.com/mlange-42/arche/ecs"

// MatrixToLayers creates an observer that serves as adapter from multiple [Matrix] observers to a [MatrixLayers] observer.
func MatrixToLayers(obs ...Matrix) MatrixLayers {
	if len(obs) == 0 {
		panic("no observers given")
	}
	return &matrixToLayers{
		Observers: obs,
	}
}

// matrixToLayers is an observer that serves as adapter from multiple [Matrix] observers to a [MatrixLayers] observer.
type matrixToLayers struct {
	Observers []Matrix
	values    [][]float64
}

// Initialize the child observer.
func (o *matrixToLayers) Initialize(w *ecs.World) {
	for i, obs := range o.Observers {
		obs.Initialize(w)
		if i == 0 {
			continue
		}
		obs0 := o.Observers[0]
		w1, h1 := obs0.Dims()
		w2, h2 := obs.Dims()

		if w1 != w2 || h1 != h2 {
			panic("grids for layers have different dimensions")
		}
	}

	o.values = make([][]float64, len(o.Observers))
}

// Update the child observer.
func (o *matrixToLayers) Update(w *ecs.World) {
	for _, obs := range o.Observers {
		obs.Update(w)
	}
}

// Dims returns the matrix dimensions.
func (o *matrixToLayers) Dims() (int, int) {
	return o.Observers[0].Dims()
}

// Layers returns the number of layers.
func (o *matrixToLayers) Layers() int {
	return len(o.Observers)
}

// Values for the current model tick.
func (o *matrixToLayers) Values(w *ecs.World) [][]float64 {
	for i, obs := range o.Observers {
		o.values[i] = obs.Values(w)
	}
	return o.values
}
