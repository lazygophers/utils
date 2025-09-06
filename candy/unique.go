package candy

import (
	"golang.org/x/exp/constraints"
)

// Unique 返回切片中的唯一元素，保持原始顺序
// 使用 map 来跟踪已出现的元素，确保返回的切片中每个元素只出现一次
// 使用 make 初始化，确保返回空切片而非 nil
func Unique[T constraints.Ordered](ss []T) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0)
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}

	return
}
