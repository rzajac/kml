package kml_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rzajac/kml"
)

func Test_InPlaceElement_Rendering(t *testing.T) {
	tt := []struct {
		elm *kml.Element
		exp string
	}{
		{kml.Altitude(1.234), `<altitude>1.234</altitude>`},
		{kml.BalloonStyle(), `<BalloonStyle></BalloonStyle>`},
		{kml.BgColor("ffffffff"), `<bgColor>ffffffff</bgColor>`},
		{kml.Camera(), `<Camera></Camera>`},
		{kml.Color("ffffffff"), `<color>ffffffff</color>`},
		{kml.Coordinates("0.1,0.2,0.3 1.1,1.2,1.3"), `<coordinates>0.1,0.2,0.3 1.1,1.2,1.3</coordinates>`},
		{kml.Data("name"), `<data name="name"></data>`},
		{kml.Description("desc"), `<description>desc</description>`},
		{kml.DisplayName("name"), `<displayName>name</displayName>`},
		{kml.Document(), `<Document></Document>`},
		{kml.ExtendedData(), `<ExtendedData></ExtendedData>`},
		{kml.Folder(), `<Folder></Folder>`},
		{kml.GxOption("sunlight", true), `<gx:option name="sunlight" enabled="1"></gx:option>`},
		{kml.GxTimeStamp(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)), `<gx:TimeStamp><when>2020-01-01T00:00:00Z</when></gx:TimeStamp>`},
		{kml.GxViewerOptions(), `<gx:ViewerOptions></gx:ViewerOptions>`},
		{kml.Heading(1.234), `<heading>1.234</heading>`},
		{kml.LabelStyle(), `<LabelStyle></LabelStyle>`},
		{kml.Latitude(1.234), `<latitude>1.234</latitude>`},
		{kml.LineString(), `<LineString></LineString>`},
		{kml.Longitude(1.234), `<longitude>1.234</longitude>`},
		{kml.MultiGeometry(), `<MultiGeometry></MultiGeometry>`},
		{kml.Name("value"), `<name>value</name>`},
		{kml.Outline(true), `<outline>1</outline>`},
		{kml.Placemark(), `<Placemark></Placemark>`},
		{kml.PolyStyle(), `<PolyStyle></PolyStyle>`},
		{kml.Roll(1.234), `<roll>1.234</roll>`},
		{kml.Scale(1.234), `<scale>1.234</scale>`},
		{kml.Schema("id", "name"), `<Schema name="name" id="id"></Schema>`},
		{kml.SchemaData("#schema"), `<SchemaData schemaUrl="#schema"></SchemaData>`},
		{kml.SimpleData("name", "value"), `<SimpleData name="name">value</SimpleData>`},
		{kml.SimpleField(kml.SFTypeString, "name"), `<SimpleField type="string" name="name"></SimpleField>`},
		{kml.Snippet("value"), `<Snippet>value</Snippet>`},
		{kml.Style("sty_id"), `<Style id="sty_id"></Style>`},
		{kml.StyleURL("#value"), `<styleUrl>#value</styleUrl>`},
		{kml.Tessellate(false), `<tessellate>0</tessellate>`},
		{kml.Text("value"), `<text>value</text>`},
		{kml.Tilt(1.234), `<tilt>1.234</tilt>`},
		{kml.Width(1.234), `<width>1.234</width>`},
	}

	for _, tc := range tt {
		t.Run(tc.elm.LocalName(), func(t *testing.T) {
			// --- When ---
			data, err := xml.Marshal(tc.elm)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_InPlaceElement_SimpleField(t *testing.T) {
	tt := []struct {
		typ  string
		name string
		exp  string
	}{
		{kml.SFTypeString, "name1", `<SimpleField type="string" name="name1"></SimpleField>`},
		{kml.SFTypeInt, "name2", `<SimpleField type="int" name="name2"></SimpleField>`},
		{kml.SFTypeUInt, "name3", `<SimpleField type="uint" name="name3"></SimpleField>`},
		{kml.SFTypeShort, "name4", `<SimpleField type="short" name="name4"></SimpleField>`},
		{kml.SFTypeUShort, "name5", `<SimpleField type="ushort" name="name5"></SimpleField>`},
		{kml.SFTypeFloat, "name6", `<SimpleField type="float" name="name6"></SimpleField>`},
		{kml.SFTypeDouble, "name7", `<SimpleField type="double" name="name7"></SimpleField>`},
		{kml.SFTypeBool, "name8", `<SimpleField type="bool" name="name8"></SimpleField>`},
	}

	for _, tc := range tt {
		t.Run(tc.typ, func(t *testing.T) {
			// --- Given ---
			k := kml.SimpleField(tc.typ, tc.name)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_InPlaceElement_AltitudeMode(t *testing.T) {
	tt := []struct {
		value string
		exp   string
	}{
		{kml.AltitudeMoreRel, "<altitudeMode>relativeToGround</altitudeMode>"},
		{kml.AltitudeMoreCla, "<altitudeMode>clampToGround</altitudeMode>"},
		{kml.AltitudeMoreAbs, "<altitudeMode>absolute</altitudeMode>"},
	}

	for _, tc := range tt {
		t.Run(tc.value, func(t *testing.T) {
			// --- Given ---
			k := kml.AltitudeMode(tc.value)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_InPlaceElement_DisplayMode(t *testing.T) {
	tt := []struct {
		value string
		exp   string
	}{
		{kml.DisplayModeDefault, "<displayMode>default</displayMode>"},
		{kml.DisplayModeHide, "<displayMode>hide</displayMode>"},
	}

	for _, tc := range tt {
		t.Run(tc.value, func(t *testing.T) {
			// --- Given ---
			k := kml.DisplayMode(tc.value)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_InPlaceElement_KML(t *testing.T) {
	// --- When ---
	k := kml.KML()
	data, err := xml.Marshal(k)

	// --- Then ---
	assert.NoError(t, err)
	exp := `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom"></kml>`
	assert.Exactly(t, exp, string(data))
}

func Test_InPlaceElement_ComplexWithID(t *testing.T) {
	// --- Given ---
	k := kml.Placemark(
		kml.AttrID("123"),
		kml.Name("some name", kml.AttrID("456")),
	)

	// --- When ---
	data, err := xml.Marshal(k)

	// --- Then ---
	assert.NoError(t, err)
	exp := `<Placemark id="123"><name id="456">some name</name></Placemark>`
	assert.Exactly(t, exp, string(data))
}

func Test_StringElement(t *testing.T) {
	tt := []struct {
		name    string
		content string
		exp     string
	}{
		{"elem", "content", "<elem>content</elem>"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// --- Given ---
			k := kml.StringElement(tc.name, tc.content)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_IntElement(t *testing.T) {
	tt := []struct {
		name  string
		value int
		exp   string
	}{
		{"elem", 123, "<elem>123</elem>"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// --- Given ---
			k := kml.IntElement(tc.name, tc.value)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_FloatElement(t *testing.T) {
	tt := []struct {
		name  string
		value float64
		exp   string
	}{
		{"elem", 1.2300, "<elem>1.23</elem>"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// --- Given ---
			k := kml.FloatElement(tc.name, tc.value)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_BoolElement(t *testing.T) {
	tt := []struct {
		name  string
		value bool
		exp   string
	}{
		{"elemT", true, "<elemT>1</elemT>"},
		{"elemF", false, "<elemF>0</elemF>"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// --- Given ---
			k := kml.BoolElement(tc.name, tc.value)

			// --- When ---
			data, err := xml.Marshal(k)

			// --- Then ---
			assert.NoError(t, err)
			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func ExampleKML() {
	// Create KML.
	k := kml.KML(
		kml.Document(
			kml.Name("document name"),
			kml.Style(
				"sty_0",
				kml.LabelStyle(
					kml.Color("01020304"),
					kml.Scale(1),
				),
				kml.LineStyle(
					kml.Color("05060708"),
					kml.Width(3.5),
				),
				kml.PolyStyle(
					kml.Color("090a0b0c"),
					kml.Outline(true),
				),
			),
			kml.Folder(
				kml.AttrID("fld_0"),
				kml.Name("folder name"),
				kml.Snippet("snip", kml.AttrMaxLines(1)),
				kml.Placemark(
					kml.AttrID("pm_0"),
					kml.Name("placemark name"),
					kml.Description("<b>placemark description</b>"),
					kml.StyleURL("#sty_0"),
					kml.MultiGeometry(
						kml.LineString(
							kml.Tessellate(true),
							kml.Coordinates("0.1,0.2,0.3 1.1,1.2,1.3"),
						),
					),
				),
			),
		),
	)

	// Encode and print created KML.
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Indent("", " ")
	if err := enc.Encode(k); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom">
	//  <Document>
	//   <name>document name</name>
	//   <Style id="sty_0">
	//    <LabelStyle>
	//     <color>01020304</color>
	//     <scale>1</scale>
	//    </LabelStyle>
	//    <LineStyle>
	//     <color>05060708</color>
	//     <width>3.5</width>
	//    </LineStyle>
	//    <PolyStyle>
	//     <color>090a0b0c</color>
	//     <outline>1</outline>
	//    </PolyStyle>
	//   </Style>
	//   <Folder id="fld_0">
	//    <name>folder name</name>
	//    <Snippet maxLines="1">snip</Snippet>
	//    <Placemark id="pm_0">
	//     <name>placemark name</name>
	//     <description><![CDATA[<b>placemark description</b>]]></description>
	//     <styleUrl>#sty_0</styleUrl>
	//     <MultiGeometry>
	//      <LineString>
	//       <tessellate>1</tessellate>
	//       <coordinates>0.1,0.2,0.3 1.1,1.2,1.3</coordinates>
	//      </LineString>
	//     </MultiGeometry>
	//    </Placemark>
	//   </Folder>
	//  </Document>
	// </kml>
}
