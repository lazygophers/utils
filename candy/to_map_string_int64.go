package candy

import "reflect"

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