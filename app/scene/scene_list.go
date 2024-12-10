package scene

type SceneList_t struct {
	nextSceneId SceneId_t
	Scenes      map[SceneId_t]*Scene_t
}

func (sl *SceneList_t) Add(s *Scene_t) {
	if sl.Scenes == nil {
		sl.Scenes = make(map[SceneId_t]*Scene_t)
	}

	sl.Scenes[s.Id] = s
	sl.nextSceneId++
}

var SceneList SceneList_t
