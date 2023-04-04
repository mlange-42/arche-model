package system

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// FixedTermination system.
//
// Terminates a model run after a fixed number of ticks.
//
// Expects a resource of type [model.Termination].
type FixedTermination struct {
	Steps   int64 // Number of simulation ticks to run.
	termRes generic.Resource[resource.Termination]
	step    int64
}

// Initialize the system
func (s *FixedTermination) Initialize(w *ecs.World) {
	s.termRes = generic.NewResource[resource.Termination](w)
	s.step = 0
}

// Update the system
func (s *FixedTermination) Update(w *ecs.World) {
	term := s.termRes.Get()

	if s.step+1 >= s.Steps {
		term.Terminate = true
	}
	s.step++
}

// Finalize the system
func (s *FixedTermination) Finalize(w *ecs.World) {}
