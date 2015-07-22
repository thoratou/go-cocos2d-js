package cc

import "github.com/gopherjs/gopherjs/js"

///////////////
// SAXParser //
///////////////

type SAXParser interface {
	Parse(string) *js.Object
}

type saxParser struct{ *js.Object }

func NewSAXPaser() SAXParser {
	return &saxParser{pcc.Get("SAXParser").New()}
}

func (s *saxParser) Parse(xmlTxt string) *js.Object {
	return s.Call("parse", xmlTxt)
}

/////////////////
// PlistParser //
/////////////////

type PlistParser interface {
	SAXParser
}

type plistParser struct{ saxParser }

func (p *plistParser) Parse(xmlTxt string) *js.Object {
	return p.Call("parse", xmlTxt)
}
