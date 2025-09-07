// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

import (
	"math/rand"
)

// Random 从切片中随机返回一个元素
//
// 类型参数：
//   - T: 任意类型
//
// 参数：
//   - ss: 输入切片
//
// 返回值：
//   - ret: 随机选择的元素，如果切片为空则返回类型的零值
//
// 特点：
//   - 支持任意类型的切片
//   - 空切片安全，返回类型零值
//   - 使用 math/rand 包生成随机数
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	randomNum := Random(numbers)
//	// randomNum 可能是 1, 2, 3, 4, 5 中的任意一个值
//
//	strings := []string{"apple", "banana", "cherry"}
//	randomStr := Random(strings)
//	// randomStr 可能是 "apple", "banana", "cherry" 中的任意一个
func Random[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[rand.Intn(len(ss))]
}
