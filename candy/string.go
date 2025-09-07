package candy

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// String 将任意有序类型转换为字符串
// 该函数提供了通用的类型转字符串功能，支持所有实现了 constraints.Ordered 接口的类型
// 包括整数、浮点数和字符串等基本类型
func String[T constraints.Ordered](s T) string {
	return fmt.Sprintf("%v", s)
}
