package systems

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SnapshotCSV reporter.
//
// Writes a CSV file per step.
type SnapshotCSV struct {
	Observer       model.MatrixObserver
	FilePattern    string
	Sep            string
	UpdateInterval int
	header         []string
	builder        strings.Builder
	timeRes        generic.Resource[model.Time]
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

	s.timeRes = generic.NewResource[model.Time](w)
}

// Update the system
func (s *SnapshotCSV) Update(w *ecs.World) {
	time := s.timeRes.Get()

	s.Observer.Update(w)
	if s.UpdateInterval == 0 || time.Tick%int64(s.UpdateInterval) == 0 {
		file, err := os.Create(fmt.Sprintf(s.FilePattern, time.Tick))
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
}

// Finalize the system
func (s *SnapshotCSV) Finalize(w *ecs.World) {}
