package object

import (
	"govima/app/misc"

	"github.com/ungerik/go-cairo"
)

type Object_i interface {
	GetId() misc.Id_t
	Render(surf *cairo.Surface, x float64, y float64)
	Clean()
	GetWidth() float64
	GetHeight() float64
}
