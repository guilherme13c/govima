package scene

import (
	"fmt"
	"govima/app/resource/config"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/ungerik/go-cairo"
)

type SceneId_t uint32

type Scene_t struct {
	Id              SceneId_t                                     // Scene Identifier
	NumberOfFrames  uint32                                        // Total number of frames
	frameRate       uint32                                        // Scene frame rate
	duration        uint32                                        // Duration of the scene in seconds
	frameDir        string                                        // Path to temporary directory to store frames
	outputVideoPath string                                        // Path of the generated video
	renderFunc      func(surf *cairo.Surface, args []interface{}) // function to handle the renderization of the scene
}

func NewScene(frameRate uint32, duration uint32, renderFunc func(surf *cairo.Surface, args []any)) {
	outputVideoPath := fmt.Sprintf("%s/scene_%04d.mp4", config.Config.OutputDir, SceneList.nextSceneId)

	s := &Scene_t{
		Id:              SceneList.nextSceneId,
		frameRate:       frameRate,
		duration:        duration,
		NumberOfFrames:  frameRate * duration,
		outputVideoPath: outputVideoPath,
		renderFunc:      renderFunc,
	}
	frameDir := s.createSceneDir()
	s.frameDir = frameDir

	SceneList.Add(s)
}

// Render a single frame using Cairo
func (s *Scene_t) RenderFrame(frameId uint32, args []interface{}) {
	filename := fmt.Sprintf("%s/frame_%08d.png", s.frameDir, frameId)
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(config.Config.Width), int(config.Config.Height))

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
func (s *Scene_t) GenerateVideo() error {
	framePattern := fmt.Sprintf("%s/frame_%%08d.png", s.frameDir)
	return ffmpeg.Input(framePattern, ffmpeg.KwArgs{"framerate": s.frameRate}).
		Output(s.outputVideoPath, config.Config.FFmpegArgs).
		OverWriteOutput().
		Run()
}

func (s *Scene_t) Clean() {
	if err := os.RemoveAll(s.frameDir); err != nil {
		log.Fatalf("Failed to remove frame directory: %v", err)
	}
}

func (s *Scene_t) createSceneDir() string {
	path := fmt.Sprintf("%s/scene_%04d", config.Config.FrameDir, s.Id)
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("Failed to create frame directory: %v", err)
	}
	return path
}
