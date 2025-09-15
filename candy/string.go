package candy

import (
	"fmt"

	"cmp"
)

// String 将任意有序类型转换为字符串
// 该函数提供了通用的类型转字符串功能，支持所有实现了 cmp.Ordered 接口的类型
// 包括整数、浮点数和字符串等基本类型
func String[T cmp.Ordered](s T) string {
	return fmt.Sprintf("%v", s)
}
