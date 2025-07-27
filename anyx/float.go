package anyx

import "strconv"

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
		val, err := strconv.ParseFloat(x, 64)
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
