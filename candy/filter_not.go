// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

// FilterNot 对切片进行反向过滤，保留不满足指定条件的元素
//
// 参数:
//   - ss: 输入切片，支持任意类型
//   - f: 谓词函数，接收一个元素并返回布尔值
//
// 返回值:
//   - []T: 包含所有不满足谓词函数条件的新切片
//
// 特点:
//   - 支持任意类型的切片
//   - 是 Filter 函数的补集，用于语义上更清晰的反向过滤操作
//   - 不修改原切片，返回新切片
//   - 空切片输入返回空切片
//
// 示例:
//
//	// 过滤偶数，保留奇数
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	result := FilterNot(numbers, func(n int) bool {
//	    return n % 2 == 0
//	})
//	// result = [1, 3, 5]
//
//	// 过滤空字符串，保留非空字符串
//	strings := []string{"hello", "", "world", ""}
//	result := FilterNot(strings, func(s string) bool {
//	    return s == ""
//	})
//	// result = ["hello", "world"]
//
//	// 过滤负数，保留非负数
//	nums := []int{-1, 0, 1, -2, 2}
//	result := FilterNot(nums, func(n int) bool {
//	    return n < 0
//	})
//	// result = [0, 1, 2]
func FilterNot[T any](ss []T, f func(T) bool) []T {
	// 使用 make 初始化，确保返回空切片而非 nil
	us := make([]T, 0)
	for _, s := range ss {
		if !f(s) {
			us = append(us, s)
		}
	}
	return us
}
