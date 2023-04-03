package systems

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/generic"
)

// PerfTimer system
type PerfTimer struct {
	UpdateInterval int
	Stats          bool
	start          time.Time
	step           int
	timeRes        generic.Resource[model.Time]
}

// Initialize the system
func (s *PerfTimer) Initialize(m *model.Model) {
	s.timeRes = generic.NewResource[model.Time](&m.World)
}

// Update the system
func (s *PerfTimer) Update(m *model.Model) {
	tm := s.timeRes.Get()

	t := time.Now()
	if tm.Tick == 0 {
		s.start = t
	}
	if tm.Tick%int64(s.UpdateInterval) == 0 {
		if tm.Tick > 0 {
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
