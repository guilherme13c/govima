package shape

import (
	"govima/app/misc"
	"govima/app/misc/color"
	styleconst "govima/app/misc/constants/style"
	"log"

	"math"

	"github.com/ungerik/go-cairo"
)

type RegularPolygon_t struct {
	id       misc.Id_t
	height   float64
	width    float64
	vertices [][2]float64

	Fill        bool
	StrokeWidth float64
	Sides       uint64
	Radius      float64
	Color       color.Color_i
}

func NewRegularPolygonObject(sides uint64, radius float64, color color.Color_i) *RegularPolygon_t {
	if sides < 3 {
		log.Fatalf("Regular Polygon must have at least 3 sides")
	}

	id := misc.NextId()

	minX := math.Inf(1)
	minY := math.Inf(1)
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)

	angleStep := 2 * math.Pi / float64(sides)
	points := make([][2]float64, sides)
	for i := uint64(0); i < sides; i++ {
		angle := angleStep*float64(i) - math.Pi/2
		x := radius * math.Cos(angle)
		y := radius * math.Sin(angle)

		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, x)
		maxY = max(maxY, y)

		points[i] = [2]float64{x, y}
	}

	return &RegularPolygon_t{
		id:          id,
		width:       maxX - minX,
		height:      maxY - minY,
		vertices:    points,
		Fill:        false,
		StrokeWidth: styleconst.DefaultStrokeWidth,
		Sides:       sides,
		Radius:      radius,
		Color:       color,
	}
}

func (o *RegularPolygon_t) GetId() misc.Id_t {
	return o.id
}

func (o *RegularPolygon_t) Render(surf *cairo.Surface, x float64, y float64) {
	color := o.Color.AsFloat64RGBA()

	surf.SetSourceRGBA(color.R, color.G, color.B, color.A)

	surf.MoveTo(x+o.vertices[0][0], y+o.vertices[0][1])
	for i := uint64(1); i < o.Sides; i++ {
		surf.LineTo(x+o.vertices[i][0], y+o.vertices[i][1])
	}
	surf.ClosePath()

	if o.Fill {
		surf.Fill()
	} else {
		surf.SetLineWidth(o.StrokeWidth)
		surf.Stroke()
	}
}

func (o *RegularPolygon_t) GetWidth() float64 {
	return o.width
}

func (o *RegularPolygon_t) GetHeight() float64 {
	return o.height
}
