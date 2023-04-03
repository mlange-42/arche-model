package reporters

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlange-42/arche-models/model"
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
}

// Initialize the system
func (s *CSV) Initialize(m *model.Model) {
	s.Observer.Initialize(m)
	s.header = s.Observer.Header(m)

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
}

// Update the system
func (s *CSV) Update(m *model.Model) {
	s.Observer.Update(m)
	if s.UpdateInterval == 0 || m.Step%int64(s.UpdateInterval) == 0 {
		values := s.Observer.Values(m)
		s.builder.Reset()
		fmt.Fprintf(&s.builder, "%d%s", m.Step, s.Sep)
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
func (s *CSV) Finalize(m *model.Model) {
	if err := s.file.Close(); err != nil {
		panic(err)
	}
}
