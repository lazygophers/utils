package candy

import "golang.org/x/exp/constraints"

// Average 计算数值切片的平均值
func Average[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	var sum float64
	for _, s := range ss {
		sum += float64(s)
	}
	return T(sum / float64(len(ss)))
}
