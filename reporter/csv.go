package reporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlange-42/arche-model/observer"
	"github.com/mlange-42/arche/ecs"
)

// CSV reporter.
//
// Writes one row to a CSV file per step.
type CSV struct {
	Observer       observer.Row // Observer to get data from.
	File           string       // Path to the output file.
	Sep            string       // Column separator. Default ",".
	UpdateInterval int          // Update interval in model ticks.
	file           *os.File
	header         []string
	builder        strings.Builder
	step           int64
}

// Initialize the system
func (s *CSV) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	s.header = s.Observer.Header()
	if s.Sep == "" {
		s.Sep = ","
	}

	err := os.MkdirAll(filepath.Dir(s.File), os.ModePerm)
	if err != nil {
		panic(err)
	}

	s.file, err = os.Create(s.File)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(s.file, "t%s%s\n", s.Sep, strings.Join(s.header, s.Sep))
	if err != nil {
		panic(err)
	}

	s.step = 0
}

// Update the system
func (s *CSV) Update(w *ecs.World) {
	s.Observer.Update(w)
	if s.UpdateInterval == 0 || s.step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		s.builder.Reset()
		fmt.Fprintf(&s.builder, "%d%s", s.step, s.Sep)
		for i, v := range values {
			fmt.Fprintf(&s.builder, "%f", v)
			if i < len(values)-1 {
				fmt.Fprint(&s.builder, s.Sep)
			}
		}
		_, err := fmt.Fprintf(s.file, "%s\n", s.builder.String())
		if err != nil {
			panic(err)
		}
	}
	s.step++
}

// Finalize the system
func (s *CSV) Finalize(w *ecs.World) {
	if err := s.file.Close(); err != nil {
		panic(err)
	}
}
