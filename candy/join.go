package candy

import (
	"strings"

	"cmp"
)

// Join 将有序类型的切片按指定分隔符连接成字符串
// 该函数提供了通用的切片连接功能，支持所有实现了 cmp.Ordered 接口的类型
// 包括整数、浮点数和字符串等基本类型
//
// 参数:
//   - ss: 输入切片，类型为 []T，其中 T 必须实现 cmp.Ordered 接口
//   - glue: 可选参数，指定连接分隔符，默认为 ","
//
// 返回:
//   - string: 连接后的字符串
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := Join(numbers, "-")
//	// result 为 "1-2-3-4-5"
//
//	words := []string{"Hello", "World", "Go"}
//	result := Join(words, " ")
//	// result 为 "Hello World Go"
func Join[T cmp.Ordered](ss []T, glue ...string) string {
	// 设置默认分隔符
	seq := ","
	if len(glue) > 0 {
		seq = glue[0]
	}

	// 使用 Map 函数将切片元素转换为字符串，然后用 strings.Join 连接
	return strings.Join(Map(ss, func(s T) string {
		return ToString(s)
	}), seq)
}
