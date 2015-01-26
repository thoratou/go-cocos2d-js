// Package cc provides all high level functions and structures for interacting with Cocos native JavaScript APIs.
package scenes

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	pcc = js.Global.Get("cc")
)

type LoaderScene struct{ js.Object }

// Init: contructor of cc.LoaderScene
func (ls *LoaderScene) Init() bool {
	return ls.Call("init").Bool()
}

// OnEnter: custom ??
func (ls *LoaderScene) OnEnter() {
	ls.Call("onEnter")
}

// OnExit: custom ??
func (ls *LoaderScene) OnExit() {
	ls.Call("onExit")
}

// InitWithResources initiqlizes with resources
func (ls *LoaderScene) InitWithResources(resources []interface{}, cb func()) {
	ls.Call("initWithResources", resources, cb)
}

// Preload: cc.LoaderScene.preload can present a loaderScene with download progress.
// when all the resource are downloaded it will invoke call function.
// Not script implementation as resources shouldn't be a map but an array
// However, more in line with classical resource declaration
func (ls *LoaderScene) Preload(resources map[string]interface{}, cb func()) *LoaderScene {
	resArray := []interface{}{}
	for _, res := range resources {
		resArray = append(resArray, res)
	}
	return &LoaderScene{ls.Call("preload", resArray, cb)}
}
