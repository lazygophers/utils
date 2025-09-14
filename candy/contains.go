package candy

import (
	"golang.org/x/exp/constraints"
)

// Contains 检查切片中是否包含指定元素
func Contains[T constraints.Ordered](ss []T, s T) bool {
	return ContainsUsing(ss, func(v T) bool {
		return s == v
	})
}
