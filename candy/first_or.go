// Package candy 包含语法糖工具函数，提供便捷的编程辅助功能
package candy

// FirstOr 返回切片的第一个元素，如果切片为空则返回指定的默认值
//
// 该函数使用泛型支持任意类型的切片，提供了安全的空切片处理机制。
// 在访问切片第一个元素之前，会先检查切片长度，避免 panic。
//
// 参数:
//   - ss: 任意类型的切片
//   - or: 当切片为空时返回的默认值
//
// 返回:
//   - 切片的第一个元素，如果切片为空则返回默认值
//
// 示例:
//
//	numbers := []int{1, 2, 3}
//	first := FirstOr(numbers, 0)     // 返回 1
//
//	empty := []int{}
//	defaultVal := FirstOr(empty, 0) // 返回 0
func FirstOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[0]
}
