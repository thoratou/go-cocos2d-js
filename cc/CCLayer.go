package cc

///////////
// Layer //
///////////

// The Layer class
type Layer interface {
	Node
	//Init() defined as part of Node
	Bake()
	Unbake()
	IsBaked() bool
	//AddChild() defined as part of Node
}

type layer struct{ node }

// NewLayer is the constructor for Layer.
func NewLayer() Layer {
	return &layer{node{pcc.Get("Layer").New()}}
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

////////////////
// LayerColor //
////////////////

///////////////////
// LayerGradient //
///////////////////

////////////////////
// LayerMultiplex //
////////////////////
