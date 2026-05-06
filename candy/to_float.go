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
// 优化版本：将最常用的类型放在前面，减少类型分支开销
//
// 支持的输入类型包括：
//   - bool: true 转换为 1.0, false 转换为 0.0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为浮点数。若解析失败，会进一步尝试解析为整数。如果两种解析都失败，则返回 0.0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0.0。
func ToFloat64(val interface{}) float64 {
	// 快速路径：nil 检查
	if val == nil {
		return 0
	}

	switch x := val.(type) {
	// 常见类型优先
	case float64: // 最常见
		return x
	case float32: // 常见
		return float64(x)
	case int: // 常见
		return float64(x)
	case int64: // 常见
		return float64(x)
	case string: // 常见
		v := strings.TrimSpace(x)
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return val
		}

		intVal, intErr := strconv.ParseInt(v, 0, 64)
		if intErr == nil {
			return float64(intVal)
		}

		return 0
	case bool:
		if x {
			return 1
		}
		return 0
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case uint:
		return float64(x)
	case uint8:
		return float64(x)
	case uint16:
		return float64(x)
	case uint32:
		return float64(x)
	case uint64:
		return float64(x)
	case []byte:
		v := strings.TrimSpace(string(x))
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return val
		}

		intVal, intErr := strconv.ParseInt(v, 0, 64)
		if intErr == nil {
			return float64(intVal)
		}

		return 0
	default:
		return 0
	}
}
