package systems

import "github.com/mlange-42/arche-models/model"

// FixedTermination system
type FixedTermination struct {
	Steps int
}

// Initialize the system
func (s *FixedTermination) Initialize(m *model.Model) {}

// Update the system
func (s *FixedTermination) Update(m *model.Model) {
	if m.Step >= int64(s.Steps) {
		m.Finished = true
	}
}

// Finalize the system
func (s *FixedTermination) Finalize(m *model.Model) {}
