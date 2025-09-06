// Package min 提供 Min 函数的实现，用于查找切片中的最小值
package candy

import (
	"golang.org/x/exp/constraints"
)

// Min 返回有序类型切片中的最小值
// 如果切片为空，返回类型的零值
// 支持所有实现了 constraints.Ordered 接口的类型
func Min[T constraints.Ordered](ss []T) (min T) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}