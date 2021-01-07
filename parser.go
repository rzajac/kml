package kml

import (
	"encoding/xml"
	"errors"
	"io"
)

// ErrInvalidKML is returned when invalid KML was passed to the Parser.
var ErrInvalidKML = errors.New("not KML")

// ErrUnexpectedElement is returned from UnmarshalXML when unexpected
// start element is passed.
var ErrUnexpectedElement = errors.New("unexpected element")

// Parse parses KML and returns instance of Parser.
func Parse(r io.Reader) (*Element, error) {
	dec := xml.NewDecoder(r)

	kml := KML()
	if err := dec.Decode(kml); err != nil {
		return nil, err
	}

	if kml.se.Name.Local != ElemKML {
		return nil, ErrInvalidKML
	}

	return kml, nil
}
