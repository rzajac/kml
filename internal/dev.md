```
sudo apt install libxml2-utils
xmllint --format --valid in.kml > out.kml
xmllint --schema http://schemas.opengis.net/kml/2.2.0/ogckml22.xsd --format in.kml > out.km
```