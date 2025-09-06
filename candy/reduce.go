// Package reduce 提供切片归约操作功能
package candy

// Reduce 对切片进行归约操作，使用指定的二元函数将切片元素合并为单个值
// 对于空切片，返回类型的零值
//
// 参数:
//   - ss: 输入切片，可以是任意类型
//   - f: 归约函数，接受两个同类型参数，返回同类型结果
//
// 返回:
//   - 归约后的单个值
//
// 示例:
//   sum := Reduce([]int{1, 2, 3, 4, 5}, func(a, b int) int { return a + b }) // 返回 15
//   product := Reduce([]int{1, 2, 3, 4}, func(a, b int) int { return a * b }) // 返回 24
//   max := Reduce([]int{3, 1, 4, 1, 5}, func(a, b int) int { 
//       if b > a { return b }; return a 
//   }) // 返回 5
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