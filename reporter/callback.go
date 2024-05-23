package reporter

import (
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// RowCallback reporter calling a function on each update, using an [observer.Row].
type RowCallback struct {
	Observer       observer.Row                  // Observer to get data from.
	UpdateInterval int                           // Update interval in model ticks.
	HeaderCallback func(header []string)         // Called with the header of the observer during initialization.
	Callback       func(step int, row []float64) // Called with step and data row on each update (subject to UpdateInterval).
	Final          bool                          // Whether Callback should be called on finalization only, instead of on every tick.
	step           int64
}

// Initialize the system
func (s *RowCallback) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	if s.UpdateInterval == 0 {
		s.UpdateInterval = 1
	}
	if s.HeaderCallback != nil {
		s.HeaderCallback(s.Observer.Header())
	}
	s.step = 0
}

// Update the system
func (s *RowCallback) Update(w *ecs.World) {
	s.Observer.Update(w)

	if !s.Final && s.step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		s.Callback(int(s.step), values)
	}

	s.step++
}

// Finalize the system
func (s *RowCallback) Finalize(w *ecs.World) {
	if !s.Final {
		return
	}
	values := s.Observer.Values(w)
	s.Callback(int(s.step), values)
}

// RowCallback reporter calling a function on each update, using an [observer.Table].
type TableCallback struct {
	Observer       observer.Table                    // Observer to get data from.
	UpdateInterval int                               // Update interval in model ticks.
	HeaderCallback func(header []string)             // Called with the header of the observer during initialization.
	Callback       func(step int, table [][]float64) // Called with step and data table on each update (subject to UpdateInterval).
	Final          bool                              // Whether Callback should be called on finalization only, instead of on every tick.
	step           int64
}

// Initialize the system
func (s *TableCallback) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	if s.UpdateInterval == 0 {
		s.UpdateInterval = 1
	}
	if s.HeaderCallback != nil {
		s.HeaderCallback(s.Observer.Header())
	}
	s.step = 0
}

// Update the system
func (s *TableCallback) Update(w *ecs.World) {
	s.Observer.Update(w)

	if !s.Final && s.step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		s.Callback(int(s.step), values)
	}

	s.step++
}

// Finalize the system
func (s *TableCallback) Finalize(w *ecs.World) {
	if !s.Final {
		return
	}
	values := s.Observer.Values(w)
	s.Callback(int(s.step), values)
}
