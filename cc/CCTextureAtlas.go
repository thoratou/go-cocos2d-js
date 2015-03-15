package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

// The touch event class
type TextureAtlas interface {
	GetTotalQuads() int
	GetCapacity() int
	GetTexture() Image
	SetTexture(Image)
	IsDirty() bool
	GetQuads() []V3F_C4B_T2F_Quad
	SetQuads([]V3F_C4B_T2F_Quad)
	Description() string
	InitWithFile(string, int)
	InitWithTexture(Image, int)
	UpdateQuad(V3F_C4B_T2F_Quad, int)
	InsertQuad(V3F_C4B_T2F_Quad, int)
	InsertQuads([]V3F_C4B_T2F_Quad, int, int)
	InsertQuadFromIndex(int, int)
	RemoveQuadAtIndex(int)
	RemoveQuadsAtIndex(int, int)
	RemoveAllQuads()
	ResizeCapacity(int) bool
	IncreaseTotalQuadsWith(int)
	MoveQuadsFromIndex(int, int, int)
	FillWithEmptyQuadsFromIndex(int, int)
	DrawQuads()
}

type textureAtlas struct{ *js.Object }

// NewTextureAtlas is the constructor for TextureAtlas.
func NewTextureAtlas(texture Texture2D, capacity int) TextureAtlas {
	return &textureAtlas{pcc.Get("TextureAtlas").New(texture, capacity)}
}
