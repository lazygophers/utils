package candy

import "golang.org/x/exp/constraints"

// Spare 返回在 against 中但不在 ss 中的元素
// 该函数实现了集合差集操作，返回在 against 切片中存在但在 ss 切片中不存在的所有元素
// 注意：该函数与 Remove 函数功能相同，都是返回差集结果
//
// 参数：
//   - ss: 作为参考集合的切片
//   - against: 作为被比较集合的切片
//
// 返回：
//   - []T: 在 against 中但不在 ss 中的元素组成的切片
//
// 示例：
//
//	ss := []int{1, 2, 3}
//	against := []int{2, 3, 4, 5}
//	result := Spare(ss, against) // 返回 [4, 5]
func Spare[T constraints.Ordered](ss []T, against []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	// 将 ss 中的所有元素添加到 map 中用于快速查找
	for _, s := range ss {
		set[s] = struct{}{}
	}

	// 遍历 against 切片，找出不在 ss 中的元素
	for _, s := range against {
		if _, ok := set[s]; !ok {
			result = append(result, s)
		}
	}
	return
}
