package system

import (
	"github.com/mlange-42/arche-model/resource"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// CallbackTermination system.
//
// Terminates a model run according to the return value of a callback function.
//
// Expects a resource of type [model.Termination].
type CallbackTermination struct {
	Callback func(t int64) bool // The callback. ends the simulation when it returns true.
	termRes  generic.Resource[resource.Termination]
	step     int64
}

// Initialize the system
func (s *CallbackTermination) Initialize(w *ecs.World) {
	s.termRes = generic.NewResource[resource.Termination](w)
	s.step = 0
}

// Update the system
func (s *CallbackTermination) Update(w *ecs.World) {
	term := s.termRes.Get()

	if s.Callback(s.step) {
		term.Terminate = true
	}
	s.step++
}

// Finalize the system
func (s *CallbackTermination) Finalize(w *ecs.World) {}
