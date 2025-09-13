package candy

import "strconv"

// ToUint8 将任意类型的值转换为 uint8 类型
// 支持的类型包括：bool、所有整数类型、浮点数、字符串、字节切片
// 对于不支持的类型或转换失败的情况，返回 0
func ToUint8(val interface{}) uint8 {
	switch x := val.(type) {
	case bool:
		// 布尔值转换：true -> 1，false -> 0
		if x {
			return 1
		}
		return 0
	case int:
		// 有符号整数直接转换为 uint8
		return uint8(x)
	case int8:
		return uint8(x)
	case int16:
		return uint8(x)
	case int32:
		return uint8(x)
	case int64:
		return uint8(x)
	case uint:
		// 无符号整数转换为 uint8，可能会截断高位
		return uint8(x)
	case uint8:
		// uint8 类型直接返回
		return x
	case uint16:
		// uint16 转换为 uint8，可能会截断高位
		return uint8(x)
	case uint32:
		// uint32 转换为 uint8，可能会截断高位
		return uint8(x)
	case uint64:
		// uint64 转换为 uint8，可能会截断高位
		return uint8(x)
	case float32:
		// 浮点数转换为 uint8，会截断小数部分
		return uint8(x)
	case float64:
		// 浮点数转换为 uint8，会截断小数部分
		return uint8(x)
	case string:
		// 字符串解析为 uint8，使用十进制格式
		// 解析失败时返回 0
		val, err := strconv.ParseUint(x, 10, 8)
		if err != nil {
			return 0
		}
		return uint8(val)
	case []byte:
		// 字节切片转换为字符串后再解析为 uint8
		// 解析失败时返回 0
		val, err := strconv.ParseUint(string(x), 10, 8)
		if err != nil {
			return 0
		}
		return uint8(val)
	default:
		// 不支持的类型返回 0
		return 0
	}
}
