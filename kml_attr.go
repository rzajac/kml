package kml

import (
	"encoding/xml"
	"strconv"
)

// Attr returns XML attribute with name and value.
func Attr(name, value string) xml.Attr {
	return xml.Attr{
		Name: xml.Name{
			Local: name,
		},
		Value: value,
	}
}

// AttrID returns id attribute with value.
func AttrID(value string) xml.Attr {
	return Attr("id", value)
}

// AttrMaxLines returns maxLines attribute with value n.
func AttrMaxLines(n int) xml.Attr {
	return Attr("maxLines", strconv.Itoa(n))
}
