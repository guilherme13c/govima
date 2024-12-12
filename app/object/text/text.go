package text

import (
	"govima/app/misc"
	"govima/app/misc/color"

	"github.com/ungerik/go-cairo"
)

type Text_t struct {
	id     misc.Id_t
	width  float64
	height float64

	Text        string
	FontSize    float64
	Color       color.Color_i
	FontFace    string
	StrokeWidth float64
}

func NewTextObject(text string, fontSize float64, color color.Color_i, fontFace string) *Text_t {
	id := misc.NextId()

	surf := cairo.NewSurface(cairo.FORMAT_ARGB32, 1, 1) // Temporary surface
	defer surf.Finish()

	surf.SelectFontFace(fontFace, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	surf.SetFontSize(fontSize)

	extents := surf.TextExtents(text)

	return &Text_t{
		id:          id,
		width:       extents.Width,
		height:      extents.Height,
		Text:        text,
		FontSize:    fontSize,
		Color:       color,
		FontFace:    fontFace,
		StrokeWidth: 1,
	}
}

func (o *Text_t) GetId() misc.Id_t {
	return o.id
}

func (o *Text_t) Render(surf *cairo.Surface, x float64, y float64) {
	color := o.Color.AsFloat64RGBA()

	surf.SelectFontFace(o.FontFace, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	surf.SetFontSize(o.FontSize)

	surf.SetSourceRGBA(color.R, color.G, color.B, color.A)

	extents := surf.TextExtents(o.Text)
	textX := x + extents.Width/2.0
	textY := y + extents.Height/2.0

	surf.MoveTo(textX, textY)
	surf.ShowText(o.Text)
	surf.Stroke()
}

func (o *Text_t) GetWidth() float64 {
	return o.width
}

func (o *Text_t) GetHeight() float64 {
	return o.height
}
