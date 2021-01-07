package kml

import (
	"bytes"
)

// needsCDATA returns true if s needs to be wrapped in CDATA directive.
func needsCDATA(s []byte) bool {
	return bytes.ContainsAny(s, "<>")
}
