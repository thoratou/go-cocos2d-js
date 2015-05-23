package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	ALIGN_CENTER       = 0x33
	ALIGN_TOP          = 0x13
	ALIGN_TOP_RIGHT    = 0x12
	ALIGN_RIGHT        = 0x32
	ALIGN_BOTTOM_RIGHT = 0x22
	ALIGN_BOTTOM       = 0x23
	ALIGN_BOTTOM_LEFT  = 0x21
	ALIGN_LEFT         = 0x31
	ALIGN_TOP_LEFT     = 0x11
)

// The touch event class
type Texture2D interface {
	GetPixelsWide() int
	GetPixelsHigh() int
	GetContentSize() Size
	GetContentSizeInPixels() Size
	InitWithElement(*js.Object)
	GetHtmlElementObj() *js.Object
	IsLoaded() bool
	HandleLoadedTexture()
	Description() string
	//Only supported in WEBGL, will see later if supported by the binding
	//InitWithData()
	//InitWithImage()
	//InitWithString()
	ReleaseTexture()
	//Only supported in WEBGL, will see later if supported by the binding
	//GetName()
	//GetMaxS()
	//SetMaxS()
	//GetMaxT()
	//SetMaxT()
	//...
	AddLoadedEventListener(func(...interface{}), Node)
	RemoveLoadedEventListener(Node)
}

type texture2D struct{ *js.Object }

// NewTexture2D is the constructor for Texture2D.
func NewTexture2D() Texture2D {
	return &texture2D{pcc.Get("Texture2D").New()}
}

func (t *texture2D) GetPixelsWide() int {
	return t.Call("getPixelsWide").Int()
}

func (t *texture2D) GetPixelsHigh() int {
	return t.Call("getPixelsHigh").Int()
}

func (t *texture2D) GetContentSize() Size {
	return &size{t.Call("getContentSize")}
}

func (t *texture2D) GetContentSizeInPixels() Size {
	return &size{t.Call("getContentSizeInPixels")}
}

func (t *texture2D) InitWithElement(element *js.Object) {
	t.Call("initWithElement", element)
}

func (t *texture2D) GetHtmlElementObj() *js.Object {
	return t.Call("getHtmlElementObj")
}

func (t *texture2D) IsLoaded() bool {
	return t.Call("isLoaded").Bool()
}

func (t *texture2D) HandleLoadedTexture() {
	t.Call("handleLoadedTexture")
}

func (t *texture2D) Description() string {
	return t.Call("description").String()
}

func (t *texture2D) ReleaseTexture() {
	t.Call("releaseTexture")
}

func (t *texture2D) AddLoadedEventListener(callback func(...interface{}), target Node) {
	t.Call("addLoadedEventListener", callback, target)
}

func (t *texture2D) RemoveLoadedEventListener(target Node) {
	t.Call("removeLoadedEventListener", target)
}
