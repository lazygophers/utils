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
//
// 性能优化：
//   - 快速路径：空切片或无需移除直接返回
//   - 两遍扫描：第一遍计数，第二遍填充（避免 append 重新分配）
//   - 使用 map[T]bool 替代 map[T]struct{}（略微更快的查找）
func Remove[T constraints.Ordered](ss []T, toRemove []T) (result []T) {
	// 快速路径：空切片
	if len(ss) == 0 {
		return make([]T, 0)
	}

	// 快速路径：无需移除
	if len(toRemove) == 0 {
		result = make([]T, len(ss))
		copy(result, ss)
		return result
	}

	// 构建移除集合
	removeSet := make(map[T]bool, len(toRemove))
	for _, item := range toRemove {
		removeSet[item] = true
	}

	// 第一遍扫描：计算保留元素数量
	count := 0
	for _, item := range ss {
		if !removeSet[item] {
			count++
		}
	}

	// 快速路径：所有元素都被移除
	if count == 0 {
		return make([]T, 0)
	}

	// 第二遍扫描：直接填充预分配的切片
	result = make([]T, count)
	idx := 0
	for _, item := range ss {
		if !removeSet[item] {
			result[idx] = item
			idx++
		}
	}
	return result
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
//
// 性能优化：
//   - 使用单一 map 构建差异集合，避免两次 Remove 调用
//   - 预分配结果切片容量，减少内存重新分配
//   - 使用 map[T]int 计数，正确处理重复元素
func Diff[T constraints.Ordered](ss []T, against []T) (added, removed []T) {
	// 快速路径：空切片
	if len(ss) == 0 && len(against) == 0 {
		return []T{}, []T{}
	}

	// 构建两个 map 计数
	mapSS := make(map[T]int, len(ss))
	for _, v := range ss {
		mapSS[v]++
	}

	mapAgainst := make(map[T]int, len(against))
	for _, v := range against {
		mapAgainst[v]++
	}

	// 预分配结果切片
	removed = make([]T, 0, len(mapSS))
	added = make([]T, 0, len(mapAgainst))

	// 找出在 ss 中但不在 against 中的
	for k, v := range mapSS {
		if mapAgainst[k] == 0 {
			for i := 0; i < v; i++ {
				removed = append(removed, k)
			}
		}
	}

	// 找出在 against 中但不在 ss 中的
	for k, v := range mapAgainst {
		if mapSS[k] == 0 {
			for i := 0; i < v; i++ {
				added = append(added, k)
			}
		}
	}

	return added, removed
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
//
// 性能优化：
//   - 预分配结果切片容量，减少 append 重新分配
//   - 使用 map[T]struct{} 作为集合，内存占用最小
//   - 快速路径：处理空切片情况
func Same[T constraints.Ordered](against []T, ss []T) (result []T) {
	// 快速路径：空切片
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	// 构建集合
	set := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	// 预分配结果切片：最大可能大小是 min(len(against), len(ss))
	maxSize := len(against)
	if len(ss) < maxSize {
		maxSize = len(ss)
	}
	result = make([]T, 0, maxSize)

	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return
}
