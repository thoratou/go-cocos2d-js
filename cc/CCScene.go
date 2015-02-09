package cc

// The Scene class
type Scene interface {
	Node
}

type scene struct{ node }

// NewScene is the constructor for Scene.
func NewScene() Scene {
	return &scene{node{pcc.Get("Scene").New()}}
}
