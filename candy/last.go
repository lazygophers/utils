// Package candy 包含实用的语法糖函数，提供简洁高效的编程工具
package candy

// Last 返回切片中的最后一个元素
// 如果切片为空，返回类型的零值
// 该函数使用泛型实现，支持任意类型的切片
//
// 参数:
//   - ss: 任意类型的切片
//
// 返回:
//   - T: 切片中的最后一个元素，如果切片为空则返回类型零值
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last := Last(numbers) // 返回 5
//
//	empty := []string{}
//	result := Last(empty) // 返回 ""
func Last[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[len(ss)-1]
}
