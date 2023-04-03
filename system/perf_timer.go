package system

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// PerfTimer system for printing elapsed time per model step, and optional world statistics.
type PerfTimer struct {
	UpdateInterval int  // Update/print interval in model ticks.
	Stats          bool // Whether to print world stats.
	start          time.Time
	step           int
	timeRes        generic.Resource[model.Time]
}

// Initialize the system
func (s *PerfTimer) Initialize(w *ecs.World) {
	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system
func (s *PerfTimer) Update(w *ecs.World) {
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
			fmt.Println(w.Stats().String())
		}
		s.step = 0
		s.start = t
	}
	s.step++
}

// Finalize the system
func (s *PerfTimer) Finalize(w *ecs.World) {}
