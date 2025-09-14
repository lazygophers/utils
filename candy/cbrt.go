package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Cbrt 计算数值的立方根
func Cbrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Cbrt(float64(s)))
}
