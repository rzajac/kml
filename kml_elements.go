package kml

import (
	"encoding/xml"
	"strconv"
)

// KML element names.
const (
	ElemAltitudeMode  = "altitudeMode"
	ElemBalloonStyle  = "BalloonStyle"
	ElemBgColor       = "bgColor"
	ElemCamera        = "camera"
	ElemColor         = "color"
	ElemCoordinates   = "coordinates"
	ElemData          = "data"
	ElemDescription   = "description"
	ElemDisplayMode   = "displayMode"
	ElemDocument      = "Document"
	ElemExtendedData  = "ExtendedData"
	ElemFolder        = "Folder"
	ElemHeading       = "heading"
	ElemKML           = "kml"
	ElemLabelStyle    = "LabelStyle"
	ElemLatitude      = "latitude"
	ElemLineStyle     = "LineStyle"
	ElemLineString    = "LineString"
	ElemLongitude     = "longitude"
	ElemMultiGeometry = "MultiGeometry"
	ElemName          = "name"
	ElemOutline       = "outline"
	ElemPlacemark     = "Placemark"
	ElemPolyStyle     = "PolyStyle"
	ElemRoll          = "roll"
	ElemScale         = "scale"
	ElemSnippet       = "Snippet"
	ElemStyle         = "Style"
	ElemStyleURL      = "styleUrl"
	ElemTessellate    = "tessellate"
	ElemText          = "text"
	ElemTilt          = "tilt"
	ElemWidth         = "width"
)

// ----------------------------------- A ---------------------------------------

// AltitudeMode mode valid values.
const (
	AltitudeMoreRel = "relativeToGround"
	AltitudeMoreCla = "clampToGround"
	AltitudeMoreAbs = "absolute"
)

// AltitudeMode returns new altitudeMode element.
func AltitudeMode(value string, xes ...interface{}) *Element {
	return StringElement(ElemAltitudeMode, value, xes...)
}

// ----------------------------------- B ---------------------------------------

// BalloonStyle returns new BalloonStyle element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#balloonstyle
func BalloonStyle(xes ...interface{}) *Element {
	return NewElement(ElemBalloonStyle, xes...)
}

// BgColor returns new bgColor element.
func BgColor(value string, xes ...interface{}) *Element {
	return StringElement(ElemBgColor, value, xes...)
}

// ----------------------------------- C ---------------------------------------

// Camera returns new camera element.
func Camera(xes ...interface{}) *Element {
	return NewElement(ElemCamera, xes...)
}

// Color returns new color element.
func Color(value string, xes ...interface{}) *Element {
	return StringElement(ElemColor, value, xes...)
}

// Coordinates returns new coordinates element.
func Coordinates(value string, xes ...interface{}) *Element {
	return StringElement(ElemCoordinates, value, xes...)
}

// ----------------------------------- D ---------------------------------------

// Data returns new data element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#extendeddata
func Data(name string, xes ...interface{}) *Element {
	return NewElement(
		ElemData,
		append([]interface{}{Attr("name", name)}, xes...)...,
	)
}

// Description returns new description element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#feature
func Description(value string, xes ...interface{}) *Element {
	return StringElement(ElemDescription, value, xes...)
}

// Display mode valid values.
const (
	DisplayModeDefault = "default"
	DisplayModeHide    = "hide"
)

// DisplayMode returns new displayMode element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#balloonstyle
func DisplayMode(value string, xes ...interface{}) *Element {
	return StringElement(ElemDisplayMode, value, xes...)
}

// Document returns new Document element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#document
func Document(xes ...interface{}) *Element {
	return NewElement(ElemDocument, xes...)
}

// ----------------------------------- E ---------------------------------------

// ExtendedData returns new ExtendedData element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#extendeddata
func ExtendedData(xes ...interface{}) *Element {
	return NewElement(ElemExtendedData, xes...)
}

// ----------------------------------- F ---------------------------------------

// Folder returns new Folder element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#folder
func Folder(xes ...interface{}) *Element {
	return NewElement(ElemFolder, xes...)
}

// ----------------------------------- G ---------------------------------------
// ----------------------------------- H ---------------------------------------

// Heading returns new heading element.
func Heading(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemHeading, value, xes...)
}

// ----------------------------------- I ---------------------------------------
// ----------------------------------- J ---------------------------------------
// ----------------------------------- K ---------------------------------------

// KML returns a new kml element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#kml
func KML(xes ...interface{}) *Element {
	xel := NewElement("kml", xes...)
	xel.se.Name.Space = "http://www.opengis.net/kml/2.2"

	xel.SetAttribute(xml.Attr{
		Name:  xml.Name{Local: "xmlns:gx"},
		Value: "http://www.google.com/kml/ext/2.2",
	})

	xel.SetAttribute(xml.Attr{
		Name:  xml.Name{Local: "xmlns:kml"},
		Value: "http://www.opengis.net/kml/2.2",
	})

	xel.SetAttribute(xml.Attr{
		Name:  xml.Name{Local: "xmlns:atom"},
		Value: "http://www.w3.org/2005/Atom",
	})

	return xel
}

// ----------------------------------- L ---------------------------------------

// LabelStyle returns new LabelStyle element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#labelstyle
func LabelStyle(xes ...interface{}) *Element {
	return NewElement(ElemLabelStyle, xes...)
}

// Latitude returns new latitude element.
func Latitude(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemLatitude, value, xes...)
}

// LineString returns new LineString element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#linestring
func LineString(xes ...interface{}) *Element {
	return NewElement(ElemLineString, xes...)
}

// LineStyle returns new LineStyle element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#linestyle
func LineStyle(xes ...interface{}) *Element {
	return NewElement(ElemLineStyle, xes...)
}

// Longitude returns new longitude element.
func Longitude(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemLongitude, value, xes...)
}

// ----------------------------------- M ---------------------------------------

// MultiGeometry returns new MultiGeometry element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#multigeometry
func MultiGeometry(xes ...interface{}) *Element {
	return NewElement(ElemMultiGeometry, xes...)
}

// ----------------------------------- N ---------------------------------------

// Name returns new Name element.
func Name(value string, xes ...interface{}) *Element {
	return StringElement(ElemName, value, xes...)
}

// ----------------------------------- O ---------------------------------------

// Outline returns new outline element.
func Outline(value bool, xes ...interface{}) *Element {
	return BoolElement(ElemOutline, value, xes...)
}

// ----------------------------------- P ---------------------------------------

// Placemark returns new Placemark element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#placemark
func Placemark(xes ...interface{}) *Element {
	return NewElement(ElemPlacemark, xes...)
}

// PolyStyle returns new PolyStyle element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#polystyle
func PolyStyle(xes ...interface{}) *Element {
	return NewElement(ElemPolyStyle, xes...)
}

// ----------------------------------- Q ---------------------------------------
// ----------------------------------- R ---------------------------------------

// Roll returns new roll element.
func Roll(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemRoll, value, xes...)
}

// ----------------------------------- S ---------------------------------------

// Scale returns new scale element.
func Scale(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemScale, value, xes...)
}

// Snippet returns new Snippet element.
func Snippet(value string, xes ...interface{}) *Element {
	return StringElement(ElemSnippet, value, xes...)
}

// Style returns new Style element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#style
func Style(xes ...interface{}) *Element {
	return NewElement(ElemStyle, xes...)
}

// StyleURL returns new styleURL element.
func StyleURL(value string, xes ...interface{}) *Element {
	return StringElement(ElemStyleURL, value, xes...)
}

// ----------------------------------- T ---------------------------------------

// Tessellate returns new tessellate element.
func Tessellate(value bool, xes ...interface{}) *Element {
	return BoolElement(ElemTessellate, value, xes...)
}

// Text returns new text element.
// Reference: https://developers.google.com/kml/documentation/kmlreference#balloonstyle
func Text(value string, xes ...interface{}) *Element {
	return StringElement(ElemText, value, xes...)
}

// Tilt returns new tilt element.
func Tilt(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemTilt, value, xes...)
}

// ----------------------------------- U ---------------------------------------
// ----------------------------------- V ---------------------------------------
// ----------------------------------- W ---------------------------------------

// Width returns new width element.
func Width(value float64, xes ...interface{}) *Element {
	return FloatElement(ElemWidth, value, xes...)
}

// ----------------------------------- X ---------------------------------------
// ----------------------------------- Y ---------------------------------------
// ----------------------------------- Z ---------------------------------------

// ------------------------------- helpers -------------------------------------

// StringElement returns new KML element with name and content.
func StringElement(name, content string, xes ...interface{}) *Element {
	xel := NewElement(name, xes...)
	xel.content = []byte(content)
	return xel
}

// IntElement returns new KML element with name and value.
func IntElement(name string, value int, xes ...interface{}) *Element {
	xel := NewElement(name, xes...)
	xel.content = []byte(strconv.Itoa(value))
	return xel
}

// FloatElement returns new KML element with name and value.
func FloatElement(name string, value float64, xes ...interface{}) *Element {
	xel := NewElement(name, xes...)
	xel.content = []byte(strconv.FormatFloat(value, 'f', -1, 64))
	return xel
}

// BoolElement returns new KML element with name and value.
func BoolElement(name string, value bool, xes ...interface{}) *Element {
	xel := NewElement(name, xes...)
	if value {
		xel.content = []byte("1")
	} else {
		xel.content = []byte("0")
	}
	return xel
}
