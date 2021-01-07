package kml

import (
	"testing"

	kit "github.com/rzajac/testkit"
	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	// --- Given ---
	f := kit.OpenFile(t, "testdata/example.kml")

	// --- When ---
	k, err := Parse(f)

	// --- Then ---
	assert.NoError(t, err)

	// kml.
	assert.Exactly(t, "kml", k.LocalName())

	assert.True(t, k.HasAttribute("xmlns:kml"))
	assert.Exactly(t, "xmlns:kml", k.Attribute("xmlns:kml").Name.Local)
	assert.Exactly(t, "", k.Attribute("xmlns:gx").Name.Space)
	assert.Exactly(t, "http://www.opengis.net/kml/2.2", k.Attribute("xmlns:kml").Value)

	assert.True(t, k.HasAttribute("xmlns:gx"))
	assert.Exactly(t, "xmlns:gx", k.Attribute("xmlns:gx").Name.Local)
	assert.Exactly(t, "", k.Attribute("xmlns:gx").Name.Space)
	assert.Exactly(t, "http://www.google.com/kml/ext/2.2", k.Attribute("xmlns:gx").Value)

	assert.True(t, k.HasAttribute("xmlns:atom"))
	assert.Exactly(t, "xmlns:atom", k.Attribute("xmlns:atom").Name.Local)
	assert.Exactly(t, "", k.Attribute("xmlns:gx").Name.Space)
	assert.Exactly(t, "http://www.w3.org/2005/Atom", k.Attribute("xmlns:atom").Value)

	assert.True(t, k.HasChild(ElemDocument))
	assert.Exactly(t, 1, k.ChildCnt())

	// kml > Document.
	doc := k.Child(0)
	assert.Exactly(t, ElemDocument, doc.LocalName())
	assert.Exactly(t, 0, doc.AttributeCnt())
	assert.Exactly(t, 3, doc.ChildCnt())

	// kml > Document > name.
	name := doc.Child(0)
	assert.Exactly(t, ElemName, name.LocalName())
	assert.Exactly(t, 0, name.AttributeCnt())
	assert.Exactly(t, 0, name.ChildCnt())
	assert.Exactly(t, "document name", name.Content())

	// kml > Document > Style.
	style := doc.Child(1)
	assert.Exactly(t, ElemStyle, style.LocalName())
	assert.Exactly(t, 1, style.AttributeCnt())
	assert.Exactly(t, "sty_0", style.Attribute("id").Value)
	assert.Exactly(t, 3, style.ChildCnt())

	// kml > Document > Style > LabelStyle.
	ls := style.Child(0)
	assert.Exactly(t, ElemLabelStyle, ls.LocalName())
	assert.Exactly(t, 0, ls.AttributeCnt())
	assert.Exactly(t, 2, ls.ChildCnt())

	// kml > Document > Style > LabelStyle > color.
	color := ls.Child(0)
	assert.Exactly(t, ElemColor, color.LocalName())
	assert.Exactly(t, 0, color.AttributeCnt())
	assert.Exactly(t, 0, color.ChildCnt())
	assert.Exactly(t, "01020304", color.Content())

	// kml > Document > Style > LabelStyle > scale.
	scale := ls.Child(1)
	assert.Exactly(t, ElemScale, scale.LocalName())
	assert.Exactly(t, 0, scale.AttributeCnt())
	assert.Exactly(t, 0, scale.ChildCnt())
	assert.Exactly(t, "1", scale.Content())

	// kml > Document > Style > LineStyle.
	ls = style.Child(1)
	assert.Exactly(t, ElemLineStyle, ls.LocalName())
	assert.Exactly(t, 0, ls.AttributeCnt())
	assert.Exactly(t, 2, ls.ChildCnt())

	// kml > Document > Style > LineStyle > color.
	color = ls.Child(0)
	assert.Exactly(t, ElemColor, color.LocalName())
	assert.Exactly(t, 0, color.AttributeCnt())
	assert.Exactly(t, 0, color.ChildCnt())
	assert.Exactly(t, "05060708", color.Content())

	// kml > Document > Style > LineStyle > width.
	width := ls.Child(1)
	assert.Exactly(t, ElemWidth, width.LocalName())
	assert.Exactly(t, 0, width.AttributeCnt())
	assert.Exactly(t, 0, width.ChildCnt())
	assert.Exactly(t, "3.5", width.Content())

	// kml > Document > Style > PolyStyle.
	pl := style.Child(2)
	assert.Exactly(t, ElemPolyStyle, pl.LocalName())
	assert.Exactly(t, 0, pl.AttributeCnt())
	assert.Exactly(t, 2, pl.ChildCnt())

	// kml > Document > Style > PolyStyle > color.
	color = pl.Child(0)
	assert.Exactly(t, ElemColor, color.LocalName())
	assert.Exactly(t, 0, color.AttributeCnt())
	assert.Exactly(t, 0, color.ChildCnt())
	assert.Exactly(t, "090a0b0c", color.Content())

	// kml > Document > Style > PolyStyle > outline.
	outline := pl.Child(1)
	assert.Exactly(t, ElemOutline, outline.LocalName())
	assert.Exactly(t, 0, outline.AttributeCnt())
	assert.Exactly(t, 0, outline.ChildCnt())
	assert.Exactly(t, "1", outline.Content())

	// kml > Document > Folder.
	folder := doc.Child(2)
	assert.Exactly(t, ElemFolder, folder.LocalName())
	assert.Exactly(t, 1, folder.AttributeCnt())
	assert.Exactly(t, "fld_0", folder.Attribute("id").Value)
	assert.Exactly(t, 3, folder.ChildCnt())

	// kml > Document > Folder > name.
	name = folder.Child(0)
	assert.Exactly(t, ElemName, name.LocalName())
	assert.Exactly(t, 0, name.AttributeCnt())
	assert.Exactly(t, 0, name.ChildCnt())
	assert.Exactly(t, "folder name", name.Content())

	// kml > Document > Folder > Snippet.
	snippet := folder.Child(1)
	assert.Exactly(t, ElemSnippet, snippet.LocalName())
	assert.Exactly(t, 1, snippet.AttributeCnt())
	assert.Exactly(t, "1", snippet.Attribute("maxLines").Value)
	assert.Exactly(t, 0, snippet.ChildCnt())

	// kml > Document > Folder > Placemark
	pm := folder.Child(2)
	assert.Exactly(t, ElemPlacemark, pm.LocalName())
	assert.Exactly(t, 1, pm.AttributeCnt())
	assert.Exactly(t, "pm_0", pm.Attribute("id").Value)
	assert.Exactly(t, 4, pm.ChildCnt())

	// kml > Document > Folder > Placemark > name.
	name = pm.Child(0)
	assert.Exactly(t, ElemName, name.LocalName())
	assert.Exactly(t, 0, name.AttributeCnt())
	assert.Exactly(t, 0, name.ChildCnt())
	assert.Exactly(t, "placemark name", name.Content())

	// kml > Document > Folder > Placemark > description.
	desc := pm.Child(1)
	assert.Exactly(t, ElemDescription, desc.LocalName())
	assert.Exactly(t, 0, desc.AttributeCnt())
	assert.Exactly(t, 0, desc.ChildCnt())
	assert.Exactly(t, "<b>placemark description</b>", desc.Content())

	// kml > Document > Folder > Placemark > styleUrl.
	styleUrl := pm.Child(2)
	assert.Exactly(t, ElemStyleURL, styleUrl.LocalName())
	assert.Exactly(t, 0, styleUrl.AttributeCnt())
	assert.Exactly(t, 0, styleUrl.ChildCnt())
	assert.Exactly(t, "#sty_0", styleUrl.Content())

	// kml > Document > Folder > Placemark > MultiGeometry.
	mg := pm.Child(3)
	assert.Exactly(t, ElemMultiGeometry, mg.LocalName())
	assert.Exactly(t, 0, mg.AttributeCnt())
	assert.Exactly(t, 1, mg.ChildCnt())

	// kml > Document > Folder > Placemark > MultiGeometry > LineString.
	ls = mg.Child(0)
	assert.Exactly(t, ElemLineString, ls.LocalName())
	assert.Exactly(t, 0, ls.AttributeCnt())
	assert.Exactly(t, 2, ls.ChildCnt())

	// kml > Document > Folder > Placemark > MultiGeometry > LineString > tessellate.
	ts := ls.Child(0)
	assert.Exactly(t, ElemTessellate, ts.LocalName())
	assert.Exactly(t, 0, ts.AttributeCnt())
	assert.Exactly(t, 0, ts.ChildCnt())
	assert.Exactly(t, "1", ts.Content())

	// kml > Document > Folder > Placemark > MultiGeometry > LineString > coordinates.
	cor := ls.Child(1)
	assert.Exactly(t, ElemCoordinates, cor.LocalName())
	assert.Exactly(t, 0, cor.AttributeCnt())
	assert.Exactly(t, 0, cor.ChildCnt())
	assert.Exactly(t, "0.1,0.2,0.3 1.1,1.2,1.3", cor.Content())
}
