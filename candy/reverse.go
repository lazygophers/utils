// Package reverse 提供了反转切片的工具函数
package candy

// Reverse 返回一个反转后的切片，原切片保持不变
// 该函数使用泛型支持任意类型的切片，返回一个新的反转后的切片
// 使用 make 预分配容量以确保性能最优
func Reverse[T any](ss []T) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0, len(ss))
	for i := len(ss) - 1; i >= 0; i-- {
		ret = append(ret, ss[i])
	}

	return
}
