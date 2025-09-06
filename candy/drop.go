// Package drop 提供 Drop 函数，用于丢弃切片前 n 个元素
package candy

// Drop 丢弃切片前 n 个元素
// 如果 n 为负数，当作 0 处理
// 如果 n 大于切片长度，返回空切片
// 该函数不修改原切片，而是返回一个新的切片
func Drop[T any](ss []T, n int) []T {
	if n < 0 {
		n = 0
	}

	if n > len(ss) {
		n = len(ss)
	}

	return ss[n:]
}
