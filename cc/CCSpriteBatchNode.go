package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type SpriteBatchNode interface {
	Node
	AddSpriteWithoutQuad(Sprite, int, int) SpriteBatchNode
	GetTextureAtlas() TextureAtlas
	SetTextureAtlas(TextureAtlas)
	GetDescendants() []SpriteBatchNode
	InitWithFile(string, int) bool
	//Also got an issue with this method, same one than InitWithFile
	//Init(string, int) bool defined as part of Node
	IncreaseAtlasCapacity()
	RemoveChildAtIndex(int, bool)
	RebuildIndexInOrder(Sprite, int) int
	HighestAtlasIndexInChild(Sprite) int
	LowestAtlasIndexInChild(Sprite) int
	AtlasIndexForChild(Sprite, int) int
	ReorderBatch(bool)
	SetBlendFunc(BlendFunc)
	GetBlendFunc() BlendFunc
	//ReorderChild(Sprite, int) defined as part of Node
	//RemoveChild(Sprite, bool) defined as part of Node
	UpdateQuadFromSprite(Sprite, int)
	InsertQuadFromSprite(Sprite, int)
	InitWithTexture(Texture2D, int) bool
	InsertChild(Sprite, int)
	AppendChild(Sprite)
	RemoveSpriteFromAtlas(Sprite)
	GetTexture() Texture2D
	SetTexture(Texture2D)
	//AddChild(Node) defined as part of Node
	//AddChildWithOrder(Node, int) defined as part of Node
	//AddChildWithTag(Node, int) defined as part of Node
	//AddChildWithOrderAndTag(Node, int, int) defined as part of Node
	//RemoveAllChildren(bool) defined as part of Node
	//SortAllChildren() defined as part of Node
}

type spriteBatchNode struct{ node }

func NewSpriteBatchNode(fileImage string, capacity int) SpriteBatchNode {
	return &spriteBatchNode{node{pcc.Get("SpriteBatchNode").New(fileImage, capacity)}}
}

const (
	SpriteBatchNode_DEFAULT_CAPACITY = 29
)

func (s *spriteBatchNode) AddSpriteWithoutQuad(child Sprite, z int, aTag int) SpriteBatchNode {
	return &spriteBatchNode{node{s.Call("addSpriteWithoutQuad", child, z, aTag)}}
}

func (s *spriteBatchNode) GetTextureAtlas() TextureAtlas {
	return &textureAtlas{s.Call("getTextureAtlas")}
}

func (s *spriteBatchNode) SetTextureAtlas(textureAtlas TextureAtlas) {
	s.Call("setTextureAtlas", textureAtlas)
}

func (s *spriteBatchNode) GetDescendants() []SpriteBatchNode {
	descendants := s.Call("getDescendants")
	length := descendants.Length()
	out := make([]SpriteBatchNode, length, length)
	for i := 0; i < length; i++ {
		out[i] = &spriteBatchNode{node{descendants.Index(i)}}
	}
	return out
}

func (s *spriteBatchNode) InitWithFile(fileImage string, capacity int) bool {
	return s.Call("initWithFile", fileImage, capacity).Bool()
}

func (s *spriteBatchNode) IncreaseAtlasCapacity() {
	s.Call("increaseAtlasCapacity")
}

func (s *spriteBatchNode) RemoveChildAtIndex(index int, doCleanup bool) {
	s.Call("removeChildAtIndex", index, doCleanup)
}

func (s *spriteBatchNode) RebuildIndexInOrder(pobParent Sprite, index int) int {
	return s.Call("rebuildIndexInOrder", pobParent, index).Int()
}

func (s *spriteBatchNode) HighestAtlasIndexInChild(sprite Sprite) int {
	return s.Call("highestAtlasIndexInChild", sprite).Int()
}

func (s *spriteBatchNode) LowestAtlasIndexInChild(sprite Sprite) int {
	return s.Call("lowestAtlasIndexInChild", sprite).Int()
}

func (s *spriteBatchNode) AtlasIndexForChild(sprite Sprite, nZ int) int {
	return s.Call("atlasIndexForChild", sprite, nZ).Int()
}

func (s *spriteBatchNode) ReorderBatch(reorder bool) {
	s.Call("reorderBatch", reorder)
}

func (s *spriteBatchNode) SetBlendFunc(blendFunc BlendFunc) {
	s.Call("setBlendFunc", blendFunc)
}

func (s *spriteBatchNode) GetBlendFunc() BlendFunc {
	return &blendFunc{s.Call("getBlendFunc")}
}

//child should be Sprite type
func (s *spriteBatchNode) ReorderChild(child Node, zOrder int) {
	s.Call("reorderChild", child, zOrder)
}

//child should be Sprite type
func (s *spriteBatchNode) RemoveChild(child Node, cleanup bool) {
	s.Call("removeChild", child, cleanup)
}

func (s *spriteBatchNode) UpdateQuadFromSprite(sprite Sprite, index int) {
	s.Call("updateQuadFromSprite", sprite, index)
}

func (s *spriteBatchNode) InsertQuadFromSprite(sprite Sprite, index int) {
	s.Call("insertQuadFromSprite", sprite, index)
}

func (s *spriteBatchNode) InitWithTexture(tex Texture2D, capacity int) bool {
	return s.Call("initWithTexture", tex, capacity).Bool()
}

func (s *spriteBatchNode) InsertChild(sprite Sprite, index int) {
	s.Call("insertChild", sprite, index)
}

func (s *spriteBatchNode) AppendChild(sprite Sprite) {
	s.Call("appendChild", sprite)
}

func (s *spriteBatchNode) RemoveSpriteFromAtlas(sprite Sprite) {
	s.Call("removeSpriteFromAtlas", sprite)
}

func (s *spriteBatchNode) GetTexture() Texture2D {
	return &texture2D{s.Call("getTexture")}
}

func (s *spriteBatchNode) SetTexture(texture Texture2D) {
	s.Call("setTexture", texture)
}

func (s *spriteBatchNode) AddChild(child Node) {
	s.Call("addChild", child, js.Undefined, js.Undefined)
}

func (s *spriteBatchNode) AddChildWithOrder(child Node, localZOrder int) {
	s.Call("addChild", child, localZOrder, js.Undefined)
}

func (s *spriteBatchNode) AddChildWithTag(child Node, tag int) {
	s.Call("addChild", child, js.Undefined, tag)
}

func (s *spriteBatchNode) AddChildWithOrderAndTag(child Node, localZOrder int, tag int) {
	s.Call("addChild", child, localZOrder, tag)
}

func (s *spriteBatchNode) RemoveAllChildren(cleanup bool) {
	s.Call("removeAllChildren", cleanup)
}

func (s *spriteBatchNode) SortAllChildren() {
	s.Call("sortAllChildren")
}
