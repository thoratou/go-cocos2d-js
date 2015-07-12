package cc

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
	Rect() Rect
	Selected()
	Unselected()
	SetCallback(func(), Node)
	Activate()
}

type menuItem struct{ node }

func NewMenuItem(callback func(), target Node) MenuItem {
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

func (mi *menuItem) Rect() Rect {
	return &rect{mi.Call("rect")}
}

func (mi *menuItem) Selected() {
	mi.Call("selected")
}

func (mi *menuItem) Unselected() {
	mi.Call("unselected")
}

func (mi *menuItem) SetCallback(callback func(), target Node) {
	mi.Call("setCallback", callback, target)
}

func (mi *menuItem) Activate() {
	mi.Call("activate")
}

///////////////////
// MenuItemLabel //
///////////////////

type MenuItemLabel interface {
	MenuItem
	GetDisabledColor() Color
	SetDisabledColor(Color)
	GetLabel() Node
	SetLabel(Node)
	//SetEnabled(bool) defined as part of Node
	//GetOpacity() int defined as part of Node
	//SetOpacity(int) defined as part of Node
	//GetColor() Color defined as part of Node
	//SetColor(Color) defined as part of Node
	SetString(string)
	GetString() string
	//Activate() defined as part of Node
	//Selected() defined as part of Node
	//Unselected() defined as part of Node
}

type menuItemLabel struct{ menuItem }

func NewMenuItemLabelAllArgs(label Node, selector func(), target Node) MenuItemLabel {
	return &menuItemLabel{menuItem{node{pcc.Get("MenuItemLabel").New(label, selector, target)}}}
}

func NewMenuItemLabel(args ...interface{}) MenuItemLabel {
	return &menuItemLabel{menuItem{node{pcc.Get("MenuItemLabel").New(args...)}}}
}

func (m *menuItemLabel) GetDisabledColor() Color {
	return &color{m.Call("getDisabledColor")}
}

func (m *menuItemLabel) SetDisabledColor(color Color) {
	m.Call("setDisabledColor", color)
}

func (m *menuItemLabel) GetLabel() Node {
	return &node{m.Call("getLabel")}
}

func (m *menuItemLabel) SetLabel(label Node) {
	m.Call("setLabel", label)
}

func (m *menuItemLabel) SetEnabled(enabled bool) {
	m.Call("setEnabled", enabled)
}

func (m *menuItemLabel) GetOpacity() int {
	return m.Call("getOpacity").Int()
}

func (m *menuItemLabel) SetOpacity(opacity int) {
	m.Call("setOpacity", opacity)
}

func (m *menuItemLabel) GetColor() Color {
	return &color{m.Call("getColor")}
}

func (m *menuItemLabel) SetColor(color Color) {
	m.Call("setColor", color)
}

func (m *menuItemLabel) SetString(label string) {
	m.Call("SetString", label)
}

func (m *menuItemLabel) GetString() string {
	return m.Call("getString").String()
}

func (m *menuItemLabel) Activate() {
	m.Call("activate")
}

func (m *menuItemLabel) Selected() {
	m.Call("selected")
}

func (m *menuItemLabel) Unselected() {
	m.Call("unselected")
}

///////////////////////
// MenuItemAtlasFont //
///////////////////////

type MenuItemAtlasFont interface {
	MenuItemLabel
}

type menuItemAtlasFont struct{ menuItemLabel }

func NewMenuItemAtlasFont(value string, charMapFile string, itemWidth float64, itemHeight float64, startCharMap string, callback func(), target Node) MenuItemAtlasFont {
	return &menuItemAtlasFont{menuItemLabel{menuItem{node{pcc.Get("MenuItemAtlasFont").New(value, charMapFile, itemWidth, itemHeight, startCharMap, callback, target)}}}}
}

//////////////////
// MenuItemFont //
//////////////////

type MenuItemFont interface {
	MenuItemLabel
	SetFontSize(int)
	GetFontSize() int
	SetFontName(string)
	GetFontName() string
}

type menuItemFont struct{ menuItemLabel }

func NewMenuItemFontAllArgs(value string, callback func(), target Node) MenuItemFont {
	return &menuItemFont{menuItemLabel{menuItem{node{pcc.Get("MenuItemFont").New(value, callback, target)}}}}
}

func NewMenuItemFontWithString(value string) MenuItemFont {
	return &menuItemFont{menuItemLabel{menuItem{node{pcc.Get("MenuItemFont").New(value)}}}}
}

func NewMenuItemFont(args ...interface{}) MenuItemFont {
	return &menuItemFont{menuItemLabel{menuItem{node{pcc.Get("MenuItemFont").New(args...)}}}}
}

func (m *menuItemFont) SetFontSize(s int) {
	m.Call("setFontSize", s)
}

func (m *menuItemFont) GetFontSize() int {
	return m.Call("getFontSize").Int()
}

func (m *menuItemFont) SetFontName(name string) {
	m.Call("setFontName", name)
}

func (m *menuItemFont) GetFontName() string {
	return m.Call("getFontName").String()
}

////////////////////
// MenuItemSprite //
////////////////////

type MenuItemSprite interface {
	MenuItem
	GetNormalImage() Sprite
	SetNormalImage(Sprite)
	GetSelectedImage() Sprite
	SetSelectedImage(Sprite)
	GetDisabledImage() Sprite
	SetDisabledImage(Sprite)
	//SetColor(Color) defined as part of Node
	//GetColor() Color defined as part of Node
	//SetOpacity(int) defined as part of Node
	//GetOpacity() int defined as part of Node
	//Selected() defined as part of Node
	//Unselected() defined as part of Node
	//SetEnabled(bool) defined as part of Node
}

type menuItemSprite struct{ menuItem }

// NewMenuItemSprite is the constructor of cc.MenuItemSprite
// {Image|Null} normalSprite normal state image
// {Image|Null} selectedSprite selected state image
// {Image|Null} disabled state image
// {function|Null} callback
// {cc.Node|Null} target Node
func NewMenuItemSpriteAllArgs(normalImage *string, selectedImage *string, disabledImage *string, callback *func(), target Node) MenuItemSprite {
	return &menuItemSprite{menuItem{node{pcc.Get("MenuItemSprite").New(normalImage, selectedImage, disabledImage, callback, target)}}}
}

func NewMenuItemSprite(args ...interface{}) MenuItemSprite {
	return &menuItemSprite{menuItem{node{pcc.Get("MenuItemSprite").New(args...)}}}
}

func (m *menuItemSprite) GetNormalImage() Sprite {
	return &sprite{node{m.Call("getNormalImage")}}
}

func (m *menuItemSprite) SetNormalImage(normalImage Sprite) {
	m.Call("setNormalImage", normalImage)
}

func (m *menuItemSprite) GetSelectedImage() Sprite {
	return &sprite{node{m.Call("getSelectedImage")}}
}

func (m *menuItemSprite) SetSelectedImage(selectedImage Sprite) {
	m.Call("setSelectedImage", selectedImage)
}

func (m *menuItemSprite) GetDisabledImage() Sprite {
	return &sprite{node{m.Call("getDisabledImage")}}
}

func (m *menuItemSprite) SetDisabledImage(disabledImage Sprite) {
	m.Call("setDisabledImage", disabledImage)
}

func (m *menuItemSprite) SetColor(color Color) {
	m.Call("setColor", color)
}

func (m *menuItemSprite) GetColor() Color {
	return &color{m.Call("getColor")}
}

func (m *menuItemSprite) SetOpacity(opacity int) {
	m.Call("setOpacity", opacity)
}

func (m *menuItemSprite) GetOpacity() int {
	return m.Call("getOpacity").Int()
}

func (m *menuItemSprite) Selected() {
	m.Call("selected")
}

func (m *menuItemSprite) Unselected() {
	m.Call("unselected")
}

func (m *menuItemSprite) SetEnabled(enabled bool) {
	m.Call("setEnabled", enabled)
}

///////////////////
// MenuItemImage //
///////////////////

type MenuItemImage interface {
	MenuItemSprite
	SetNormalSpriteFrame(SpriteFrame)
	SetSelectedSpriteFrame(SpriteFrame)
	SetDisabledSpriteFrame(SpriteFrame)
}

type menuItemImage struct{ menuItemSprite }

// NewMenuItemImage is the constructor of cc.MenuItemImage
// {Image|Null} normalSprite normal state image
// {Image|Null} selectedSprite selected state image
// {Image|Null} disabled state image
// {function|Null} callback
// {cc.Node|Null} target Node
func NewMenuItemImageAllArgs(normalImage *string, selectedImage *string, disabledImage *string, callback *func(), target Node) MenuItemImage {
	return &menuItemImage{menuItemSprite{menuItem{node{pcc.Get("MenuItemImage").New(normalImage, selectedImage, disabledImage, callback, target)}}}}
}

func NewMenuItemImage(args ...interface{}) MenuItemImage {
	return &menuItemImage{menuItemSprite{menuItem{node{pcc.Get("MenuItemImage").New(args...)}}}}
}

func (m *menuItemImage) SetNormalSpriteFrame(frame SpriteFrame) {
	m.Call("setNormalSpriteFrame", frame)
}

func (m *menuItemImage) SetSelectedSpriteFrame(frame SpriteFrame) {
	m.Call("setSelectedSpriteFrame", frame)
}

func (m *menuItemImage) SetDisabledSpriteFrame(frame SpriteFrame) {
	m.Call("setDisabledSpriteFrame", frame)
}

////////////////////
// MenuItemToggle //
////////////////////

type MenuItemToggle interface {
	MenuItem
	//GetOpacity() int defined as part of Node
	//SetOpacity(int) defined as part of Node
	//GetColor() Color defined as part of Node
	//SetColor(Color) defined as part of Node
	GetSelectedIndex() int
	SetSelectedIndex(int)
	GetSubItems() []MenuItem
	SetSubItems([]MenuItem)
	AddSubItem(MenuItem)
	//Activate() defined as part of Node
	//Selected() defined as part of Node
	//Unselected() defined as part of Node
	//SetEnabled(bool) defined as part of Node
	GetSelectedItem() MenuItem
	//SetOnEnter(func()) defined as part of Node
	//OnEnter() defined as part of Node
	//OnEnterSuper() defined as part of Node
}

type menuItemToggle struct{ menuItem }

func NewMenuItemToggle(items ...MenuItem) MenuItemToggle {
	//MenuItemToggle only accepts multiple parameters
	//has to convert it to ...interface{}
	length := len(items)
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = items[i]
	}
	return &menuItemToggle{menuItem{node{pcc.Get("MenuItemToggle").New(out...)}}}
}

func (m *menuItemToggle) GetOpacity() int {
	return m.Call("getOpacity").Int()
}

func (m *menuItemToggle) SetOpacity(opacity int) {
	m.Call("setOpacity", opacity)
}

func (m *menuItemToggle) GetColor() Color {
	return &color{m.Call("getColor")}
}

func (m *menuItemToggle) SetColor(color Color) {
	m.Call("setColor", color)
}

func (m *menuItemToggle) GetSelectedIndex() int {
	return m.Call("getSelectedIndex").Int()
}

func (m *menuItemToggle) SetSelectedIndex(selectedIndex int) {
	m.Call("setSelectedIndex", selectedIndex)
}

func (m *menuItemToggle) GetSubItems() []MenuItem {
	subItems := m.Call("getSubItems")
	length := subItems.Length()
	out := make([]MenuItem, length, length)
	for i := 0; i < length; i++ {
		out[i] = &menuItem{node{subItems.Index(i)}}
	}
	return out
}

func (m *menuItemToggle) SetSubItems(subItems []MenuItem) {
	m.Call("setSubItems", subItems)
}
func (m *menuItemToggle) AddSubItem(item MenuItem) {
	m.Call("addSubItem", item)
}

func (m *menuItemToggle) Activate() {
	m.Call("activate")
}

func (m *menuItemToggle) Selected() {
	m.Call("selected")
}

func (m *menuItemToggle) Unselected() {
	m.Call("unselected")
}

func (m *menuItemToggle) SetEnabled(enabled bool) {
	m.Call("setEnabled", enabled)
}

func (m *menuItemToggle) GetSelectedItem() MenuItem {
	return &menuItem{node{m.Call("getSelectedItem")}}
}

func (m *menuItemToggle) SetOnEnter(cb func()) {
	BackupFunc(m.Object, "onEnter")
	m.Set("onEnter", cb)
}

func (m *menuItemToggle) OnEnter() {
	m.Call("onEnter")
}

func (m *menuItemToggle) OnEnterSuper() {
	SuperCall(m.Object, "onEnter")
}
