package candy

import (
	"reflect"

	"github.com/lazygophers/utils/json"
)

// ToMap 将任何类型的值转换为 map[string]interface{}。
//
// 支持的输入类型包括：
//   - []byte: 尝试JSON反序列化为map，失败则通过ToMapStringAny转换
//   - string: 尝试JSON反序列化为map，失败则通过ToMapStringAny转换
//   - 其他类型: 通过ToMapStringAny函数转换
//
// 如果输入为 nil 或转换失败，将返回相应的默认值。
func ToMap(v interface{}) map[string]interface{} {
	switch x := v.(type) {
	case []byte:
		var m map[string]interface{}
		err := json.Unmarshal(x, &m)
		if err == nil {
			return m
		}

	case string:
		var m map[string]interface{}
		err := json.UnmarshalString(x, &m)
		if err == nil {
			return m
		}

	}
	return ToMapStringAny(v)
}

// ToMapInt32String 将任何类型的 map 转换为 map[int32]string。
//
// 如果输入不是 map 类型，将返回一个空的 map[int32]string{}。
// map 的 key 会通过 ToInt32 函数转换为 int32，value 会通过 ToString 函数转换为字符串。
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

// ToMapInt64String 将任何类型的 map 转换为 map[int64]string。
//
// 如果输入不是 map 类型，将返回一个空的 map[int64]string{}。
// map 的 key 会通过 ToInt64 函数转换为 int64，value 会通过 ToString 函数转换为字符串。
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

// ToMapStringAny 将任何类型的 map 转换为 map[string]interface{}。
//
// 如果输入为 nil，将返回 nil。
// 如果输入不是 map 类型，将返回一个空的 map[string]interface{}{}。
// map 的 key 会通过 ToString 函数转换为字符串，value 保持原始类型。
func ToMapStringAny(v interface{}) map[string]interface{} {
	if v == nil {
		return nil
	}

	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]interface{}{}
	}

	m := make(map[string]interface{})
	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = mg.Value().Interface()
	}

	return m
}

// ToMapStringArrayString 将任何类型的 map 转换为 map[string][]string。
//
// 如果输入为 nil，将返回 nil。
// 如果输入不是 map 类型，将会 panic。
// map 的 key 会通过 ToString 函数转换为字符串，value 会通过 ToStringSlice 函数转换为 []string。
func ToMapStringArrayString(v interface{}) map[string][]string {
	if v == nil {
		return nil
	}

	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		panic("required map type")
	}

	m := make(map[string][]string)
	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToStringSlice(mg.Value().Interface())
	}

	return m
}

// ToMapStringInt64 将任何类型的 map 转换为 map[string]int64。
//
// 如果输入不是 map 类型，将返回一个空的 map[string]int64{}。
// map 的 key 会通过 ToString 函数转换为字符串，value 会通过 ToInt64 函数转换为 int64。
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

// ToMapStringString 将任何类型的 map 转换为 map[string]string。
//
// 如果输入不是 map 类型，将返回一个空的 map[string]string{}。
// map 的 key 和 value 都会通过 ToString 函数转换为字符串。
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
