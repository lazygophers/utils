package candy

import (
	"cmp"
)

// Index 返回元素 sub 在切片 ss 中的索引位置
// 如果未找到，返回 -1
// 这是一个泛型函数，支持所有可排序的类型
func Index[T cmp.Ordered](ss []T, sub T) int {
	if len(ss) == 0 {
		return -1
	}

	for i, s := range ss {
		if s == sub {
			return i
		}
	}

	return -1
}
