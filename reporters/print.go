package reporters

import (
	"fmt"

	"github.com/mlange-42/arche-models/model"
)

// Print reporter
type Print struct {
	UpdateInterval int
	Observer       Observer
	header         []string
}

// Initialize the system
func (s *Print) Initialize(m *model.Model) {
	s.Observer.Initialize(m)
	s.header = s.Observer.Header(m)
}

// Update the system
func (s *Print) Update(m *model.Model) {
	s.Observer.Update(m)
	if m.Step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(m)
		fmt.Printf("%v\n%v\n", s.header, values)
	}
}

// Finalize the system
func (s *Print) Finalize(m *model.Model) {}
