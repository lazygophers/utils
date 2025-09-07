// Package candy 提供了常用的工具函数和语法糖，简化日常 Go 开发中的常见操作
package candy

import (
	"golang.org/x/exp/constraints"
)

// Contains 检查切片中是否包含指定元素
//
// 泛型约束：T 必须实现 constraints.Ordered 接口（支持比较操作）
// 参数：
//   - ss: 要搜索的切片
//   - s: 要查找的元素
//
// 返回值：
//   - bool: 如果找到元素返回 true，否则返回 false
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found := Contains(numbers, 3) // 返回 true
//	found = Contains(numbers, 6)  // 返回 false
func Contains[T constraints.Ordered](ss []T, s T) bool {
	return ContainsUsing(ss, func(v T) bool {
		return s == v
	})
}
