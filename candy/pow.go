package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Pow 返回 x 的 y 次幂
// 支持整数和浮点数类型，使用 math.Pow 进行计算并转换回原类型
func Pow[T constraints.Integer | constraints.Float](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}
