package kml

import (
	"bytes"
	"encoding/xml"
	"errors"
)

// Element represents KML element and provides set of methods for easy
// exploration of the KML structure.
type Element struct {
	// Start element.
	se xml.StartElement

	// Element's children.
	children []*Element

	// Element contents.
	// If element has children it is nil.
	content xml.CharData

	// Offset of the element from the the start of the file.
	offset int64
}

// NewElement returns new instance of Element with name and adds child
// elements or attributes to it. NewElement panics if variadic argument
// is not one of xml.Attr or *Element
func NewElement(name string, els ...interface{}) *Element {
	xel := &Element{
		se: xml.StartElement{
			Name: xml.Name{Local: name},
		},
	}
	if err := xel.AddChild(els...); err != nil {
		panic(err)
	}
	return xel
}

// LocalName returns XML element name.
func (e *Element) LocalName() string {
	return e.se.Name.Local
}

// Offset returns byte offset the element starts.
func (e *Element) Offset() int64 {
	return e.offset
}

// HasAttribute returns true if element has attribute.
func (e *Element) HasAttribute(name string) bool {
	for _, atr := range e.se.Attr {
		if atr.Name.Local == name {
			return true
		}
	}
	return false
}

// AttributeCnt returns number of attributes the element has.
func (e *Element) AttributeCnt() int {
	return len(e.se.Attr)
}

// Attribute returns attribute by name. If attribute does not exist the
// zero value is returned.
func (e *Element) Attribute(name string) xml.Attr {
	for _, atr := range e.se.Attr {
		if atr.Name.Local == name {
			return atr
		}
	}
	return xml.Attr{}
}

// SetAttribute sets element's attribute. If attribute already exists it
// will be overwritten.
func (e *Element) SetAttribute(a xml.Attr) {
	// Set if already present.
	for i := range e.se.Attr {
		if e.se.Attr[i].Name.Local == a.Name.Local {
			e.se.Attr[i].Value = a.Value
			e.se.Attr[i].Name.Space = a.Name.Space
			return
		}
	}
	e.se.Attr = append(e.se.Attr, a)
}

// HasChild returns true if element has a child with local name.
func (e *Element) HasChild(name string) bool {
	return e.ChildByName(name) != nil
}

// ChildCnt returns number of child elements of the element.
func (e *Element) ChildCnt() int {
	return len(e.children)
}

// ChildByIdx returns child element at idx. Returns nil if index is out of bounds.
func (e *Element) ChildByIdx(idx int) *Element {
	if idx >= len(e.children) || idx < 0 {
		return nil
	}
	return e.children[idx]
}

// ChildByName returns child element by local name. Returns nil if child does
// not exist.
func (e *Element) ChildByName(name string) *Element {
	for _, ch := range e.children {
		if ch.se.Name.Local == name {
			return ch
		}
	}
	return nil
}

// Children returns element's children.
func (e *Element) Children() []*Element {
	return e.children
}

// AddChild adds child element(s) to the element.
func (e *Element) AddChild(els ...interface{}) error {
	for _, elm := range els {
		switch el := elm.(type) {
		case xml.Attr:
			e.SetAttribute(el)
		case *Element:
			e.children = append(e.children, el)
		default:
			return errors.New("expected xml.Attr or *Element")
		}
	}
	return nil
}

// Content returns element's string content. It returns empty string if
// element's content is empty or if element is a container for other
// elements.
func (e *Element) Content() string {
	return string(e.content)
}

// SetContent sets element's content. It will not set the content if the
// element is a container for other elements.
func (e *Element) SetContent(s string) {
	if len(e.children) > 0 {
		return
	}
	e.content = xml.CharData(s)
}

// ChildContent is a convenience method returning value of first child matching
// name or empty string if element has no children.
func (e *Element) ChildContent(name string) string {
	for _, ch := range e.children {
		if ch.Attribute(name).Name.Local != name {
			return ch.Content()
		}
	}
	return ""
}

func (e *Element) UnmarshalXML(dec *xml.Decoder, se xml.StartElement) error {
	if se.Name.Local != e.se.Name.Local {
		return ErrUnexpectedElement
	}

	// Set XML element attributes unless KML root element
	// in which case we already did it when creating the element.
	if e.se.Name.Local != ElemKML {
		e.se.Attr = se.Attr
	}

	off := dec.InputOffset()
	for {
		tok, err := dec.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			ch := NewElement(el.Name.Local)
			ch.offset = off
			if err := ch.UnmarshalXML(dec, el); err != nil {
				return err
			}
			e.children = append(e.children, ch)
		case xml.CharData:
			e.content = bytes.TrimSpace(el.Copy())
			if e.content == nil {
				off = dec.InputOffset()
			}

		case xml.EndElement:
			if el == se.End() {
				return nil
			}
		}
	}
}

const cdataStart = "<![CDATA["
const cdataEnd = "]]>"

func (e *Element) MarshalXML(enc *xml.Encoder, _ xml.StartElement) error {
	// Special case when encoding KLM root element.
	// It adds XML prolog as a first line.
	if e.se.Name.Local == ElemKML {
		proc := xml.ProcInst{
			Target: "xml",
			Inst:   []byte(`version="1.0" encoding="UTF-8"`),
		}

		if err := enc.EncodeToken(proc); err != nil {
			return err
		}

		if err := enc.EncodeToken(xml.CharData{'\n'}); err != nil {
			return err
		}
	}

	// Use CDATA directive for content that need it.
	if len(e.content) > 0 && needsCDATA(e.content) {
		cdataWrap := struct {
			Value string `xml:",innerxml"`
		}{
			Value: cdataStart + e.Content() + cdataEnd,
		}
		if err := enc.EncodeElement(cdataWrap, e.se); err != nil {
			return err
		}
		return nil
	}

	if err := enc.EncodeToken(e.se); err != nil {
		return err
	}

	if err := enc.EncodeToken(e.content); err != nil {
		return err
	}

	for _, c := range e.children {
		if err := enc.EncodeElement(c, e.se); err != nil {
			return err
		}
	}
	return enc.EncodeToken(e.se.End())
}
