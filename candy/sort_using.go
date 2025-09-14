package candy

import (
	"sort"
)

// SortUsing 使用自定义比较函数对切片进行排序
func SortUsing[T any](slice []T, less func(T, T) bool) []T {
	if len(slice) < 2 {
		return slice
	}

	sorted := make([]T, len(slice))
	copy(sorted, slice)

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}
