package anyx

import "strconv"

// ToFloat32 将任意类型转换为float32类型
// 转换规则与ToFloat64相同，但目标类型为float32
// 特殊情况：
// - JSON序列化失败返回0
// - 整数类型转换为浮点数形式
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
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return float32(val)
	case []byte:
		val, err := strconv.ParseFloat(string(x), 64)
		if err != nil {
			return 0
		}
		return float32(val)
	default:
		return 0
	}
}

// ToFloat64 将任意类型转换为float64类型
// 字符串处理：
// - 优先尝试解析为浮点数
// - 整数转换为浮点数形式
// []byte同字符串处理逻辑
// JSON序列化失败返回0
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
		val, err := strconv.ParseFloat(x, 64)
		if err != nil {
			val, err := strconv.ParseInt(x, 10, 64)
			if err != nil {
				return 0
			}
			return float64(val)
		}
		return val
	case []byte:
		val, err := strconv.ParseFloat(string(x), 64)
		if err != nil {
			val, err := strconv.ParseInt(string(x), 10, 64)
			if err != nil {
				return 0
			}
			return float64(val)
		}
		return val
	default:
		return 0
	}
}

func ToFloat64Slice(val interface{}) []float64 {
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
		var v []float64
		for _, val := range x {
			v = append(v, ToFloat64(val))
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
