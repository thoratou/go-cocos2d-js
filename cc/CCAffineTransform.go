package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type AffineTransform interface {
	GetA() float64
	GetB() float64
	GetC() float64
	GetD() float64
	GetTX() float64
	GetTY() float64
	SetA(float64)
	SetB(float64)
	SetC(float64)
	SetD(float64)
	SetTX(float64)
	SetTY(float64)
	PointApply(Point) Point
	SizeApply(Size) Size
	RectApply(Rect) Rect
	Translate(float64, float64) AffineTransform
	Scale(float64, float64) AffineTransform
	Rotate(float64) AffineTransform
	Concat(AffineTransform) AffineTransform
	EqualsTo(AffineTransform) bool
	Invert() AffineTransform
}

type affineTransform struct{ js.Object }

// NewAffineTransform is the constructor for AffineTransform.
func NewAffineTransform(a, b, c, d, tx, ty float64) AffineTransform {
	return &affineTransform{pcc.Call("affineTransformMake", a, b, c, d, tx, ty)}
}

// NewAffineTransformIdentity is the constructor for AffineTransform identity.
func NewAffineTransformIdentity() AffineTransform {
	return &affineTransform{pcc.Call("affineTransformIdentity")}
}

func (at *affineTransform) GetA() float64 {
	return at.Get("a").Float()
}

func (at *affineTransform) GetB() float64 {
	return at.Get("b").Float()
}

func (at *affineTransform) GetC() float64 {
	return at.Get("c").Float()
}

func (at *affineTransform) GetD() float64 {
	return at.Get("d").Float()
}

func (at *affineTransform) GetTX() float64 {
	return at.Get("tx").Float()
}

func (at *affineTransform) GetTY() float64 {
	return at.Get("ty").Float()
}

func (at *affineTransform) SetA(a float64) {
	at.Set("a", a)
}

func (at *affineTransform) SetB(b float64) {
	at.Set("b", b)
}

func (at *affineTransform) SetC(c float64) {
	at.Set("c", c)
}

func (at *affineTransform) SetD(d float64) {
	at.Set("d", d)
}

func (at *affineTransform) SetTX(tx float64) {
	at.Set("tx", tx)
}

func (at *affineTransform) SetTY(ty float64) {
	at.Set("ty", ty)
}

// PointApply applies the affine transformation on a point.
func (at *affineTransform) PointApply(p Point) Point {
	return &point{pcc.Call("pointApplyAffineTransform", p, at)}
}

// SizeApply applies the affine transformation on a size.
func (at *affineTransform) SizeApply(s Size) Size {
	return &size{pcc.Call("sizeApplyAffineTransform", s, at)}
}

// RectApply applies the affine transformation on a rect.
func (at *affineTransform) RectApply(r Rect) Rect {
	return &rect{pcc.Call("rectApplyAffineTransform ", r, at)}
}

// Translate creates a new affine transformation with a base transformation matrix and a translation based on it.
func (at *affineTransform) Translate(tx float64, ty float64) AffineTransform {
	return &affineTransform{pcc.Call("affineTransformTranslate", at, tx, ty)}
}

// Scale creates a new affine transformation with a base transformation matrix and a scale based on it.
func (at *affineTransform) Scale(sx float64, sy float64) AffineTransform {
	return &affineTransform{pcc.Call("affineTransformScale", at, sx, sy)}
}

// Rotate creates a new affine transformation with a base transformation matrix and a rotation based on it.
func (at *affineTransform) Rotate(anAngle float64) AffineTransform {
	return &affineTransform{pcc.Call("affineTransformRotate", at, anAngle)}
}

// Concat concatenates a transform matrix to another and return the result:
func (at *affineTransform) Concat(t2 AffineTransform) AffineTransform {
	return &affineTransform{pcc.Call("affineTransformConcat", at, t2)}
}

// EqualsTo returns true if an affine transform equals to another, false otherwise.
func (at *affineTransform) EqualsTo(t2 AffineTransform) bool {
	return pcc.Call("affineTransformEqualToTransform", at, t2).Bool()
}

// Invert creates the invert transform of an AffineTransform object
func (at *affineTransform) Invert() AffineTransform {
	return &affineTransform{pcc.Call("affineTransformInvert", at)}
}
