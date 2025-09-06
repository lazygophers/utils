// Package abs 提供绝对值计算功能
// 
// 该包包含 Abs 函数，支持整数和浮点数类型的绝对值计算
package candy

import (
	"golang.org/x/exp/constraints"
)

// Abs 计算数值的绝对值
// 
// 泛型函数，支持所有整数和浮点数类型
// 对于负数返回其相反数，对于正数和零返回其本身
// 
// 类型参数:
//   - T: 约束为 Integer 或 Float 的数值类型
// 
// 参数:
//   - s: 输入的数值
// 
// 返回值:
//   - T: 输入数值的绝对值
// 
// 示例:
//   result := Abs(-42)     // 返回 42
//   result := Abs(3.14)   // 返回 3.14
//   result := Abs(0)      // 返回 0
func Abs[T constraints.Integer | constraints.Float](s T) T {
	if s < 0 {
		return -s
	}

	return s
}