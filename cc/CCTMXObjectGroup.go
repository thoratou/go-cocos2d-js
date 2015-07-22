package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

//////////////////
// TMXLayerInfo //
//////////////////

type TMXObjectGroup interface {
	GetPositionOffset() Point
	SetPositionOffset(Point)
	GetProperties() []*js.Object
	SetProperties(*js.Object)
	GetGroupName() string
	SetGroupName(string)
	PropertyNamed(string) *js.Object
	GetObject(string) *js.Object
	GetObjects() []*js.Object
	SetObjects(*js.Object)
}

type tmxObjectGroup struct{ *js.Object }

func NewTMXObjectGroup() TMXObjectGroup {
	return &tmxObjectGroup{pcc.Get("TMXObjectGroup").New()}
}

func (t *tmxObjectGroup) GetPositionOffset() Point {
	return &point{t.Call("getPositionOffset")}
}

func (t *tmxObjectGroup) SetPositionOffset(offset Point) {
	t.Call("setPositionOffset", offset)
}

func (t *tmxObjectGroup) GetProperties() []*js.Object {
	properties := t.Call("getProperties")
	length := properties.Length()
	out := make([]*js.Object, length, length)
	for i := 0; i < length; i++ {
		out[i] = properties.Index(i)
	}
	return out
}

func (t *tmxObjectGroup) SetProperties(properties *js.Object) {
	t.Call("setProperties", properties)
}

func (t *tmxObjectGroup) GetGroupName() string {
	return t.Call("getGroupName").String()
}

func (t *tmxObjectGroup) SetGroupName(groupName string) {
	t.Call("setGroupName", groupName)
}

func (t *tmxObjectGroup) PropertyNamed(propertyName string) *js.Object {
	return t.Call("propertyNamed", propertyName)
}

func (t *tmxObjectGroup) GetObject(objectName string) *js.Object {
	return t.Call("getObject", objectName)
}

func (t *tmxObjectGroup) GetObjects() []*js.Object {
	objects := t.Call("getObjects")
	length := objects.Length()
	out := make([]*js.Object, length, length)
	for i := 0; i < length; i++ {
		out[i] = objects.Index(i)
	}
	return out
}

func (t *tmxObjectGroup) SetObjects(objects *js.Object) {
	t.Call("setObjects", objects)
}
