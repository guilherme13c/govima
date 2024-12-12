package shape

import (
	"govima/app/misc"
	"govima/app/misc/color"
	styleconst "govima/app/misc/constants/style"

	"github.com/ungerik/go-cairo"
)

type Rectangle_t struct {
	id          misc.Id_t
	Fill        bool
	StrokeWidth float64
	Width       float64
	Height      float64
	Color       color.Color_i
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
	}
}

func (o *Rectangle_t) GetId() misc.Id_t {
	return o.id
}

func (o *Rectangle_t) Render(surf *cairo.Surface, x float64, y float64) {
	color := o.Color.AsFloat64RGBA()

	surf.SetSourceRGBA(color.R, color.G, color.B, color.A)

	surf.Rectangle(x-o.Width/2, y-o.Height/2, o.Width, o.Height)

	if o.Fill {
		surf.Fill()
	} else {
		surf.SetLineWidth(o.StrokeWidth)
		surf.Stroke()
	}
}

func (o *Rectangle_t) GetWidth() float64 {
	return o.Width
}

func (o *Rectangle_t) GetHeight() float64 {
	return o.Height
}
