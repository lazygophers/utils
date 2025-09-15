package candy

import "cmp"

// Max 返回切片中的最大值
//
// 参数:
//   - ss: 输入切片，支持所有可排序类型
//
// 返回值:
//   - T: 切片中的最大值，如果切片为空则返回类型零值
//
// 特点:
//   - 支持整数、浮点数、字符串等所有可排序类型
//   - 空切片输入返回类型零值
//   - 时间复杂度 O(n)，单次遍历
//   - 不修改原切片
//
// 示例:
//
//	// 查找整数切片的最大值
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	result := Max(numbers)
//	// result = 9
//
//	// 查找字符串切片的最大值（按字典序）
//	strings := []string{"apple", "banana", "cherry", "date"}
//	result := Max(strings)
//	// result = "date"
//
//	// 查找浮点数切片的最大值
//	floats := []float64{3.14, 1.618, 2.718, 1.414}
//	result := Max(floats)
//	// result = 3.14
//
//	// 空切片处理
//	empty := []int{}
//	result := Max(empty)
//	// result = 0 (int类型的零值)
func Max[T cmp.Ordered](ss []T) (max T) {
	if len(ss) == 0 {
		return
	}
	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}
	return
}
