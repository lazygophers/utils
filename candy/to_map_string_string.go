package candy

import "reflect"

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
