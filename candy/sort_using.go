// Package candy 提供了常用的工具函数和语法糖，简化日常 Go 开发中的常见操作
package candy

import (
	"sort"
)

// SortUsing 使用自定义比较函数对切片进行排序
//
// 泛型约束：T 可以是任意类型
// 参数：
//   - slice: 要排序的切片
//   - less: 比较函数，如果函数返回 true，则表示 a 应该排在 b 前面
//
// 返回值：
//   - []T: 排序后的切片副本（原始切片不会被修改）
//
// 示例：
//
//	// 降序排序
//	numbers := []int{1, 2, 3, 4, 5}
//	sorted := SortUsing(numbers, func(a, b int) bool {
//	    return a > b
//	}) // 返回 []int{5, 4, 3, 2, 1}
//
//	// 按字符串长度排序
//	words := []string{"apple", "banana", "cherry", "date"}
//	sorted := SortUsing(words, func(a, b string) bool {
//	    return len(a) < len(b)
//	}) // 返回 []string{"date", "apple", "banana", "cherry"}
func SortUsing[T any](slice []T, less func(T, T) bool) []T {
	// 如果切片长度小于2，直接返回副本
	if len(slice) < 2 {
		return slice
	}

	// 创建新的切片用于排序
	sorted := make([]T, len(slice))
	copy(sorted, slice)

	// 使用 sort.Slice 进行排序
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}