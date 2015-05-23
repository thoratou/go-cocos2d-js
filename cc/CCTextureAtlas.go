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

func (t *textureAtlas) GetTotalQuads() int {
	return t.Call("getTotalQuads").Int()
}

func (t *textureAtlas) GetCapacity() int {
	return t.Call("getCapacity").Int()
}

func (t *textureAtlas) GetTexture() Image {
	return &image{t.Call("getTexture")}
}

func (t *textureAtlas) SetTexture(texture Image) {
	t.Call("setTexture", texture)
}

func (t *textureAtlas) IsDirty() bool {
	return t.Call("isDirty").Bool()
}

func (t *textureAtlas) GetQuads() []V3F_C4B_T2F_Quad {
	quads := t.Call("getQuads")
	length := quads.Length()
	out := make([]V3F_C4B_T2F_Quad, length, length)
	for i := 0; i < length; i++ {
		out[i] = quads.Index(i)
	}
	return out
}

func (t *textureAtlas) SetQuads(quads []V3F_C4B_T2F_Quad) {
	t.Call("setQuads", quads)
}

func (t *textureAtlas) Description() string {
	return t.Call("description").String()
}

func (t *textureAtlas) InitWithFile(file string, capacity int) {
	t.Call("initWithFile", file, capacity)
}

func (t *textureAtlas) InitWithTexture(texture Image, capacity int) {
	t.Call("initWithTexture", texture, capacity)
}

func (t *textureAtlas) UpdateQuad(quad V3F_C4B_T2F_Quad, index int) {
	t.Call("updateQuad", quad, index)
}

func (t *textureAtlas) InsertQuad(quad V3F_C4B_T2F_Quad, index int) {
	t.Call("insertQuad", quad, index)
}
func (t *textureAtlas) InsertQuads(quads []V3F_C4B_T2F_Quad, index int, amount int) {
	t.Call("insertQuads", quads, index, amount)
}

func (t *textureAtlas) InsertQuadFromIndex(fromIndex int, newIndex int) {
	t.Call("insertQuadFromIndex", fromIndex, newIndex)
}

func (t *textureAtlas) RemoveQuadAtIndex(index int) {
	t.Call("removeQuadAtIndex", index)
}

func (t *textureAtlas) RemoveQuadsAtIndex(index int, amount int) {
	t.Call("removeQuadsAtIndex", index, amount)
}

func (t *textureAtlas) RemoveAllQuads() {
	t.Call("removeAllQuads")
}

func (t *textureAtlas) ResizeCapacity(newCapacity int) bool {
	return t.Call("resizeCapacity", newCapacity).Bool()
}

func (t *textureAtlas) IncreaseTotalQuadsWith(amount int) {
	t.Call("increaseTotalQuadsWith", amount)
}

func (t *textureAtlas) MoveQuadsFromIndex(oldIndex int, amount int, newIndex int) {
	t.Call("moveQuadsFromIndex", oldIndex, amount, newIndex)
}

func (t *textureAtlas) FillWithEmptyQuadsFromIndex(index int, amount int) {
	t.Call("fillWithEmptyQuadsFromIndex", index, amount)
}
func (t *textureAtlas) DrawQuads() {
	t.Call("drawQuads")
}
