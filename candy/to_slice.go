package candy

import (
	"reflect"
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
func ToFloat64Slice(val interface{}) []float64 {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []bool:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []int:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []int8:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []int16:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []int32:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []int64:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []uint:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []uint8:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []uint16:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []uint32:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []uint64:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []float32:
		v := make([]float64, len(x))
		for i, val := range x {
			v[i] = ToFloat64(val)
		}
		return v
	case []float64:
		var v []float64
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []string:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case [][]byte:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
		}
		return v
	case []interface{}:
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
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
func ToInt64Slice(val interface{}) []int64 {
	switch x := val.(type) {
	case []bool:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	case []int:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int8:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int16:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int32:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []int64:
		var v []int64
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint8:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint16:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint32:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []uint64:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []float32:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []float64:
		var v []int64
		for _, val := range x {
			v = append(v, int64(val))
		}
		return v
	case []string:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	case [][]byte:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	case []interface{}:
		var v []int64
		for _, val := range x {
			v = append(v, ToInt64(val))
		}
		return v
	default:
		return []int64{}
	}
}

// ToStringSlice 将任意类型转换为字符串切片
// 对于切片类型，将每个元素转换为字符串
// 对于字符串类型，如果包含逗号则按逗号分割，否则返回单个元素的切片
// 对于其他类型，转换为字符串后返回单个元素的切片
func ToStringSlice(v interface{}) []string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		// 处理非切片类型
		switch x := v.(type) {
		case string:
			if strings.Contains(x, ",") {
				// 如果包含逗号，按逗号分割
				return strings.Split(x, ",")
			}
			// 否则返回单个元素的切片
			return []string{x}
		case nil:
			return nil
		default:
			// 其他非切片类型，转换为字符串后返回单个元素的切片
			return []string{ToString(x)}
		}
	}

	// 处理 nil 切片
	if vv.IsNil() {
		return nil
	}

	ss := make([]string, 0, vv.Len())
	for i := 0; i < vv.Len(); i++ {
		ss = append(ss, ToString(vv.Index(i).Interface()))
	}

	return ss
}

// ToArrayString 是 ToStringSlice 的别名，保持向后兼容
// Deprecated: 请使用 ToStringSlice 代替
func ToArrayString(v interface{}) []string {
	return ToStringSlice(v)
}
