package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

///////////
// Layer //
///////////

// The Layer class
type Layer interface {
	Node
	//Init() bool defined as part of Node
	Bake()
	Unbake()
	IsBaked() bool
	//AddChild(Node) defined as part of Node
	//AddChildWithOrder(Node, int) defined as part of Node
	//AddChildWithTag(Node, int) defined as part of Node
	//AddChildWithOrderAndTag(Node, int, int) defined as part of Node

}

type layer struct{ node }

// NewLayer is the constructor for Layer.
func NewLayer() Layer {
	return &layer{node{pcc.Get("Layer").New()}}
}

func (l *layer) Init() bool {
	return l.Call("init").Bool()
}

func (l *layer) Bake() {
	l.Call("bake")
}

func (l *layer) Unbake() {
	l.Call("unbake")
}

func (l *layer) IsBaked() bool {
	return l.Call("isBaked").Bool()
}

func (l *layer) AddChild(child Node) {
	l.Call("addChild", child, js.Undefined, js.Undefined)
}

func (l *layer) AddChildWithOrder(child Node, localZOrder int) {
	l.Call("addChild", child, localZOrder, js.Undefined)
}

func (l *layer) AddChildWithTag(child Node, tag int) {
	l.Call("addChild", child, js.Undefined, tag)
}

func (l *layer) AddChildWithOrderAndTag(child Node, localZOrder int, tag int) {
	l.Call("addChild", child, localZOrder, tag)
}

////////////////
// LayerColor //
////////////////

// The LayerColor class
type LayerColor interface {
	Layer
	GetBlendFunc() BlendFunc
	ChangeWidthAndHeight(float64, float64)
	ChangeWidth(float64)
	ChangeHeight(float64)
	//SetOpacityModifyRGB(bool) defined as part of Node
	//IsOpacityModifyRGB() bool defined as part of Node
	SetBlendFunc(BlendFunc)
}

type layerColor struct{ layer }

func NewLayerColor(color Color, width float64, height float64) LayerColor {
	return &layerColor{layer{node{pcc.Get("LayerColor").New(color, width, height)}}}
}

func (l *layerColor) GetBlendFunc() BlendFunc {
	return &blendFunc{l.Call("getBlendFunc")}
}

func (l *layerColor) ChangeWidthAndHeight(w float64, h float64) {
	l.Call("changeWidthAndHeight", w, h)
}

func (l *layerColor) ChangeWidth(w float64) {
	l.Call("changeWidth", w)
}

func (l *layerColor) ChangeHeight(h float64) {
	l.Call("changeHeight", h)
}

func (l *layerColor) SetOpacityModifyRGB(value bool) {
	l.Call("setOpacityModifyRGB", value)
}

func (l *layerColor) IsOpacityModifyRGB() bool {
	return l.Call("isOpacityModifyRGB").Bool()
}

func (l *layerColor) SetBlendFunc(blendFunc BlendFunc) {
	l.Call("setBlendFunc", blendFunc)
}

///////////////////
// LayerGradient //
///////////////////

type LayerGradientColorStop struct {
	P     float64 `js:"p"`
	Color Color   `js:"color"`
}

// The LayerGradient class
type LayerGradient interface {
	LayerColor
	//SetContentSize(Size) defined as part of Node
	GetStartColor() Color
	SetStartColor(Color)
	GetEndColor() Color
	SetEndColor(Color)
	SetStartOpacity(int)
	GetStartOpacity() int
	SetEndOpacity(int)
	GetEndOpacity() int
	SetVector(Point)
	GetVector() Point
	IsCompressedInterpolation() bool
	SetCompressedInterpolation(bool)
	GetColorStops() []LayerGradientColorStop
	SetColorStops([]LayerGradientColorStop)
}

type layerGradient struct{ layerColor }

func NewLayerGradient(startColor Color, endColor Color) LayerGradient {
	return &layerGradient{layerColor{layer{node{pcc.Get("LayerGradient").New(startColor, endColor)}}}}
}

func NewLayerGradientWithVector(startColor Color, endColor Color, v Point) LayerGradient {
	return &layerGradient{layerColor{layer{node{pcc.Get("LayerGradient").New(startColor, endColor, v)}}}}
}

func NewLayerGradientWithVectorAndStops(startColor Color, endColor Color, v Point, stops LayerGradientColorStop) LayerGradient {
	return &layerGradient{layerColor{layer{node{pcc.Get("LayerGradient").New(startColor, endColor, v, stops)}}}}
}

func (l *layerGradient) SetContentSize(size Size) {
	l.Call("setContentSize", size)
}

func (l *layerGradient) GetStartColor() Color {
	return &color{l.Call("getStartColor")}
}

func (l *layerGradient) SetStartColor(color Color) {
	l.Call("setStartColor", color)
}

func (l *layerGradient) GetEndColor() Color {
	return &color{l.Call("getEndColor")}
}

func (l *layerGradient) SetEndColor(color Color) {
	l.Call("setEndColor", color)
}

func (l *layerGradient) SetStartOpacity(o int) {
	l.Call("setStartOpacity", o)
}

func (l *layerGradient) GetStartOpacity() int {
	return l.Call("getStartOpacity").Int()
}

func (l *layerGradient) SetEndOpacity(o int) {
	l.Call("setEndOpacity", o)
}

func (l *layerGradient) GetEndOpacity() int {
	return l.Call("getEndOpacity").Int()
}

func (l *layerGradient) SetVector(v Point) {
	l.Call("setVector", v)
}

func (l *layerGradient) GetVector() Point {
	return &point{l.Call("getVector")}
}

func (l *layerGradient) IsCompressedInterpolation() bool {
	return l.Call("isCompressedInterpolation").Bool()
}

func (l *layerGradient) SetCompressedInterpolation(compress bool) {
	l.Call("setCompressedInterpolation", compress)
}

func (l *layerGradient) GetColorStops() []LayerGradientColorStop {
	stops := l.Call("getColorStops")
	length := stops.Length()
	out := make([]LayerGradientColorStop, length, length)
	for i := 0; i < length; i++ {
		stop := stops.Index(i)
		out[i] = LayerGradientColorStop{P: stop.Get("p").Float(), Color: &color{stop.Get("color")}}
	}
	return out
}

func (l *layerGradient) SetColorStops(colorStops []LayerGradientColorStop) {
	l.Call("setColorStops", colorStops)
}

////////////////////
// LayerMultiplex //
////////////////////
