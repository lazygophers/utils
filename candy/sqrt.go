package candy

import (
	"math"
)

// Sqrt 计算数值的平方根
func Sqrt[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}](s T) T {
	return T(math.Sqrt(float64(s)))
}
