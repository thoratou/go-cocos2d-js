package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

///////////
// POINT //
///////////

// Point is the class for point object.
type Point interface {
	X() int
	Y() int
	EqualsTo(Point) bool
}

type point struct{ js.Object }

// NewPoint is the constructor for Point.
func NewPoint(x int, y int) Point {
	return &point{pcc.Call("p", x, y)}
}

// X returns the Point x value.
func (p *point) X() int {
	return p.Get("x").Int()
}

// Y returns the Point x value.
func (p *point) Y() int {
	return p.Get("y").Int()
}

// EqualsTo checks if points are equal.
func (p1 *point) EqualsTo(p2 Point) bool {
	return pcc.Call("pointEqualToPoint", p1, p2).Bool()
}

//////////
// SIZE //
//////////

// Size is the class for size object.
type Size interface {
	Width() int
	Height() int
	EqualsTo(Size) bool
}

type size struct{ js.Object }

// NewSize is the constructor for Size.
func NewSize(width int, height int) Size {
	return &size{pcc.Call("size", width, height)}
}

// Width returns the Size width.
func (s *size) Width() int {
	return s.Get("width").Int()
}

// Height returns the Size height.
func (s *size) Height() int {
	return s.Get("height").Int()
}

// EqualsTo checks if points are equal.
func (s1 *size) EqualsTo(s2 Size) bool {
	return pcc.Call("sizeEqualToSize", s1, s2).Bool()
}

//////////
// RECT //
//////////

// Rect is the class for rectangle object.
type Rect interface {
	X() int
	Y() int
	Width() int
	Height() int
	EqualsTo(Rect) bool
	Contains(Rect) bool
	RectGetMaxX() int
	RectGetMidX() int
	RectGetMinX() int
	RectGetMaxY() int
	RectGetMidY() int
	RectGetMinY() int
	ContainsPoint(Point) bool
	Intersects(Rect) bool
	Overlaps(Rect) bool
	Union(Rect) Rect
	Intersection(Rect) Rect
}

type rect struct{ js.Object }

// NewRect is the constructor for Rect.
func NewRect(x int, y int, width int, height int) Rect {
	return &rect{pcc.Call("rect", x, y, width, height)}
}

// X returns the Rect x value.
func (r *rect) X() int {
	return r.Get("x").Int()
}

// Y returns the Rect x value.
func (r *rect) Y() int {
	return r.Get("y").Int()
}

// Width returns the Rect width.
func (r *rect) Width() int {
	return r.Get("width").Int()
}

// Height returns the Rect height.
func (r *rect) Height() int {
	return r.Get("height").Int()
}

// EqualsTo checks if points are equal.
func (r1 *rect) EqualsTo(r2 Rect) bool {
	return pcc.Call("rectEqualToRect", r1, r2).Bool()
}

// Contains checks if the first Rect contains the second one.
func (r1 *rect) Contains(r2 Rect) bool {
	return pcc.Call("rectContainsRect", r1, r2).Bool()
}

// RectGetMaxX returns the rightmost x-value of a rect.
func (r *rect) RectGetMaxX() int {
	return pcc.Call("rectGetMaxX", r).Int()
}

// RectGetMidX returns the midpoint x-value of a rect.
func (r *rect) RectGetMidX() int {
	return pcc.Call("rectGetMidX", r).Int()
}

// RectGetMinX returns the leftmost x-value of a rect.
func (r *rect) RectGetMinX() int {
	return pcc.Call("rectGetMinX", r).Int()
}

// RectGetMaxY returns the topmost y-value of a rect.
func (r *rect) RectGetMaxY() int {
	return pcc.Call("rectGetMaxY", r).Int()
}

// RectGetMidY returns the midpoint y-value of a rect.
func (r *rect) RectGetMidY() int {
	return pcc.Call("rectGetMidY", r).Int()
}

// RectGetMinY returns the bottommost y-value of a rect.
func (r *rect) RectGetMinY() int {
	return pcc.Call("rectGetMinY", r).Int()
}

// ContainsPoint checks whether a rect contains a point.
func (r *rect) ContainsPoint(p Point) bool {
	return pcc.Call("rectContainsPoint", r, p).Bool()
}

// Intersects checks whether a rect intersect with another.
func (r1 *rect) Intersects(r2 Rect) bool {
	return pcc.Call("rectIntersectsRect", r1, r2).Bool()
}

// Overlaps checks whether a rect overlaps another.
func (r1 *rect) Overlaps(r2 Rect) bool {
	return pcc.Call("rectOverlapsRect", r1, r2).Bool()
}

// Union returns the smallest rectangle that contains the two source rectangles.
func (r1 *rect) Union(r2 Rect) Rect {
	return &rect{pcc.Call("rectUnion", r1, r2)}
}

// Intersection Returns the overlapping portion of 2 rectangles.
func (r1 *rect) Intersection(r2 Rect) Rect {
	return &rect{pcc.Call("rectIntersection", r1, r2)}
}
