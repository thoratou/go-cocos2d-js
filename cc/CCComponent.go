package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// The Component class
type Component interface {
	js.Object
	SetInit(func() bool)
	Init() bool
	SetOnEnter(func())
	OnEnter()
	SetOnExit(func())
	OnExit()
	SetUpdate(func(float64))
	Update(float64)
	//TODO Serialize(func(Reader))
	IsEnabled() bool
	SetEnabled(bool)
	GetName() string
	SetName(string)
	SetOwner(js.Object)
	GetOwner() js.Object
}

type component struct{ js.Object }

// NewComponent is the constructor for Component.
func NewComponent() Component {
	return &component{pcc.Get("Component").New()}
}

// Init initializes the component.
func (c *component) SetInit(cb func() bool) {
	BackupFunc(c, "init")
	c.Set("init", cb)
}

// Init initializes the component.
func (c *component) Init() bool {
	return c.Call("init").Bool()
}

// OnEnter is the callback when a component enter stage.
func (c *component) SetOnEnter(cb func()) {
	BackupFunc(c, "onEnter")
	c.Set("onEnter", cb)
}

// OnEnter is the callback when a component enter stage.
func (c *component) OnEnter() {
	c.Call("onEnter")
}

// OnExit is the callback when a component exit stage.
func (c *component) SetOnExit(cb func()) {
	BackupFunc(c, "onExit")
	c.Set("onExit", cb)
}

// OnExit is the callback when a component exit stage.
func (c *component) OnExit() {
	c.Call("onExit")
}

// Update is the callback per every frame if it schedules update.
func (c *component) SetUpdate(cb func(float64)) {
	BackupFunc(c, "update")
	c.Set("update", cb)
}

// Update is the callback per every frame if it schedules update.
func (c *component) Update(dt float64) {
	c.Call("update", dt)
}

// IsEnabled returns component whether is enabled.
func (c *component) IsEnabled() bool {
	return c.Call("isEnabled").Bool()
}

// SetEnabled sets component whether is enabled.
func (c *component) SetEnabled(enable bool) {
	c.Call("setEnabled", enable)
}

// GetName returns the Component name.
func (c *component) GetName() string {
	return c.Call("getName").String()
}

// SetName sets the Component name.
func (c *component) SetName(name string) {
	c.Call("setName", name)
}

// SetOwner sets the Component owner.
func (c *component) SetOwner(owner js.Object) {
	c.Call("setOwner", owner)
}

// GetOwner returns the Component owner.
func (c *component) GetOwner() js.Object {
	return c.Call("getOwner")
}
