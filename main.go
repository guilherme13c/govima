package main

import (
	"fmt"
	"sync"

	"govima/app/object/latex"
	"govima/app/resource/config"
	imagescene "govima/app/scene/image_scene"
	videoscene "govima/app/scene/video_scene"

	"github.com/ungerik/go-cairo"
)

func main() {
	config.Init()

	s1 := videoscene.NewVideoScene(800, 600, 60, 3, scene1Func, map[string]interface{}{
		"totalFrames": 3 * 60,
		"frameId":     0,
	})
	s2 := videoscene.NewVideoScene(800, 600, 60, 1, scene2Func, map[string]interface{}{})
	s3 := imagescene.NewImageScene(800, 600, scene3Func, map[string]interface{}{})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		s1.Save()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s2.Save()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s3.Save()
	}()

	wg.Wait()
}

func scene1Func(surf *cairo.Surface, state map[string]interface{}) {
	frameId := state["frameId"].(int)
	totalFrames := state["totalFrames"].(int)
	width := float64(state["width"].(uint32))
	height := float64(state["height"].(uint32))

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)
	surf.SetSourceRGB(0.8, 0.2, 0.2)
	rectWidth := 100.0
	rectHeight := 100.0
	x := float64(frameId) / float64(totalFrames-1) * (width - rectWidth)
	y := height/2 - rectHeight/2
	surf.Rectangle(x, y, rectWidth, rectHeight)
	surf.Fill()

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

	surf.Rectangle(0, 0, width, height)
	surf.SetSourceRGB(1, 1, 1)
	surf.Fill()
	latexObj := latex.NewLatexObject(`$f(x) = \frac{\sqrt{x}}{2\pi}$`, 12, 300, nil)
	latexObj.Compile()
	latexObj.Render(surf, width/2-latexObj.GetWidth(), height/2-latexObj.GetHeight())
	latexObj.Clean()
}
