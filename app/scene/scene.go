package scene

import "govima/app/misc"

type Scene_i interface {
	Save()
	GetId() misc.Id_t
}
