package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// The Scene class
type Scene interface {
	Node
	Super() js.Object
}

type scene struct{ node }

// NewScene is the constructor for Scene.
func NewScene() Scene {
	return &scene{node{pcc.Get("Scene").New()}}
}

func (s *scene) Super() js.Object {
	return func(sc Scene) js.Object {
		return sc.Get("node")
	}(s)
}
