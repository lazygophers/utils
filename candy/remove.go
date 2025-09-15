package candy

import (
	"cmp"
)

// Remove 移除 ss 存在也 against 存在的部分
// 返回在 against 中但不在 ss 中的元素
func Remove[T cmp.Ordered](ss []T, against []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	for _, s := range ss {
		set[s] = struct{}{}
	}

	for _, s := range against {
		if _, ok := set[s]; !ok {
			result = append(result, s)
		}
	}
	return
}
