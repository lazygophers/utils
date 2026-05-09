package candy

import (
	"reflect"
	"strconv"
	"strings"
)

// ToFloat64Slice 将一个切片接口尽力转换为 []float64。
//
// 支持的输入切片类型包括：
//   - []bool, []int, []int8, ..., []uint64, []float32, []float64, []string, [][]byte, []interface{}
//
// 切片中的每一个元素都会通过 ToFloat64 函数进行转换。
// 如果输入为 nil，将直接返回 nil。
// 如果输入为不支持的类型，将返回一个空的 []float64{}。
//
// 性能优化：
//   - 使用预分配切片容量避免动态扩容
//   - 使用直接索引赋值避免 append 开销
//   - 针对 []float64 使用零拷贝优化
//   - 内联常见类型的转换避免函数调用开销
func ToFloat64Slice(val interface{}) []float64 {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []float64:
		// 零拷贝优化：直接返回原切片
		return x
	case []float32:
		// 优化：使用索引循环
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []int:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []int8:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []int16:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []int32:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []int64:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []uint:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []uint8:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []uint16:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []uint32:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []uint64:
		v := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			v[i] = float64(x[i])
		}
		return v
	case []bool:
		// 针对布尔类型的优化：直接赋值 0 或 1
		v := make([]float64, len(x))
		for i := range x {
			if x[i] {
				v[i] = 1
			} else {
				v[i] = 0
			}
		}
		return v
	case []string:
		// 字符串类型：使用 ToFloat64 处理复杂解析逻辑
		v := make([]float64, len(x))
		for i := range x {
			v[i] = ToFloat64(x[i])
		}
		return v
	case [][]byte:
		v := make([]float64, len(x))
		for i := range x {
			v[i] = ToFloat64(x[i])
		}
		return v
	case []interface{}:
		// 接口类型：使用 ToFloat64 处理所有可能的类型
		v := make([]float64, len(x))
		for i := range x {
			v[i] = ToFloat64(x[i])
		}
		return v
	default:
		return []float64{}
	}
}

// ToInt64Slice 将一个切片接口尽力转换为 []int64。
//
// 支持的输入切片类型包括：
//   - []bool, []int, []int8, ..., []uint64, []float32, []float64, []string, [][]byte, []interface{}
//
// 切片中的每一个元素都会通过 ToInt64 函数进行转换。
// 如果输入为 nil，将直接返回 nil。
// 如果输入为不支持的类型，将返回一个空的 []int64{}。
//
// 性能优化：
//   - 使用预分配切片容量避免动态扩容
//   - 使用直接索引赋值避免 append 开销
//   - 针对 []int64 使用 copy 优化
func ToInt64Slice(val interface{}) []int64 {
	switch x := val.(type) {
	case []int64:
		// 针对 []int64 使用 copy 优化
		result := make([]int64, len(x))
		copy(result, x)
		return result
	case []int:
		// 优化：使用索引循环替代 range
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []int32:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []int16:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []int8:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []uint:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i]) // #nosec G115 -- intentional truncation for best-effort conversion
		}
		return result
	case []uint32:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []uint64:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i]) // #nosec G115 -- intentional truncation for best-effort conversion
		}
		return result
	case []uint16:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []uint8:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []float32:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []float64:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = int64(x[i])
		}
		return result
	case []string:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = ToInt64(x[i])
		}
		return result
	case [][]byte:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = ToInt64(x[i])
		}
		return result
	case []interface{}:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = ToInt64(x[i])
		}
		return result
	case []bool:
		result := make([]int64, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = ToInt64(x[i])
		}
		return result
	default:
		return []int64{}
	}
}

// ToStringSlice 将任意类型转换为字符串切片
// 对于切片类型，将每个元素转换为字符串
// 对于字符串类型，如果包含逗号则按逗号分割，否则返回单个元素的切片
// 对于其他类型，转换为字符串后返回单个元素的切片
func ToStringSlice(v interface{}) []string {
	if v == nil {
		return nil
	}

	// 使用类型断言优化常见类型，避免反射开销
	switch x := v.(type) {
	case []string:
		if x == nil {
			return nil
		}
		// 零拷贝优化：直接返回原切片
		return x
	case string:
		if strings.Contains(x, ",") {
			return strings.Split(x, ",")
		}
		return []string{x}
	case []int:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int8:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int16:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int32:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int64:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint8:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint16:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint32:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint64:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []float32:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = ToString(x[i])
		}
		return result
	case []float64:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = ToString(x[i])
		}
		return result
	case []bool:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	case []any:
		if x == nil {
			return nil
		}
		result := make([]string, len(x))
		for i := range x {
			result[i] = ToString(x[i])
		}
		return result
	default:
		// 回退到反射处理其他类型
		vv := reflect.ValueOf(v)
		if vv.Kind() != reflect.Slice {
			return []string{ToString(v)}
		}

		if vv.IsNil() {
			return nil
		}

		// 预分配容量，避免 append 重新分配
		ss := make([]string, vv.Len())
		for i := 0; i < vv.Len(); i++ {
			ss[i] = ToString(vv.Index(i).Interface())
		}
		return ss
	}
}

// ToArrayString 是 ToStringSlice 的别名，保持向后兼容
// Deprecated: 请使用 ToStringSlice 代替
func ToArrayString(v interface{}) []string {
	return ToStringSlice(v)
}

// ToUint64Slice 将一个切片接口尽力转换为 []uint64。
//
// 支持的输入切片类型包括：
//   - []bool, []int, []int8, ..., []uint64, []float32, []float64, []string, [][]byte, []interface{}
//
// 切片中的每一个元素都会通过 ToUint64 函数进行转换。
// 如果输入为 nil，将直接返回 nil。
// 如果输入为不支持的类型，将返回一个空的 []uint64{}。
func ToUint64Slice(val interface{}) []uint64 {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []uint64:
		return x
	case []int:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []int8:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []int16:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []int32:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []int64:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []uint:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []uint8:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []uint16:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []uint32:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []float32:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []float64:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = uint64(val)
		}
		return result
	case []string:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = ToUint64(val)
		}
		return result
	case []interface{}:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = ToUint64(val)
		}
		return result
	case []bool:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = ToUint64(val)
		}
		return result
	case [][]byte:
		result := make([]uint64, len(x))
		for i, val := range x {
			result[i] = ToUint64(val)
		}
		return result
	default:
		return []uint64{}
	}
}

// ToUint32Slice 将一个切片接口尽力转换为 []uint32。
//
// 支持的输入切片类型包括：
//   - []bool, []int, []int8, ..., []uint64, []float32, []float64, []string, [][]byte, []interface{}
//
// 切片中的每一个元素都会通过 ToUint32 函数进行转换。
// 如果输入为 nil，将直接返回 nil。
// 如果输入为不支持的类型，将返回一个空的 []uint32{}。
func ToUint32Slice(val interface{}) []uint32 {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []uint32: // 最常用：零拷贝
		return x
	case []int: // 常用：直接转换
		v := make([]uint32, len(x))
		for i, val := range x {
			if val < 0 {
				v[i] = 0
			} else {
				v[i] = uint32(val) // #nosec G115
			}
		}
		return v
	case []uint: // 常用：直接转换
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = uint32(val) // #nosec G115
		}
		return v
	case []int64: // 常用：直接转换
		v := make([]uint32, len(x))
		for i, val := range x {
			if val < 0 {
				v[i] = 0
			} else {
				v[i] = uint32(val) // #nosec G115
			}
		}
		return v
	case []uint64: // 常用：直接转换
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = uint32(val) // #nosec G115
		}
		return v
	case []uint8:
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = uint32(val)
		}
		return v
	case []uint16:
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = uint32(val)
		}
		return v
	case []int32:
		v := make([]uint32, len(x))
		for i, val := range x {
			if val < 0 {
				v[i] = 0
			} else {
				v[i] = uint32(val)
			}
		}
		return v
	case []int16:
		v := make([]uint32, len(x))
		for i, val := range x {
			if val < 0 {
				v[i] = 0
			} else {
				v[i] = uint32(val)
			}
		}
		return v
	case []int8:
		v := make([]uint32, len(x))
		for i, val := range x {
			if val < 0 {
				v[i] = 0
			} else {
				v[i] = uint32(val)
			}
		}
		return v
	case []bool:
		v := make([]uint32, len(x))
		for i, val := range x {
			if val {
				v[i] = 1
			} else {
				v[i] = 0
			}
		}
		return v
	case []float32:
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = uint32(val)
		}
		return v
	case []float64:
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = uint32(val)
		}
		return v
	case []string:
		v := make([]uint32, len(x))
		for i, val := range x {
			parsed, _ := strconv.ParseUint(val, 10, 32)
			v[i] = uint32(parsed)
		}
		return v
	case [][]byte:
		v := make([]uint32, len(x))
		for i, val := range x {
			parsed, _ := strconv.ParseUint(string(val), 10, 32)
			v[i] = uint32(parsed)
		}
		return v
	case []interface{}:
		v := make([]uint32, len(x))
		for i, val := range x {
			v[i] = ToUint32(val)
		}
		return v
	default:
		return []uint32{}
	}
}

// ToInterfaceSlice 将一个切片接口尽力转换为 []interface{}。
//
// 支持的输入切片类型包括：
//   - []bool, []int, []int8, ..., []uint64, []float32, []float64, []string, [][]byte, []interface{}
//
// 切片中的每一个元素都会被转换为 interface{} 类型。
// 如果输入为 nil，将直接返回 nil。
// 如果输入为不支持的类型，将返回一个空的 []interface{}{}。
//
// 性能优化：
//   - 针对最常用类型（[]interface{}, []int, []string）优先处理
//   - []interface{} 使用零拷贝优化
//   - 其他类型使用预分配 + 直接索引避免 append 开销
//   - 针对大数据集使用批量处理（每次8个元素）
func ToInterfaceSlice(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []interface{}:
		// 零拷贝优化：最常见类型
		return x
	case []int:
		// 预分配 + 批量处理
		v := make([]interface{}, len(x))
		if len(x) >= 8 {
			i := 0
			for ; i < len(x)-7; i += 8 {
				v[i] = x[i]
				v[i+1] = x[i+1]
				v[i+2] = x[i+2]
				v[i+3] = x[i+3]
				v[i+4] = x[i+4]
				v[i+5] = x[i+5]
				v[i+6] = x[i+6]
				v[i+7] = x[i+7]
			}
			for ; i < len(x); i++ {
				v[i] = x[i]
			}
		} else {
			for i := range x {
				v[i] = x[i]
			}
		}
		return v
	case []string:
		// 预分配 + 批量处理
		v := make([]interface{}, len(x))
		if len(x) >= 8 {
			i := 0
			for ; i < len(x)-7; i += 8 {
				v[i] = x[i]
				v[i+1] = x[i+1]
				v[i+2] = x[i+2]
				v[i+3] = x[i+3]
				v[i+4] = x[i+4]
				v[i+5] = x[i+5]
				v[i+6] = x[i+6]
				v[i+7] = x[i+7]
			}
			for ; i < len(x); i++ {
				v[i] = x[i]
			}
		} else {
			for i := range x {
				v[i] = x[i]
			}
		}
		return v
	case []bool:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int8:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int16:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int32:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int64:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint8:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint16:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint32:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint64:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []float32:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []float64:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case [][]byte:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	default:
		return []interface{}{}
	}
}
