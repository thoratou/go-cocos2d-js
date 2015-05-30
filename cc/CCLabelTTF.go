package cc

// The Scene class
type LabelTTF interface {
	Sprite
	InitWithString(string, string, int, Size, int, int) bool
	//Init() bool defined as part of Node
	Description() string
	GetLineHeight() int
	SetLineHeight(int)
	GetString() string
	GetHorizontalAlignment() int
	GetVerticalAlignment() int
	GetDimensions() Size
	GetFontSize() int
	GetFontName() string
	//TODO FontDefinition
	//InitWithStringAndTextDefinition(string, FontDefinition) bool
	//SetTextDefinition(FontDefinition)
	//GetTextDefinition() FontDefinition
	EnableShadow(Color, Size, float64)
	DisableShadow()
	EnableStroke(Color, float64)
	DisableStroke()
	SetFontFillColor(Color)
	SetString(string)
	SetHorizontalAlignment(int)
	SetVerticalAlignment(int)
	SetDimensions(Size)
	SetFontSize(int)
	SetFontName(string)
	//GetContentSize() Size defined as part of Sprite
	SetTextureRectNoConvert(Rect, bool, Size)
}

type labelTTF struct{ sprite }

// NewLabelTTF is the constructor for Scene.
func NewLabelTTF(text string, fontName string, fontSize int) LabelTTF {
	return &labelTTF{sprite{node{pcc.Get("LabelTTF").New(text, fontName, fontSize)}}}
}

// NewLabelTTF is the constructor for Scene.
func NewLabelTTFWithDimensionAndAlignment(text string, fontName string, fontSize int, dimensions Size, hAlignment int, vAlignment int) LabelTTF {
	return &labelTTF{sprite{node{pcc.Get("LabelTTF").New(text, fontName, fontSize, dimensions, hAlignment, vAlignment)}}}
}

func (l *labelTTF) InitWithString(label string, fontName string, fontSize int, dimensions Size, hAlignment int, vAlignment int) bool {
	return l.Call("initWithString", label, fontName, fontSize, dimensions, hAlignment, vAlignment).Bool()
}

func (l *labelTTF) Init() bool {
	return l.Call("init").Bool()
}

func (l *labelTTF) Description() string {
	return l.Call("description").String()
}

func (l *labelTTF) GetLineHeight() int {
	return l.Call("getLineHeight").Int()
}

func (l *labelTTF) SetLineHeight(lineHeight int) {
	l.Call("setLineHeight", lineHeight)
}

func (l *labelTTF) GetString() string {
	return l.Call("getString").String()
}

func (l *labelTTF) GetHorizontalAlignment() int {
	return l.Call("getHorizontalAlignment").Int()
}

func (l *labelTTF) GetVerticalAlignment() int {
	return l.Call("getVerticalAlignment").Int()
}

func (l *labelTTF) GetDimensions() Size {
	return &size{l.Call("getDimensions")}
}

func (l *labelTTF) GetFontSize() int {
	return l.Call("getFontSize").Int()
}

func (l *labelTTF) GetFontName() string {
	return l.Call("getFontName").String()
}

func (l *labelTTF) EnableShadow(shadowColor Color, offset Size, blurRadius float64) {
	l.Call("enableShadow", shadowColor, offset, blurRadius)
}

func (l *labelTTF) DisableShadow() {
	l.Call("disableShadow")
}

func (l *labelTTF) EnableStroke(strokeColor Color, strokeSize float64) {
	l.Call("enableStroke", strokeColor, strokeSize)
}

func (l *labelTTF) DisableStroke() {
	l.Call("disableStroke")
}

func (l *labelTTF) SetFontFillColor(fillColor Color) {
	l.Call("setFontFillColor", fillColor)
}

func (l *labelTTF) SetString(text string) {
	l.Call("setString", text)
}

func (l *labelTTF) SetHorizontalAlignment(alignment int) {
	l.Call("setHorizontalAlignment", alignment)
}

func (l *labelTTF) SetVerticalAlignment(verticalAlignment int) {
	l.Call("setVerticalAlignment", verticalAlignment)
}

func (l *labelTTF) SetDimensions(dim Size) {
	l.Call("setDimensions", dim)
}

func (l *labelTTF) SetFontSize(fontSize int) {
	l.Call("setFontSize", fontSize)
}

func (l *labelTTF) SetFontName(fontName string) {
	l.Call("setFontName", fontName)
}

func (l *labelTTF) GetContentSize() Size {
	return &size{l.Call("getContentSize")}
}

func (l *labelTTF) SetTextureRectNoConvert(rect Rect, rotated bool, untrimmedSize Size) {
	l.Call("setTextureRect", rect, rotated, untrimmedSize)
}
