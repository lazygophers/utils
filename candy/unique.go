// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

import (
	"cmp"
)

// Unique 返回切片中的唯一元素，保持原始顺序
//
// 类型参数：
//   - T: 可排序的类型，支持所有可比较的内置类型
//
// 参数：
//   - ss: 输入切片，包含可排序类型的元素
//
// 返回值：
//   - []T: 去重后的切片，保持原始顺序
//
// 特点：
//   - 保持原始元素的顺序
//   - 使用 map 高效去重，时间复杂度 O(n)
//   - 空切片安全，返回空切片而非 nil
//   - 只支持可排序类型，确保类型安全
//
// 示例：
//
//	// 对整数切片去重
//	numbers := []int{1, 2, 2, 3, 4, 4, 5}
//	unique := Unique(numbers)
//	// unique 的值为 []int{1, 2, 3, 4, 5}
//
//	// 对字符串切片去重
//	names := []string{"Alice", "Bob", "Alice", "Charlie", "Bob"}
//	uniqueNames := Unique(names)
//	// uniqueNames 的值为 []string{"Alice", "Bob", "Charlie"}
//
//	// 对浮点数切片去重
//	floats := []float64{1.1, 2.2, 1.1, 3.3, 2.2}
//	uniqueFloats := Unique(floats)
//	// uniqueFloats 的值为 []float64{1.1, 2.2, 3.3}
func Unique[T cmp.Ordered](ss []T) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0)
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}

	return
}
