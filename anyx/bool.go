package anyx

import (
	"bytes"
	"strings"
)

func ToBool(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		return x
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return x != 0
	case string:
		switch strings.ToLower(x) {
		case "true", "1", "t", "y", "yes", "on":
			return true
		case "false", "0", "f", "n", "no", "off":
			return false
		default:
			return strings.TrimSpace(x) != ""
		}
	case []byte:
		switch string(bytes.ToLower(x)) {
		case "true", "1", "t", "y", "yes", "on":
			return true
		case "false", "0", "f", "n", "no", "off":
			return false
		default:
			return len(bytes.TrimSpace(x)) != 0
		}
	default:
		return val == nil
	}
}
