package shape

import (
	"govima/app/misc"
	"govima/app/misc/color"
	styleconst "govima/app/misc/constants/style"

	"github.com/ungerik/go-cairo"
)

type Rectangle_t struct {
	id misc.Id_t

	Fill        bool
	StrokeWidth float64
	Width       float64
	Height      float64
	Color       color.Color_i
	x           float64
	y           float64
}

func NewRectangleObject(width float64, height float64, color color.Color_i) *Rectangle_t {
	id := misc.NextId()

	return &Rectangle_t{
		id:          id,
		Fill:        false,
		StrokeWidth: styleconst.DefaultStrokeWidth,
		Width:       width,
		Height:      height,
		Color:       color,
		x:           0,
		y:           0,
	}
}

func (o *Rectangle_t) GetId() misc.Id_t {
	return o.id
}

func (o *Rectangle_t) Render(surf *cairo.Surface) {
	color := o.Color.AsFloat64RGBA()

	surf.SetSourceRGBA(color.R, color.G, color.B, color.A)

	surf.Rectangle(o.x-o.Width/2, o.y-o.Height/2, o.Width, o.Height)

	if o.Fill {
		surf.Fill()
	} else {
		surf.SetLineWidth(o.StrokeWidth)
		surf.Stroke()
	}
}

func (o *Rectangle_t) GetDim() (float64, float64) {
	return o.Width, o.Height
}

func (o *Rectangle_t) GetPos() (float64, float64) {
	return o.x, o.y
}

func (o *Rectangle_t) SetPos(x float64, y float64) {
	o.x = x
	o.y = y
}
