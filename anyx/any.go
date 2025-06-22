package anyx

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/lazygophers/utils/json"
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

// ToInt 将任意类型转换为int
// 布尔值处理：
// - true -> 1
// - false -> 0
// 数值类型直接转换
// 字符串处理：
// - 优先尝试解析为uint64后转换
// - 解析失败返回0
// - 超出int范围时返回最大/最小值
// - 解析失败返回0
// []byte同字符串处理逻辑
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

// ToInt8 将任意类型转换为int8
// 转换规则与ToInt相同，但目标类型为int8
// 可能发生截断，使用者需注意数值范围
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

// ToInt16 将任意类型转换为int16
// 转换规则与ToInt相同，但目标类型为int16
// 可能发生截断，使用者需注意数值范围
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

// ToInt32 将任意类型转换为int32
// 转换规则与ToInt相同，但目标类型为int32
// 可能发生截断，使用者需注意数值范围
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

// ToInt64 将任意类型转换为int64类型
// 处理逻辑：
// - 布尔值：true -> 1，false -> 0
// - 所有整数类型直接转换
// - 浮点数取整转换
// - 字符串处理：
//   - 优先尝试解析为int64
//   - 解析失败返回0
//
// - []byte同字符串处理逻辑
// - time.Duration返回纳秒值
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

func ToBool(val interface{}) bool {
	switch x := val.(type) {
	case bool:
		return x
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return x != 0
	case string:
		switch strings.ToLower(x) {
		case "true", "1", "t", "y", "yes", "on":
			return true
		case "false", "0", "f", "n", "no", "off":
			return false
		default:
			return strings.TrimSpace(x) != ""
		}
	case []byte:
		switch string(bytes.ToLower(x)) {
		case "true", "1", "t", "y", "yes", "on":
			return true
		case "false", "0", "f", "n", "no", "off":
			return false
		default:
			return len(bytes.TrimSpace(x)) != 0
		}
	default:
		return val == nil
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

// ToMapStringAny 将任意map转换为string-key的interface{} map
// 处理逻辑：
// - 使用反射遍历map键值对
// - 键转换为字符串
// - 值保持为interface{}类型
// - 非map类型返回空map
func ToMapStringAny(v interface{}) map[string]interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]interface{}{}
	}

	m := make(map[string]any)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = mg.Value().Interface()
	}

	return m
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

// ToMap 将任意类型转换为map[string]interface{}
// 支持转换类型：
// - []byte：尝试JSON解析
// - string：尝试JSON解析
// - 其他类型：使用反射遍历map
// 性能注意事项：
// - 使用反射时需注意类型校验开销
// - JSON解析失败时需避免panic
// - 空map返回而非nil
// 返回nil表示转换失败
func ToMap(v interface{}) map[string]interface{} {
	switch x := v.(type) {
	case []byte:
		var m map[string]any
		err := json.Unmarshal(x, &m)
		if err == nil {
			return m
		}

	case string:
		var m map[string]any
		err := json.UnmarshalString(x, &m)
		if err == nil {
			return m
		}

	}

	return ToMapStringAny(v)
}

// ToMapStringString 将任意map转换为string-key的string map
// 处理逻辑：
// - 使用反射遍历map键值对
// - 键转换为字符串
// - 值转换为字符串
// - 非map类型返回空map
func ToMapStringString(v interface{}) map[string]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]string{}
	}

	m := make(map[string]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToString(mg.Value().Interface())
	}

	return m
}

func ToMapStringInt64(v interface{}) map[string]int64 {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]int64{}
	}

	m := make(map[string]int64)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToInt64(mg.Value().Interface())
	}

	return m
}

func ToMapInt64String(v interface{}) map[int64]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[int64]string{}
	}

	m := make(map[int64]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToInt64(mg.Key().Interface())] = ToString(mg.Value().Interface())
	}

	return m
}

func ToMapInt32String(v interface{}) map[int32]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[int32]string{}
	}

	m := make(map[int32]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToInt32(mg.Key().Interface())] = ToString(mg.Value().Interface())
	}

	return m
}

// ToMapStringArrayString 将任意map转换为string-key的[]string map
// 处理逻辑：
// - 使用反射遍历map键值对
// - 键转换为字符串
// - 值转换为字符串切片
// - 非map类型返回空map
func ToMapStringArrayString(v interface{}) map[string][]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string][]string{}
	}

	m := make(map[string][]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToArrayString(mg.Value().Interface())
	}

	return m
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

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
