package candy

import (
	"strconv"
	"strings"
)

// ToFloat32 将任何类型的值尽力转换为 float32。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1.0, false 转换为 0.0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为浮点数，若解析失败则返回 0.0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0.0。
func ToFloat32(val interface{}) float32 {
	// 快速路径：nil 检查
	if val == nil {
		return 0
	}

	switch x := val.(type) {
	case float32:
		return x // 零拷贝：直接返回相同类型
	case float64:
		return float32(x)
	case int:
		return float32(x)
	case int64:
		return float32(x)
	case string:
		v := strings.TrimSpace(x)
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return float32(val)
	case bool:
		if x {
			return 1
		}
		return 0
	case int8:
		return float32(x)
	case int16:
		return float32(x)
	case int32:
		return float32(x)
	case uint:
		return float32(x)
	case uint8:
		return float32(x)
	case uint16:
		return float32(x)
	case uint32:
		return float32(x)
	case uint64:
		return float32(x)
	case []byte:
		v := strings.TrimSpace(string(x))
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return float32(val)
	default:
		return 0
	}
}

// ToFloat64 将任何类型的值尽力转换为 float64。
// 优化版本：完全展开所有类型分支，减少嵌套，提升性能
//
// 支持的输入类型包括：
//   - bool: true 转换为 1.0, false 转换为 0.0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为浮点数。若解析失败，会进一步尝试解析为整数。如果两种解析都失败，则返回 0.0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0.0。
func ToFloat64(val interface{}) float64 {
	switch v := val.(type) {
	case nil:
		return 0
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		str := strings.TrimSpace(v)
		val, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return val
		}

		intVal, intErr := strconv.ParseInt(str, 0, 64)
		if intErr == nil {
			return float64(intVal)
		}

		return 0
	case []byte:
		s := strings.TrimSpace(string(v))
		val, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return val
		}

		intVal, intErr := strconv.ParseInt(s, 0, 64)
		if intErr == nil {
			return float64(intVal)
		}

		return 0
	default:
		return 0
	}
}
