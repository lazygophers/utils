package candy

import (
	"sort"

	"cmp"
)

// Sort 对有序类型的切片进行排序
// 接受一个实现了 cmp.Ordered 接口的切片，返回一个新的已排序切片
// 原始切片不会被修改，返回的是排序后的副本
func Sort[T cmp.Ordered](ss []T) []T {
	// 如果切片长度小于2，直接返回副本
	if len(ss) < 2 {
		return ss
	}

	// 创建新的切片用于排序
	sorted := make([]T, len(ss))
	copy(sorted, ss)

	// 使用 sort.Slice 进行排序
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}
