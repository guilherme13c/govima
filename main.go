package main

import (
	"fmt"
	"sync"

	colorconst "govima/app/misc/constants/color"
	"govima/app/object/latex"
	"govima/app/object/shape"
	"govima/app/resource/config"
	"govima/app/scene"
	imagescene "govima/app/scene/image_scene"
	videoscene "govima/app/scene/video_scene"

	"github.com/ungerik/go-cairo"
)

func main() {
	config.Init()

	videoscene.NewVideoScene(1920, 1080, 60, 3, scene1Func, map[string]interface{}{
		"totalFrames": 3 * 60,
		"frameId":     0,
	})
	videoscene.NewVideoScene(1080, 1080, 60, 1, scene2Func, map[string]interface{}{})
	imagescene.NewImageScene(1080, 1920, scene3Func, map[string]interface{}{})

	wg := sync.WaitGroup{}

	for _, s := range scene.SceneList.Scenes {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Save()
		}()
	}

	wg.Wait()
}

func scene1Func(surf *cairo.Surface, state map[string]interface{}) {
	frameId := state["frameId"].(int)
	totalFrames := state["totalFrames"].(int)
	width := float64(state["width"].(uint32))
	height := float64(state["height"].(uint32))

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	square := shape.NewRectangeObject(100, 100, colorconst.Blue)
	xSquare := float64(frameId)/float64(totalFrames-1)*(width-square.GetWidth()) + square.GetWidth()/2
	ySquare := height/2 - square.GetHeight()/2 - 150
	square.Render(surf, xSquare, ySquare)

	circle := shape.NewCircleObject(50, colorconst.Red)
	xCircle := float64(frameId)/float64(totalFrames-1)*(width-circle.GetWidth()) + circle.GetWidth()/2
	yCircle := height/2 - circle.GetHeight()/2 - 50
	circle.Render(surf, xCircle, yCircle)

	regPoly := shape.NewRegularPolygonObject(3, 50, colorconst.Green)
	xRegPoly := float64(frameId)/float64(totalFrames-1)*(width-regPoly.GetWidth()) + regPoly.GetWidth()/2
	yRegPoly := height/2 - regPoly.GetHeight()/2 + 50
	regPoly.Render(surf, xRegPoly, yRegPoly)

	poly := shape.NewPolygonObject([][2]float64{{50, 50}, {50, -50}, {-50, -50}, {-50, 50}}, colorconst.White)
	xPoly := float64(frameId)/float64(totalFrames-1)*(width-poly.GetWidth()) + poly.GetWidth()/2
	yPoly := height/2 - regPoly.GetHeight()/2 + 150
	poly.Render(surf, xPoly, yPoly)

	state["frameId"] = frameId + 1
}

func scene2Func(surf *cairo.Surface, state map[string]interface{}) {
	width := float64(state["width"].(uint32))
	height := float64(state["height"].(uint32))

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.Rectangle(0, 0, width, height)
	surf.SetSourceRGB(0, 0, 0)
	surf.Fill()

	surf.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surf.SetFontSize(42)
	surf.SetSourceRGB(1, 1, 1)
	text := fmt.Sprintf("%s", "Govima")
	extents := surf.TextExtents(text)
	surf.MoveTo(width/2-extents.Width/2, height/2+extents.Height/2)
	surf.ShowText(text)
}

func scene3Func(surf *cairo.Surface, state map[string]interface{}) {
	width := float64(state["width"].(uint32))
	height := float64(state["height"].(uint32))

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.Rectangle(0, 0, width, height)
	surf.SetSourceRGB(1, 1, 1)
	surf.Fill()
	latexObj := latex.NewLatexObject(`$f(x) = \frac{\sqrt{x}}{2\pi}$`, 12, 300, nil)
	latexObj.Compile()
	latexObj.Render(surf, width/2-latexObj.GetWidth()/2, height/2-latexObj.GetHeight()/2)
	latexObj.Clean()
}
