package candy

import "reflect"

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