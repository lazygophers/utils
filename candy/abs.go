package candy

import (
	"golang.org/x/exp/constraints"
)

// Abs 计算数值的绝对值
func Abs[T constraints.Integer | constraints.Float](s T) T {
	if s < 0 {
		return -s
	}

	return s
}
