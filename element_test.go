package kml

import (
	"bytes"
	"encoding/xml"
	"testing"

	kit "github.com/rzajac/testkit"
	"github.com/stretchr/testify/assert"
)

func Test_Element_Unmarshal_Marshal(t *testing.T) {
	// --- Given ---
	data := kit.ReadAll(t, kit.OpenFile(t, "testdata/example.kml"))

	// --- When ---
	root := KML()
	assert.NoError(t, xml.Unmarshal(data, root))

	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Indent("", "  ")
	err := enc.Encode(root)

	// --- Then ---
	assert.NoError(t, err)
	assert.Exactly(t, string(data), buf.String())
}

func Test_Element_ChildByName(t *testing.T) {
	// --- Given ---
	data := kit.ReadAll(t, kit.OpenFile(t, "testdata/example.kml"))
	root := KML()
	assert.NoError(t, xml.Unmarshal(data, root))

	// --- When ---
	fld := root.ChildByIdx(0).ChildByName(ElemFolder)

	// --- Then ---
	assert.Exactly(t, "fld_0", fld.Attribute("id").Value)
}

func Test_Element_ChildByName_notExisting(t *testing.T) {
	// --- Given ---
	data := kit.ReadAll(t, kit.OpenFile(t, "testdata/example.kml"))
	root := KML()
	assert.NoError(t, xml.Unmarshal(data, root))

	// --- When ---
	fld := root.ChildByIdx(0).ChildByName(ElemSnippet)

	// --- Then ---
	assert.Nil(t, fld)
}

func Test_Element_AddChild(t *testing.T) {
	// --- Given ---
	doc := Document()

	// --- When ---
	err := doc.AddChild(Name("name"), Description("desc"))

	// --- Then ---
	assert.NoError(t, err)

	data, err := xml.Marshal(doc)
	assert.NoError(t, err)

	exp := `<Document><name>name</name><description>desc</description></Document>`
	assert.Exactly(t, exp, string(data))
}
