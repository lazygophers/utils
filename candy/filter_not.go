// Package candy 提供实用的语法糖函数，简化常见的编程操作
package candy

// FilterNot 对切片进行反向过滤，保留不满足指定条件的元素
// 该函数是 Filter 函数的补集，用于语义上更清晰的反向过滤操作
//
// 参数:
//   - ss: 输入切片，可以是任意类型的切片
//   - f: 谓词函数，接收一个元素并返回布尔值
//
// 返回值:
//   - []T: 包含所有不满足谓词函数条件的新切片
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
