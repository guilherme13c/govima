package scene

import "govima/app/misc"

type SceneList_t struct {
	Scenes []Scene_i
}

func (sl *SceneList_t) Add(s Scene_i) {
	if sl.Scenes == nil {
		sl.Scenes = make([]Scene_i, 0)
	}

	sl.Scenes = append(sl.Scenes, s)
	misc.NextId()
}

var SceneList SceneList_t
