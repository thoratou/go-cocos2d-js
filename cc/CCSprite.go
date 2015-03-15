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
	//Init() defined as part of Node
	InitWithFile(string, Rect)
	InitWithTexture(Texture2D, Rect, bool, bool, bool)
	SetTextureRect(Rect, bool, Size)
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

type sprite struct{ *js.Object }

func NewSprite(fileName string, rect Rect, rotated bool) Node {
	return &sprite{pcc.Get("Sprite").New(fileName, rect, rotated)}
}

const (
	INDEX_NOT_INITIALIZED = -1
)
