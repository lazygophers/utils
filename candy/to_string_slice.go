package candy

import (
	"reflect"
	"strings"
)

// ToStringSlice 将任意类型转换为字符串切片
// 对于切片类型，将每个元素转换为字符串
// 对于字符串类型，如果包含逗号则按逗号分割，否则返回单个元素的切片
// 对于其他类型，转换为字符串后返回单个元素的切片
func ToStringSlice(v interface{}) []string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Slice {
		// 处理非切片类型
		switch x := v.(type) {
		case string:
			if strings.Contains(x, ",") {
				// 如果包含逗号，按逗号分割
				return strings.Split(x, ",")
			}
			// 否则返回单个元素的切片
			return []string{x}
		case nil:
			return nil
		default:
			// 其他非切片类型，转换为字符串后返回单个元素的切片
			return []string{ToString(x)}
		}
	}

	// 处理 nil 切片
	if vv.IsNil() {
		return nil
	}

	ss := make([]string, 0, vv.Len())
	for i := 0; i < vv.Len(); i++ {
		ss = append(ss, ToString(vv.Index(i).Interface()))
	}

	return ss
}

// ToArrayString 是 ToStringSlice 的别名，保持向后兼容
// Deprecated: 请使用 ToStringSlice 代替
func ToArrayString(v interface{}) []string {
	return ToStringSlice(v)
}
