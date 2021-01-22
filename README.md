[![Go Report Card](https://goreportcard.com/badge/github.com/rzajac/kml)](https://goreportcard.com/report/github.com/rzajac/kml)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rzajac/kml)


# KML

Package `kml` provides methods for constructing and exploring KML files.

# Installation

```
go get github.com/rzajac/kml
```

# Usage

## Unmarshal, explore, edit, marshal

```
data, err := os.Open("example.kml")
checkErr(err)

root := KML()
err := xml.Unmarshal(data, root)
checkErr(err)

// Explore. See documentation for available methods. 
id := k.ChildAtIdx(0).Attribute("id").Value

// Edit.
root.ChildAtIdx(0).ChildAtIdx(2).ChildAtIdx(0).SetContent("new value")

// Marshal.
data, err := xml.Marshal(root)
checkErr(err)
```

## In place KML construction and writing.

```
// Create KML.
k := kml.KML(
    kml.Document(
        kml.Name("document name"),
        kml.Style(
            kml.AttrID("sty_0"),
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
```

will print:

```
<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2" xmlns:gx="http://www.google.com/kml/ext/2.2" xmlns:kml="http://www.opengis.net/kml/2.2" xmlns:atom="http://www.w3.org/2005/Atom">
 <Document>
  <name>document name</name>
  <Style id="sty_0">
   <LabelStyle>
    <color>01020304</color>
    <scale>1</scale>
   </LabelStyle>
   <LineStyle>
    <color>05060708</color>
    <width>3.5</width>
   </LineStyle>
   <PolyStyle>
    <color>090a0b0c</color>
    <outline>1</outline>
   </PolyStyle>
  </Style>
  <Folder id="fld_0">
   <name>folder name</name>
   <Snippet maxLines="1">snip</Snippet>
   <Placemark id="pm_0">
    <name>placemark name</name>
    <description><![CDATA[<b>placemark description</b>]]></description>
    <styleUrl>#sty_0</styleUrl>
    <MultiGeometry>
     <LineString>
      <tessellate>1</tessellate>
      <coordinates>0.1,0.2,0.3 1.1,1.2,1.3</coordinates>
     </LineString>
    </MultiGeometry>
   </Placemark>
  </Folder>
 </Document>
</kml>
```

## TODO

- Add missing methods for in place KML construction.

## License

Apache License, Version 2.0
