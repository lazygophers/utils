// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sqrt 计算数值的平方根
//
// 类型参数：
//   - T: 数值类型，支持所有整数和浮点数类型
//
// 参数：
//   - s: 要计算平方根的数值
//
// 返回值：
//   - T: 输入数值的平方根
//
// 特点：
//   - 支持整数和浮点数类型，通过泛型约束确保类型安全
//   - 内部使用 math.Sqrt 进行计算，然后将结果转换回原始类型
//   - 对于负数输入，math.Sqrt 会返回 NaN
//
// 示例：
//
//	// 计算浮点数的平方根
//	result := Sqrt(16.0)
//	// result 的值为 4.0
//
//	// 计算整数的平方根
//	result := Sqrt(int(16))
//	// result 的值为 4
//
//	// 计算浮点数的平方根
//	result := Sqrt(25.5)
//	// result 约等于 5.04975246918104
func Sqrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Sqrt(float64(s)))
}
