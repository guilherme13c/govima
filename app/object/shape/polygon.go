package shape

import (
	"govima/app/misc"
	"govima/app/misc/color"
	styleconst "govima/app/misc/constants/style"
	"log"

	"math"

	"github.com/ungerik/go-cairo"
)

type Polygon_t struct {
	id     misc.Id_t
	height float64
	width  float64
	sides  uint64

	Vertices    [][2]float64
	Fill        bool
	StrokeWidth float64
	Color       color.Color_i
}

func NewPolygonObject(vertices [][2]float64, color color.Color_i) *Polygon_t {
	sides := len(vertices)
	if sides < 3 {
		log.Fatalf("Regular Polygon must have at least 3 sides")
	}

	id := misc.NextId()

	minX := math.Inf(1)
	minY := math.Inf(1)
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)

	for _, v := range vertices {
		minX = min(minX, v[0])
		minY = min(minX, v[1])
		maxX = max(maxX, v[0])
		maxY = max(maxX, v[1])
	}

	return &Polygon_t{
		id:          id,
		width:       maxX - minX,
		height:      maxY - minY,
		Vertices:    vertices,
		Fill:        false,
		StrokeWidth: styleconst.DefaultStrokeWidth,
		sides:       uint64(sides),
		Color:       color,
	}
}

func (o *Polygon_t) GetId() misc.Id_t {
	return o.id
}

func (o *Polygon_t) Render(surf *cairo.Surface, x float64, y float64) {
	color := o.Color.AsFloat64RGBA()

	surf.SetSourceRGBA(color.R, color.G, color.B, color.A)

	surf.MoveTo(x+o.Vertices[0][0], y+o.Vertices[0][1])
	for i := uint64(1); i < o.sides; i++ {
		surf.LineTo(x+o.Vertices[i][0], y+o.Vertices[i][1])
	}
	surf.ClosePath()

	if o.Fill {
		surf.Fill()
	} else {
		surf.SetLineWidth(o.StrokeWidth)
		surf.Stroke()
	}
}

func (o *Polygon_t) GetWidth() float64 {
	return o.width
}

func (o *Polygon_t) GetHeight() float64 {
	return o.height
}
