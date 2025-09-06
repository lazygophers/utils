// Package max 提供查找切片最大值的功能
package candy

import "golang.org/x/exp/constraints"

// Max 返回切片中的最大值
// 如果切片为空，返回类型的零值
// 支持所有可比较的类型（整数、浮点数、字符串等）
func Max[T constraints.Ordered](ss []T) (max T) {
	if len(ss) == 0 {
		return
	}
	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}
	return
}
