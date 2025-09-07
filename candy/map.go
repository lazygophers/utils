package candy

// Map 对切片中的每个元素应用指定的函数，返回新的切片
// 该函数使用泛型支持任意类型的输入和输出
//
// 参数:
//   - ss: 输入切片，类型为 []T
//   - f: 映射函数，接收类型 T 的参数，返回类型 U 的结果
//
// 返回:
//   - []U: 应用映射函数后的新切片
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := Map(numbers, func(n int) int {
//	    return n * 2
//	})
//	// doubled 为 []int{2, 4, 6, 8, 10}
func Map[T, U any](ss []T, f func(T) U) (ret []U) {
	// 预分配切片容量，避免频繁扩容
	ret = make([]U, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, f(s))
	}

	return
}
