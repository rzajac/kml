package kml

import (
	"bytes"
	"encoding/xml"
	"testing"

	kit "github.com/rzajac/testkit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	fld := root.ChildAtIdx(0).ChildByName(ElemFolder)

	// --- Then ---
	assert.Exactly(t, "fld_0", fld.Attribute("id").Value)
}

func Test_Element_ChildByName_notExisting(t *testing.T) {
	// --- Given ---
	data := kit.ReadAll(t, kit.OpenFile(t, "testdata/example.kml"))
	root := KML()
	assert.NoError(t, xml.Unmarshal(data, root))

	// --- When ---
	fld := root.ChildAtIdx(0).ChildByName(ElemSnippet)

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

func Test_Element_PrependChild(t *testing.T) {
	// --- Given ---
	doc := Document(
		Name("name"),
		Description("desc"),
	)

	// --- When ---
	err := doc.PrependChild(Folder(AttrID("f1")), Folder(AttrID("f2")))

	// --- Then ---
	assert.NoError(t, err)

	data, err := xml.Marshal(doc)
	assert.NoError(t, err)

	exp := `<Document><Folder id="f1"></Folder><Folder id="f2"></Folder><name>name</name><description>desc</description></Document>`
	assert.Exactly(t, exp, string(data))
}

func Test_Element_RemoveChildren(t *testing.T) {
	// --- Given ---
	doc := Document(
		Name("name"),
		Description("desc"),
	)

	// --- When ---
	doc.RemoveChildren()

	// --- Then ---
	assert.Exactly(t, 0, doc.ChildCnt())
}

func Test_Element_RemoveChildAtIdx(t *testing.T) {
	tt := []struct {
		testN string

		elm *Element
		idx int
		rem string
		exp string
	}{
		{
			"start",
			Document(Name("name"), Description("desc"), Folder(AttrID("f1"))),
			0,
			ElemName,
			`<Document><description>desc</description><Folder id="f1"></Folder></Document>`,
		},
		{
			"middle",
			Document(Name("name"), Description("desc"), Folder(AttrID("f1"))),
			1,
			ElemDescription,
			`<Document><name>name</name><Folder id="f1"></Folder></Document>`,
		},
		{
			"end",
			Document(Name("name"), Description("desc"), Folder(AttrID("f1"))),
			2,
			ElemFolder,
			`<Document><name>name</name><description>desc</description></Document>`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			got := tc.elm.RemoveChildAtIdx(tc.idx)

			// --- Then ---
			require.NotNil(t, got)
			assert.Exactly(t, tc.rem, got.LocalName())

			data, err := xml.Marshal(tc.elm)
			assert.NoError(t, err)

			assert.Exactly(t, tc.exp, string(data))
		})
	}
}

func Test_Element_RemoveChildAtIdx_NotExisting(t *testing.T) {
	// --- Given ---
	doc := Document(
		Name("name"),
		Folder(AttrID("f1")),
	)

	// --- When ---
	got := doc.RemoveChildAtIdx(2)

	// --- Then ---
	require.Nil(t, got)

	data, err := xml.Marshal(doc)
	assert.NoError(t, err)

	exp := `<Document><name>name</name><Folder id="f1"></Folder></Document>`
	assert.Exactly(t, exp, string(data))
}

func Test_Element_ChildByID(t *testing.T) {
	// --- Given ---
	doc := Document(
		Name("name"),
		Folder(AttrID("f1")),
		Folder(AttrID("f2")),
	)

	// --- When ---
	ch := doc.ChildByID("f2")

	// --- Then ---
	require.NotNil(t, ch)
	assert.Exactly(t, "f2", ch.ID())
}
