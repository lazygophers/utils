// Package candy 包含实用的语法糖函数，提供简洁高效的编程工具
package candy

// LastOr 返回切片中的最后一个元素
// 如果切片为空，返回指定的默认值
// 该函数使用泛型实现，支持任意类型的切片
//
// 参数:
//   - ss: 任意类型的切片
//   - or: 当切片为空时返回的默认值
//
// 返回:
//   - T: 切片中的最后一个元素，如果切片为空则返回指定的默认值
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last := LastOr(numbers, 0) // 返回 5
//
//	empty := []string{}
//	result := LastOr(empty, "default") // 返回 "default"
func LastOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[len(ss)-1]
}
