package system

import (
	"fmt"
	"time"

	"github.com/mlange-42/arche/ecs"
)

// PerfTimer system for printing elapsed time per model step, and optional world statistics.
type PerfTimer struct {
	UpdateInterval int  // Update/print interval in model ticks.
	Stats          bool // Whether to print world stats.
	start          time.Time
	step           int64
}

// Initialize the system
func (s *PerfTimer) Initialize(w *ecs.World) {
	s.step = 0
}

// Update the system
func (s *PerfTimer) Update(w *ecs.World) {
	t := time.Now()
	if s.step == 0 {
		s.start = t
	}
	if s.step%int64(s.UpdateInterval) == 0 {
		if s.step > 0 {
			dur := t.Sub(s.start)
			usec := float64(dur.Microseconds()) / float64(s.UpdateInterval)
			fmt.Printf("%d updates, %0.2f us/update\n", s.UpdateInterval, usec)
		}
		if s.Stats {
			fmt.Println(w.Stats().String())
		}
		s.start = t
	}
	s.step++
}

// Finalize the system
func (s *PerfTimer) Finalize(w *ecs.World) {}
