package systems

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche-model/model"
)

// PerfTimer system
type PerfTimer struct {
	UpdateInterval int
	Stats          bool
	start          time.Time
	step           int
}

// Initialize the system
func (s *PerfTimer) Initialize(m *model.Model) {}

// Update the system
func (s *PerfTimer) Update(m *model.Model) {
	t := time.Now()
	if m.Step == 0 {
		s.start = t
	}
	if m.Step%int64(s.UpdateInterval) == 0 {
		if m.Step > 0 {
			dur := t.Sub(s.start)
			usec := float64(dur.Microseconds()) / float64(s.step)
			fmt.Printf("%d updates, %0.2f us/update\n", s.UpdateInterval, usec)
		}
		if s.Stats {
			fmt.Println(m.World.Stats().String())
		}
		s.step = 0
		s.start = t
	}
	s.step++
}

// Finalize the system
func (s *PerfTimer) Finalize(m *model.Model) {}
