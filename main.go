package main

import (
	"fmt"
	"govima/app/resource/config"
	"govima/app/scene"
	"log"
	"os"
	"sync"

	"github.com/ungerik/go-cairo"
)

func main() {
	createBaseFolders()

	scene.NewScene(60, 3, scene1)
	scene.NewScene(60, 1, scene2)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for frameId := uint32(0); frameId < scene.SceneList.Scenes[0].NumberOfFrames; frameId++ {
			scene.SceneList.Scenes[0].RenderFrame(frameId, []interface{}{frameId, scene.SceneList.Scenes[0].NumberOfFrames})
		}
		scene.SceneList.Scenes[0].GenerateVideo()
		scene.SceneList.Scenes[0].Clean()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for frameId := uint32(0); frameId < scene.SceneList.Scenes[1].NumberOfFrames; frameId++ {
			scene.SceneList.Scenes[1].RenderFrame(frameId, []interface{}{frameId, "Govima"})
		}
		scene.SceneList.Scenes[1].GenerateVideo()
		scene.SceneList.Scenes[1].Clean()
	}()

	wg.Wait()
}

func scene1(surf *cairo.Surface, args []interface{}) {
	frame := args[0].(uint32)
	totalFrames := args[1].(uint32)

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.SetSourceRGB(0.8, 0.2, 0.2)
	rectWidth := 100.0
	rectHeight := 100.0
	x := float64(frame) / float64(totalFrames-1) * float64(config.Config.Width-rectWidth)
	y := float64(config.Config.Height)/2 - rectHeight/2
	surf.Rectangle(x, y, rectWidth, rectHeight)
	surf.Fill()
}

func scene2(surf *cairo.Surface, args []interface{}) {
	frame := args[0].(uint32)
	videoTitle := args[1].(string)

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.Rectangle(0, 0, config.Config.Width, config.Config.Height)
	surf.SetSourceRGB(0, 0, 0)
	surf.Fill()

	surf.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surf.SetFontSize(42)
	surf.SetSourceRGB(1, 1, 1)
	text := fmt.Sprintf("%s - Frame %d", videoTitle, frame)
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
