package videoscene

import (
	"fmt"
	"govima/app/misc"
	"govima/app/resource/config"
	"govima/app/scene"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/ungerik/go-cairo"
)

type VideoScene_t struct {
	id              misc.Id_t                                               // Scene Identifier
	numberOfFrames  uint32                                                  // Total number of frames
	state           map[string]interface{}                                  // Scene state
	frameRate       uint32                                                  // Scene frame rate
	duration        uint32                                                  // Duration of the scene in seconds
	frameDir        string                                                  // Path to temporary directory to store frames
	outputVideoPath string                                                  // Path of the generated video
	renderFunc      func(surf *cairo.Surface, state map[string]interface{}) // function to handle the renderization of the scene
}

func NewVideoScene(frameRate uint32, duration uint32, renderFunc func(surf *cairo.Surface, state map[string]interface{}), initState map[string]interface{}) *VideoScene_t {
	id := misc.NextId()
	outputVideoPath := fmt.Sprintf("%s/scene_%04d.mp4", config.Config.OutputDir, id)

	s := VideoScene_t{
		id:              id,
		frameRate:       frameRate,
		duration:        duration,
		numberOfFrames:  frameRate * duration,
		outputVideoPath: outputVideoPath,
		renderFunc:      renderFunc,
		state:           initState,
	}
	frameDir := s.createSceneDir()
	s.frameDir = frameDir

	scene.SceneList.Add(s)

	return &s
}

func (s VideoScene_t) Save() {
	for frameId := uint32(0); frameId < s.numberOfFrames; frameId++ {
		s.renderFrame(frameId)
	}
	s.generateVideo()
	s.clean()
}

func (s VideoScene_t) GetId() misc.Id_t {
	return s.id
}

// Render a single frame using Cairo
func (s VideoScene_t) renderFrame(frameId uint32) {
	filename := fmt.Sprintf("%s/frame_%08d.png", s.frameDir, frameId)
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(config.Config.Width), int(config.Config.Height))

	s.renderFunc(surface, s.state)

	// Save the frame as PNG
	status := surface.WriteToPNG(filename)
	if status != cairo.STATUS_SUCCESS {
		log.Fatalf("Failed to save frame to PNG: %v", status)
	}

	// Clean up
	surface.Finish()
}

// Generate a video from the scene frames using FFmpeg
func (s VideoScene_t) generateVideo() error {
	framePattern := fmt.Sprintf("%s/frame_%%08d.png", s.frameDir)
	return ffmpeg.Input(framePattern, ffmpeg.KwArgs{"framerate": s.frameRate}).
		Output(s.outputVideoPath, config.Config.FFmpegArgs).
		GlobalArgs("-loglevel", "quiet").
		OverWriteOutput().
		Run()
}

func (s VideoScene_t) clean() {
	if err := os.RemoveAll(s.frameDir); err != nil {
		log.Fatalf("Failed to remove frame directory: %v", err)
	}
}

func (s VideoScene_t) createSceneDir() string {
	path := fmt.Sprintf("%s/scene_%04d", config.Config.FrameDir, s.id)
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalf("Failed to create frame directory: %v", err)
	}
	return path
}
