package reporter

import (
	"fmt"

	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// Print reporter to print a table row per time step.
type Print struct {
	Observer       observer.Row // Observer to get data from.
	UpdateInterval int          // Update/print interval in model ticks.
	header         []string
	step           int64
}

// Initialize the system
func (s *Print) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	s.header = s.Observer.Header()
	s.step = 0
}

// Update the system
func (s *Print) Update(w *ecs.World) {
	s.Observer.Update(w)
	if s.step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		fmt.Printf("%v\n%v\n", s.header, values)
	}
	s.step++
}

// Finalize the system
func (s *Print) Finalize(w *ecs.World) {}
