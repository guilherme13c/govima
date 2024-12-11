package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"govima/app/resource/config"
	imagescene "govima/app/scene/image_scene"
	videoscene "govima/app/scene/video_scene"

	"github.com/ungerik/go-cairo"
)

func main() {
	createBaseFolders()

	s1 := videoscene.NewVideoScene(60, 3, scene1Func, map[string]interface{}{
		"totalFrames": 3 * 60,
		"frameId":     0,
	})
	s2 := videoscene.NewVideoScene(60, 1, scene2Func, map[string]interface{}{})
	s3 := imagescene.NewImageScene(scene3Func, map[string]interface{}{})

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

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)
	surf.SetSourceRGB(0.8, 0.2, 0.2)
	rectWidth := 100.0
	rectHeight := 100.0
	x := float64(frameId) / float64(totalFrames-1) * float64(config.Config.Width-rectWidth)
	y := float64(config.Config.Height)/2 - rectHeight/2
	surf.Rectangle(x, y, rectWidth, rectHeight)
	surf.Fill()

	state["frameId"] = frameId + 1
}

func scene2Func(surf *cairo.Surface, state map[string]interface{}) {
	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.Rectangle(0, 0, config.Config.Width, config.Config.Height)
	surf.SetSourceRGB(0, 0, 0)
	surf.Fill()

	surf.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surf.SetFontSize(42)
	surf.SetSourceRGB(1, 1, 1)
	text := fmt.Sprintf("%s", "Govima")
	extents := surf.TextExtents(text)
	surf.MoveTo(config.Config.Width/2-extents.Width/2, config.Config.Height/2+extents.Height/2)
	surf.ShowText(text)
}

func scene3Func(surf *cairo.Surface, state map[string]interface{}) {
	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.Rectangle(0, 0, config.Config.Width, config.Config.Height)
	surf.SetSourceRGB(0, 0, 0)
	surf.Fill()

	surf.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surf.SetFontSize(42)
	surf.SetSourceRGB(1, 1, 1)
	text := fmt.Sprintf("%s", "Test Image Scene")
	extents := surf.TextExtents(text)
	surf.MoveTo(config.Config.Width/2-extents.Width/2, config.Config.Height/2+extents.Height/2)
	surf.ShowText(text)
}

func createBaseFolders() {
	if err := os.MkdirAll(config.Config.FrameDir, 0755); err != nil {
		log.Fatalf("Failed to create frame directory: %v", err)
	}

	if err := os.MkdirAll(config.Config.OutputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}
}
