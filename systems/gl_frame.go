package systems

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
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
	Initialize(w *ecs.World, win *pixelgl.Window)
	Draw(w *ecs.World, win *pixelgl.Window)
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

// Add adds a drawer
func (s *GlFrame) Add(d GlDrawer) {
	s.drawers = append(s.drawers, d)
}

// InitializeUI the system
func (s *GlFrame) InitializeUI(w *ecs.World) {
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

	defer func() {
		if err := recover(); err != nil {
			txt := fmt.Sprint(err)
			if txt == "mainthread: did not call Run" {
				log.Fatal("ERROR: when using graphics via the pixel engine, run the model like this:\n    pixelgl.Run(model.Run)")
			}
			panic(err)
		}
	}()

	var err error
	s.window, err = pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	for _, d := range s.drawers {
		d.Initialize(w, s.window)
	}
	s.timeRes = generic.NewResource[model.Time](w)
}

// UpdateUI the system
func (s *GlFrame) UpdateUI(w *ecs.World) {
	if s.window.Closed() {
		time := s.timeRes.Get()
		time.Finished = true
		return
	}
	if s.DrawInterval <= 1 || s.step%s.DrawInterval == 0 {
		s.window.Clear(colornames.Black)

		for _, d := range s.drawers {
			d.Draw(w, s.window)
		}
	}
	s.step++
}

// PostUpdateUI updates the GL window.
func (s *GlFrame) PostUpdateUI(w *ecs.World) {
	s.window.Update()
}

// FinalizeUI the system
func (s *GlFrame) FinalizeUI(w *ecs.World) {}
