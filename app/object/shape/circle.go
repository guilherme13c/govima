package shape

import (
	"govima/app/misc"
	"govima/app/misc/color"
	styleconst "govima/app/misc/constants/style"
	"math"

	"github.com/ungerik/go-cairo"
)

type Circle_t struct {
	id misc.Id_t

	Fill        bool
	StrokeWidth float64
	Radius      float64
	Color       color.Color_i
	x           float64
	y           float64
}

func NewCircleObject(radius float64, color color.Color_i) *Circle_t {
	id := misc.NextId()

	return &Circle_t{
		id:          id,
		Fill:        false,
		StrokeWidth: styleconst.DefaultStrokeWidth,
		Radius:      radius,
		Color:       color,
		x:           0,
		y:           0,
	}
}

func (o *Circle_t) GetId() misc.Id_t {
	return o.id
}

func (o *Circle_t) Render(surf *cairo.Surface) {
	color := o.Color.AsFloat64RGBA()

	surf.SetSourceRGBA(color.R, color.G, color.B, color.A)
	surf.Arc(o.x, o.y, o.Radius, 0, 2*math.Pi)

	if o.Fill {
		surf.Fill()
	} else {
		surf.SetLineWidth(o.StrokeWidth)
		surf.Stroke()
	}
}

func (o *Circle_t) GetDim() (float64, float64) {
	return 2 * o.Radius, 2 * o.Radius
}

func (o *Circle_t) GetPos() (float64, float64) {
	return o.x, o.y
}

func (o *Circle_t) SetPos(x float64, y float64) {
	o.x = x
	o.y = y
}
