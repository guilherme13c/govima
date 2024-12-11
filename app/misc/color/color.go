package color

import (
	colortypes "govima/app/misc/color/type"
)

type Color_i interface {
	AsUint8RGBA() *colortypes.ColorRGBAUint8_t
	AsFloat64RGBA() *colortypes.ColorRGBAFloat64_t
}
