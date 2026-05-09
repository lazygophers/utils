package candy

import (
	"strconv"
	"time"
)

// ToInt 将任何类型的值尽力转换为 int。
// 优化版本：完全展开所有类型分支，避免函数调用开销
//
// 支持的输入类型包括：
//   - bool: true 转换为 1, false 转换为 0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为整数，若解析失败则返回 0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0。
func ToInt(val interface{}) int {
	switch v := val.(type) {
	case nil:
		return 0
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v) // #nosec G115 -- intentional truncation for best-effort conversion
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

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
	// 快速路径：nil 检查
	if val == nil {
		return 0
	}

	switch x := val.(type) {
	case int8:
		return x // 零拷贝：直接返回相同类型
	case int:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case int64:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case string:
		val, err := strconv.ParseInt(x, 10, 8)
		if err != nil {
			return 0
		}
		return int8(val)
	case bool:
		if x {
			return 1
		}
		return 0
	case int16:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case int32:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint8:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint16:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint32:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint64:
		return int8(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case float32:
		return int8(x)
	case float64:
		return int8(x)
	case []byte:
		val, err := strconv.ParseInt(string(x), 10, 8)
		if err != nil {
			return 0
		}
		return int8(val)
	default:
		return 0
	}
}

// ToInt16 将任何类型的值尽力转换为 int16。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1, false 转换为 0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为无符号整数，若解析失败则返回 0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0。
func ToInt16(val interface{}) int16 {
	// 快速路径：nil 检查
	if val == nil {
		return 0
	}

	switch x := val.(type) {
	case int16:
		return x // 零拷贝：直接返回相同类型
	case int:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case int64:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case string:
		val, err := strconv.ParseInt(x, 10, 16)
		if err != nil {
			return 0
		}
		return int16(val)
	case bool:
		if x {
			return 1
		}
		return 0
	case int8:
		return int16(x)
	case int32:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint8:
		return int16(x)
	case uint16:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint32:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint64:
		return int16(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case float32:
		return int16(x)
	case float64:
		return int16(x)
	case []byte:
		val, err := strconv.ParseInt(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return int16(val)
	default:
		return 0
	}
}

// ToInt32 将任何类型的值尽力转换为 int32。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1, false 转换为 0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为无符号整数，若解析失败则返回 0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0。
func ToInt32(val interface{}) int32 {
	// 快速路径：nil 检查
	if val == nil {
		return 0
	}

	switch x := val.(type) {
	case int32:
		return x // 零拷贝：直接返回相同类型
	case int:
		return int32(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case int64:
		return int32(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case string:
		val, err := strconv.ParseInt(x, 10, 32)
		if err != nil {
			return 0
		}
		return int32(val)
	case bool:
		if x {
			return 1
		}
		return 0
	case int8:
		return int32(x)
	case int16:
		return int32(x)
	case uint:
		return int32(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint8:
		return int32(x)
	case uint16:
		return int32(x)
	case uint32:
		return int32(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint64:
		return int32(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case float32:
		return int32(x)
	case float64:
		return int32(x)
	case []byte:
		val, err := strconv.ParseInt(string(x), 10, 32)
		if err != nil {
			return 0
		}
		return int32(val)
	default:
		return 0
	}
}

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
	// 快速路径：nil 检查
	if val == nil {
		return 0
	}

	switch x := val.(type) {
	case int64:
		return x // 零拷贝：直接返回相同类型
	case int:
		return int64(x)
	case string:
		val, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return 0
		}
		return val
	case bool:
		if x {
			return 1
		}
		return 0
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case uint:
		return int64(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x) // #nosec G115 -- intentional truncation for best-effort conversion
	case time.Duration:
		return int64(x)
	case float32:
		return int64(x)
	case float64:
		return int64(x)
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
