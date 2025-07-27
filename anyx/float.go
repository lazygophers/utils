package anyx

import (
	"strconv"
	"strings"
)

// ToFloat32 将任何类型的值尽力转换为 float32。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1.0, false 转换为 0.0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为浮点数，若解析失败则返回 0.0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0.0。
func ToFloat32(val interface{}) float32 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return float32(x)
	case int8:
		return float32(x)
	case int16:
		return float32(x)
	case int32:
		return float32(x)
	case int64:
		return float32(x)
	case uint:
		return float32(x)
	case uint8:
		return float32(x)
	case uint16:
		return float32(x)
	case uint32:
		return float32(x)
	case uint64:
		return float32(x)
	case float32:
		return x
	case float64:
		return float32(x)
	case string:
		v := strings.TrimSpace(x)
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return float32(val)
	case []byte:
		v := strings.TrimSpace(string(x))
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return float32(val)
	default:
		return 0
	}
}

// ToFloat64 将任何类型的值尽力转换为 float64。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1.0, false 转换为 0.0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为浮点数。若解析失败，会进一步尝试解析为整数。如果两种解析都失败，则返回 0.0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0.0。
func ToFloat64(val interface{}) float64 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return float64(x)
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	case uint:
		return float64(x)
	case uint8:
		return float64(x)
	case uint16:
		return float64(x)
	case uint32:
		return float64(x)
	case uint64:
		return float64(x)
	case float32:
		return float64(x)
	case float64:
		return x
	case string:
		v := strings.TrimSpace(x)
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return val
		}

		intVal, intErr := strconv.ParseInt(v, 0, 64)
		if intErr == nil {
			return float64(intVal)
		}

		return 0
	case []byte:
		v := strings.TrimSpace(string(x))
		val, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return val
		}

		intVal, intErr := strconv.ParseInt(v, 0, 64)
		if intErr == nil {
			return float64(intVal)
		}

		return 0
	default:
		return 0
	}
}

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
