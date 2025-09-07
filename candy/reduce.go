// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

// Reduce 对切片进行归约操作，使用指定的二元函数将切片元素合并为单个值
//
// 类型参数：
//   - T: 任意类型
//
// 参数：
//   - ss: 输入切片，可以是任意类型
//   - f: 归约函数，接受两个同类型参数，返回同类型结果
//
// 返回值：
//   - 归约后的单个值，对于空切片返回类型的零值
//
// 特点：
//   - 支持任意类型的切片归约操作
//   - 空切片安全，返回类型零值
//   - 使用泛型确保类型安全
//
// 示例：
//
//	// 计算切片元素的和
//	sum := Reduce([]int{1, 2, 3, 4, 5}, func(a, b int) int { return a + b })
//	// sum 的值为 15
//
//	// 计算切片元素的乘积
//	product := Reduce([]int{1, 2, 3, 4}, func(a, b int) int { return a * b })
//	// product 的值为 24
//
//	// 找出切片中的最大值
//	max := Reduce([]int{3, 1, 4, 1, 5}, func(a, b int) int {
//	    if b > a { return b }
//	    return a
//	})
//	// max 的值为 5
func Reduce[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}

	result := ss[0]
	for _, s := range ss[1:] {
		result = f(result, s)
	}
	return result
}
