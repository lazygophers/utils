package anyx

import "strconv"

// ToUint 将任意类型转换为uint类型
// 布尔值处理同整数转换
// 数值类型直接转换为无符号整数
// 字符串处理：
// - 尝试解析为uint64后转换
// - 解析失败返回0
// []byte同字符串处理逻辑
func ToUint(val interface{}) uint {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
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
		return uint(x)
	case float64:
		return uint(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return uint(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return uint(val)
	default:
		return 0
	}
}

// ToUint8 将任意类型转换为uint8类型
// 转换规则与ToUint相同，但目标类型为uint8
// 可能发生截断，使用者需注意数值范围
func ToUint8(val interface{}) uint8 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
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
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return uint8(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return uint8(val)
	default:
		return 0
	}
}

func ToUint16(val interface{}) uint16 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint16(x)
	case int8:
		return uint16(x)
	case int16:
		return uint16(x)
	case int32:
		return uint16(x)
	case int64:
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
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return uint16(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return uint16(val)
	default:
		return 0
	}
}

func ToUint32(val interface{}) uint32 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint32(x)
	case int8:
		return uint32(x)
	case int16:
		return uint32(x)
	case int32:
		return uint32(x)
	case int64:
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
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return uint32(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return uint32(val)
	default:
		return 0
	}
}

func ToUint64(val interface{}) uint64 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint64(x)
	case int8:
		return uint64(x)
	case int16:
		return uint64(x)
	case int32:
		return uint64(x)
	case int64:
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
		return uint64(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return uint64(val)
	default:
		return 0
	}
}
