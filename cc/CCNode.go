package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type Node interface {
	js.Object
	Init() bool
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
	AddChild(Node)
	AddChildWithOrder(Node, int)
	AddChildWithTag(Node, int)
	AddChildWithOrderAndTag(Node, int, int)
	RemoveFromParent(bool)
	RemoveChild(Node, bool)
	RemoveChildByTag(int, bool)
	RemoveAllChildren(bool)
	SetNodeDirty()
	ReorderChild(Node, int)
	SortAllChildren()
	SetOnEnter(func())
	OnEnter()
	OnEnterSuper()
	SetOnEnterTransitionDidFinish(func())
	OnEnterTransitionDidFinish()
	OnEnterTransitionDidFinishSuper()
	SetOnExitTransitionDidStart(func())
	OnExitTransitionDidStart()
	OnExitTransitionDidStartSuper()
	SetOnExit(func())
	OnExit()
	OnExitSuper()
	RunAction(Action)
	StopAllActions()
	StopAction(Action)
	StopActionByTag(int)
	GetActionByTag(int) Action
	GetNumberOfRunningActions() int
	ScheduleUpdate()
	ScheduleUpdateWithPriority(int)
	UnscheduleUpdate()
	Schedule(func(...interface{}), float64, int, float64)
	ScheduleOnce(func(...interface{}), float64)
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
	GetNodeToParentTransform() AffineTransform
	GetShaderProgram() GLProgram
	SetShaderProgram(GLProgram)
	GetBoundingBoxToWorld() Rect
	GetOpacity() int
	GetDisplayedOpacity() int
	SetOpacity(int)
	UpdateDisplayedOpacity(int)
	IsCascadeOpacityEnabled() bool
	SetCascadeOpacityEnabled(bool)
	GetColor() Color
	GetDisplayedColor() Color
	SetColor(Color)
	UpdateDisplayedColor(Color)
	IsCascadeColorEnabled() bool
	SetCascadeColorEnabled(bool)
	SetOpacityModifyRGB(bool)
	IsOpacityModifyRGB() bool
}

type node struct{ js.Object }

func NewNode() Node {
	return &node{pcc.Get("Node").New()}
}

//TODO: tmp
func NewNodeJs(o js.Object) Node {
	return &node{o}
}

func (n *node) Init() bool {
	return n.Call("init").Bool()
}

func (n *node) Attr(attrs map[string]interface{}) {
	n.Call("attr", attrs)
}

func (n *node) GetSkewX() float64 {
	return n.Call("getSkewX").Float()
}

func (n *node) SetSkewX(newSkewX float64) {
	n.Call("setSkewX", newSkewX)
}

func (n *node) GetSkewY() float64 {
	return n.Call("getSkewY").Float()
}

func (n *node) SetSkewY(newSkewY float64) {
	n.Call("setSkewY", newSkewY)
}

func (n *node) GetLocalZOrder() int {
	return n.Call("getLocalZOrder").Int()
}

func (n *node) SetLocalZOrder(localZOrder int) {
	n.Call("setLocalZOrder", localZOrder)
}

func (n *node) GetGlobalZOrder() int {
	return n.Call("getGlobalZOrder").Int()
}

func (n *node) SetGlobalZOrder(globalZOrder int) {
	n.Call("setGlobalZOrder", globalZOrder)
}

func (n *node) GetVertexZ() float64 {
	return n.Call("getVertexZ").Float()
}

func (n *node) SetVertexZ(z float64) {
	n.Call("setVertexZ", z)
}

func (n *node) GetRotation() float64 {
	return n.Call("getRotation").Float()
}

func (n *node) SetRotation(newRotation float64) {
	n.Call("setRotation", newRotation)
}

func (n *node) GetRotationX() float64 {
	return n.Call("getRotationX").Float()
}

func (n *node) SetRotationX(rotationX float64) {
	n.Call("setRotationX", rotationX)
}

func (n *node) GetRotationY() float64 {
	return n.Call("getRotationY").Float()
}

func (n *node) SetRotationY(rotationY float64) {
	n.Call("setRotationY", rotationY)
}

func (n *node) GetScale() float64 {
	return n.Call("getScale").Float()
}

func (n *node) SetScale(scale float64) {
	n.Call("setScale", scale)
}

func (n *node) GetScaleX() float64 {
	return n.Call("getScaleX").Float()
}

func (n *node) SetScaleX(scaleX float64) {
	n.Call("setScaleX", scaleX)
}

func (n *node) GetScaleY() float64 {
	return n.Call("getScaleY").Float()
}

func (n *node) SetScaleY(scaleY float64) {
	n.Call("setScaleY", scaleY)
}

func (n *node) GetPosition() Point {
	return &point{n.Call("getPosition")}
}

func (n *node) SetPosition(newPos Point) {
	n.Call("setPosition", newPos)
}

func (n *node) GetNormalizedPosition() Point {
	return &point{n.Call("getNormalizedPosition")}
}

func (n *node) SetNormalizedPosition(newPos Point) {
	n.Call("setNormalizedPosition", newPos)
}

func (n *node) GetPositionX() int {
	return n.Call("getPositionX").Int()
}

func (n *node) SetPositionX(x int) {
	n.Call("setPositionX", x)
}

func (n *node) GetPositionY() int {
	return n.Call("getPositionY").Int()
}

func (n *node) SetPositionY(y int) {
	n.Call("setPositionY", y)
}

func (n *node) GetChildrenCount() int {
	return n.Call("getChildrenCount").Int()
}

func (n *node) GetChildren() []Node {
	array := n.Call("getChildren")
	length := array.Length()
	result := make([]Node, length)

	for i := 0; i < length; i++ {
		result[i] = &node{array.Index(i)}
	}

	return result
}

func (n *node) IsVisible() bool {
	return n.Call("isVisible").Bool()
}

func (n *node) SetVisible(visible bool) {
	n.Call("setVisible", visible)
}

func (n *node) GetAnchorPoint() Point {
	return &point{n.Call("getAnchorPoint")}
}

func (n *node) SetAnchorPoint(point Point) {
	n.Call("setAnchorPoint", point)
}

func (n *node) GetAnchorPointInPoints() Point {
	return &point{n.Call("getAnchorPointInPoints")}
}

func (n *node) GetContentSize() Size {
	return &size{n.Call("getContentSize")}
}

func (n *node) SetContentSize(size Size) {
	n.Call("setContentSize", size)
}

func (n *node) IsRunning() bool {
	return n.Call("isRunning").Bool()
}

func (n *node) GetParent() Node {
	return &node{n.Call("getParent")}
}

func (n *node) SetParent(parent Node) {
	n.Call("setParent", parent)
}

func (n *node) IsIgnoreAnchorPointForPosition() bool {
	return n.Call("isIgnoreAnchorPointForPosition").Bool()
}

func (n *node) IgnoreAnchorPointForPosition(newValue bool) {
	n.Call("ignoreAnchorPointForPosition", newValue)
}

func (n *node) GetTag() int {
	return n.Call("getTag").Int()
}

func (n *node) SetTag(tag int) {
	n.Call("setTag", tag)
}

func (n *node) GetName() string {
	return n.Call("getName").String()
}

func (n *node) SetName(name string) {
	n.Call("setName", name)
}

func (n *node) GetUserData() js.Object {
	return n.Call("getUserData")
}

func (n *node) SetUserData(data js.Object) {
	n.Call("setUserData", data)
}

func (n *node) GetUserObject() js.Object {
	return n.Call("getUserObject")
}

func (n *node) SetUserObject(object js.Object) {
	n.Call("setUserObject", object)
}

func (n *node) GetOrderOfArrival() int {
	return n.Call("getOrderOfArrival").Int()
}

func (n *node) SetOrderOfArrival(order int) {
	n.Call("setOrderOfArrival", order)
}

func (n *node) GetActionManager() ActionManager {
	return &actionManager{n.Call("getActionManager")}
}

func (n *node) SetActionManager(am ActionManager) {
	n.Call("setActionManager", am)
}

func (n *node) GetScheduler() Scheduler {
	return &scheduler{n.Call("getScheduler")}
}

func (n *node) SetScheduler(s Scheduler) {
	n.Call("setScheduler", s)
}

func (n *node) GetBoundingBox() Rect {
	return &rect{n.Call("getBoundingBox")}
}

func (n *node) Cleanup() {
	n.Call("cleanup")
}

func (n *node) GetChildByTag(tag int) Node {
	return &node{n.Call("getChildByTag", tag)}
}

func (n *node) GetChildByName(name string) Node {
	return &node{n.Call("getChildByName", name)}
}

func (n *node) AddChild(child Node) {
	n.Call("addChild", child, js.Undefined, js.Undefined)
}

func (n *node) AddChildWithOrder(child Node, localZOrder int) {
	n.Call("addChild", child, localZOrder, js.Undefined)
}

func (n *node) AddChildWithTag(child Node, tag int) {
	n.Call("addChild", child, js.Undefined, tag)
}

func (n *node) AddChildWithOrderAndTag(child Node, localZOrder int, tag int) {
	n.Call("addChild", child, localZOrder, tag)
}

func (n *node) RemoveFromParent(cleanup bool) {
	n.Call("removeFromParent", cleanup)
}

func (n *node) RemoveChild(child Node, cleanup bool) {
	n.Call("removeChild", child, cleanup)
}

func (n *node) RemoveChildByTag(tag int, cleanup bool) {
	n.Call("removeChildByTag", tag, cleanup)
}

func (n *node) RemoveAllChildren(cleanup bool) {
	n.Call("removeAllChildren", cleanup)
}

func (n *node) SetNodeDirty() {
	n.Call("setNodeDirty")
}

func (n *node) ReorderChild(child Node, zOrder int) {
	n.Call("reorderChild", child, zOrder)
}

func (n *node) SortAllChildren() {
	n.Call("sortAllChildren")
}

func (n *node) SetOnEnter(cb func()) {
	BackupFunc(n, "onEnter")
	n.Set("onEnter", cb)
}

func (n *node) OnEnter() {
	n.Call("onEnter")
}

func (n *node) OnEnterSuper() {
	SuperCall(n, "onEnter")
}

func (n *node) SetOnEnterTransitionDidFinish(cb func()) {
	BackupFunc(n, "onEnterTransitionDidFinish")
	n.Set("onEnterTransitionDidFinish", cb)
}

func (n *node) OnEnterTransitionDidFinish() {
	n.Call("onEnterTransitionDidFinish")
}

func (n *node) OnEnterTransitionDidFinishSuper() {
	SuperCall(n, "onEnterTransitionDidFinish")
}

func (n *node) SetOnExitTransitionDidStart(cb func()) {
	BackupFunc(n, "onExitTransitionDidStart")
	n.Set("onExitTransitionDidStart", cb)
}

func (n *node) OnExitTransitionDidStart() {
	n.Call("onExitTransitionDidStart")
}

func (n *node) OnExitTransitionDidStartSuper() {
	SuperCall(n, "onExitTransitionDidStart")
}

func (n *node) SetOnExit(cb func()) {
	BackupFunc(n, "onExit")
	n.Set("onExit", cb)
}

func (n *node) OnExit() {
	n.Call("onExit")
}

func (n *node) OnExitSuper() {
	SuperCall(n, "onExit")
}

func (n *node) RunAction(action Action) {
	n.Call("runAction", action)
}

func (n *node) StopAllActions() {
	n.Call("stopAllActions")
}

func (n *node) StopAction(action Action) {
	n.Call("stopAction", action)
}

func (n *node) StopActionByTag(tag int) {
	n.Call("stopActionByTag", tag)
}

func (n *node) GetActionByTag(tag int) Action {
	return &action{n.Call("getActionByTag", tag)}
}

func (n *node) GetNumberOfRunningActions() int {
	return n.Call("getNumberOfRunningActions").Int()
}

func (n *node) ScheduleUpdate() {
	n.Call("scheduleUpdate")
}

func (n *node) ScheduleUpdateWithPriority(priority int) {
	n.Call("scheduleUpdateWithPriority", priority)
}

func (n *node) UnscheduleUpdate() {
	n.Call("unscheduleUpdate")
}

func (n *node) Schedule(cb func(...interface{}), interval float64, repeat int, delay float64) {
	n.Call("scheduleUpdateWithPriority", cb, interval, repeat, delay)
}

func (n *node) ScheduleOnce(cb func(...interface{}), delay float64) {
	n.Call("scheduleOnce", cb, delay)
}

func (n *node) Unschedule(cb func(...interface{})) {
	n.Call("unschedule", cb)
}

func (n *node) UnscheduleAllCallbacks() {
	n.Call("unscheduleAllCallbacks")
}

func (n *node) Resume() {
	n.Call("resume")
}

func (n *node) Pause() {
	n.Call("pause")
}

func (n *node) SetAdditionalTransform(at AffineTransform) {
	n.Call("setAdditionalTransform", at)
}

func (n *node) GetParentToNodeTransform() AffineTransform {
	return &affineTransform{n.Call("getParentToNodeTransform")}
}

func (n *node) GetNodeToWorldTransform() AffineTransform {
	return &affineTransform{n.Call("getNodeToWorldTransform")}
}

func (n *node) GetWorldToNodeTransform() AffineTransform {
	return &affineTransform{n.Call("getWorldToNodeTransform")}
}

func (n *node) ConvertToNodeSpace(worldPoint Point) Point {
	return &point{n.Call("convertToNodeSpace", worldPoint)}
}

func (n *node) ConvertToWorldSpace(nodePoint Point) Point {
	return &point{n.Call("convertToWorldSpace", nodePoint)}
}

func (n *node) ConvertToNodeSpaceAR(worldPoint Point) Point {
	return &point{n.Call("convertToNodeSpaceAR", worldPoint)}
}

func (n *node) ConvertToWorldSpaceAR(nodePoint Point) Point {
	return &point{n.Call("convertToWorldSpaceAR", nodePoint)}
}

func (n *node) ConvertTouchToNodeSpace(touch Touch) Point {
	return &point{n.Call("convertTouchToNodeSpace", touch)}
}

func (n *node) ConvertTouchToNodeSpaceAR(touch Touch) Point {
	return &point{n.Call("convertTouchToNodeSpaceAR", touch)}
}

func (n *node) Update(dt float64) {
	n.Call("update", dt)
}

func (n *node) UpdateTransform() {
	n.Call("updateTransform")
}

func (n *node) Retain() {
	n.Call("retain")
}

func (n *node) Release() {
	n.Call("release")
}

func (n *node) GetComponent(name string) Component {
	return &component{n.Call("getComponent", name)}
}

func (n *node) AddComponent(c Component) {
	n.Call("addComponent", c)
}

func (n *node) RemoveComponent(name string) {
	n.Call("removeComponent", name)
}

func (n *node) RemoveAllComponents() {
	n.Call("removeAllComponents")
}

//TODO RenderCmd
//func (n *node) Visit(RenderCmd)
//func (n *node) Transform(RenderCmd, bool)

func (n *node) GetNodeToParentTransform() AffineTransform {
	return &affineTransform{n.Call("getNodeToParentTransform")}
}

func (n *node) GetShaderProgram() GLProgram {
	return &glProgram{n.Call("getShaderProgram")}
}

func (n *node) SetShaderProgram(newShaderProgram GLProgram) {
	n.Call("setShaderProgram", newShaderProgram)
}

func (n *node) GetBoundingBoxToWorld() Rect {
	return &rect{n.Call("getBoundingBoxToWorld")}
}

func (n *node) GetOpacity() int {
	return n.Call("getOpacity").Int()
}

func (n *node) GetDisplayedOpacity() int {
	return n.Call("getDisplayedOpacity").Int()
}

func (n *node) SetOpacity(opacity int) {
	n.Call("setOpacity", opacity)
}

func (n *node) UpdateDisplayedOpacity(parentOpacity int) {
	n.Call("updateDisplayedOpacity", parentOpacity)
}

func (n *node) IsCascadeOpacityEnabled() bool {
	return n.Call("isCascadeOpacityEnabled").Bool()
}

func (n *node) SetCascadeOpacityEnabled(cascadeOpacityEnabled bool) {
	n.Call("setCascadeOpacityEnabled", cascadeOpacityEnabled)
}

func (n *node) GetColor() Color {
	return &color{n.Call("getColor")}
}

func (n *node) GetDisplayedColor() Color {
	return &color{n.Call("getDisplayedColor")}
}

func (n *node) SetColor(c Color) {
	n.Call("setColor", c)
}

func (n *node) UpdateDisplayedColor(c Color) {
	n.Call("updateDisplayedColor", c)
}

func (n *node) IsCascadeColorEnabled() bool {
	return n.Call("isCascadeColorEnabled").Bool()
}

func (n *node) SetCascadeColorEnabled(cascadeColorEnabled bool) {
	n.Call("setCascadeColorEnabled", cascadeColorEnabled)
}

func (n *node) SetOpacityModifyRGB(opacityValue bool) {
	n.Call("setOpacityModifyRGB", opacityValue)
}

func (n *node) IsOpacityModifyRGB() bool {
	return n.Call("isOpacityModifyRGB").Bool()
}
