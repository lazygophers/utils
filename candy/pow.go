package candy

import (
	"math"
)

// Pow 返回 x 的 y 次幂
//
// 参数:
//   - x: 底数，支持整数和浮点数类型
//   - y: 指数，支持整数和浮点数类型
//
// 返回值:
//   - T: x 的 y 次幂结果，保持原类型
//
// 特点:
//   - 支持整数和浮点数类型的幂运算
//   - 底层使用 math.Pow 进行高精度计算
//   - 自动处理类型转换，保持类型一致性
//   - 支持负指数和分数指数
//
// 示例:
//
//	// 整数幂运算
//	result := Pow(2, 3)
//	// result = 8
//
//	// 浮点数幂运算
//	result := Pow(2.5, 2.0)
//	// result = 6.25
//
//	// 负指数运算
//	result := Pow(2.0, -2.0)
//	// result = 0.25
//
//	// 分数指数运算
//	result := Pow(4.0, 0.5)
//	// result = 2.0 (平方根)
//
//	// 大数幂运算
//	result := Pow(10, 6)
//	// result = 1000000
func Pow[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}
