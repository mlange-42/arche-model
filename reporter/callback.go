package reporter

import (
	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// Callback reporter calling a function on each update.
type Callback struct {
	Observer       observer.Row // Observer to get data from.
	UpdateInterval int          // Update/print interval in model ticks.
	HeaderCallback func(header []string)
	Callback       func(step int, row []float64)
	step           int64
}

// Initialize the system
func (s *Callback) Initialize(w *ecs.World) {
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
func (s *Callback) Update(w *ecs.World) {
	s.Observer.Update(w)
	if s.step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		s.Callback(int(s.step), values)
	}
	s.step++
}

// Finalize the system
func (s *Callback) Finalize(w *ecs.World) {}
