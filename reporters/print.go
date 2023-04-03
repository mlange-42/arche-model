package reporters

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// Print reporter
type Print struct {
	UpdateInterval int
	Observer       Observer
	header         []string
	timeRes        generic.Resource[model.Time]
}

// Initialize the system
func (s *Print) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	s.header = s.Observer.Header(w)

	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system
func (s *Print) Update(w *ecs.World) {
	time := s.timeRes.Get()

	s.Observer.Update(w)
	if time.Tick%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		fmt.Printf("%v\n%v\n", s.header, values)
	}
}

// Finalize the system
func (s *Print) Finalize(w *ecs.World) {}
