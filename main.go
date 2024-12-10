package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/ungerik/go-cairo"
)

const (
	Width      = 1920
	Height     = 1080
	FrameRate  = 60
	Duration   = 2
	FrameDir   = "frames"
	OutputDir  = "output"
	VideoTitle = "Govima"
)

var ffmpegArgs = ffmpeg.KwArgs{
	"c:v":     "libx264",
	"pix_fmt": "yuv420p",
	"preset":  "fast",
	"crf":     6,
}

type sceneId_t uint32

type SceneList_t struct {
	nextSceneId sceneId_t
	Scenes      map[sceneId_t]*Scene_t
}

func (sl *SceneList_t) Add(s *Scene_t) {
	if sl.Scenes == nil {
		sl.Scenes = make(map[sceneId_t]*Scene_t)
	}

	sl.Scenes[s.Id] = s
	sl.nextSceneId++
}

var sceneList SceneList_t

type Scene_t struct {
	Id              sceneId_t
	FrameRate       uint32 // Scene frame rate
	Duration        uint32 // Duration of the scene in seconds
	NumberOfFrames  uint32 // Total number of frames
	FrameDir        string
	OutputVideoPath string
	renderFunc      func(surf *cairo.Surface, args []interface{})
}

func NewScene(frameRate uint32, duration uint32, renderFunc func(surf *cairo.Surface, args []any)) {
	frameDir := createSceneDir(sceneList.nextSceneId)
	outputVideoPath := fmt.Sprintf("%s/scene_%04d.mp4", OutputDir, sceneList.nextSceneId)

	s := &Scene_t{
		Id:              sceneList.nextSceneId,
		FrameRate:       frameRate,
		Duration:        duration,
		NumberOfFrames:  frameRate * duration,
		FrameDir:        frameDir,
		OutputVideoPath: outputVideoPath,
		renderFunc:      renderFunc,
	}
	sceneList.Add(s)
}

// Render a single frame using Cairo
func (s *Scene_t) RenderFrame(frameId uint32, args []interface{}) {
	filename := fmt.Sprintf("%s/frame_%08d.png", s.FrameDir, frameId)
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, Width, Height)

	s.renderFunc(surface, args)

	// Save the frame as PNG
	status := surface.WriteToPNG(filename)
	if status != cairo.STATUS_SUCCESS {
		log.Fatalf("Failed to save frame to PNG: %v", status)
	}

	// Clean up
	surface.Finish()
}

// Generate a video from the scene frames using FFmpeg
func (s *Scene_t) generateVideo() error {
	framePattern := fmt.Sprintf("%s/frame_%%08d.png", s.FrameDir)
	return ffmpeg.Input(framePattern, ffmpeg.KwArgs{"framerate": s.FrameRate}).
		Output(s.OutputVideoPath, ffmpegArgs).
		OverWriteOutput().
		Run()
}

func removeDir(sceneFrameDir string) {
	if err := os.RemoveAll(sceneFrameDir); err != nil {
		log.Fatalf("Failed to remove frame directory: %v", err)
	}
}

func createSceneDir(sceneId sceneId_t) string {
	path := fmt.Sprintf("%s/scene_%04d", FrameDir, sceneId)
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("Failed to create frame directory: %v", err)
	}
	return path
}

func (s *Scene_t) clean() {
	removeDir(s.FrameDir)
}

func createBaseFolders() {
	if err := os.MkdirAll(FrameDir, 0755); err != nil {
		log.Fatalf("Failed to create frame directory: %v", err)
	}

	if err := os.MkdirAll(OutputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}
}

func main() {
	createBaseFolders()

	NewScene(FrameRate, 2, scene1)
	NewScene(FrameRate, 2, scene2)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for frameId := uint32(0); frameId < sceneList.Scenes[0].NumberOfFrames; frameId++ {
			sceneList.Scenes[0].RenderFrame(frameId, []interface{}{frameId, sceneList.Scenes[0].NumberOfFrames})
		}
		sceneList.Scenes[0].generateVideo()
		sceneList.Scenes[0].clean()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for frameId := uint32(0); frameId < sceneList.Scenes[1].NumberOfFrames; frameId++ {
			sceneList.Scenes[1].RenderFrame(frameId, []interface{}{frameId, "Govima"})
		}
		sceneList.Scenes[1].generateVideo()
		sceneList.Scenes[1].clean()
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
	x := float64(frame) / float64(totalFrames) * float64(Width-rectWidth)
	y := float64(Height)/2 - rectHeight/2
	surf.Rectangle(x, y, rectWidth, rectHeight)
	surf.Fill()
}

func scene2(surf *cairo.Surface, args []interface{}) {
	frame := args[0].(uint32)
	videoTitle := args[1].(string)

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	surf.Rectangle(0, 0, Width, Height)
	surf.SetSourceRGB(0, 0, 0)
	surf.Fill()

	surf.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surf.SetFontSize(42)
	surf.SetSourceRGB(1, 1, 1)
	text := fmt.Sprintf("%s - Frame %d", videoTitle, frame)
	extents := surf.TextExtents(text)
	surf.MoveTo(float64(Width)/2-extents.Width/2, float64(Height)/2+extents.Height/2)
	surf.ShowText(text)
}
