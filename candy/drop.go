// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

// Drop 丢弃切片前 n 个元素
//
// 参数:
//   - ss: 输入切片，支持任意类型
//   - n: 要丢弃的元素个数
//
// 返回值:
//   - []T: 丢弃前 n 个元素后的新切片
//
// 特点:
//   - 支持任意类型的切片
//   - 如果 n 为负数，当作 0 处理
//   - 如果 n 大于切片长度，返回空切片
//   - 不修改原切片，返回新切片
//
// 示例:
//
//	result := Drop([]int{1, 2, 3, 4, 5}, 2)      // 返回 [3, 4, 5]
//	result := Drop([]string{"a", "b", "c"}, 0)  // 返回 ["a", "b", "c"]
//	result := Drop([]int{1, 2, 3}, 5)      // 返回 []
//	result := Drop([]int{1, 2, 3}, -1)     // 返回 [1, 2, 3]
func Drop[T any](ss []T, n int) []T {
	if n < 0 {
		n = 0
	}

	if n > len(ss) {
		n = len(ss)
	}

	return ss[n:]
}
