package candy

import "github.com/lazygophers/utils/json"

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
