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
	root, err := Parse(f)

	// --- Then ---
	assert.NoError(t, err)

	// kml.
	assert.Exactly(t, "kml", root.LocalName())
	assert.Exactly(t, int64(0), root.Offset())
	assert.Exactly(t, 3, root.AttributeCnt())
	assert.Exactly(t, 1, root.ChildCnt())
	assert.True(t, root.HasChild(ElemDocument))

	assert.True(t, root.HasAttribute("xmlns:kml"))
	assert.Exactly(t, "xmlns:kml", root.Attribute("xmlns:kml").Name.Local)
	assert.Exactly(t, "", root.Attribute("xmlns:gx").Name.Space)
	assert.Exactly(t, "http://www.opengis.net/kml/2.2", root.Attribute("xmlns:kml").Value)

	assert.True(t, root.HasAttribute("xmlns:gx"))
	assert.Exactly(t, "xmlns:gx", root.Attribute("xmlns:gx").Name.Local)
	assert.Exactly(t, "", root.Attribute("xmlns:gx").Name.Space)
	assert.Exactly(t, "http://www.google.com/kml/ext/2.2", root.Attribute("xmlns:gx").Value)

	assert.True(t, root.HasAttribute("xmlns:atom"))
	assert.Exactly(t, "xmlns:atom", root.Attribute("xmlns:atom").Name.Local)
	assert.Exactly(t, "", root.Attribute("xmlns:gx").Name.Space)
	assert.Exactly(t, "http://www.w3.org/2005/Atom", root.Attribute("xmlns:atom").Value)

	// kml > Document.
	doc := root.ChildAtIdx(0)
	assert.Exactly(t, ElemDocument, doc.LocalName())
	assert.Exactly(t, int64(215), doc.Offset())
	assert.Exactly(t, 0, doc.AttributeCnt())
	assert.Exactly(t, 3, doc.ChildCnt())
	assert.True(t, doc.HasChild(ElemName))
	assert.True(t, doc.HasChild(ElemStyle))
	assert.True(t, doc.HasChild(ElemFolder))

	// kml > Document > name.
	name := doc.ChildAtIdx(0)
	assert.Exactly(t, ElemName, name.LocalName())
	assert.Exactly(t, int64(230), name.Offset())
	assert.Exactly(t, 0, name.AttributeCnt())
	assert.Exactly(t, 0, name.ChildCnt())
	assert.Exactly(t, "document name", name.Content())

	// kml > Document > Style.
	style := doc.ChildAtIdx(1)
	assert.Exactly(t, ElemStyle, style.LocalName())
	assert.Exactly(t, int64(261), style.Offset())
	assert.Exactly(t, 1, style.AttributeCnt())
	assert.Exactly(t, 3, style.ChildCnt())
	assert.True(t, style.HasChild(ElemLabelStyle))
	assert.True(t, style.HasChild(ElemLineStyle))
	assert.True(t, style.HasChild(ElemPolyStyle))

	assert.Exactly(t, "sty_0", style.Attribute("id").Value)

	// kml > Document > Style > LabelStyle.
	labSty := style.ChildAtIdx(0)
	assert.Exactly(t, ElemLabelStyle, labSty.LocalName())
	assert.Exactly(t, int64(286), labSty.Offset())
	assert.Exactly(t, 0, labSty.AttributeCnt())
	assert.Exactly(t, 2, labSty.ChildCnt())
	assert.True(t, labSty.HasChild(ElemColor))
	assert.True(t, labSty.HasChild(ElemScale))

	// kml > Document > Style > LabelStyle > color.
	color := labSty.ChildAtIdx(0)
	assert.Exactly(t, ElemColor, color.LocalName())
	assert.Exactly(t, int64(307), color.Offset())
	assert.Exactly(t, 0, color.AttributeCnt())
	assert.Exactly(t, 0, color.ChildCnt())
	assert.Exactly(t, "01020304", color.Content())

	// kml > Document > Style > LabelStyle > scale.
	scale := labSty.ChildAtIdx(1)
	assert.Exactly(t, ElemScale, scale.LocalName())
	assert.Exactly(t, int64(339), scale.Offset())
	assert.Exactly(t, 0, scale.AttributeCnt())
	assert.Exactly(t, 0, scale.ChildCnt())
	assert.Exactly(t, "1", scale.Content())

	// kml > Document > Style > LineStyle.
	linSty := style.ChildAtIdx(1)
	assert.Exactly(t, ElemLineStyle, linSty.LocalName())
	assert.Exactly(t, int64(382), linSty.Offset())
	assert.Exactly(t, 0, linSty.AttributeCnt())
	assert.Exactly(t, 2, linSty.ChildCnt())
	assert.True(t, linSty.HasChild(ElemColor))
	assert.True(t, linSty.HasChild(ElemWidth))

	// kml > Document > Style > LineStyle > color.
	color = linSty.ChildAtIdx(0)
	assert.Exactly(t, ElemColor, color.LocalName())
	assert.Exactly(t, int64(382), linSty.Offset())
	assert.Exactly(t, 0, color.AttributeCnt())
	assert.Exactly(t, 0, color.ChildCnt())
	assert.Exactly(t, "05060708", color.Content())

	// kml > Document > Style > LineStyle > width.
	width := linSty.ChildAtIdx(1)
	assert.Exactly(t, ElemWidth, width.LocalName())
	assert.Exactly(t, int64(434), width.Offset())
	assert.Exactly(t, 0, width.AttributeCnt())
	assert.Exactly(t, 0, width.ChildCnt())
	assert.Exactly(t, "3.5", width.Content())

	// kml > Document > Style > PolyStyle.
	polSty := style.ChildAtIdx(2)
	assert.Exactly(t, ElemPolyStyle, polSty.LocalName())
	assert.Exactly(t, int64(478), polSty.Offset())
	assert.Exactly(t, 0, polSty.AttributeCnt())
	assert.Exactly(t, 2, polSty.ChildCnt())
	assert.True(t, polSty.HasChild(ElemColor))
	assert.True(t, polSty.HasChild(ElemOutline))

	// kml > Document > Style > PolyStyle > color.
	color = polSty.ChildAtIdx(0)
	assert.Exactly(t, ElemColor, color.LocalName())
	assert.Exactly(t, int64(498), color.Offset())
	assert.Exactly(t, 0, color.AttributeCnt())
	assert.Exactly(t, 0, color.ChildCnt())
	assert.Exactly(t, "090a0b0c", color.Content())

	// kml > Document > Style > PolyStyle > outline.
	outline := polSty.ChildAtIdx(1)
	assert.Exactly(t, ElemOutline, outline.LocalName())
	assert.Exactly(t, int64(530), outline.Offset())
	assert.Exactly(t, 0, outline.AttributeCnt())
	assert.Exactly(t, 0, outline.ChildCnt())
	assert.Exactly(t, "1", outline.Content())

	// kml > Document > Folder.
	folder := doc.ChildAtIdx(2)
	assert.Exactly(t, ElemFolder, folder.LocalName())
	assert.Exactly(t, int64(587), folder.Offset())
	assert.Exactly(t, 1, folder.AttributeCnt())
	assert.Exactly(t, 3, folder.ChildCnt())
	assert.True(t, folder.HasChild(ElemName))
	assert.True(t, folder.HasChild(ElemSnippet))
	assert.True(t, folder.HasChild(ElemPlacemark))

	assert.Exactly(t, "fld_0", folder.Attribute("id").Value)

	// kml > Document > Folder > name.
	name = folder.ChildAtIdx(0)
	assert.Exactly(t, ElemName, name.LocalName())
	assert.Exactly(t, int64(613), name.Offset())
	assert.Exactly(t, 0, name.AttributeCnt())
	assert.Exactly(t, 0, name.ChildCnt())
	assert.Exactly(t, "folder name", name.Content())

	// kml > Document > Folder > Snippet.
	snippet := folder.ChildAtIdx(1)
	assert.Exactly(t, ElemSnippet, snippet.LocalName())
	assert.Exactly(t, int64(644), snippet.Offset())
	assert.Exactly(t, 1, snippet.AttributeCnt())
	assert.Exactly(t, 0, snippet.ChildCnt())
	assert.Exactly(t, "snip", snippet.Content())

	assert.Exactly(t, "1", snippet.Attribute("maxLines").Value)

	// kml > Document > Folder > Placemark
	pm := folder.ChildAtIdx(2)
	assert.Exactly(t, ElemPlacemark, pm.LocalName())
	assert.Exactly(t, int64(687), pm.Offset())
	assert.Exactly(t, 1, pm.AttributeCnt())
	assert.Exactly(t, "pm_0", pm.Attribute("id").Value)
	assert.Exactly(t, 4, pm.ChildCnt())

	// kml > Document > Folder > Placemark > name.
	name = pm.ChildAtIdx(0)
	assert.Exactly(t, ElemName, name.LocalName())
	assert.Exactly(t, int64(717), name.Offset())
	assert.Exactly(t, 0, name.AttributeCnt())
	assert.Exactly(t, 0, name.ChildCnt())
	assert.Exactly(t, "placemark name", name.Content())

	// kml > Document > Folder > Placemark > description.
	desc := pm.ChildAtIdx(1)
	assert.Exactly(t, ElemDescription, desc.LocalName())
	assert.Exactly(t, int64(753), desc.Offset())
	assert.Exactly(t, 0, desc.AttributeCnt())
	assert.Exactly(t, 0, desc.ChildCnt())
	assert.Exactly(t, "<b>placemark description</b>", desc.Content())

	// kml > Document > Folder > Placemark > styleUrl.
	styleUrl := pm.ChildAtIdx(2)
	assert.Exactly(t, ElemStyleURL, styleUrl.LocalName())
	assert.Exactly(t, int64(829), styleUrl.Offset())
	assert.Exactly(t, 0, styleUrl.AttributeCnt())
	assert.Exactly(t, 0, styleUrl.ChildCnt())
	assert.Exactly(t, "#sty_0", styleUrl.Content())

	// kml > Document > Folder > Placemark > MultiGeometry.
	mg := pm.ChildAtIdx(3)
	assert.Exactly(t, ElemMultiGeometry, mg.LocalName())
	assert.Exactly(t, int64(865), mg.Offset())
	assert.Exactly(t, 0, mg.AttributeCnt())
	assert.Exactly(t, 1, mg.ChildCnt())
	assert.True(t, mg.HasChild(ElemLineString))

	// kml > Document > Folder > Placemark > MultiGeometry > LineString.
	linStr := mg.ChildAtIdx(0)
	assert.Exactly(t, ElemLineString, linStr.LocalName())
	assert.Exactly(t, int64(891), linStr.Offset())
	assert.Exactly(t, 0, linStr.AttributeCnt())
	assert.Exactly(t, 2, linStr.ChildCnt())
	assert.True(t, linStr.HasChild(ElemTessellate))
	assert.True(t, linStr.HasChild(ElemCoordinates))

	// kml > Document > Folder > Placemark > MultiGeometry > LineString > tessellate.
	ts := linStr.ChildAtIdx(0)
	assert.Exactly(t, ElemTessellate, ts.LocalName())
	assert.Exactly(t, int64(916), ts.Offset())
	assert.Exactly(t, 0, ts.AttributeCnt())
	assert.Exactly(t, 0, ts.ChildCnt())
	assert.Exactly(t, "1", ts.Content())

	// kml > Document > Folder > Placemark > MultiGeometry > LineString > coordinates.
	cor := linStr.ChildAtIdx(1)
	assert.Exactly(t, ElemCoordinates, cor.LocalName())
	assert.Exactly(t, int64(955), cor.Offset())
	assert.Exactly(t, 0, cor.AttributeCnt())
	assert.Exactly(t, 0, cor.ChildCnt())
	assert.Exactly(t, "0.1,0.2,0.3 1.1,1.2,1.3", cor.Content())
}
