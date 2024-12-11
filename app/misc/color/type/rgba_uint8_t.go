package types

type ColorRGBAUint8_t struct {
	R, G, B, A uint8
}

func NewColorRGBAUint8(R uint8, G uint8, B uint8, A uint8) *ColorRGBAUint8_t {
	return &ColorRGBAUint8_t{
		R: R,
		G: G,
		B: B,
		A: A,
	}
}

func (s *ColorRGBAUint8_t) AsUint8RGBA() *ColorRGBAUint8_t {
	panic("unimplemented")
}

func (c *ColorRGBAUint8_t) AsFloat64RGBA() *ColorRGBAFloat64_t {
	return &ColorRGBAFloat64_t{
		R: float64(c.R) / 255.0,
		G: float64(c.G) / 255.0,
		B: float64(c.B) / 255.0,
		A: float64(c.A) / 255.0,
	}
}
