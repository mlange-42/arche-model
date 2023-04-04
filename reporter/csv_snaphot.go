package reporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
)

// SnapshotCSV reporter.
//
// Writes a CSV file per step.
type SnapshotCSV struct {
	Observer       model.MatrixObserver // Observer to get data from.
	FilePattern    string               // File path and pattern for output files, like out/foo-%06d.csv
	Sep            string               // Column separator. Default ",".
	UpdateInterval int                  // Update interval in model ticks.
	header         []string
	builder        strings.Builder
	step           int64
}

// Initialize the system
func (s *SnapshotCSV) Initialize(w *ecs.World) {
	s.Observer.Initialize(w)
	s.header = s.Observer.Header(w)

	if s.Sep == "" {
		s.Sep = ","
	}

	err := os.MkdirAll(filepath.Dir(fmt.Sprintf(s.FilePattern, 1)), os.ModePerm)
	if err != nil {
		panic(err)
	}

	s.step = 0
}

// Update the system
func (s *SnapshotCSV) Update(w *ecs.World) {
	s.Observer.Update(w)
	if s.UpdateInterval == 0 || s.step%int64(s.UpdateInterval) == 0 {
		file, err := os.Create(fmt.Sprintf(s.FilePattern, s.step))
		if err != nil {
			panic(err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}()

		_, err = fmt.Fprintf(file, "%s\n", strings.Join(s.header, s.Sep))
		if err != nil {
			panic(err)
		}

		values := s.Observer.Values(w)
		s.builder.Reset()
		for _, row := range values {
			for i, v := range row {
				fmt.Fprintf(&s.builder, "%f", v)
				if i < len(row)-1 {
					fmt.Fprint(&s.builder, s.Sep)
				}
			}
			fmt.Fprint(&s.builder, "\n")
		}
		_, err = fmt.Fprint(file, s.builder.String())
		if err != nil {
			panic(err)
		}
	}
	s.step++
}

// Finalize the system
func (s *SnapshotCSV) Finalize(w *ecs.World) {}
