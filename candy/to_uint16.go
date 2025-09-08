// Package candy 提供语法糖和便捷的类型转换工具函数
package candy

import "strconv"

// ToUint16 将各种类型的值转换为 uint16 类型
//
// 支持的输入类型：
//   - bool: true 转换为 1，false 转换为 0
//   - 整数类型 (int, int8, int16, int32, int64): 直接转换
//   - 无符号整数 (uint, uint8, uint16, uint32, uint64): 直接转换
//   - 浮点数 (float32, float64): 截断小数部分后转换
//   - string: 使用 strconv.ParseUint 解析十进制字符串，失败返回 0
//   - []byte: 转换为字符串后解析，失败返回 0
//   - 其他类型: 返回 0
//
// 参数:
//
//	val: 要转换的值，支持多种类型
//
// 返回:
//
//	uint16: 转换后的 uint16 值，转换失败时返回 0
//
// 示例:
//
//	ToUint16(42)        // 返回 42
//	ToUint16("100")     // 返回 100
//	ToUint16(true)      // 返回 1
//	ToUint16(3.14)      // 返回 3
//	ToUint16("invalid") // 返回 0
func ToUint16(val interface{}) uint16 {
	switch x := val.(type) {
	case bool:
		// 布尔值转换：true -> 1, false -> 0
		if x {
			return 1
		}
		return 0
	case int:
		// 有符号整数直接转换
		return uint16(x)
	case int8:
		// 8位有符号整数直接转换
		return uint16(x)
	case int16:
		// 16位有符号整数直接转换
		return uint16(x)
	case int32:
		// 32位有符号整数直接转换
		return uint16(x)
	case int64:
		// 64位有符号整数直接转换
		return uint16(x)
	case uint:
		// 无符号整数直接转换
		return uint16(x)
	case uint8:
		// 8位无符号整数直接转换
		return uint16(x)
	case uint16:
		// 如果已经是 uint16 类型，直接返回
		return x
	case uint32:
		// 32位无符号整数直接转换
		return uint16(x)
	case uint64:
		// 64位无符号整数直接转换
		return uint16(x)
	case float32:
		// 32位浮点数转换，截断小数部分
		return uint16(x)
	case float64:
		// 64位浮点数转换，截断小数部分
		return uint16(x)
	case string:
		// 字符串解析为无符号整数
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			// 解析失败返回 0
			return 0
		}
		return uint16(val)
	case []byte:
		// 字节切片转换为字符串后解析
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			// 解析失败返回 0
			return 0
		}
		return uint16(val)
	default:
		// 不支持的类型返回 0
		return 0
	}
}
