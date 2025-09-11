package candy

import "reflect"

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