package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

type SpriteFrame interface {
	TextureLoaded() bool
	GetRectInPixels() Rect
	SetRectInPixels(Rect)
	IsRotated() bool
	SetRotated(bool)
	GetRect() Rect
	SetRect(Rect)
	GetOffsetInPixels() Point
	SetOffsetInPixels(Point)
	GetOriginalSizeInPixels() Size
	SetOriginalSizeInPixels(Size)
	GetOriginalSize() Size
	SetOriginalSize(Size)
	GetTexture() Texture2D
	SetTexture(Texture2D)
	GetOffset() Point
	SetOffset(Point)
	Clone() SpriteFrame
	CopyWithZone() SpriteFrame
	Copy() SpriteFrame
	InitWithTexture(Texture2D, Rect, bool, Point, Size) bool
}

type spriteFrame struct{ *js.Object }

func NewSpriteFrame(fileName string, rect Rect, rotated bool, offset Point, originalSize Size) SpriteFrame {
	return &spriteFrame{pcc.Get("SpriteFrame").New(fileName, rect, rotated, offset, originalSize)}
}
