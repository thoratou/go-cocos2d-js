package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

//////////////
// MenuItem //
//////////////

type MenuItem interface {
	Node
	IsSelected() bool
	//SetOpacityModifyRGB(bool) defined as part of Node
	//IsOpacityModifyRGB() bool defined as part of Node
	IsEnabled() bool
	SetEnabled(bool)
	InitWithCallback(func(*js.Object), Node) bool
	Rect() Rect
	Selected()
	Unselected()
	SetCallback(func(*js.Object), Node)
	Activate()
}

type menuItem struct{ node }

func NewMenuItem(callback func(*js.Object), target Node) MenuItem {
	return &menuItem{node{pcc.Get("MenuItem").New(callback, target)}}
}

func (mi *menuItem) IsSelected() bool {
	return mi.Call("isSelected").Bool()
}

func (mi *menuItem) SetOpacityModifyRGB(value bool) {
	mi.Call("setOpacityModifyRGB", value)
}

func (mi *menuItem) IsOpacityModifyRGB() bool {
	return mi.Call("isOpacityModifyRGB").Bool()
}

func (mi *menuItem) IsEnabled() bool {
	return mi.Call("isEnabled").Bool()
}

func (mi *menuItem) SetEnabled(enable bool) {
	mi.Call("setEnabled", enable)
}

func (mi *menuItem) InitWithCallback(callback func(*js.Object), target Node) bool {
	return mi.Call("isEnabled", callback, target).Bool()
}

func (mi *menuItem) Rect() Rect {
	return &rect{mi.Call("rect")}
}

func (mi *menuItem) Selected() {
	mi.Call("selected")
}

func (mi *menuItem) Unselected() {
	mi.Call("unselected")
}

func (mi *menuItem) SetCallback(callback func(*js.Object), target Node) {
	mi.Call("setCallback", callback, target)
}

func (mi *menuItem) Activate() {
	mi.Call("activate")
}

///////////////////
// MenuItemLabel //
///////////////////

//TODO MenuItemLabel

///////////////////////
// MenuItemAtlasFont //
///////////////////////

//TODO MenuItemAtlasFont

//////////////////
// MenuItemFont //
//////////////////

//TODO MenuItemFont

////////////////////
// MenuItemSprite //
////////////////////

type MenuItemSprite interface {
	MenuItem
}

type menuItemSprite struct{ menuItem }

//func NewMenuItemSprite() MenuItemSprite {
//	return &menuItemSprite{menuItem{node{pcc.Get("MenuItemSprite").New()}}}
//}

///////////////////
// MenuItemImage //
///////////////////

type MenuItemImage interface {
	MenuItemSprite
}

type menuItemImage struct{ menuItemSprite }

func NewMenuItemImage(normalImage *string, selectedImage *string, disabledImage *string, callback *func(), target Node) MenuItemImage {
	return &menuItemImage{menuItemSprite{menuItem{node{pcc.Get("MenuItemImage").New(normalImage, selectedImage, disabledImage, callback, target)}}}}
}

////////////////////
// MenuItemToggle //
////////////////////

//TODO MenuItemToggle
