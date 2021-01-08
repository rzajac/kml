package internal

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

// TODO
func xmllint(t *testing.T, s string) {
	cmd := exec.Command("xmllint", "--schema", "xsd/kml22gx.xsd", "--schema", "xsd/ogckml22.xsd", "-")
	cmd.Stdin = strings.NewReader(s)

	out := &bytes.Buffer{}
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = out

	if err := cmd.Run(); err != nil {
		// t.Fatal(err)
	}
	t.Log(out.String())
}
