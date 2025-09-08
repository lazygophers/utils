package candy

import "strconv"

// ToUint 将任意类型的值转换为 uint 类型
// 支持的类型包括：bool、所有整数类型、浮点数、字符串、字节切片
// 对于不支持的类型或转换失败的情况，返回 0
func ToUint(val interface{}) uint {
	switch x := val.(type) {
	case bool:
		// 布尔值转换：true -> 1，false -> 0
		if x {
			return 1
		}
		return 0
	case int:
		// 有符号整数直接转换为 uint
		return uint(x)
	case int8:
		return uint(x)
	case int16:
		return uint(x)
	case int32:
		return uint(x)
	case int64:
		return uint(x)
	case uint:
		// 无符号整数直接返回
		return x
	case uint8:
		return uint(x)
	case uint16:
		return uint(x)
	case uint32:
		return uint(x)
	case uint64:
		return uint(x)
	case float32:
		// 浮点数转换为 uint，会截断小数部分
		return uint(x)
	case float64:
		return uint(x)
	case string:
		// 字符串解析为 uint，使用十进制格式
		// 解析失败时返回 0
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return uint(val)
	case []byte:
		// 字节切片转换为字符串后再解析为 uint
		// 解析失败时返回 0
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return uint(val)
	default:
		// 不支持的类型返回 0
		return 0
	}
}
