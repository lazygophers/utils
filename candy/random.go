package candy

import (
	"math/rand"
)

// Random 从切片中随机返回一个元素
// 如果切片为空，返回类型的零值
// 这是一个泛型函数，可以处理任何类型的切片
func Random[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[rand.Intn(len(ss))]
}