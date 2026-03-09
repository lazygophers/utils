package candy

import "golang.org/x/exp/constraints"

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
func Remove[T constraints.Ordered](ss []T, toRemove []T) (result []T) {
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

// Diff 计算两个有序切片之间的差异
//
// 参数:
//   - ss: 第一个切片
//   - against: 第二个切片
//
// 返回值:
//   - added: 在 against 中存在但不在 ss 中的元素
//   - removed: 在 ss 中存在但不在 against 中的元素
//
// 示例:
//
//	ss := []int{1, 2, 3}
//	against := []int{2, 3, 4}
//	added, removed := Diff(ss, against)
//	// added = [4]
//	// removed = [1]
func Diff[T constraints.Ordered](ss []T, against []T) (added, removed []T) {
	removed = Remove(ss, against)
	added = Remove(against, ss)

	return
}

// Index 返回元素 sub 在切片 ss 中的索引位置
// 如果未找到，返回 -1
// 这是一个泛型函数，支持所有可排序的类型
func Index[T constraints.Ordered](ss []T, sub T) int {
	if len(ss) == 0 {
		return -1
	}

	for i, s := range ss {
		if s == sub {
			return i
		}
	}

	return -1
}

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
