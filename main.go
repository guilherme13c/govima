package main

import (
	"sync"

	colorconst "govima/app/misc/constants/color"
	"govima/app/object"
	"govima/app/object/group"
	"govima/app/object/shape"
	"govima/app/resource/config"
	"govima/app/scene"
	videoscene "govima/app/scene/video_scene"

	"github.com/ungerik/go-cairo"
)

func main() {
	config.Init()

	square := shape.NewRectangleObject(100, 100, colorconst.Orange)
	square.Fill = true

	circle := shape.NewRegularPolygonObject(3, 80, colorconst.Blue)
	circle.Fill = true
	circle.SetPos(0, -50)

	g := group.NewGroupObject(square, circle)
	g.SetPos(1920/2, 1080/2)

	fps := uint32(60)
	duration := uint32(3)
	videoscene.NewVideoScene(1920, 1080, fps, duration, scene1Func, map[string]interface{}{
		"totalFrames": int(duration * fps),
		"frameId":     0,
		"objects": map[string]object.Object_i{
			"group": g,
		},
	})

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
	group := state["objects"].(map[string]object.Object_i)["group"].(object.Object_i)

	surf.SetAntialias(cairo.ANTIALIAS_GRAY)

	x := float64(frameId) * width / (float64(totalFrames) - 1)
	group.SetPos(x, height/2)
	group.Render(surf)

	state["frameId"] = frameId + 1
}
