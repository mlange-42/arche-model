package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/generic"
	"golang.org/x/image/colornames"
)

// Bounds define a bounding box for a window.
type Bounds struct {
	X      int
	Y      int
	Width  int
	Height int
}

// B created a new Bounds object.
func B(x, y, w, h int) Bounds {
	return Bounds{x, y, w, h}
}

// GlDrawer interface
type GlDrawer interface {
	Initialize(m *model.Model, w *pixelgl.Window)
	Draw(m *model.Model, w *pixelgl.Window)
}

// GlFrame provides an OpenGL window for drawing.
type GlFrame struct {
	Bounds       Bounds
	DrawInterval int
	window       *pixelgl.Window
	drawers      []GlDrawer
	step         int
	timeRes      generic.Resource[model.Time]
}

// Window returns the window of this system.
func (s *GlFrame) Window() *pixelgl.Window {
	return s.window
}

// Add adds a drawer
func (s *GlFrame) Add(d GlDrawer) {
	s.drawers = append(s.drawers, d)
}

// InitializeUI the system
func (s *GlFrame) InitializeUI(m *model.Model) {
	if s.Bounds.Width <= 0 {
		s.Bounds.Width = 1024
	}
	if s.Bounds.Height <= 0 {
		s.Bounds.Height = 768
	}
	cfg := pixelgl.WindowConfig{
		Title:    "Arche",
		Bounds:   pixel.R(0, 0, float64(s.Bounds.Width), float64(s.Bounds.Height)),
		Position: pixel.V(float64(s.Bounds.X), float64(s.Bounds.Y)),
	}

	var err error
	s.window, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	for _, d := range s.drawers {
		d.Initialize(m, s.window)
	}
	s.timeRes = generic.NewResource[model.Time](&m.World)
}

// UpdateUI the system
func (s *GlFrame) UpdateUI(m *model.Model) {
	if s.window.Closed() {
		time := s.timeRes.Get()
		time.Finished = true
		return
	}
	if s.DrawInterval <= 1 || s.step%s.DrawInterval == 0 {
		s.window.Clear(colornames.Black)

		for _, d := range s.drawers {
			d.Draw(m, s.window)
		}
	}
	s.step++
}

// FinalizeUI the system
func (s *GlFrame) FinalizeUI(m *model.Model) {}
