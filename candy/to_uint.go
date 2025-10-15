package candy

import "strconv"

// ToUint 将任意类型的值转换为 uint 类型
// 支持的类型包括：bool、所有整数类型、浮点数、字符串、字节切片
// 对于不支持的类型或转换失败的情况，返回 0
// 注意：负数会被转换为 0
func ToUint(val interface{}) uint {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		if x < 0 {
			return 0
		}
		return uint(x)
	case int8:
		if x < 0 {
			return 0
		}
		return uint(x)
	case int16:
		if x < 0 {
			return 0
		}
		return uint(x)
	case int32:
		if x < 0 {
			return 0
		}
		return uint(x)
	case int64:
		if x < 0 {
			return 0
		}
		return uint(x)
	case uint:
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
		if x < 0 {
			return 0
		}
		return uint(x)
	case float64:
		if x < 0 {
			return 0
		}
		return uint(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 0)
		if err != nil {
			return 0
		}
		return uint(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 0)
		if err != nil {
			return 0
		}
		return uint(val)
	default:
		return 0
	}
}

// ToUint8 将任意类型的值转换为 uint8 类型
// 支持的类型包括：bool、所有整数类型、浮点数、字符串、字节切片
// 对于不支持的类型或转换失败的情况，返回 0
// 注意：负数会被转换为 0
func ToUint8(val interface{}) uint8 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		if x < 0 {
			return 0
		}
		return uint8(x)
	case int8:
		if x < 0 {
			return 0
		}
		return uint8(x)
	case int16:
		if x < 0 {
			return 0
		}
		return uint8(x)
	case int32:
		if x < 0 {
			return 0
		}
		return uint8(x)
	case int64:
		if x < 0 {
			return 0
		}
		return uint8(x)
	case uint:
		return uint8(x)
	case uint8:
		return x
	case uint16:
		return uint8(x)
	case uint32:
		return uint8(x)
	case uint64:
		return uint8(x)
	case float32:
		return uint8(x)
	case float64:
		return uint8(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 8)
		if err != nil {
			return 0
		}
		return uint8(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 8)
		if err != nil {
			return 0
		}
		return uint8(val)
	default:
		return 0
	}
}

// ToUint16 将各种类型的值转换为 uint16 类型
//
// 支持的输入类型：
//   - bool: true 转换为 1，false 转换为 0
//   - 整数类型 (int, int8, int16, int32, int64): 直接转换，负数返回 0
//   - 无符号整数 (uint, uint8, uint16, uint32, uint64): 直接转换
//   - 浮点数 (float32, float64): 截断小数部分后转换，负数返回 0
//   - string: 使用 strconv.ParseUint 解析十进制字符串，失败返回 0
//   - []byte: 转换为字符串后解析，失败返回 0
//   - 其他类型: 返回 0
func ToUint16(val interface{}) uint16 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		if x < 0 {
			return 0
		}
		return uint16(x)
	case int8:
		if x < 0 {
			return 0
		}
		return uint16(x)
	case int16:
		if x < 0 {
			return 0
		}
		return uint16(x)
	case int32:
		if x < 0 {
			return 0
		}
		return uint16(x)
	case int64:
		if x < 0 {
			return 0
		}
		return uint16(x)
	case uint:
		return uint16(x)
	case uint8:
		return uint16(x)
	case uint16:
		return x
	case uint32:
		return uint16(x)
	case uint64:
		return uint16(x)
	case float32:
		return uint16(x)
	case float64:
		return uint16(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return uint16(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return uint16(val)
	default:
		return 0
	}
}

// ToUint32 将任意类型转换为 uint32 类型
// 支持的类型包括：bool、所有整数类型、浮点数、字符串、字节数组
// 转换失败时返回 0
// 注意：负数会被转换为 0
func ToUint32(val interface{}) uint32 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		if x < 0 {
			return 0
		}
		return uint32(x)
	case int8:
		if x < 0 {
			return 0
		}
		return uint32(x)
	case int16:
		if x < 0 {
			return 0
		}
		return uint32(x)
	case int32:
		if x < 0 {
			return 0
		}
		return uint32(x)
	case int64:
		if x < 0 {
			return 0
		}
		return uint32(x)
	case uint:
		return uint32(x)
	case uint8:
		return uint32(x)
	case uint16:
		return uint32(x)
	case uint32:
		return x
	case uint64:
		return uint32(x)
	case float32:
		return uint32(x)
	case float64:
		return uint32(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 32)
		if err != nil {
			return 0
		}
		return uint32(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 32)
		if err != nil {
			return 0
		}
		return uint32(val)
	default:
		return 0
	}
}

// ToUint64 将各种类型的值转换为 uint64 类型
//
// 支持的输入类型：
//   - bool: true 转换为 1，false 转换为 0
//   - 整数类型 (int, int8, int16, int32, int64): 直接转换，负数返回 0
//   - 无符号整数 (uint, uint8, uint16, uint32, uint64): 直接转换
//   - 浮点数 (float32, float64): 截断小数部分后转换，负数返回 0
//   - string: 使用 strconv.ParseUint 解析十进制字符串，失败返回 0
//   - []byte: 转换为字符串后解析，失败返回 0
//   - 其他类型: 返回 0
func ToUint64(val interface{}) uint64 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		if x < 0 {
			return 0
		}
		return uint64(x)
	case int8:
		if x < 0 {
			return 0
		}
		return uint64(x)
	case int16:
		if x < 0 {
			return 0
		}
		return uint64(x)
	case int32:
		if x < 0 {
			return 0
		}
		return uint64(x)
	case int64:
		if x < 0 {
			return 0
		}
		return uint64(x)
	case uint:
		return uint64(x)
	case uint8:
		return uint64(x)
	case uint16:
		return uint64(x)
	case uint32:
		return uint64(x)
	case uint64:
		return x
	case float32:
		return uint64(x)
	case float64:
		return uint64(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return val
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}
