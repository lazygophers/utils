package candy

import (
	"math"
)

// Cbrt 计算数值的立方根
func Cbrt[T interface{ ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 }](s T) T {
	return T(math.Cbrt(float64(s)))
}
