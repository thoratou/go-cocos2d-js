package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type SpriteBatchNode interface {
	Node
	AddSpriteWithoutQuad(Sprite, int int) SpriteBatchNode
	GetTextureAtlas() TextureAtlas
	SetTextureAtlas(TextureAtlas)
	GetDescendants() []SpriteBatchNode
	InitWithFile(string, int) bool
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

type spriteBatchNode struct{ *js.Object }

func NewSpriteBatchNode(fileImage string, capacity int) SpriteBatchNode {
	return &spriteBatchNode{pcc.Get("SpriteBatchNode").New(fileImage, capacity)}
}

const (
	SpriteBatchNode_DEFAULT_CAPACITY = 29
)
