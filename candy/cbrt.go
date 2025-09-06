// Package candy 包含数学基础函数的实现
package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Cbrt 计算数值的立方根
// 支持整数和浮点数类型，使用 math.Cbrt 进行计算
func Cbrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Cbrt(float64(s)))
}