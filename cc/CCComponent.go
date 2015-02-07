package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// The Component class
type Component interface {
	js.Object
	Init() bool
	InitCb(func() bool)
	OnEnter(func())
	OnExit(func())
	Update(func(float64))
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
func (c *component) Init() bool {
	return c.Call("init").Bool()
}

func (c *component) InitCb(cb func() bool) {
	c.Set("init", cb)
}

// OnEnter is the callback when a component enter stage.
func (c *component) OnEnter(cb func()) {
	c.Set("onEnter", cb)
}

// OnExit is the callback when a component exit stage.
func (c *component) OnExit(cb func()) {
	c.Set("onExit", cb)
}

// Update is the callback per every frame if it schedules update.
func (c *component) Update(cb func(float64)) {
	c.Set("update", cb)
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
