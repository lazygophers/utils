package candy

import (
	"reflect"
	"strconv"
	"strings"
)

// Numeric 定义数值类型约束
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Convertible 定义可转换类型约束
type Convertible interface {
	~bool | ~string | ~[]byte | Numeric
}

// Convert 通用类型转换函数，使用泛型替代所有 to_* 函数
func Convert[T Convertible, U Numeric](val T) U {
	v := any(val)
	switch x := v.(type) {
	case bool:
		if x {
			return U(1)
		}
		return U(0)
	case int:
		return U(x)
	case int8:
		return U(x)
	case int16:
		return U(x)
	case int32:
		return U(x)
	case int64:
		return U(x)
	case uint:
		return U(x)
	case uint8:
		return U(x)
	case uint16:
		return U(x)
	case uint32:
		return U(x)
	case uint64:
		return U(x)
	case float32:
		return U(x)
	case float64:
		return U(x)
	case string:
		// 根据目标类型选择解析方法
		var zero U
		switch any(zero).(type) {
		case int, int8, int16, int32, int64:
			if parsed, err := strconv.ParseInt(x, 10, 64); err == nil {
				return U(parsed)
			}
		case uint, uint8, uint16, uint32, uint64:
			if parsed, err := strconv.ParseUint(x, 10, 64); err == nil {
				return U(parsed)
			}
		case float32, float64:
			if parsed, err := strconv.ParseFloat(x, 64); err == nil {
				return U(parsed)
			}
		}
		return U(0)
	case []byte:
		return Convert[string, U](string(x))
	default:
		return U(0)
	}
}

// ConvertWithDefault 带默认值的类型转换
func ConvertWithDefault[T any, U Numeric](val T, defaultVal U) U {
	v := reflect.ValueOf(val)
	if !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return defaultVal
	}

	switch x := any(val).(type) {
	case bool:
		if x {
			return U(1)
		}
		return U(0)
	case int:
		return U(x)
	case int8:
		return U(x)
	case int16:
		return U(x)
	case int32:
		return U(x)
	case int64:
		return U(x)
	case uint:
		return U(x)
	case uint8:
		return U(x)
	case uint16:
		return U(x)
	case uint32:
		return U(x)
	case uint64:
		return U(x)
	case float32:
		return U(x)
	case float64:
		return U(x)
	case string:
		var zero U
		switch any(zero).(type) {
		case int, int8, int16, int32, int64:
			if parsed, err := strconv.ParseInt(x, 10, 64); err == nil {
				return U(parsed)
			}
		case uint, uint8, uint16, uint32, uint64:
			if parsed, err := strconv.ParseUint(x, 10, 64); err == nil {
				return U(parsed)
			}
		case float32, float64:
			if parsed, err := strconv.ParseFloat(x, 64); err == nil {
				return U(parsed)
			}
		}
		return defaultVal
	case []byte:
		return ConvertWithDefault[string, U](string(x), defaultVal)
	default:
		return defaultVal
	}
}

// ToBoolGeneric 通用布尔转换函数
func ToBoolGeneric[T any](val T) bool {
	return ToBoolInterface(val)
}

// ToBoolInterface 兼容 interface{} 的布尔转换函数
func ToBoolInterface(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		return x
	case int:
		return x != 0
	case int8:
		return x != 0
	case int16:
		return x != 0
	case int32:
		return x != 0
	case int64:
		return x != 0
	case uint:
		return x != 0
	case uint8:
		return x != 0
	case uint16:
		return x != 0
	case uint32:
		return x != 0
	case uint64:
		return x != 0
	case float32:
		return x != 0 && !isNaN32(x)
	case float64:
		return x != 0 && !isNaN64(x)
	case string:
		return stringToBool(x)
	case []byte:
		return stringToBool(string(x))
	default:
		// 对于所有其他类型，保持与原始实现一致，返回 false
		return false
	}
}

// 辅助函数
func isNaN32(f float32) bool {
	return f != f
}

func isNaN64(f float64) bool {
	return f != f
}

func stringToBool(s string) bool {
	s = strings.TrimSpace(strings.ToLower(s))
	switch s {
	case "true", "1", "t", "y", "yes", "on":
		return true
	case "false", "0", "f", "n", "no", "off", "":
		return false
	default:
		return true
	}
}

// ToStringGeneric 通用字符串转换函数
func ToStringGeneric[T any](val T) string {
	switch x := any(val).(type) {
	case string:
		return x
	case []byte:
		return string(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	case int:
		return strconv.FormatInt(int64(x), 10)
	case int8:
		return strconv.FormatInt(int64(x), 10)
	case int16:
		return strconv.FormatInt(int64(x), 10)
	case int32:
		return strconv.FormatInt(int64(x), 10)
	case int64:
		return strconv.FormatInt(x, 10)
	case uint:
		return strconv.FormatUint(uint64(x), 10)
	case uint8:
		return strconv.FormatUint(uint64(x), 10)
	case uint16:
		return strconv.FormatUint(uint64(x), 10)
	case uint32:
		return strconv.FormatUint(uint64(x), 10)
	case uint64:
		return strconv.FormatUint(x, 10)
	case float32:
		return strconv.FormatFloat(float64(x), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	default:
		v := reflect.ValueOf(val)
		if !v.IsValid() {
			return ""
		}
		if v.Kind() == reflect.Ptr && v.IsNil() {
			return ""
		}
		return ""
	}
}

// ToSlice 通用切片转换函数
func ToSlice[T any, U any](val T, converter func(any) U) []U {
	v := reflect.ValueOf(val)
	if !v.IsValid() {
		return nil
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]U, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = converter(v.Index(i).Interface())
		}
		return result
	default:
		// 单个值转换为单元素切片
		return []U{converter(val)}
	}
}