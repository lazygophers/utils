// Package candy 提供计算平均值的函数
package candy

import "golang.org/x/exp/constraints"

// Average 计算数值切片的平均值
// 支持整数和浮点数类型，使用 float64 计算以保持精度
// 如果切片为空，返回类型的零值
func Average[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	// 使用 float64 计算以保持精度
	var sum float64
	for _, s := range ss {
		sum += float64(s)
	}
	return T(sum / float64(len(ss)))
}
