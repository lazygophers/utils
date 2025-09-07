// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Cbrt 计算数值的立方根
//
// 参数:
//   - s: 输入数值，支持整数和浮点数
//
// 返回值:
//   - T: 输入数值的立方根
//
// 特点:
//   - 支持整数和浮点数类型
//   - 通过泛型实现类型安全
//   - 对于负数也返回正确结果
//
// 示例:
//
//	result := Cbrt(8.0)      // 返回 2.0
//	result := Cbrt(27)       // 返回 3
//	result := Cbrt(-8.0)     // 返回 -2.0
func Cbrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Cbrt(float64(s)))
}
