package imagescene

import (
	"fmt"
	"govima/app/misc"
	"govima/app/resource/config"
	"govima/app/scene"
	"log"

	"github.com/ungerik/go-cairo"
)

type ImageScene_t struct {
	id              misc.Id_t                                               // Scene Identifier
	state           map[string]interface{}                                  // Scene state
	outputImagePath string                                                  // Path of the generated video
	renderFunc      func(surf *cairo.Surface, state map[string]interface{}) // function to handle the renderization of the scene
	width           uint32                                                  // Canvas width
	height          uint32                                                  // Canvas height
}

func NewImageScene(width uint32, height uint32, renderFunc func(surf *cairo.Surface, state map[string]interface{}), initState map[string]interface{}) *ImageScene_t {
	id := misc.NextId()
	outputImagePath := fmt.Sprintf("%s/scene_%04d.png", config.Config.OutputDir, id)

	initState["width"] = width
	initState["height"] = height

	s := ImageScene_t{
		id:              id,
		renderFunc:      renderFunc,
		state:           initState,
		outputImagePath: outputImagePath,
		width:           width,
		height:          height,
	}
	scene.SceneList.Add(&s)

	return &s
}

func (s *ImageScene_t) Save() {
	s.renderFrame()
}

func (s *ImageScene_t) GetId() misc.Id_t {
	return s.id
}

func (s *ImageScene_t) GetWidth() uint32 {
	return s.width
}

func (s *ImageScene_t) GetHeight() uint32 {
	return s.width
}

// Render a single frame using Cairo
func (s *ImageScene_t) renderFrame() {
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, int(s.width), int(s.height))

	s.renderFunc(surface, s.state)

	// Save the frame as PNG
	status := surface.WriteToPNG(s.outputImagePath)
	if status != cairo.STATUS_SUCCESS {
		log.Fatalf("Failed to save frame to PNG: %v", status)
	}

	// Clean up
	surface.Finish()
}
