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

func (s *spriteFrame) TextureLoaded() bool {
	return s.Call("textureLoaded").Bool()
}

func (s *spriteFrame) GetRectInPixels() Rect {
	return &rect{s.Call("getRectInPixels")}
}

func (s *spriteFrame) SetRectInPixels(rectInPixels Rect) {
	s.Call("setRectInPixels", rectInPixels)
}

func (s *spriteFrame) IsRotated() bool {
	return s.Call("isRotated").Bool()
}

func (s *spriteFrame) SetRotated(bRotated bool) {
	s.Call("setRotated", bRotated)
}

func (s *spriteFrame) GetRect() Rect {
	return &rect{s.Call("getRect")}
}

func (s *spriteFrame) SetRect(rect Rect) {
	s.Call("setRect", rect)
}

func (s *spriteFrame) GetOffsetInPixels() Point {
	return &point{s.Call("getOffsetInPixels")}
}

func (s *spriteFrame) SetOffsetInPixels(offsetInPixels Point) {
	s.Call("setOffsetInPixels", offsetInPixels)
}

func (s *spriteFrame) GetOriginalSizeInPixels() Size {
	return &size{s.Call("getOriginalSizeInPixels")}
}

func (s *spriteFrame) SetOriginalSizeInPixels(sizeInPixels Size) {
	s.Call("setOriginalSizeInPixels", sizeInPixels)
}

func (s *spriteFrame) GetOriginalSize() Size {
	return &size{s.Call("getOriginalSize")}
}

func (s *spriteFrame) SetOriginalSize(sizeInPixels Size) {
	s.Call("setOriginalSize", sizeInPixels)
}

func (s *spriteFrame) GetTexture() Texture2D {
	return &texture2D{s.Call("getTexture")}
}

func (s *spriteFrame) SetTexture(texture Texture2D) {
	s.Call("setTexture", texture)
}

func (s *spriteFrame) GetOffset() Point {
	return &point{s.Call("getOffset")}
}

func (s *spriteFrame) SetOffset(offsets Point) {
	s.Call("setOffset", offsets)
}

func (s *spriteFrame) Clone() SpriteFrame {
	return &spriteFrame{s.Call("clone")}
}

func (s *spriteFrame) CopyWithZone() SpriteFrame {
	return &spriteFrame{s.Call("copyWithZone")}
}

func (s *spriteFrame) Copy() SpriteFrame {
	return &spriteFrame{s.Call("copy")}
}

func (s *spriteFrame) InitWithTexture(texture Texture2D, rect Rect, rotated bool, offset Point, originalSize Size) bool {
	return s.Call("initWithTexture", texture, rect, rotated, offset, originalSize).Bool()
}
