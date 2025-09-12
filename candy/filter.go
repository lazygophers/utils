// Package candy 提供了 Go 语言中常用的语法糖函数和工具方法
package candy

// Filter 函数用于过滤切片中的元素，返回满足条件的所有元素
//
// 参数：
//   - ss: 要过滤的切片，可以是任意类型的切片
//   - f: 过滤函数，接收一个元素并返回布尔值，true 表示保留该元素
//
// 返回值：
//   - []T: 包含所有满足条件的元素的新切片
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	evenNumbers := Filter(numbers, func(n int) bool {
//	    return n%2 == 0
//	})
//	// evenNumbers = [2, 4]
//
// 注意事项：
//   - 使用 make 初始化返回切片，确保返回空切片而非 nil
//   - 该函数不会修改原始切片
//   - 时间复杂度为 O(n)，其中 n 为切片长度
func Filter[T any](ss []T, f func(T) bool) []T {
	if len(ss) == 0 {
		return []T{}
	}
	
	// 使用原始长度的1/4作为初始容量预估，减少重新分配
	ret := make([]T, 0, len(ss)/4+1)
	for _, s := range ss {
		if f(s) {
			ret = append(ret, s)
		}
	}

	return ret
}
