package anyx

import (
	"strconv"
	"time"
)

func ToInt(val interface{}) int {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

func ToInt8(val interface{}) int8 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int8(x)
	case int8:
		return x
	case int16:
		return int8(x)
	case int32:
		return int8(x)
	case int64:
		return int8(x)
	case uint:
		return int8(x)
	case uint8:
		return int8(x)
	case uint16:
		return int8(x)
	case uint32:
		return int8(x)
	case uint64:
		return int8(x)
	case float32:
		return int8(x)
	case float64:
		return int8(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return int8(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return int8(val)
	default:
		return 0
	}
}

func ToInt16(val interface{}) int16 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int16(x)
	case int8:
		return int16(x)
	case int16:
		return x
	case int32:
		return int16(x)
	case int64:
		return int16(x)
	case uint:
		return int16(x)
	case uint8:
		return int16(x)
	case uint16:
		return int16(x)
	case uint32:
		return int16(x)
	case uint64:
		return int16(x)
	case float32:
		return int16(x)
	case float64:
		return int16(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return int16(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return int16(val)
	default:
		return 0
	}
}

func ToInt32(val interface{}) int32 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int32(x)
	case int8:
		return int32(x)
	case int16:
		return int32(x)
	case int32:
		return x
	case int64:
		return int32(x)
	case uint:
		return int32(x)
	case uint8:
		return int32(x)
	case uint16:
		return int32(x)
	case uint32:
		return int32(x)
	case uint64:
		return int32(x)
	case float32:
		return int32(x)
	case float64:
		return int32(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return 0
		}
		return int32(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return int32(val)
	default:
		return 0
	}
}

func ToInt64(val interface{}) int64 {
	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	case uint:
		return int64(x)
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x)
	case time.Duration:
		return int64(x)
	case float32:
		return int64(x)
	case float64:
		return int64(x)
	case string:
		val, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			return 0
		}
		return val
	case []byte:
		val, err := strconv.ParseInt(string(x), 10, 64)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}

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
			v = append(v, int64(val))
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
