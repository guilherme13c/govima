package object

import (
	"govima/app/misc"

	"github.com/ungerik/go-cairo"
)

type Object_i interface {
	GetId() misc.Id_t
	Render(surf *cairo.Surface)
	GetDim() (float64, float64)
	GetPos() (float64, float64)
	SetPos(float64, float64)
}
