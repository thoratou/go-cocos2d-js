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

///////////////////
// LayerGradient //
///////////////////

////////////////////
// LayerMultiplex //
////////////////////
