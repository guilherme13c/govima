package colorconst

import (
	colortypes "govima/app/misc/color/type"
)

var (
	White     = &colortypes.ColorRGBAUint8_t{R: 255, G: 255, B: 255, A: 255}
	Black     = &colortypes.ColorRGBAUint8_t{R: 0, G: 0, B: 0, A: 255}
	Red       = &colortypes.ColorRGBAUint8_t{R: 255, G: 0, B: 0, A: 255}
	Green     = &colortypes.ColorRGBAUint8_t{R: 0, G: 255, B: 0, A: 255}
	Blue      = &colortypes.ColorRGBAUint8_t{R: 0, G: 0, B: 255, A: 255}
	Yellow    = &colortypes.ColorRGBAUint8_t{R: 255, G: 255, B: 0, A: 255}
	Cyan      = &colortypes.ColorRGBAUint8_t{R: 0, G: 255, B: 255, A: 255}
	Magenta   = &colortypes.ColorRGBAUint8_t{R: 255, G: 0, B: 255, A: 255}
	Gray      = &colortypes.ColorRGBAUint8_t{R: 128, G: 128, B: 128, A: 255}
	LightGray = &colortypes.ColorRGBAUint8_t{R: 192, G: 192, B: 192, A: 255}
	DarkGray  = &colortypes.ColorRGBAUint8_t{R: 64, G: 64, B: 64, A: 255}
	Orange    = &colortypes.ColorRGBAUint8_t{R: 255, G: 165, B: 0, A: 255}
	Purple    = &colortypes.ColorRGBAUint8_t{R: 128, G: 0, B: 128, A: 255}
	Brown     = &colortypes.ColorRGBAUint8_t{R: 165, G: 42, B: 42, A: 255}
	Pink      = &colortypes.ColorRGBAUint8_t{R: 255, G: 192, B: 203, A: 255}
)
