package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// The image class
// not defined into cocos2d-js framework
type Image interface{}

type image struct{ *js.Object }

func NewImage(object *js.Object) Image {
	return &image{object}
}
