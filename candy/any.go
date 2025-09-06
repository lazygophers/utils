package candy

// Any 检查切片中是否存在至少一个元素满足指定条件
//
// 参数:
//   - ss: 输入切片
//   - f: 判断函数，对每个元素执行判断
//
// 返回值:
//   - bool: 如果存在任一元素满足条件则返回 true，否则返回 false
//
// 示例:
//   numbers := []int{1, 2, 3, 4, 5}
//   hasEven := Any(numbers, func(n int) bool {
//       return n%2 == 0
//   }) // 返回 true
func Any[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}

	return false
}