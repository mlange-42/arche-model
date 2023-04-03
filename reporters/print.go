package reporters

import (
	"fmt"

	"github.com/mlange-42/arche-model/model"
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
func (s *Print) Initialize(m *model.Model) {
	s.Observer.Initialize(m)
	s.header = s.Observer.Header(m)

	s.timeRes = generic.NewResource[model.Time](&m.World)
}

// Update the system
func (s *Print) Update(m *model.Model) {
	time := s.timeRes.Get()

	s.Observer.Update(m)
	if time.Tick%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(m)
		fmt.Printf("%v\n%v\n", s.header, values)
	}
}

// Finalize the system
func (s *Print) Finalize(m *model.Model) {}
