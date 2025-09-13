package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sqrt 计算数值的平方根
func Sqrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Sqrt(float64(s)))
}
