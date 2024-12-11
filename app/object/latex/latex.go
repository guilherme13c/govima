package latex

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/font/ttf"
	"github.com/go-latex/latex/mtex"
	"github.com/ungerik/go-cairo"

	"govima/app/misc"
	"govima/app/resource/config"
)

type Latex_t struct {
	id          misc.Id_t
	tmpFilePath string
	width       float64
	height      float64

	Expr string
	Size float64
	Dpi  float64
	Font *ttf.Fonts
}

func NewLatexObject(expr string, size float64, dpi float64, font *ttf.Fonts) *Latex_t {
	id := misc.NextId()
	path := fmt.Sprintf("%s/latex_%08d.png", config.Config.LatexDir, id)

	return &Latex_t{
		id:          id,
		tmpFilePath: path,
		Expr:        expr,
		Size:        size,
		Dpi:         dpi,
		Font:        font,
		width:       0,
		height:      0,
	}
}

func (o *Latex_t) Compile() {
	f, errOpen := os.Create(o.tmpFilePath)
	if errOpen != nil {
		log.Fatalf("Failed to create latex temporary file: %s", o.tmpFilePath)
	}
	defer f.Close()

	dst := drawimg.NewRenderer(f)
	errRender := mtex.Render(dst, o.Expr, o.Size, o.Dpi, o.Font)
	if errRender != nil {
		log.Fatalf("Failed to render latex expression: %s", errRender)
	}

	// send cursor to the begining of the file
	if _, err := f.Seek(0, 0); err != nil {
		log.Fatalf("Failed to reset file pointer: %v", err)
	}

	img, _, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatalf("Failed to decode PNG file: %v", err)
	}

	o.width = float64(img.Width)
	o.height = float64(img.Height)
}

func (o *Latex_t) GetId() misc.Id_t {
	return o.id
}

func (o *Latex_t) Render(surf *cairo.Surface, x float64, y float64) {
	img, status := cairo.NewSurfaceFromPNG(o.tmpFilePath)
	if status != cairo.STATUS_SUCCESS {
		log.Fatalf("Failed to load PNG file: %s", o.tmpFilePath)
	}
	defer img.Finish()

	surf.SetSourceSurface(img, x, y)
	surf.Paint()
}

func (o *Latex_t) Clean() {
	os.RemoveAll(o.tmpFilePath)
}

func (o *Latex_t) GetWidth() float64 {
	return o.width
}

func (o *Latex_t) GetHeight() float64 {
	return o.height
}
