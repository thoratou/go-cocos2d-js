package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type Node interface {
	js.Object
	Init()
	Attr(map[string]interface{})
	GetSkewX() float64
	SetSkewX(float64)
	GetSkewY() float64
	SetSkewY(float64)
	GetLocalZOrder() int
	SetLocalZOrder(int)
	GetGlobalZOrder() int
	SetGlobalZOrder(int)
	GetVertexZ() float64
	SetVertexZ(float64)
	GetRotation() float64
	SetRotation(float64)
	GetRotationX() float64
	SetRotationX(float64)
	GetRotationY() float64
	SetRotationY(float64)
	GetScale() float64
	SetScale(float64)
	GetScaleX() float64
	SetScaleX(float64)
	GetScaleY() float64
	SetScaleY(float64)
	GetPosition() *Point
	SetPosition(*Point)
	GetNormalizedPosition() *Point
	SetNormalizedPosition(*Point)
	GetPositionX() int
	SetPositionX(int)
	GetPositionY() int
	SetPositionY(int)
	GetChildrenCount() int
	GetChildren() []Node
	IsVisible() bool
	SetVisible(bool)
	GetAnchorPoint() *Point
	SetAnchorPoint(*Point)
	GetAnchorPointInPoints() *Point
	GetContentSize() *Size
	SetContentSize(*Size)
	IsRunning() bool
	GetParent() Node
	SetParent(Node)
	IsIgnoreAnchorPointForPosition() bool
	IgnoreAnchorPointForPosition(bool)
	GetTag() int
	SetTag(int)
	GetName() string
	SetName(string)
	GetUserData() js.Object
	SetUserData(js.Object)
	GetUserObject() js.Object
	SetUserObject(js.Object)
	GetOrderOfArrival() int
	SetOrderOfArrival(int)
	GetActionManager() ActionManager
	SetActionManager(ActionManager)
	//TODO scheduler
}

type node struct{ js.Object }

func NewNode() Node {
	return &node{pcc.Get("Node").New()}
}
