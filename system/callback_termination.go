package system

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// CallbackTermination system.
//
// Terminates a model run according to the return value of a callback function.
type CallbackTermination struct {
	Callback func(t int64) bool // The callback. ends the simulation when it returns true.
	timeRes  generic.Resource[model.Time]
}

// Initialize the system
func (s *CallbackTermination) Initialize(w *ecs.World) {
	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system
func (s *CallbackTermination) Update(w *ecs.World) {
	time := s.timeRes.Get()

	if s.Callback(time.Tick) {
		time.Finished = true
	}
}

// Finalize the system
func (s *CallbackTermination) Finalize(w *ecs.World) {}
