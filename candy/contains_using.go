// Package candy 提供了常用的工具函数和语法糖，简化日常 Go 开发中的常见操作
package candy

// ContainsUsing 使用自定义函数检查切片中是否包含满足条件的元素
//
// 泛型约束：T 可以是任意类型
// 参数：
//   - ss: 要搜索的切片
//   - f: 判断函数，如果函数返回 true，则表示找到匹配的元素
//
// 返回值：
//   - bool: 如果找到满足条件的元素返回 true，否则返回 false
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found := ContainsUsing(numbers, func(v int) bool {
//	    return v > 3
//	}) // 返回 true
func ContainsUsing[T any](ss []T, f func(v T) bool) bool {
	// 遍历切片中的每个元素
	for _, v := range ss {
		// 使用自定义函数判断当前元素是否满足条件
		if f(v) {
			return true
		}
	}

	// 遍历完所有元素都没有找到满足条件的，返回 false
	return false
}
