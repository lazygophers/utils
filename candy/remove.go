package candy

import (
	"cmp"
)

// Remove 从第一个切片中移除第二个切片中存在的元素
// 返回第一个切片中但不在第二个切片中的元素
//
// 参数:
//   - ss: 源切片
//   - toRemove: 要移除的元素列表
//
// 返回值:
//   - result: ss 中移除了 toRemove 中存在元素后的结果
//
// 示例:
//
//	ss := []int{1, 2, 3, 4, 5}
//	toRemove := []int{2, 4, 6}
//	result := Remove(ss, toRemove)
//	// result = [1, 3, 5]
func Remove[T cmp.Ordered](ss []T, toRemove []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	removeSet := make(map[T]struct{}, len(toRemove))

	// 构建要移除元素的集合
	for _, item := range toRemove {
		removeSet[item] = struct{}{}
	}

	// 遍历原切片，只保留不在移除集合中的元素
	for _, item := range ss {
		if _, shouldRemove := removeSet[item]; !shouldRemove {
			result = append(result, item)
		}
	}
	return
}
