package candy

import "reflect"

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
