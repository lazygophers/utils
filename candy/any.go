// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

// Any 检查切片中是否存在至少一个元素满足指定条件
//
// 参数:
//   - ss: 输入切片，可以是任意类型的切片
//   - f: 判断函数，对每个元素执行判断
//
// 返回值:
//   - bool: 如果存在任一元素满足条件则返回 true，否则返回 false
//
// 特点:
//   - 空切片或 nil 切片返回 false
//   - 遇到第一个满足条件的元素时立即返回，提高效率
//   - 支持任意类型的切片，通过泛型实现类型安全
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := Any(numbers, func(n int) bool {
//	    return n > 3
//	}) // 返回 true
//
//	result := Any(numbers, func(n int) bool {
//	    return n > 10
//	}) // 返回 false
func Any[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}

	return false
}
