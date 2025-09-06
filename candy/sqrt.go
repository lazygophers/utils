// Package candy 提供 Go 语法糖工具函数
// 
// 本文件包含数学基础函数 Sqrt，用于计算各种数值类型的平方根
package candy

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sqrt 计算数值的平方根
// 
// 该函数支持整数和浮点数类型，通过泛型约束确保类型安全
// 内部使用 math.Sqrt 进行计算，然后将结果转换回原始类型
//
// 类型参数:
//   T: 数值类型，支持所有整数和浮点数类型
//
// 参数:
//   s: 要计算平方根的数值
//
// 返回值:
//   T: 输入数值的平方根
//
// 示例:
//   result := Sqrt(16.0)  // 返回 4.0
//   result := Sqrt(int(16))  // 返回 4
func Sqrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Sqrt(float64(s)))
}