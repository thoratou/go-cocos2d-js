// Package cc provides all high level functions and structures for interacting with Cocos native JavaScript APIs.
package cc

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/thoratou/go-cocos2d-js/cc/core/platform/view"
)

var (
	pcc = js.Global.Get("cc")
)

func View() *view.EGLView {
	return &view.EGLView{pcc.Get("view")}
}

//func NewElement(name string) js.Object {
//	return pcc.Call("newElement", name)
//}

// Each iterates over an object or an array, executing a function for each matched element.
// obj could be a js.Object or []js.Object.
func Each(obj interface{}, iterator func(...interface{}), context js.Object) {
	pcc.Call("each", obj, iterator, context)
}

//Extend copies all of the properties in source objects to target object and returns the target object.
func Extend(target js.Object) js.Object {
	return pcc.Call("extend", target)
}

//IsFunction checks the obj whether is function or not.
func IsFunction(obj interface{}) bool {
	return pcc.Call("isFunction", obj).Bool()
}

//IsNumber checks the obj whether is number or not.
func IsNumber(obj interface{}) bool {
	return pcc.Call("isNumber", obj).Bool()
}

//IsString checks the obj whether is string or not.
func IsString(obj interface{}) bool {
	return pcc.Call("isString", obj).Bool()
}

//IsArray checks the obj whether is array or not.
func IsArray(obj interface{}) bool {
	return pcc.Call("isArray", obj).Bool()
}

//IsUndefined checks the obj whether is undefined or not.
func IsUndefined(obj interface{}) bool {
	return pcc.Call("isUndefined", obj).Bool()
}

//IsObject checks the obj whether is object or not.
func IsObject(obj interface{}) bool {
	return pcc.Call("isObject", obj).Bool()
}

//IsCrossOrigin checks the url whether cross origin.
func IsCrossOrigin(url string) bool {
	return pcc.Call("isCrossOrigin", url).Bool()
}

///////////
// ASYNC //
///////////

//AsyncPool function is a helper of Async.
//first paremeter should be a js.Object or []js.Object.
type AsyncPool func(interface{}, int, func(...interface{}), func(...interface{}), js.Object)

type async struct{ js.Object }

var (
	//Async is the global instance for asynchrone processing helper.
	Async = async{pcc.Get("async")}
)

// Series does tasks series.
// tasks could be a js.Object or []js.Object.
func (a *async) Series(tasks interface{}, cb func(...interface{}), target js.Object) AsyncPool {
	return a.Call("series", tasks, cb, target).Interface().(AsyncPool)
}

// Parallel does tasks parallel.
// tasks could be a js.Object or []js.Object.
func (a *async) Parallel(tasks interface{}, cb func(...interface{}), target js.Object) AsyncPool {
	return a.Call("parallel", tasks, cb, target).Interface().(AsyncPool)
}

// Waterfall does tasks waterfall.
// tasks could be a js.Object or []js.Object.
func (a *async) Waterfall(tasks interface{}, cb func(...interface{}), target js.Object) AsyncPool {
	return a.Call("waterfall", tasks, cb, target).Interface().(AsyncPool)
}

// Map does tasks by iterator.
// tasks could be a js.Object or []js.Object.
// iterator could be a js.Object or func(...interface{}).
func (a *async) Map(tasks interface{}, iterator interface{}, cb func(...interface{}), target js.Object) AsyncPool {
	return a.Call("map", tasks, iterator, cb, target).Interface().(AsyncPool)
}

// MapLimit does tasks by iterator limit.
// tasks could be a js.Object or []js.Object.
func (a *async) MapLimit(tasks interface{}, limit int, iterator func(...interface{}), cb func(...interface{}), target js.Object) AsyncPool {
	return a.Call("mapLimit", tasks, limit, iterator, cb, target).Interface().(AsyncPool)
}

//////////
// PATH //
//////////

////////////
// LOADER //
////////////

///////////////////
// WINDOW EVENTS //
///////////////////

/////////
// LOG //
/////////

func cclog(jsMethod string, args []interface{}) {
	switch len(args) {
	case 0:
		return
	case 1:
		for _, toLog := range args {
			pcc.Call(jsMethod, toLog)
		}
		return
	default:
		pcc.Call(jsMethod, args)
		return
	}
}

//Log logs to console/page depending on the user project.
func Log(args ...interface{}) {
	cclog("log", args)
}

//Log logs warning to console/page depending on the user project.
func Warn(args ...interface{}) {
	cclog("warn", args)
}

//Log logs error to console/page depending on the user project.
func Error(args ...interface{}) {
	cclog("error", args)
}

//Assert tests the assertion, stop if not verified and logs error to console/page depending on the user project.
func Assert(args ...interface{}) {
	cclog("assert", args)
}

//////////////
// SYS INIT //
//////////////

/////////////
// CC GAME //
/////////////

const (
	ORIENTATION_PORTRAIT             = 0
	ORIENTATION_PORTRAIT_UPSIDE_DOWN = 1
	ORIENTATION_LANDSCAPE_LEFT       = 2
	ORIENTATION_LANDSCAPE_RIGHT      = 3
)

type game struct{ js.Object }

var (
	//Game is an object to boot the game.
	Game = game{pcc.Get("game")}
)

const (
	DEBUG_MODE_NONE               = 0
	DEBUG_MODE_INFO               = 1
	DEBUG_MODE_WARN               = 2
	DEBUG_MODE_ERROR              = 3
	DEBUG_MODE_INFO_FOR_WEB_PAGE  = 4
	DEBUG_MODE_WARN_FOR_WEB_PAGE  = 5
	DEBUG_MODE_ERROR_FOR_WEB_PAGE = 6
	EVENT_HIDE                    = "game_on_hide"
	EVENT_SHOW                    = "game_on_show"
)

// Onstart is the callback when the scripts of engine have been load.
func (g *game) OnStart(cb func()) {
	g.Set("onStart", cb)
}

// Onstop is the callback when the game .
func (g *game) OnStop(cb func()) {
	g.Set("onStop", cb)
}

// Run runs the game.
func (g *game) Run() {
	g.Call("run")
}

// Run runs the game.
func (g *game) RunWithId(id int) {
	g.Call("run", id)
}
