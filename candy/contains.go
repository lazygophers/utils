package candy

import (
	"cmp"
)

// Contains 检查切片中是否包含指定元素
func Contains[T cmp.Ordered](ss []T, s T) bool {
	return ContainsUsing(ss, func(v T) bool {
		return s == v
	})
}
