package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

///////////
// Color //
///////////

// The Color class
type Color interface {
	js.Object
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

type color struct{ js.Object }

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

//////////////
// Vertex2F //
//////////////

//////////////
// Vertex3F //
//////////////

///////////
// Tex2F //
///////////

///////////////
// BlendFunc //
///////////////

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
