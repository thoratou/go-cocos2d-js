// Package cc provides all high level functions and structures for interacting with Cocos native JavaScript APIs.
package cocoa

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	pcc = js.Global.Get("cc")
)

///////////
// POINT //
///////////

// Point is the class for point object.
type Point struct{ js.Object }

// NewPoint is the constructor for Point.
func NewPoint(x int, y int) *Point {
	return &Point{pcc.Call("p", x, y)}
}

// X returns the Point x value.
func (p *Point) X() int {
	return p.Get("x").Int()
}

// Y returns the Point x value.
func (p *Point) Y() int {
	return p.Get("y").Int()
}

// EqualsTo checks if points are equal.
func (p1 *Point) EqualsTo(p2 *Point) bool {
	return pcc.Call("pointEqualToPoint", p1, p2).Bool()
}

//////////
// SIZE //
//////////

// Size is the class for size object.
type Size struct{ js.Object }

// NewSize is the constructor for Size.
func NewSize(width int, height int) *Size {
	return &Size{pcc.Call("size", width, height)}
}

// Width returns the Size width.
func (s *Size) Width() int {
	return s.Get("width").Int()
}

// Height returns the Size height.
func (s *Size) Height() int {
	return s.Get("height").Int()
}

// EqualsTo checks if points are equal.
func (s1 *Size) EqualsTo(s2 *Size) bool {
	return pcc.Call("sizeEqualToSize", s1, s2).Bool()
}

//////////
// RECT //
//////////

// Rect is the class for rectangle object.
type Rect struct{ js.Object }

// NewRect is the constructor for Rect.
func NewRect(x int, y int, width int, height int) *Rect {
	return &Rect{pcc.Call("rect", x, y, width, height)}
}

// X returns the Rect x value.
func (r *Rect) X() int {
	return r.Get("x").Int()
}

// Y returns the Rect x value.
func (r *Rect) Y() int {
	return r.Get("y").Int()
}

// Width returns the Rect width.
func (r *Rect) Width() int {
	return r.Get("width").Int()
}

// Height returns the Rect height.
func (r *Rect) Height() int {
	return r.Get("height").Int()
}

// EqualsTo checks if points are equal.
func (r1 *Rect) EqualsTo(r2 *Rect) bool {
	return pcc.Call("rectEqualToRect", r1, r2).Bool()
}

// Contains checks if the first Rect contains the second one.
func (r1 *Rect) Contains(r2 *Rect) bool {
	return pcc.Call("rectContainsRect", r1, r2).Bool()
}

// RectGetMaxX returns the rightmost x-value of a rect.
func (r *Rect) RectGetMaxX() int {
	return pcc.Call("rectGetMaxX", r).Int()
}

// RectGetMidX returns the midpoint x-value of a rect.
func (r *Rect) RectGetMidX() int {
	return pcc.Call("rectGetMidX", r).Int()
}

// RectGetMinX returns the leftmost x-value of a rect.
func (r *Rect) RectGetMinX() int {
	return pcc.Call("rectGetMinX", r).Int()
}

// RectGetMaxY returns the topmost y-value of a rect.
func (r *Rect) RectGetMaxY() int {
	return pcc.Call("rectGetMaxY", r).Int()
}

// RectGetMidY returns the midpoint y-value of a rect.
func (r *Rect) RectGetMidY() int {
	return pcc.Call("rectGetMidY", r).Int()
}

// RectGetMinY returns the bottommost y-value of a rect.
func (r *Rect) RectGetMinY() int {
	return pcc.Call("rectGetMinY", r).Int()
}

// ContainsPoint checks whether a rect contains a point.
func (r *Rect) ContainsPoint(p *Point) bool {
	return pcc.Call("rectContainsPoint", r, p).Bool()
}

// Intersects checks whether a rect intersect with another.
func (r1 *Rect) Intersects(r2 *Rect) bool {
	return pcc.Call("rectIntersectsRect", r1, r2).Bool()
}

// Overlaps checks whether a rect overlaps another.
func (r1 *Rect) Overlaps(r2 *Rect) bool {
	return pcc.Call("rectOverlapsRect", r1, r2).Bool()
}

// Union returns the smallest rectangle that contains the two source rectangles.
func (r1 *Rect) Union(r2 *Rect) *Rect {
	return &Rect{pcc.Call("rectUnion", r1, r2)}
}

// Intersection Returns the overlapping portion of 2 rectangles.
func (r1 *Rect) Intersection(r2 *Rect) *Rect {
	return &Rect{pcc.Call("rectIntersection", r1, r2)}
}
