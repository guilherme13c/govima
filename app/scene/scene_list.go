package scene

type SceneList_t struct {
	nextSceneId SceneId_t
	Scenes      []Scene_i
}

func (sl *SceneList_t) Add(s Scene_i) {
	if sl.Scenes == nil {
		sl.Scenes = make([]Scene_i, 0)
	}

	sl.Scenes = append(sl.Scenes, s)
	sl.nextSceneId++
}

func (sl *SceneList_t) GetNextId() SceneId_t {
	return sl.nextSceneId
}

var SceneList SceneList_t
