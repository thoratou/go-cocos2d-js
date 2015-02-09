package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type director struct{ js.Object }

func (d *director) Init() {
	d.Call("init")
}

func (d *director) CalculateDeltaTime() {
	d.Call("calculateDeltaTime")
}

func (d *director) ConvertToGL(p Point) Point {
	return &point{d.Call("convertToGL", p)}
}

func (d *director) ConvertToUI(p Point) Point {
	return &point{d.Call("convertToUI", p)}
}

func (d *director) DrawScene() {
	d.Call("drawScene")
}

func (d *director) End() {
	d.Call("end")
}

func (d *director) GetContentScaleFactor() float64 {
	return d.Call("getContentScaleFactor").Float()
}

func (d *director) GetNotificationNode() Node {
	return &node{d.Call("getNotificationNode")}
}

func (d *director) GetWinSize() Size {
	return &size{d.Call("getWinSize")}
}

func (d *director) GetWinSizeInPixels() Size {
	return &size{d.Call("getWinSizeInPixels")}
}

func (d *director) GetVisibleSize() Size {
	return &size{d.Call("getVisibleSize")}
}

func (d *director) GetVisibleOrigin() Point {
	return &point{d.Call("getVisibleOrigin")}
}

func (d *director) GetZEye() float64 {
	return d.Call("getZEye").Float()
}

func (d *director) Pause() {
	d.Call("pause")
}

func (d *director) PopScene() {
	d.Call("popScene")
}

func (d *director) PurgeCachedData() {
	d.Call("purgeCachedData")
}

func (d *director) PurgeDirector() {
	d.Call("purgeDirector")
}

func (d *director) PushScene(s Scene) {
	d.Call("pushScene", s)
}

func (d *director) RunScene(s Scene) {
	d.Call("runScene", s)
}

func (d *director) Resume() {
	d.Call("resume")
}

func (d *director) SetContentScaleFactor(scaleFactor float64) {
	d.Call("setContentScaleFactor", scaleFactor)
}

func (d *director) SetDepthTest(on bool) {
	d.Call("setDepthTest", on)
}

func (d *director) SetDefaultValues() {
	d.Call("setDefaultValues")
}

func (d *director) SetNextDeltaTimeZero(nextDeltaTimeZero bool) {
	d.Call("setNextDeltaTimeZero", nextDeltaTimeZero)
}

func (d *director) SetNextScene() {
	d.Call("setNextScene")
}

func (d *director) SetNotificationNode(n Node) {
	d.Call("setNotificationNode", n)
}

//TODO DirectorDelegate

//func (d *director) GetDelegate() DirectorDelegate {
//	return &directorDelegate{d.Call("getDelegate")}
//}

//func (d *director) SetDelegate(d DirectorDelegate) {
//	d.Call("setDelegate", d)
//}

func (d *director) SetOpenGLView(openGLView *eGLView) {
	d.Call("setOpenGLView", openGLView)
}

func (d *director) SetProjection(projection float64) {
	d.Call("setProjection", projection)
}

//TODO Viewport ??

//func (d *director) SetViewport() Viewport {
//	return &viewport{d.Call("setViewport")}
//}

func (d *director) GetOpenGLView() *eGLView {
	return &eGLView{d.Call("getOpenGLView")}
}

func (d *director) GetProjection() float64 {
	return d.Call("getProjection").Float()
}

func (d *director) SetAlphaBlending(on bool) {
	d.Call("setAlphaBlending", on)
}

func (d *director) IsSendCleanupToScene() bool {
	return d.Call("isSendCleanupToScene").Bool()
}

func (d *director) GetRunningScene() Scene {
	return &scene{node{d.Call("getRunningScene")}}
}

func (d *director) GetAnimationInterval() float64 {
	return d.Call("getAnimationInterval").Float()
}

func (d *director) IsDisplayStats() bool {
	return d.Call("isDisplayStats").Bool()
}

func (d *director) SetDisplayStats(displayStats bool) {
	d.Call("setDisplayStats", displayStats)
}

func (d *director) GetSecondsPerFrame() float64 {
	return d.Call("getSecondsPerFrame").Float()
}

func (d *director) IsNextDeltaTimeZero() bool {
	return d.Call("isNextDeltaTimeZero").Bool()
}

func (d *director) IsPaused() bool {
	return d.Call("isPaused").Bool()
}

func (d *director) GetTotalFrames() int {
	return d.Call("getTotalFrames").Int()
}

func (d *director) PopToRootScene() {
	d.Call("popToRootScene")
}

func (d *director) PopToSceneStackLevel(level int) {
	d.Call("popToSceneStackLevel", level)
}

func (d *director) GetScheduler() Scheduler {
	return &scheduler{d.Call("getScheduler")}
}

func (d *director) SetScheduler(s Scheduler) {
	d.Call("setScheduler", s)
}

func (d *director) GetActionManager() ActionManager {
	return &actionManager{d.Call("getActionManager")}
}

func (d *director) SetActionManager(am ActionManager) {
	d.Call("setActionManager", am)
}

func (d *director) GetDeltaTime() float64 {
	return d.Call("getDeltaTime").Float()
}

const (
	EVENT_PROJECTION_CHANGED = "director_projection_changed"
	EVENT_AFTER_DRAW         = "director_after_draw"
	EVENT_AFTER_VISIT        = "director_after_visit"
	EVENT_AFTER_UPDATE       = "director_after_update"
)

const (
	PROJECTION_2D      = 0
	PROJECTION_3D      = 1
	PROJECTION_CUSTOM  = 3
	PROJECTION_DEFAULT = PROJECTION_3D
)
