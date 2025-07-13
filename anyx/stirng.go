package anyx

import (
	"bytes"
	"fmt"
	"github.com/lazygophers/utils/json"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// ToString 将任意类型转换为字符串表示
// 支持类型：
// - 布尔值：true -> "1"，false -> "0"
// - 整数类型：直接格式化为十进制字符串
// - 浮点数：整数部分无小数时返回整数形式，否则保留小数
// - time.Duration：使用其String()方法
// - 字符串/[]byte：直接返回
// - error：返回错误信息
// - 其他类型：使用JSON序列化
// 返回空字符串表示转换失败
func ToString(val interface{}) string {
	switch x := val.(type) {
	case bool:
		if x {
			return "1"
		}
		return "0"
	case int:
		return fmt.Sprintf("%d", x)
	case int8:
		return fmt.Sprintf("%d", x)
	case int16:
		return fmt.Sprintf("%d", x)
	case int32:
		return fmt.Sprintf("%d", x)
	case int64:
		return fmt.Sprintf("%d", x)
	case uint:
		return fmt.Sprintf("%d", x)
	case uint8:
		return fmt.Sprintf("%d", x)
	case uint16:
		return fmt.Sprintf("%d", x)
	case uint32:
		return fmt.Sprintf("%d", x)
	case uint64:
		return fmt.Sprintf("%d", x)
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return fmt.Sprintf("%.0f", x)
		}

		return fmt.Sprintf("%f", x)
	case float64:
		if math.Floor(x) == x {
			return fmt.Sprintf("%.0f", x)
		}

		return fmt.Sprintf("%f", x)
	case time.Duration:
		return x.String()
	case string:
		return x
	case []byte:
		return string(x)
	case nil:
		return ""
	case error:
		return x.Error()

	default:
		buf, err := json.Marshal(x)
		if err != nil {
			return ""
		}

		return string(buf)
	}
}

// ToBytes 将任意类型转换为字节数组
// 转换规则与ToString相同，但返回[]byte类型
// 特殊情况：
// - nil返回nil
// - error类型返回错误信息的字节表示
// - JSON序列化失败时返回nil
func ToBytes(val interface{}) []byte {
	switch x := val.(type) {
	case bool:
		if x {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(fmt.Sprintf("%d", x))
	case int8:
		return []byte(fmt.Sprintf("%d", x))
	case int16:
		return []byte(fmt.Sprintf("%d", x))
	case int32:
		return []byte(fmt.Sprintf("%d", x))
	case int64:
		return []byte(fmt.Sprintf("%d", x))
	case uint:
		return []byte(fmt.Sprintf("%d", x))
	case uint8:
		return []byte(fmt.Sprintf("%d", x))
	case uint16:
		return []byte(fmt.Sprintf("%d", x))
	case uint32:
		return []byte(fmt.Sprintf("%d", x))
	case uint64:
		return []byte(fmt.Sprintf("%d", x))
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return []byte(fmt.Sprintf("%.0f", x))
		}

		return []byte(fmt.Sprintf("%f", x))
	case float64:
		if math.Floor(x) == x {
			return []byte(fmt.Sprintf("%.0f", x))
		}

		return []byte(fmt.Sprintf("%f", x))
	case time.Duration:
		return []byte(x.String())
	case string:
		return []byte(x)
	case []byte:
		return x
	case nil:
		return nil
	case error:
		return []byte(x.Error())

	default:
		buf, err := json.Marshal(x)
		if err != nil {
			return nil
		}

		return buf
	}
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// ToArrayString 将任意切片转换为字符串切片
// 处理规则：
// - 使用反射遍历切片元素
// - 每个元素转换为字符串
// - 非切片类型返回空切片
func ToArrayString(v interface{}) []string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		return []string{}
	}

	ss := make([]string, 0, vv.Len())
	for i := 0; i < vv.Len(); i++ {
		ss = append(ss, ToString(vv.Index(i).Interface()))
	}

	return ss
}

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

	return nil
}
