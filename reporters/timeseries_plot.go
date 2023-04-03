package reporters

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche-model/systems"
	"github.com/mlange-42/arche/generic"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

// TimeSeriesPlot reporter
type TimeSeriesPlot struct {
	Bounds         systems.Bounds
	Observer       Observer
	UpdateInterval int
	DrawInterval   int
	systems.GlFrame
	drawer  timeSeriesDrawer
	timeRes generic.Resource[model.Time]
}

// Initialize the system
func (s *TimeSeriesPlot) Initialize(m *model.Model) {
	s.timeRes = generic.NewResource[model.Time](&m.World)
}

// InitializeUI the system
func (s *TimeSeriesPlot) InitializeUI(m *model.Model) {
	s.Observer.Initialize(m)

	s.drawer = timeSeriesDrawer{}
	s.drawer.addSeries(s.Observer.Header(m))

	s.GlFrame.DrawInterval = s.DrawInterval
	s.GlFrame.Bounds = s.Bounds
	s.GlFrame.Add(&s.drawer)
	s.GlFrame.InitializeUI(m)
}

// Update the system
func (s *TimeSeriesPlot) Update(m *model.Model) {
	time := s.timeRes.Get()

	if s.UpdateInterval > 1 && time.Tick%int64(s.UpdateInterval) != 0 {
		return
	}
	s.Observer.Update(m)
	s.drawer.append(float64(time.Tick), s.Observer.Values(m))
}

// Finalize the system
func (s *TimeSeriesPlot) Finalize(m *model.Model) {}

type timeSeriesDrawer struct {
	headers []string
	series  []plotter.XYs
	scale   float64
}

func (s *timeSeriesDrawer) addSeries(names []string) {
	s.headers = names
	s.series = make([]plotter.XYs, len(s.headers))
}

func (s *timeSeriesDrawer) append(x float64, values []float64) {
	for i := 0; i < len(s.series); i++ {
		s.series[i] = append(s.series[i], plotter.XY{X: x, Y: values[i]})
	}
}

// Initialize the system
func (s *timeSeriesDrawer) Initialize(m *model.Model, w *pixelgl.Window) {
	width := 100.0
	c := vgimg.New(vg.Points(width), vg.Points(width))
	img := c.Image()
	s.scale = width / float64(img.Bounds().Dx())
}

// Draw the system
func (s *timeSeriesDrawer) Draw(m *model.Model, w *pixelgl.Window) {
	width := w.Canvas().Bounds().W()
	height := w.Canvas().Bounds().H()

	c := vgimg.New(vg.Points(width*s.scale), vg.Points(height*s.scale))

	p := plot.New()
	p.X.Tick.Label.Font.Size = 12
	p.Y.Tick.Label.Font.Size = 12

	p.Legend = plot.NewLegend()

	for i := 0; i < len(s.series); i++ {
		lines, err := plotter.NewLine(s.series[i])
		if err != nil {
			panic(err)
		}
		lines.Color = defaultColors[i%len(defaultColors)]
		p.Add(lines)
		p.Legend.Add(s.headers[i], lines)
	}

	// Workaround for rendering bug: https://github.com/gonum/plot/issues/761
	p.Legend.XOffs = -10
	p.X.Max *= 1.03
	p.Y.Max *= 1.03

	p.Draw(draw.New(c))

	img := c.Image()
	picture := pixel.PictureDataFromImage(img)

	sprite := pixel.NewSprite(picture, picture.Bounds())
	sprite.Draw(w, pixel.IM.Moved(pixel.V(picture.Rect.W()/2.0, picture.Rect.H()/2.0)))
}
