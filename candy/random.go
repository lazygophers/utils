package candy

import (
	"math/rand"
)

// Random 从切片中随机返回一个元素
func Random[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[rand.Intn(len(ss))]
}
