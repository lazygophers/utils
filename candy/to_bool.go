package candy

import (
	"math"
	"strings"
)

// ToBool 尝试将任意类型 (interface{}) 的输入值转换为布尔值 (bool)。
// 此函数现在使用泛型实现，提供更好的性能和类型安全。
//
// 转换规则如下:
//
//   - **bool**:
//     直接返回原始值。
//
//   - **整型** (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64):
//     如果值为 0，则返回 false，否则返回 true。
//
//   - **浮点型** (float32, float64):
//     如果值为 0.0 或 NaN (Not-a-Number)，则返回 false，否则返回 true。
//
//   - **字符串 (string) 和字节切片 ([]byte)**:
//     首先会转换为小写并移除首尾空白字符。
//
//   - "true", "1", "t", "y", "yes", "on" 被视为 true。
//
//   - "false", "0", "f", "n", "no", "off" 被视为 false。
//
//   - 对于其他非空字符串，返回 true。
//
//   - 对于空字符串或仅包含空白字符的字符串，返回 false。
//
//   - **nil**:
//     返回 false。
//
//   - **其他所有类型**:
//     均返回 false (例如: struct, map, slice 等)。
//
// 示例:
//
//	candy.ToBool(true)    // true
//	candy.ToBool(0)       // false
//	candy.ToBool("yes")   // true
//	candy.ToBool("off")   // false
//	candy.ToBool("hello") // true
//	candy.ToBool(nil)     // false
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
		return x != 0 && !math.IsNaN(float64(x))
	case float64:
		return x != 0 && !math.IsNaN(x)
	case string:
		s := strings.ToLower(strings.TrimSpace(x))
		switch s {
		case "true", "1", "t", "y", "yes", "on":
			return true
		case "false", "0", "f", "n", "no", "off", "":
			return false
		default:
			return true
		}
	case []byte:
		s := strings.ToLower(strings.TrimSpace(string(x)))
		switch s {
		case "true", "1", "t", "y", "yes", "on":
			return true
		case "false", "0", "f", "n", "no", "off", "":
			return false
		default:
			return true
		}
	case nil:
		return false
	default:
		return false
	}
}
