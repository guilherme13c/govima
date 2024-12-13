package group

import (
	"govima/app/misc"
	"govima/app/object"
	"math"

	"github.com/ungerik/go-cairo"
)

type Group_t struct {
	id      misc.Id_t
	width   float64
	height  float64
	objects []object.Object_i

	x float64
	y float64
}

func NewGroupObject(objs ...object.Object_i) *Group_t {
	id := misc.NextId()

	o := &Group_t{
		id:      id,
		objects: objs,
		x:       0,
		y:       0,
	}

	maxX, maxY, minX, minY := o.calculateDim()

	o.width = maxX - minX
	o.height = maxY - minY

	return o
}

func (o *Group_t) Add(objs ...object.Object_i) {
	o.objects = append(o.objects, objs...)
	o.calculateDim()
}

func (o *Group_t) GetId() misc.Id_t {
	return o.id
}

func (o *Group_t) Render(surf *cairo.Surface) {
	for _, obj := range o.objects {
		obj.Render(surf)
	}
}

func (o *Group_t) GetDim() (float64, float64) {
	return o.width, o.height
}

func (o *Group_t) GetPos() (float64, float64) {
	return o.x, o.y
}

func (o *Group_t) calculateDim() (float64, float64, float64, float64) {
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)
	minX := math.Inf(1)
	minY := math.Inf(1)

	for _, obj := range o.objects {
		x, y := obj.GetPos()
		maxX = max(maxX, x)
		maxY = max(maxY, y)
		minX = min(minX, x)
		minY = min(minY, y)
	}
	return maxX, maxY, minX, minY
}

func (o *Group_t) SetPos(x float64, y float64) {
	diffX := x - o.x
	diffY := y - o.y

	for _, obj := range o.objects {
		ox, oy := obj.GetPos()
		obj.SetPos(ox+diffX, oy+diffY)
	}

	o.x = x
	o.y = y
}
