package misc

type Id_t uint32

var nextId Id_t = 0

func NextId() Id_t {
    id := nextId
    nextId++
	return id
}
