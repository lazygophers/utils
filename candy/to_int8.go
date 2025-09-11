package candy

import "strconv"

// ToInt8 将任何类型的值尽力转换为 int8。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1, false 转换为 0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为无符号整数，若解析失败则返回 0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0。
func ToInt8(val interface{}) int8 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int8(x)
	case int8:
		return x
	case int16:
		return int8(x)
	case int32:
		return int8(x)
	case int64:
		return int8(x)
	case uint:
		return int8(x)
	case uint8:
		return int8(x)
	case uint16:
		return int8(x)
	case uint32:
		return int8(x)
	case uint64:
		return int8(x)
	case float32:
		return int8(x)
	case float64:
		return int8(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return int8(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return int8(val)
	default:
		return 0
	}
}