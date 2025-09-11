package candy

import (
	"strconv"
	"time"
)

// ToInt64 将任何类型的值尽力转换为 int64。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1, false 转换为 0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - time.Duration: 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为有符号整数，若解析失败则返回 0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0。
func ToInt64(val interface{}) int64 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	case uint:
		return int64(x)
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x)
	case time.Duration:
		return int64(x)
	case float32:
		return int64(x)
	case float64:
		return int64(x)
	case string:
		val, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return 0
		}
		return val
	case []byte:
		val, err := strconv.ParseInt(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}