package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	DENSITYDPI_DEVICE = "device-dpi"
	DENSITYDPI_HIGH   = "high-dpi"
	DENSITYDPI_MEDIUM = "medium-dpi"
	DENSITYDPI_LOW    = "low-dpi"
)

/////////////
// eGLView //
/////////////

type eGLView struct{ js.Object }

// SetTargetDensityDPI sets view's target-densitydpi for android mobile browser.
// It can be set to:
//  // cc.DENSITYDPI_DEVICE, value is "device-dpi"
//  // cc.DENSITYDPI_HIGH, value is "high-dpi"  (default value)
//  // cc.DENSITYDPI_MEDIUM, value is "medium-dpi" (browser's default value)
//  // cc.DENSITYDPI_LOW, value is "low-dpi"
//  // Custom value, e.g: "480"
func (v *eGLView) SetTargetDensityDPI(densityDPI string) {
	v.Call("setTargetDensityDPI", densityDPI)
}

// GetTargetDensityDPI returns the current target-densitydpi value of cc.view.
func (v *eGLView) GetTargetDensityDPI() string {
	return v.Call("getTargetDensityDPI").String()
}

// ResizeWithBrowserSize sets whether resize canvas automatically when browser's size changed.
// Useful only on web.
func (v *eGLView) ResizeWithBrowserSize(enabled bool) {
	v.Call("resizeWithBrowserSize", enabled)
}

// SetResizeCallback sets the callback function for cc.view's resize action,
// this callback will be invoked before applying resolution policy,
// so you can do any additional modifications within the callback.
// Useful only on web.
func (v *eGLView) SetResizeCallback(cb func()) {
	v.Call("setResizeCallback", cb)
}

// Initialize, don't know if it could be useful for the binding ??
func (v *eGLView) Initialize() {
	v.Call("initialize")
}

// AdjustViewPort sets whether the engine modify the "viewport" meta in your web page.
// It's enabled by default, we strongly suggest you not to disable it.
// And even when it's enabled, you can still set your own "viewport" meta, it won't be overridden.
// Only useful on web.
func (v *eGLView) AdjustViewPort(enabled bool) {
	v.Call("adjustViewPort", enabled)
}

// EnableRetina: Retina support is enabled by default for Apple device but disabled for other devices,
// it takes effect only when you called setDesignResolutionPolicy
// Only useful on web.
func (v *eGLView) EnableRetina(enabled bool) {
	v.Call("enableRetina", enabled)
}

// IsRetinaEnabled checks whether retina display is enabled.
// Only useful on web.
func (v *eGLView) IsRetinaEnabled() bool {
	return v.Call("isRetinaEnabled").Bool()
}

// EnableAutoFullScreen: If enabled, the application will try automatically to enter full screen mode on mobile devices.
// You can pass true as parameter to enable it and disable it by passing false.
// Only useful on web.
func (v *eGLView) EnableAutoFullScreen(enabled bool) {
	v.Call("enableAutoFullScreen", enabled)
}

// EnableAutoFullScreen checks whether auto full screen is enabled.<br/>
// Only useful on web
func (v *eGLView) IsAutoFullScreenEnabled() bool {
	return v.Call("isAutoFullScreenEnabled").Bool()
}

// End forces destroying EGL view, subclass must implement this method.
func (v *eGLView) End() {
	v.Call("end")
}

// IsOpenGLReady gets whether render system is ready(no matter opengl or canvas),
// this name is for the compatibility with cocos2d-x, subclass must implement this method.
func (v *eGLView) IsOpenGLReady() bool {
	return v.Call("isOpenGLReady").Bool()
}

// SetFrameZoomFactor sets zoom factor for frame. This method is for debugging big resolution (e.g.new ipad) app on desktop.
func (v *eGLView) SetFrameZoomFactor(zoomFactor float64) {
	v.Call("setFrameZoomFactor", zoomFactor)
}

// SwapBuffers exchanges the front and back buffers, subclass must implement this method.
func (v *eGLView) SwapBuffers() {
	v.Call("swapBuffers")
}

// SetIMEKeyboardState opens or closes IME keyboard , subclass must implement this method.
func (v *eGLView) SetIMEKeyboardState(isOpen bool) {
	v.Call("setIMEKeyboardState", isOpen)
}

// SetContentTranslateLeftTop sets the resolution translate on EGLView.
func (v *eGLView) SetContentTranslateLeftTop(offsetLeft int, offsetTop int) {
	v.Call("setContentTranslateLeftTop", offsetLeft, offsetTop)
}

// SetContentTranslateLeftTop sets the resolution translate on EGLView.
func (v *eGLView) GetContentTranslateLeftTop() Size {
	leftTop := v.Call("getContentTranslateLeftTop")
	left := leftTop.Get("left")
	top := leftTop.Get("top")
	if left != nil && top != nil {
		return NewSize(left.Int(), top.Int())
	}
	return &size{leftTop}
}

// GetFrameSize returns the frame size of the view.
// On native platforms, it returns the screen size since the view is a fullscreen view.
// On web, it returns the size of the canvas's outer DOM element.
func (v *eGLView) GetFrameSize() Size {
	return &size{v.Call("getFrameSize")}
}

// SetFrameSize sets, on native, the frame size of view.
// On web, it sets the size of the canvas's outer DOM element.
func (v *eGLView) SetFrameSize(width int, height int) {
	v.Call("setFrameSize", width, height)
}

// CenterWindow empty function
func (v *eGLView) CenterWindow() {
	v.Call("centerWindow")
}

// GetVisibleSize returns the visible area size of the view port.
func (v *eGLView) GetVisibleSize() Size {
	return &size{v.Call("getVisibleSize")}
}

// GetVisibleOrigin returns the visible origin of the view port.
func (v *eGLView) GetVisibleOrigin() Point {
	return &point{v.Call("getVisibleOrigin")}
}

// CanSetContentScaleFactor returns whether developer can set content's scale factor.
func (v *eGLView) CanSetContentScaleFactor() bool {
	return v.Call("canSetContentScaleFactor").Bool()
}

// GetResolutionPolicy returns the current resolution policy.
func (v *eGLView) GetResolutionPolicy() *ResolutionPolicy {
	return &ResolutionPolicy{v.Call("getResolutionPolicy")}
}

// SetResolutionPolicy sets the current resolution policy.
// resolutionPolicy could be *ResolutionPolicy or int
func (v *eGLView) SetResolutionPolicy(resolutionPolicy interface{}) {
	v.Call("setResolutionPolicy", resolutionPolicy)
}

// SetDesignResolutionSize sets the resolution policy with designed view size in points.
// resolutionPolicy could be *ResolutionPolicy or int
func (v *eGLView) SetDesignResolutionSize(width int, height int, resolutionPolicy interface{}) {
	v.Call("setDesignResolutionSize", width, height, resolutionPolicy)
}

// GetDesignResolutionSize returns the designed size for the view.
// Default resolution size is the same as 'getFrameSize'.
func (v *eGLView) GetDesignResolutionSize() Size {
	return &size{v.Call("getDesignResolutionSize")}
}

// SetViewPortInPoints sets view port rectangle with points.
func (v *eGLView) SetViewPortInPoints(x int, y int, w int, h int) {
	v.Call("setViewPortInPoints", x, y, w, h)
}

// SetScissorInPoints sets Scissor rectangle with points.
func (v *eGLView) SetScissorInPoints(x int, y int, w int, h int) {
	v.Call("setScissorInPoints", x, y, w, h)
}

// IsScissorEnabled returns whether GL_SCISSOR_TEST is enable.
func (v *eGLView) IsScissorEnabled() bool {
	return v.Call("isScissorEnabled").Bool()
}

// GetScissorRect returns the current scissor rectangle.
func (v *eGLView) GetScissorRect() Rect {
	return &rect{v.Call("getScissorRect")}
}

// SetViewName sets the name of the view.
func (v *eGLView) SetViewName(viewName string) {
	v.Call("setViewName", viewName)
}

// GetViewName returns the name of the view.
func (v *eGLView) GetViewName() string {
	return v.Call("getViewName").String()
}

// GetViewPortRect returns the view port rectangle.
func (v *eGLView) GetViewPortRect() Rect {
	return &rect{v.Call("getViewPortRect")}
}

// GetScaleX returns scale factor of the horizontal direction (X axis).
func (v *eGLView) GetScaleX() float64 {
	return v.Call("getScaleX").Float()
}

// GetScaleY returns scale factor of the vertical direction (Y axis).
func (v *eGLView) GetScaleY() float64 {
	return v.Call("getScaleY").Float()
}

// GetDevicePixelRatio returns device pixel ratio for retina display.
func (v *eGLView) GetDevicePixelRatio() float64 {
	return v.Call("getDevicePixelRatio").Float()
}

// ConvertToLocationInView the real location in view for a translation based on a related position
// relatedPos contains the related position object including "left", "top", "width", "height" informations
func (v *eGLView) ConvertToLocationInView(tx int, ty int, relatedPos js.Object) Point {
	return &point{v.Call("convertToLocationInView", tx, ty, relatedPos)}
}

///////////////////////
// ContainerStrategy //
///////////////////////

type ContainerStrategy interface {
	js.Object
	PreApply(v *eGLView)
	Apply(v *eGLView, designedResolution Size)
	PostApply(v *eGLView)
}

type containerStrategy struct{ js.Object }

// PreApply performs manipulation before appling the strategy
func (cs *containerStrategy) PreApply(v *eGLView) {
	cs.Call("preApply", v)
}

// Apply applies the strategy
func (cs *containerStrategy) Apply(v *eGLView, designedResolution Size) {
	cs.Call("apply", v, designedResolution)
}

// PostApply performs manipulation after appling the strategy
func (cs *containerStrategy) PostApply(v *eGLView) {
	cs.Call("postApply", v)
}

/*
var (
	CS_EQUAL_TO_FRAME      = &containerStrategy{pcc.Get("ContainerStrategy").Get("EQUAL_TO_FRAME")}
	CS_PROPORTION_TO_FRAME = &containerStrategy{pcc.Get("ContainerStrategy").Get("PROPORTION_TO_FRAME")}
	//CS_EQUAL_TO_WINDOW      = &containerStrategy{pcc.Get("ContainerStrategy").Get("EQUAL_TO_WINDOW")}           NOT STABLE
	//CS_PROPORTION_TO_WINDOW = &containerStrategy{pcc.Get("ContainerStrategy").Get("PROPORTION_TO_WINDOW")}    NOT STABLE
	CS_ORIGINAL_CONTAINER = &containerStrategy{pcc.Get("ContainerStrategy").Get("ORIGINAL_CONTAINER")}
)
*/

/////////////////////
// ContentStrategy //
/////////////////////

type ContentStrategy interface {
	js.Object
	PreApply(v *eGLView)
	Apply(v *eGLView, designedResolution Size) js.Object
	PostApply(v *eGLView)
}

type contentStrategy struct{ js.Object }

// PreApply performs manipulation before appling the strategy
func (cs *contentStrategy) PreApply(v *eGLView) {
	cs.Call("preApply", v)
}

// Apply applies the strategy
// The return value is {scale: [scaleX, scaleY], viewport: {cc.Rect}},
// The target view can then apply these value to itself, it's preferred not to modify directly its private variables
func (cs *contentStrategy) Apply(v *eGLView, designedResolution Size) js.Object {
	return cs.Call("apply", v, designedResolution)
}

// PostApply performs manipulation after appling the strategy
func (cs *contentStrategy) PostApply(v *eGLView) {
	cs.Call("postApply", v)
}

/*
var (
	CS_EXACT_FIT    = &containerStrategy{pcc.Get("ContentStrategy").Get("EXACT_FIT")}
	CS_SHOW_ALL     = &containerStrategy{pcc.Get("ContentStrategy").Get("SHOW_ALL")}
	CS_NO_BORDER    = &containerStrategy{pcc.Get("ContentStrategy").Get("NO_BORDER")}
	CS_FIXED_HEIGHT = &containerStrategy{pcc.Get("ContentStrategy").Get("FIXED_HEIGHT")}
	CS_FIXED_WIDTH  = &containerStrategy{pcc.Get("ContentStrategy").Get("FIXED_WIDTH")}
)
*/

//////////////////////
// ResolutionPolicy //
//////////////////////

type ResolutionPolicy struct{ js.Object }

// NewResolutionPolicy creates a new ResolutionPolicy
func NewResolutionPolicy(cs1 ContainerStrategy, cs2 ContentStrategy) *ResolutionPolicy {
	return &ResolutionPolicy{pcc.Get("ResolutionPolicy").Call("ctor", cs1, cs2)}
}

// PreApply performs manipulation before appling the strategy
func (rp *ResolutionPolicy) PreApply(v *eGLView) {
	rp.Call("preApply", v)
}

// Apply applies the strategy
// The return value is {scale: [scaleX, scaleY], viewport: {cc.Rect}},
// The target view can then apply these value to itself, it's preferred not to modify directly its private variables
func (rp *ResolutionPolicy) Apply(v *eGLView, designedResolution Size) js.Object {
	return rp.Call("apply", v, designedResolution)
}

// PostApply performs manipulation after appling the strategy
func (rp *ResolutionPolicy) PostApply(v *eGLView) {
	rp.Call("postApply", v)
}

// SetContainerStrategy setups the container's scale strategy
func (rp *ResolutionPolicy) SetContainerStrategy(containerStg ContainerStrategy) {
	rp.Call("setContainerStrategy", containerStg)
}

// SetContentStrategy setups the content's scale strategy
func (rp *ResolutionPolicy) SetContentStrategy(contentStg ContentStrategy) {
	rp.Call("setContentStrategy", contentStg)
}

const (
	EXACT_FIT    = 0
	NO_BORDER    = 1
	SHOW_ALL     = 2
	FIXED_HEIGHT = 3
	FIXED_WIDTH  = 4
	UNKNOWN      = 5
)
