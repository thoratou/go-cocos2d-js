package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type Sprite interface {
	Node
	TextureLoaded() bool
	IsDirty() bool
	IsTextureRectRotated() bool
	GetAtlasIndex() int
	SetAtlasIndex(int)
	GetTextureRect() Rect
	GetTextureAtlas() TextureAtlas
	SetTextureAtlas(TextureAtlas)
	GetOffsetPosition() Point
	GetBlendFunc() BlendFunc
	InitWithSpriteFrame(SpriteFrame)
	InitWithSpriteFrameName(string)
	UseBatchNode(SpriteBatchNode)
	SetVertexRect(Rect)
	//SortAllChildren() defined as part of Node
	//ReorderChild(Node, int) defined as part of Node
	//SetVisible(bool) defined as part of Node
	//RemoveAllChildren(bool) defined as part of Node
	//IgnoreAnchorPointForPosition(bool) defined as part of Node
	SetFlippedX(bool)
	SetFlippedY(bool)
	IsFlippedX() bool
	IsFlippedY() bool
	//SetOpacityModifyRGB(bool) defined as part of Node
	//IsOpacityModifyRGB() bool defined as part of Node
	SetDisplayFrameWithAnimationName(string, int)
	GetBatchNode() SpriteBatchNode
	GetTexture() Texture2D
	GetQuad() V3F_C4B_T2F_Quad
	SetBlendFunc(BlendFunc)
	//Init() bool defined as part of Node
	InitWithFile(string, Rect) bool
	InitWithTexture(Texture2D, Rect, bool, bool) bool
	SetTextureRect(Rect, bool, Size, bool)
	//UpdateTransform() defined as part of Node
	//AddChild(Node) defined as part of Node
	//AddChildWithOrder(Node, int) defined as part of Node
	//AddChildWithTag(Node, int) defined as part of Node
	//AddChildWithOrderAndTag(Node, int, int) defined as part of Node
	SetSpriteFrame(SpriteFrame)
	SetDisplayFrame(SpriteFrame)
	IsFrameDisplayed(SpriteFrame) bool
	DisplayFrame() SpriteFrame
	SetBatchNode(SpriteBatchNode)
	SetTexture(Texture2D)
}

type sprite struct{ node }

func NewSprite(fileName string) Sprite {
	return &sprite{node{pcc.Get("Sprite").New(fileName)}}
}

func NewSpriteWithRectAndRotated(fileName string, rect Rect, rotated bool) Sprite {
	return &sprite{node{pcc.Get("Sprite").New(fileName, rect, rotated)}}
}

const (
	INDEX_NOT_INITIALIZED = -1
)

func (s *sprite) TextureLoaded() bool {
	return s.Call("textureLoaded").Bool()
}

func (s *sprite) IsDirty() bool {
	return s.Call("isDirty").Bool()
}

func (s *sprite) IsTextureRectRotated() bool {
	return s.Call("isTextureRectRotated").Bool()
}

func (s *sprite) GetAtlasIndex() int {
	return s.Call("getAtlasIndex").Int()
}

func (s *sprite) SetAtlasIndex(atlasIndex int) {
	s.Call("setAtlasIndex", atlasIndex)
}

func (s *sprite) GetTextureRect() Rect {
	return &rect{s.Call("getTextureRect")}
}

func (s *sprite) GetTextureAtlas() TextureAtlas {
	return &textureAtlas{s.Call("getTextureAtlas")}
}

func (s *sprite) SetTextureAtlas(textureAtlas TextureAtlas) {
	s.Call("setTextureAtlas", textureAtlas)
}

func (s *sprite) GetOffsetPosition() Point {
	return &point{s.Call("getOffsetPosition")}
}

func (s *sprite) GetBlendFunc() BlendFunc {
	return &blendFunc{s.Call("getBlendFunc")}
}

func (s *sprite) InitWithSpriteFrame(spriteFrame SpriteFrame) {
	s.Call("initWithSpriteFrame", spriteFrame)
}

func (s *sprite) InitWithSpriteFrameName(spriteFrameName string) {
	s.Call("initWithSpriteFrameName", spriteFrameName)
}

func (s *sprite) UseBatchNode(batchNode SpriteBatchNode) {
	s.Call("useBatchNode", batchNode)
}

func (s *sprite) SetVertexRect(rect Rect) {
	s.Call("setVertexRect", rect)
}

func (s *sprite) SortAllChildren() {
	s.Call("sortAllChildren")
}

func (s *sprite) ReorderChild(child Node, zOrder int) {
	s.Call("reorderChild", child, zOrder)
}

func (s *sprite) SetVisible(visible bool) {
	s.Call("setVisible", visible)
}

func (s *sprite) RemoveAllChildren(cleanup bool) {
	s.Call("removeAllChildren", cleanup)
}

func (s *sprite) IgnoreAnchorPointForPosition(newValue bool) {
	s.Call("ignoreAnchorPointForPosition", newValue)
}

func (s *sprite) SetFlippedX(flippedX bool) {
	s.Call("setFlippedX", flippedX)
}

func (s *sprite) SetFlippedY(flippedY bool) {
	s.Call("setFlippedY", flippedY)
}

func (s *sprite) IsFlippedX() bool {
	return s.Call("isFlippedX").Bool()
}

func (s *sprite) IsFlippedY() bool {
	return s.Call("isFlippedY").Bool()
}

func (s *sprite) SetOpacityModifyRGB(opacityValue bool) {
	s.Call("setOpacityModifyRGB", opacityValue)
}

func (s *sprite) IsOpacityModifyRGB() bool {
	return s.Call("isOpacityModifyRGB").Bool()
}

func (s *sprite) SetDisplayFrameWithAnimationName(animationName string, frameIndex int) {
	s.Call("setDisplayFrameWithAnimationName", animationName, frameIndex)
}

func (s *sprite) GetBatchNode() SpriteBatchNode {
	return &spriteBatchNode{node{s.Call("getBatchNode")}}
}

func (s *sprite) GetTexture() Texture2D {
	return &texture2D{s.Call("getTexture")}
}
func (s *sprite) GetQuad() V3F_C4B_T2F_Quad {
	return s.Call("getQuad")
}

func (s *sprite) SetBlendFunc(blendFunc BlendFunc) {
	s.Call("setBlendFunc", blendFunc)
}

func (s *sprite) Init() bool {
	return s.Call("init").Bool()
}

func (s *sprite) InitWithFile(filename string, rect Rect) bool {
	return s.Call("initWithFile", filename, rect).Bool()
}
func (s *sprite) InitWithTexture(texture Texture2D, rect Rect, rotated bool, counterclockwise bool) bool {
	return s.Call("initWithTexture", texture, rect, counterclockwise).Bool()
}

func (s *sprite) SetTextureRect(rect Rect, rotated bool, untrimmedSize Size, needConvert bool) {
	s.Call("setTextureRect", rect, rotated, untrimmedSize, needConvert)
}

func (s *sprite) UpdateTransform() {
	s.Call("updateTransform")
}

func (s *sprite) AddChild(child Node) {
	s.Call("addChild", child, js.Undefined, js.Undefined)
}

func (s *sprite) AddChildWithOrder(child Node, localZOrder int) {
	s.Call("addChild", child, localZOrder, js.Undefined)
}

func (s *sprite) AddChildWithTag(child Node, tag int) {
	s.Call("addChild", child, js.Undefined, tag)
}

func (s *sprite) AddChildWithOrderAndTag(child Node, localZOrder int, tag int) {
	s.Call("addChild", child, localZOrder, tag)
}

func (s *sprite) SetSpriteFrame(newFrame SpriteFrame) {
	s.Call("setSpriteFrame", newFrame)
}

func (s *sprite) SetDisplayFrame(newFrame SpriteFrame) {
	s.Call("setDisplayFrame", newFrame)
}

func (s *sprite) IsFrameDisplayed(frame SpriteFrame) bool {
	return s.Call("isFrameDisplayed", frame).Bool()
}

func (s *sprite) DisplayFrame() SpriteFrame {
	return &spriteFrame{s.Call("displayFrame")}
}

func (s *sprite) SetBatchNode(spriteBatchNode SpriteBatchNode) {
	s.Call("setBatchNode", spriteBatchNode)
}

func (s *sprite) SetTexture(texture Texture2D) {
	s.Call("setTexture", texture)
}
