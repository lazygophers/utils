package candy

import "strconv"

// ToUint32 将任意类型转换为 uint32 类型
// 支持的类型包括：bool、所有整数类型、浮点数、字符串、字节数组
// 转换失败时返回 0
func ToUint32(val interface{}) uint32 {
	switch x := val.(type) {
	case bool:
		// 布尔值转换：true -> 1，false -> 0
		if x {
			return 1
		}
		return 0
	case int:
		// 有符号整数转换为 uint32
		return uint32(x)
	case int8:
		// 8位有符号整数转换为 uint32
		return uint32(x)
	case int16:
		// 16位有符号整数转换为 uint32
		return uint32(x)
	case int32:
		// 32位有符号整数转换为 uint32
		return uint32(x)
	case int64:
		// 64位有符号整数转换为 uint32
		return uint32(x)
	case uint:
		// 无符号整数转换为 uint32
		return uint32(x)
	case uint8:
		// 8位无符号整数转换为 uint32
		return uint32(x)
	case uint16:
		// 16位无符号整数转换为 uint32
		return uint32(x)
	case uint32:
		// 如果已经是 uint32 类型，直接返回
		return x
	case uint64:
		// 64位无符号整数转换为 uint32，可能发生截断
		return uint32(x)
	case float32:
		// 32位浮点数转换为 uint32
		return uint32(x)
	case float64:
		// 64位浮点数转换为 uint32
		return uint32(x)
	case string:
		// 字符串解析为 uint32
		val, err := strconv.ParseUint(x, 10, 32)
		if err != nil {
			return 0
		}
		return uint32(val)
	case []byte:
		// 字节数组转换为字符串后解析为 uint32
		val, err := strconv.ParseUint(string(x), 10, 32)
		if err != nil {
			return 0
		}
		return uint32(val)
	default:
		// 不支持的类型返回 0
		return 0
	}
}
