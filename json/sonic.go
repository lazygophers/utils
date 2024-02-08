//go:build linux && amd64
// +build linux,amd64

package json

import (
	"github.com/bytedance/sonic"
	"io"
)

func Marshal(v any) ([]byte, error) {
	return sonic.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return sonic.Unmarshal(data, v)
}

func MarshalString(v any) (string, error) {
	buf, err := sonic.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func UnmarshalString(data string, v any) error {
	return sonic.Unmarshal([]byte(data), v)
}

func NewEncoder(w io.Writer) sonic.Encoder {
	return sonic.ConfigDefault.NewEncoder(w)
}

func NewDecoder(r io.Reader) sonic.Decoder {
	return sonic.ConfigDefault.NewDecoder(r)
}
