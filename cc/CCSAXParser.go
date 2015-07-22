package cc

import "github.com/gopherjs/gopherjs/js"

///////////////
// SAXParser //
///////////////

type SAXParser interface {
	Parse(string) interface{}
}

type saxParser struct{ *js.Object }

func NewSAXPaser() SAXParser {
	return &saxParser{pcc.Get("SAXParser").New()}
}

func (s *saxParser) Parse(xmlTxt string) interface{} {
	return s.Call("parse", xmlTxt).Interface()
}

/////////////////
// PlistParser //
/////////////////

type PlistParser interface {
	SAXParser
}

type plistParser struct{ saxParser }

func (p *plistParser) Parse(xmlTxt string) interface{} {
	return p.Call("parse", xmlTxt).Interface()
}
