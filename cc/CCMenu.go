package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	MENU_STATE_WAITING        = 0
	MENU_STATE_TRACKING_TOUCH = 1
	MENU_HANDLER_PRIORITY     = -128
	DEFAULT_PADDING           = 5
)

type Menu interface {
	Layer
	//SetOnEnter(func()) defined as part of Node
	//OnEnter() defined as part of Node
	//OnEnterSuper() defined as part of Node
	IsEnabled() bool
	SetEnabled(bool)
	InitWithItems(...MenuItem) bool
	InitWithArray([]MenuItem) bool
	//AddChild(Node) defined as part of Node
	//AddChildWithOrder(Node, int) defined as part of Node
	//AddChildWithTag(Node, int) defined as part of Node
	//AddChildWithOrderAndTag(Node, int, int) defined as part of Node
	AlignItemsVertically()
	AlignItemsVerticallyWithPadding(int)
	AlignItemsHorizontally()
	AlignItemsHorizontallyWithPadding(int)
	AlignItemsInColumns(...int)
	AlignItemsInRows(...int)
	//RemoveChild(Node, bool) defined as part of Node
	//SetOnExit(func()) defined as part of Node
	//OnExit() defined as part of Node
	//OnExitSuper() defined as part of Node
	//SetOpacityModifyRGB(bool) defined as part of Node
	//IsOpacityModifyRGB() bool defined as part of Node

}

type menu struct{ layer }

// NewMenu is the constructor for Menu.
func NewMenu(menuItems ...interface{}) Menu {
	//Menu accepts array as parameter
	return &menu{layer{node{pcc.Get("Menu").New(menuItems)}}}
}

// Event callback that is invoked every time when Menu enters the 'stage'.
// If the Menu enters the 'stage' with a transition, this event is called when the transition starts.
// During OnEnter you can't access a "sister/brother" node.
// If you override OnEnter, you must call its parent's OnEnter function with OnEnterSuper().
func (m *menu) SetOnEnter(cb func()) {
	BackupFunc(m.Object, "onEnter")
	m.Set("onEnter", cb)
}

// Event callback that is invoked every time when Menu enters the 'stage'.
// If the Menu enters the 'stage' with a transition, this event is called when the transition starts.
// During OnEnter you can't access a "sister/brother" node.
// If you override OnEnter, you must call its parent's OnEnter function with OnEnterSuper().
func (m *menu) OnEnter() {
	m.Call("onEnter")
}

// Event callback that is invoked every time when Menu enters the 'stage'.
// If the Menu enters the 'stage' with a transition, this event is called when the transition starts.
// During OnEnter you can't access a "sister/brother" node.
// If you override OnEnter, you must call its parent's OnEnter function with OnEnterSuper().
func (m *menu) OnEnterSuper() {
	SuperCall(m.Object, "onEnter")
}

//IsEnabled returns whether or not the menu will receive events.
func (m *menu) IsEnabled() bool {
	return m.Call("isEnabled").Bool()
}

// SetEnabled sets whether or not the menu will receive events.
func (m *menu) SetEnabled(enable bool) {
	m.Call("setEnabled", enable)
}

// InitWithItems initializes a Menu with its items.
func (m *menu) InitWithItems(menuItems ...MenuItem) bool {
	return m.Call("initWithItems", menuItems).Bool()
}

// InitWithItems initializes a Menu with a Array of MenuItem objects.
func (m *menu) InitWithArray(menuItems []MenuItem) bool {
	return m.Call("initWithArray", menuItems).Bool()
}

// AddChild add a child for Menu
func (m *menu) AddChild(child Node) {
	m.Call("addChild", child, js.Undefined, js.Undefined)
}

// AddChild add a child for Menu
func (m *menu) AddChildWithOrder(child Node, localZOrder int) {
	m.Call("addChild", child, localZOrder, js.Undefined)
}

// AddChild add a child for Menu
func (m *menu) AddChildWithTag(child Node, tag int) {
	m.Call("addChild", child, js.Undefined, tag)
}

// AddChild add a child for Menu
func (m *menu) AddChildWithOrderAndTag(child Node, localZOrder int, tag int) {
	m.Call("addChild", child, localZOrder, tag)
}

// AlignItemsVertically aligns items vertically with default padding
func (m *menu) AlignItemsVertically() {
	m.Call("alignItemsVertically")
}

// AlignItemsVerticallyWithPadding aligns items vertically with specified padding
func (m *menu) AlignItemsVerticallyWithPadding(padding int) {
	m.Call("alignItemsVerticallyWithPadding", padding)
}

// AlignItemsHorizontally aligns items horizontally with default padding
func (m *menu) AlignItemsHorizontally() {
	m.Call("alignItemsHorizontally")
}

// AlignItemsHorizontallyWithPadding aligns items horizontally with specified padding
func (m *menu) AlignItemsHorizontallyWithPadding(padding int) {
	m.Call("alignItemsVerticallyWithPadding", padding)
}

// AlignItemsInColumns aligns items in columns
// Example
// menu.AlignItemsInColumns(3,2,3) // this will create 3 columns, with 3 items for first column, 2 items for second and 3 for third
// menu.AlignItemsInColumns(3,3)   // this creates 2 columns, each have 3 items
func (m *menu) AlignItemsInColumns(args ...int) {
	m.Call("alignItemsInColumns", args)
}

// AlignItemsInRows align menu items in rows
// Example
// menu.AlignItemsInRows(5,3) // this will align items to 2 rows, first row with 5 items, second row with 3
// menu.AlignItemsInRows(4,4,4,4)//this creates 4 rows each have 4 items
func (m *menu) AlignItemsInRows(args ...int) {
	m.Call("alignItemsInRows", args)
}

// RemoveChild removes a child from Menu
func (m *menu) RemoveChild(child Node, cleanup bool) {
	m.Call("removeChild", child, cleanup)
}

// Event callback that is called every time the Menu leaves the 'stage'.
// If the Menu leaves the 'stage' with a transition, this callback is called when the transition finishes.
// During OnExit you can't access a sibling node.
// If you override OnExit, you shall call its parent's OnExit with OnExitSuper().
func (m *menu) SetOnExit(cb func()) {
	BackupFunc(m.Object, "onExit")
	m.Set("onExit", cb)
}

// Event callback that is called every time the Menu leaves the 'stage'.
// If the Menu leaves the 'stage' with a transition, this callback is called when the transition finishes.
// During OnExit you can't access a sibling node.
// If you override OnExit, you shall call its parent's OnExit with OnExitSuper().
func (m *menu) OnExit() {
	m.Call("onExit")
}

// Event callback that is called every time the Menu leaves the 'stage'.
// If the Menu leaves the 'stage' with a transition, this callback is called when the transition finishes.
// During OnExit you can't access a sibling node.
// If you override OnExit, you shall call its parent's OnExit with OnExitSuper().
func (m *menu) OnExitSuper() {
	SuperCall(m.Object, "onExit")
}

// only use for jsbinding
func (m *menu) SetOpacityModifyRGB(opacityValue bool) {
	m.Call("setOpacityModifyRGB", opacityValue)
}

// only use for jsbinding
func (m *menu) IsOpacityModifyRGB() bool {
	return m.Call("isOpacityModifyRGB").Bool()
}
