package systems

import (
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// FixedTermination system
type FixedTermination struct {
	Steps   int64
	timeRes generic.Resource[model.Time]
}

// Initialize the system
func (s *FixedTermination) Initialize(w *ecs.World) {
	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system
func (s *FixedTermination) Update(w *ecs.World) {
	time := s.timeRes.Get()

	if time.Tick >= s.Steps {
		time.Finished = true
	}
}

// Finalize the system
func (s *FixedTermination) Finalize(w *ecs.World) {}
