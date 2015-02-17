package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type loaderScene struct{ *js.Object }

// Init: contructor of cc.LoaderScene
func (ls *loaderScene) Init() bool {
	return ls.Call("init").Bool()
}

// OnEnter: custom ??
func (ls *loaderScene) OnEnter() {
	ls.Call("onEnter")
}

// OnExit: custom ??
func (ls *loaderScene) OnExit() {
	ls.Call("onExit")
}

// InitWithResources initiqlizes with resources
func (ls *loaderScene) InitWithResources(resources []interface{}, cb func()) {
	ls.Call("initWithResources", resources, cb)
}

// Preload: cc.LoaderScene.preload can present a loaderScene with download progress.
// when all the resource are downloaded it will invoke call function.
// Not script implementation as resources shouldn't be a map but an array
// However, more in line with classical resource declaration
func (ls *loaderScene) Preload(resources map[string]interface{}, cb func()) *loaderScene {
	resArray := []interface{}{}
	for _, res := range resources {
		resArray = append(resArray, res)
	}
	return &loaderScene{ls.Call("preload", resArray, cb)}
}
