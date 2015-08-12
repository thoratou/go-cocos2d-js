package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	TMX_PROPERTY_NONE        = 0
	TMX_PROPERTY_MAP         = 1
	TMX_PROPERTY_LAYER       = 2
	TMX_PROPERTY_OBJECTGROUP = 3
	TMX_PROPERTY_OBJECT      = 4
	TMX_PROPERTY_TILE        = 5
	TMX_TILE_HORIZONTAL_FLAG = 0x80000000
	TMX_TILE_VERTICAL_FLAG   = 0x40000000
	TMX_TILE_DIAGONAL_FLAG   = 0x20000000
	TMX_TILE_FLIPPED_ALL     = TMX_TILE_HORIZONTAL_FLAG | TMX_TILE_VERTICAL_FLAG | TMX_TILE_DIAGONAL_FLAG
	TMX_TILE_FLIPPED_MASK    = 0x0fffffff
)

//////////////////
// TMXLayerInfo //
//////////////////

type TMXLayerInfo interface {
	GetProperties() []*js.Object
	SetProperties([]*js.Object)
}

type tmxLayerInfo struct{ *js.Object }

func NewTMXLayerInfo() TMXLayerInfo {
	return &tmxLayerInfo{pcc.Get("TMXLayerInfo").New()}
}

func (t *tmxLayerInfo) GetProperties() []*js.Object {
	properties := t.Call("getProperties")
	length := properties.Length()
	out := make([]*js.Object, length, length)
	for i := 0; i < length; i++ {
		out[i] = properties.Index(i)
	}
	return out
}

func (t *tmxLayerInfo) SetProperties(properties []*js.Object) {
	t.Call("setProperties", properties)
}

/////////////////////
// TMXTilesetInfo  //
/////////////////////

type TMXTilesetInfo interface {
	RectForGID(int) Rect
}

type tmxTilesetInfo struct{ *js.Object }

func NewTMXTilesetInfo() TMXTilesetInfo {
	return &tmxTilesetInfo{pcc.Get("TMXTilesetInfo").New()}
}

func (t *tmxTilesetInfo) RectForGID(gid int) Rect {
	return &rect{t.Call("rectForGID", gid)}
}

////////////////
// TMXMapInfo //
////////////////

type TMXMapInfo interface {
	SAXParser
	GetOrientation() int
	SetOrientation(int)
	GetMapSize() Size
	SetMapSize(Size)
	GetTileSize() Size
	SetTileSize(Size)
	GetLayers() []TMXLayerInfo
	SetLayers(TMXLayerInfo)
	GetTilesets() []TMXTilesetInfo
	SetTilesets(TMXTilesetInfo)
	//TODO: TMXObjectGroup
	//GetObjectGroups() []TMXObjectGroup
	//SetObjectGroups(TMXObjectGroup)
	GetParentElement() *js.Object
	SetParentElement(*js.Object)
	GetParentGID() int
	SetParentGID(int)
	GetLayerAttribs() *js.Object
	SetLayerAttribs(*js.Object)
	GetStoringCharacters() bool
	SetStoringCharacters(bool)
	GetProperties() []*js.Object
	SetProperties([]*js.Object)
	ParseXMLFile(string, bool) *js.Object
	ParseXMLString(string) bool
	GetTileProperties() []*js.Object
	SetTileProperties(*js.Object)
	GetCurrentString() string
	SetCurrentString(string)
	GetTMXFileName() string
	SetTMXFileName(string)
}

type tmxMapInfo struct{ saxParser }

func NewTMXMapInfoWithFile(tmxFile string) TMXMapInfo {
	return &tmxMapInfo{saxParser{pcc.Get("TMXMapInfo").New(tmxFile)}}
}

func NewTMXMapInfoWithString(contentString string, resourcePath string) TMXMapInfo {
	return &tmxMapInfo{saxParser{pcc.Get("TMXMapInfo").New(contentString, resourcePath)}}
}

func NewTMXMapInfo(args ...interface{}) TMXMapInfo {
	return &tmxMapInfo{saxParser{pcc.Get("TMXMapInfo").New(args...)}}
}

func (t *tmxMapInfo) GetOrientation() int {
	return t.Call("getOrientation").Int()
}

func (t *tmxMapInfo) SetOrientation(value int) {
	t.Call("setOrientation", value)
}

func (t *tmxMapInfo) GetMapSize() Size {
	return &size{t.Call("getMapSize")}
}

func (t *tmxMapInfo) SetMapSize(value Size) {
	t.Call("setMapSize", value)
}

func (t *tmxMapInfo) GetTileSize() Size {
	return &size{t.Call("getTileSize")}
}

func (t *tmxMapInfo) SetTileSize(value Size) {
	t.Call("setTileSize", value)
}

func (t *tmxMapInfo) GetLayers() []TMXLayerInfo {
	layers := t.Call("getLayers")
	length := layers.Length()
	out := make([]TMXLayerInfo, length, length)
	for i := 0; i < length; i++ {
		out[i] = &tmxLayerInfo{layers.Index(i)}
	}
	return out
}

func (t *tmxMapInfo) SetLayers(value TMXLayerInfo) {
	t.Call("setLayers", value)
}

func (t *tmxMapInfo) GetTilesets() []TMXTilesetInfo {
	tilesets := t.Call("getLayers")
	length := tilesets.Length()
	out := make([]TMXTilesetInfo, length, length)
	for i := 0; i < length; i++ {
		out[i] = &tmxTilesetInfo{tilesets.Index(i)}
	}
	return out
}

func (t *tmxMapInfo) SetTilesets(value TMXTilesetInfo) {
	t.Call("setTilesets", value)
}

func (t *tmxMapInfo) GetParentElement() *js.Object {
	return t.Call("getParentElement")
}

func (t *tmxMapInfo) SetParentElement(value *js.Object) {
	t.Call("setParentElement", value)
}

func (t *tmxMapInfo) GetParentGID() int {
	return t.Call("getParentGID").Int()
}

func (t *tmxMapInfo) SetParentGID(value int) {
	t.Call("setParentGID", value)
}

func (t *tmxMapInfo) GetLayerAttribs() *js.Object {
	return t.Call("getLayerAttribs")
}

func (t *tmxMapInfo) SetLayerAttribs(value *js.Object) {
	t.Call("setLayerAttribs", value)
}

func (t *tmxMapInfo) GetStoringCharacters() bool {
	return t.Call("getStoringCharacters").Bool()
}

func (t *tmxMapInfo) SetStoringCharacters(value bool) {
	t.Call("setStoringCharacters", value)
}

func (t *tmxMapInfo) GetProperties() []*js.Object {
	properties := t.Call("getProperties")
	length := properties.Length()
	out := make([]*js.Object, length, length)
	for i := 0; i < length; i++ {
		out[i] = properties.Index(i)
	}
	return out
}

func (t *tmxMapInfo) SetProperties(value []*js.Object) {
	t.Call("setProperties", value)
}

func (t *tmxMapInfo) ParseXMLFile(tmxFile string, isXmlString bool) *js.Object {
	return t.Call("parseXMLFile", tmxFile, isXmlString)
}

func (t *tmxMapInfo) ParseXMLString(xmlString string) bool {
	return t.Call("parseXMLString", xmlString).Bool()
}

func (t *tmxMapInfo) GetTileProperties() []*js.Object {
	properties := t.Call("getTileProperties")
	length := properties.Length()
	out := make([]*js.Object, length, length)
	for i := 0; i < length; i++ {
		out[i] = properties.Index(i)
	}
	return out
}

func (t *tmxMapInfo) SetTileProperties(tileProperties *js.Object) {
	t.Call("setTileProperties", tileProperties)
}

func (t *tmxMapInfo) GetCurrentString() string {
	return t.Call("getCurrentString").String()
}

func (t *tmxMapInfo) SetCurrentString(currentString string) {
	t.Call("setCurrentString", currentString)
}

func (t *tmxMapInfo) GetTMXFileName() string {
	return t.Call("getTMXFileName").String()
}

func (t *tmxMapInfo) SetTMXFileName(fileName string) {
	t.Call("setTMXFileName", fileName)
}

const (
	TMXLayerInfo_ATTRIB_NONE   = 1 << 0
	TMXLayerInfo_ATTRIB_BASE64 = 1 << 1
	TMXLayerInfo_ATTRIB_GZIP   = 1 << 2
	TMXLayerInfo_ATTRIB_ZLIB   = 1 << 3
)
