package reporters

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// CSV reporter.
//
// Writes one row to a CSV file per step.
type CSV struct {
	Observer       Observer
	File           string
	Sep            string
	UpdateInterval int
	file           *os.File
	header         []string
	builder        strings.Builder
	timeRes        generic.Resource[model.Time]
}

// Initialize the system
func (s *CSV) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	s.header = s.Observer.Header(w)
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

	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system
func (s *CSV) Update(w *ecs.World) {
	time := s.timeRes.Get()

	s.Observer.Update(w)
	if s.UpdateInterval == 0 || time.Tick%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(w)
		s.builder.Reset()
		fmt.Fprintf(&s.builder, "%d%s", time.Tick, s.Sep)
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
}

// Finalize the system
func (s *CSV) Finalize(w *ecs.World) {
	if err := s.file.Close(); err != nil {
		panic(err)
	}
}
