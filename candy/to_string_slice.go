package candy

import (
	"bytes"
	"math"
	"strconv"
	"strings"

	"github.com/lazygophers/utils/json"
)

// ToStringSlice 将任意类型转换为字符串切片
// 支持的类型包括各种基础类型的切片、字符串、字节切片等
// seqs 参数用于指定分隔符，默认为逗号
func ToStringSlice(val interface{}, seqs ...string) []string {
	var seq string
	if len(seqs) > 0 {
		seq = seqs[0]
	} else {
		seq = ","
	}

	switch x := val.(type) {
	case []bool:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			if v {
				ss = append(ss, "1")
			} else {
				ss = append(ss, "0")
			}
		}
		return ss

	case []int:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.Itoa(v))
		}
		return ss

	case []int8:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatInt(int64(v), 10))
		}
		return ss

	case []int16:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatInt(int64(v), 10))
		}
		return ss

	case []int32:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatInt(int64(v), 10))
		}
		return ss

	case []int64:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatInt(v, 10))
		}
		return ss

	case []uint:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatUint(uint64(v), 10))
		}
		return ss

	case []uint16:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatUint(uint64(v), 10))
		}
		return ss

	case []uint32:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatUint(uint64(v), 10))
		}
		return ss

	case []uint64:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, strconv.FormatUint(v, 10))
		}
		return ss

	case []float32:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			if math.Floor(float64(v)) == float64(v) {
				ss = append(ss, strconv.FormatInt(int64(v), 10))
			} else {
				ss = append(ss, strconv.FormatFloat(float64(v), 'f', -1, 32))
			}
		}
		return ss

	case []float64:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			if math.Floor(v) == (v) {
				ss = append(ss, strconv.FormatInt(int64(v), 10))
			} else {
				ss = append(ss, strconv.FormatFloat(v, 'f', -1, 32))
			}
		}
		return ss

	case []string:
		return x

	case []byte:
		if bytes.HasPrefix(x, []byte("[")) && bytes.HasSuffix(x, []byte("]")) {
			var values []any
			err := json.Unmarshal(x, &values)
			if err == nil {
				return ToStringSlice(values)
			}
		}

		if seq == "" {
			return []string{toString(x)}
		}

		return strings.Split(toString(x), seq)

	case string:
		if strings.HasPrefix(x, "[") && strings.HasSuffix(x, "]") {
			var values []any
			err := json.UnmarshalString(x, &values)
			if err == nil {
				return ToStringSlice(values)
			}
		}

		if seq == "" {
			return []string{x}
		}

		return strings.Split(x, seq)

	case []interface{}:
		ss := make([]string, 0, len(x))
		for _, v := range x {
			ss = append(ss, ToString(v))
		}
		return ss

	default:
		return nil
	}
}
