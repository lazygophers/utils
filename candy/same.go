package candy

import (
	"golang.org/x/exp/constraints"
)

// Same 返回在 against 和 ss 中都存在的元素
// 用于查找两个有序集合的交集
func Same[T constraints.Ordered](against []T, ss []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	for _, s := range ss {
		set[s] = struct{}{}
	}

	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return
}
