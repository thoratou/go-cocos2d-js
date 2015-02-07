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
	GetPosition() Point
	SetPosition(Point)
	GetNormalizedPosition() Point
	SetNormalizedPosition(Point)
	GetPositionX() int
	SetPositionX(int)
	GetPositionY() int
	SetPositionY(int)
	GetChildrenCount() int
	GetChildren() []Node
	IsVisible() bool
	SetVisible(bool)
	GetAnchorPoint() Point
	SetAnchorPoint(Point)
	GetAnchorPointInPoints() Point
	GetContentSize() Size
	SetContentSize(Size)
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
	GetScheduler() Scheduler
	SetScheduler(Scheduler)
	GetBoundingBox() Rect
	Cleanup()
	GetChildByTag(int) Node
	GetChildByName(string) Node
	AddChild(Node, int, int)
	RemoveFromParent(bool)
	RemoveChild(Node, bool)
	RemoveChildByTag(int, bool)
	RemoveAllChildren(bool)
	SetNodeDirty()
	ReorderChild(Node, int)
	SortAllChildren()
	OnEnter(func())
	OnEnterTransitionDidFinish(func())
	OnExitTransitionDidStart(func())
	OnExit(func())
	RunAction(Action)
	StopAllActions()
	StopAction(Action)
	StopActionByTag(int)
	GetActionByTag() int
	GetNumberOfRunningActions() int
	ScheduleUpdate()
	ScheduleUpdateWithPriority(int)
	UnscheduleUpdate()
	Schedule(func(...interface{}), int, int, float64)
	Unschedule(func(...interface{}))
	UnscheduleAllCallbacks()
	Resume()
	Pause()
	SetAdditionalTransform(AffineTransform)
	GetParentToNodeTransform() AffineTransform
	GetNodeToWorldTransform() AffineTransform
	GetWorldToNodeTransform() AffineTransform
	ConvertToNodeSpace(Point) Point
	ConvertToWorldSpace(Point) Point
	ConvertToNodeSpaceAR(Point) Point
	ConvertToWorldSpaceAR(Point) Point
	ConvertTouchToNodeSpace(Touch) Point
	ConvertTouchToNodeSpaceAR(Touch) Point
	Update(float64)
	UpdateTransform()
	Retain()
	Release()
	GetComponent(string) Component
	AddComponent(Component)
	RemoveComponent(string)
	RemoveAllComponents()
	//TODO RenderCmd
	//Visit(RenderCmd)
	//Transform(RenderCmd, bool)
	//GetNodeToParentTransform() AffineTransform
	//TODO GLProgram
	//GetShaderProgram() GLProgram
	//SetShaderProgram(GLProgram)
	GetBoundingBoxToWorld() Rect
	GetOpacity() int
	GetDisplayedOpacity() int
	SetOpacity(int)
	UpdateDisplayedOpacity(int)
	IsCascadeOpacityEnabled() bool
	SetCascadeOpacityEnabled(bool)
	//TODO Color
	//GetColor() Color
	//GetDisplayedColor() Color
	//SetColor(Color)
	//UpdateDisplayedColor(Color)
	IsCascadeColorEnabled() bool
	SetCascadeColorEnabled(bool)
	SetOpacityModifyRGB(bool)
	IsOpacityModifyRGB() bool
}

type node struct{ js.Object }

func NewNode() Node {
	return &node{pcc.Get("Node").New()}
}
