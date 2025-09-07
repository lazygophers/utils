// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

// All 检查切片中的所有元素是否都满足指定的条件函数
//
// 参数:
//   - ss: 输入切片，可以是任意类型的切片
//   - f: 条件函数，对每个元素进行判断
//
// 返回值:
//   - bool: 如果所有元素都满足条件则返回 true，否则返回 false
//
// 特点:
//   - 空切片或 nil 切片返回 true
//   - 遇到第一个不满足条件的元素时立即返回，提高效率
//   - 支持任意类型的切片，通过泛型实现类型安全
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := All(numbers, func(n int) bool {
//	    return n > 0
//	}) // 返回 true
//
//	result := All(numbers, func(n int) bool {
//	    return n > 5
//	}) // 返回 false
func All[T any](ss []T, f func(T) bool) bool {
	// 空切片或 nil 切片被认为是满足条件的
	if len(ss) == 0 {
		return true
	}

	// 遍历切片，对每个元素应用条件函数
	for _, s := range ss {
		if !f(s) {
			// 发现不满足条件的元素，立即返回 false
			return false
		}
	}

	// 所有元素都满足条件，返回 true
	return true
}
