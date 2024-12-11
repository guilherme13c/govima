package types

type ColorRGBAFloat64_t struct {
	R, G, B, A float64
}

func NewColorRGBAFloat64(R float64, G float64, B float64, A float64) *ColorRGBAFloat64_t {
	return &ColorRGBAFloat64_t{
		R: R,
		G: G,
		B: B,
		A: A,
	}
}

func (s *ColorRGBAFloat64_t) AsFloat64RGBA() *ColorRGBAFloat64_t {
	return s
}

func (s *ColorRGBAFloat64_t) AsUint8RGBA() *ColorRGBAUint8_t {
	return &ColorRGBAUint8_t{
		R: uint8(s.R * 255.0),
		G: uint8(s.G * 255.0),
		B: uint8(s.B * 255.0),
		A: uint8(s.A * 255.0),
	}
}
