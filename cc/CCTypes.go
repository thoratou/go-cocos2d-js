package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

///////////
// Color //
///////////

// The Color class
type Color interface {
	GetR() int
	GetG() int
	GetB() int
	GetA() int
	SetR(int)
	SetG(int)
	SetB(int)
	SetA(int)
	EqualsTo(Color) bool
	ToHex() string
}

type color struct{ *js.Object }

// NewColor is the constructor for Color.
func NewColor(r, g, b, a int) Color {
	return &color{pcc.Call("color", r, g, b, a)}
}

// NewColorFromHex is the constructor for Color.
func NewColorFromHex(hex string) Color {
	return &color{pcc.Call("color", hex)}
}

func (c *color) GetR() int {
	return c.Get("r").Int()
}

func (c *color) GetG() int {
	return c.Get("g").Int()
}

func (c *color) GetB() int {
	return c.Get("b").Int()
}

func (c *color) GetA() int {
	return c.Get("a").Int()
}

func (c *color) SetR(r int) {
	c.Set("r", r)
}

func (c *color) SetG(g int) {
	c.Set("g", g)
}

func (c *color) SetB(b int) {
	c.Set("b", b)
}

func (c *color) SetA(a int) {
	c.Set("a", a)
}

// EqualsTo returns true if both Color are equal. Otherwise it returns false.
func (c1 *color) EqualsTo(c2 Color) bool {
	return pcc.Call("colorEqual", c1, c2).Bool()
}

// EqualsTo returns true if both Color are equal. Otherwise it returns false.
func (c *color) ToHex() string {
	return pcc.Call("colorToHex", c).String()
}

//////////////////
// Acceleration //
//////////////////

// The Acceleration class
type Acceleration interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	GetTimestamp() float64
	SetX(float64)
	SetY(float64)
	SetZ(float64)
	SetTimestamp(float64)
}

type acceleration struct{ *js.Object }

// NewAcceleration is the constructor for Acceleration.
func NewAcceleration(x, y, z, timestamp float64) Acceleration {
	return &acceleration{pcc.Call("Acceleration", x, y, z, timestamp)}
}

func (a *acceleration) GetX() float64 {
	return a.Get("x").Float()
}

func (a *acceleration) GetY() float64 {
	return a.Get("y").Float()
}

func (a *acceleration) GetZ() float64 {
	return a.Get("z").Float()
}

func (a *acceleration) GetTimestamp() float64 {
	return a.Get("timestamp").Float()
}

func (a *acceleration) SetX(x float64) {
	a.Set("x", x)
}

func (a *acceleration) SetY(y float64) {
	a.Set("y", y)
}

func (a *acceleration) SetZ(z float64) {
	a.Set("z", z)
}

func (a *acceleration) SetTimestamp(timestamp float64) {
	a.Set("timestamp", timestamp)
}

//////////////
// Vertex2F //
//////////////

// The Vertex3F class
type Vertex2F interface {
	GetX() float64
	GetY() float64
	SetX(float64)
	SetY(float64)
}

type vertex2F struct{ *js.Object }

// NewVertex2F is the constructor for Vertex2F.
func NewVertex2F(x, y float64) Vertex2F {
	return &vertex2F{pcc.Call("Vertex2F", x, y)}
}

func (v *vertex2F) GetX() float64 {
	return v.Get("x").Float()
}

func (v *vertex2F) GetY() float64 {
	return v.Get("y").Float()
}

func (v *vertex2F) SetX(x float64) {
	v.Set("x", x)
}

func (v *vertex2F) SetY(y float64) {
	v.Set("y", y)
}

//////////////
// Vertex3F //
//////////////

// The Vertex3F class
type Vertex3F interface {
	GetX() float64
	GetY() float64
	GetZ() float64
	SetX(float64)
	SetY(float64)
	SetZ(float64)
}

type vertex3F struct{ *js.Object }

// NewVertex3F is the constructor for Vertex3F.
func NewVertex3F(x, y, z float64) Vertex3F {
	return &vertex3F{pcc.Call("Vertex3F", x, y, z)}
}

func (v *vertex3F) GetX() float64 {
	return v.Get("x").Float()
}

func (v *vertex3F) GetY() float64 {
	return v.Get("y").Float()
}

func (v *vertex3F) GetZ() float64 {
	return v.Get("z").Float()
}

func (v *vertex3F) SetX(x float64) {
	v.Set("x", x)
}

func (v *vertex3F) SetY(y float64) {
	v.Set("y", y)
}

func (v *vertex3F) SetZ(z float64) {
	v.Set("z", z)
}

///////////
// Tex2F //
///////////

// The Tex2F class
type Tex2F interface {
	GetU() float64
	GetV() float64
	SetU(float64)
	SetV(float64)
}

type tex2F struct{ *js.Object }

// NewTex2F is the constructor for Tex2F.
func NewTex2F(u, v float64) Tex2F {
	return &tex2F{pcc.Call("Tex2F", u, v)}
}

func (t *tex2F) GetU() float64 {
	return t.Get("u").Float()
}

func (t *tex2F) GetV() float64 {
	return t.Get("v").Float()
}

func (t *tex2F) SetU(u float64) {
	t.Set("u", u)
}

func (t *tex2F) SetV(v float64) {
	t.Set("v", v)
}

///////////////
// BlendFunc //
///////////////

// The Tex2F class
type BlendFunc interface {
	GetSrc() int
	GetDst() int
	SetSrc(int)
	SetDst(int)
}

type blendFunc struct{ *js.Object }

// NewBlendFunc is the constructor for BlendFunc.
func NewBlendFunc(src, dst int) BlendFunc {
	return &blendFunc{pcc.Call("BlendFunc", src, dst)}
}

func (b *blendFunc) GetSrc() int {
	return b.Get("src").Int()
}

func (b *blendFunc) GetDst() int {
	return b.Get("dst").Int()
}

func (b *blendFunc) SetSrc(src int) {
	b.Set("src", src)
}

func (b *blendFunc) SetDst(dst int) {
	b.Set("dst", dst)
}

const (
	TEXT_ALIGNMENT_LEFT            = 0
	TEXT_ALIGNMENT_CENTER          = 1
	TEXT_ALIGNMENT_RIGHT           = 2
	VERTICAL_TEXT_ALIGNMENT_TOP    = 0
	VERTICAL_TEXT_ALIGNMENT_CENTER = 1
	VERTICAL_TEXT_ALIGNMENT_BOTTOM = 2
)

////////////////////
// FontDefinition //
////////////////////
