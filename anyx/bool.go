package anyx

import (
	"bytes"
	"math"
	"strings"
)

// ToBool 将输入值转换为布尔类型.
//
// 支持的类型包括:
// - bool: 直接返回.
// - 整型 (int, uint): 非 0 则为 true.
// - 浮点型 (float32, float64): 非 NaN 且非 0.0 则为 true.
// - 字符串 (string) / 字节切片 ([]byte):
//   - "true", "1", "t", "y", "yes", "on" (不区分大小写) -> true
//   - "false", "0", "f", "n", "no", "off" (不区分大小写) -> false
//   - 其他非空(trim后)字符串 -> true
//
// - 其他类型: false
func ToBool(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		return x
	case int:
		return x != 0
	case int8:
		return x != 0
	case int16:
		return x != 0
	case int32:
		return x != 0
	case int64:
		return x != 0
	case uint:
		return x != 0
	case uint8:
		return x != 0
	case uint16:
		return x != 0
	case uint32:
		return x != 0
	case uint64:
		return x != 0
	case float32:
		return !math.IsNaN(float64(x)) && x != 0.0
	case float64:
		return !math.IsNaN(x) && x != 0.0
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
		return false
	}
}
