package candy

import "reflect"

// ToMapStringArrayString 将任何类型的 map 转换为 map[string][]string。
//
// 如果输入为 nil，将返回 nil。
// 如果输入不是 map 类型，将会 panic。
// map 的 key 会通过 ToString 函数转换为字符串，value 会通过 ToArrayString 函数转换为 []string。
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
		m[ToString(mg.Key().Interface())] = ToArrayString(mg.Value().Interface())
	}

	return m
}