package json

import (
	"bytes"
	"encoding/json"
)

func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}
