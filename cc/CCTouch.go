package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// The touch event class
type Touch interface {
	GetLocation() Point
	GetLocationX() float64
	GetLocationY() float64
	GetPreviousLocation() Point
	GetStartLocation() Point
	GetDelta() Point
	GetLocationInView() Point
	GetPreviousLocationInView() Point
	GetStartLocationInView() Point
	GetID() int
	SetTouchInfo(int, float64, float64)
}

type touch struct{ *js.Object }

// NewTouch is the constructor for Touch.
func NewTouch(x float64, y float64, id int) Touch {
	return &touch{pcc.Get("Touch").New(x, y, id)}
}

// GetLocation returns the current touch location in OpenGL coordinates
func (t *touch) GetLocation() Point {
	return &point{t.Call("getLocation")}
}

// GetLocationX returns X axis location value
func (t *touch) GetLocationX() float64 {
	return t.Call("getLocationX").Float()
}

// GetLocationY returns Y axis location value
func (t *touch) GetLocationY() float64 {
	return t.Call("getLocationY").Float()
}

// GetPreviousLocation returns the previous touch location in OpenGL coordinates
func (t *touch) GetPreviousLocation() Point {
	return &point{t.Call("getPreviousLocation")}
}

// GetStartLocation returns the start touch location in OpenGL coordinates
func (t *touch) GetStartLocation() Point {
	return &point{t.Call("getStartLocation")}
}

// GetDelta returns the delta distance from the previous touche to the current one in screen coordinates
func (t *touch) GetDelta() Point {
	return &point{t.Call("getDelta")}
}

// GetLocationInView returns the current touch location in screen coordinates
func (t *touch) GetLocationInView() Point {
	return &point{t.Call("getLocationInView")}
}

// GetPreviousLocationInView returns the previous touch location in screen coordinates
func (t *touch) GetPreviousLocationInView() Point {
	return &point{t.Call("getPreviousLocationInView")}
}

// GetStartLocationInView returns the start touch location in screen coordinates
func (t *touch) GetStartLocationInView() Point {
	return &point{t.Call("getStartLocationInView")}
}

// GetID returns the id of the Touch event
func (t *touch) GetID() int {
	return t.Call("getID").Int()
}

// SetTouchInfo sets information to touch
func (t *touch) SetTouchInfo(id int, x float64, y float64) {
	t.Call("setTouchInfo", id, x, y)
}
