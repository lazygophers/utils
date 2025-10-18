package candy

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

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

// Min 返回切片中的最小值
//
// 参数:
//   - ss: 输入切片，支持所有可排序类型
//
// 返回值:
//   - T: 切片中的最小值，如果切片为空则返回类型零值
//
// 特点:
//   - 支持整数、浮点数、字符串等所有可排序类型
//   - 空切片输入返回类型零值
//   - 时间复杂度 O(n)，单次遍历
//   - 不修改原切片
//
// 示例:
//
//	// 查找整数切片的最小值
//	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
//	result := Min(numbers)
//	// result = 1
//
//	// 查找字符串切片的最小值（按字典序）
//	strings := []string{"apple", "banana", "cherry", "date"}
//	result := Min(strings)
//	// result = "apple"
//
//	// 查找浮点数切片的最小值
//	floats := []float64{3.14, 1.618, 2.718, 1.414}
//	result := Min(floats)
//	// result = 1.414
//
//	// 空切片处理
//	empty := []int{}
//	result := Min(empty)
//	// result = 0 (int类型的零值)
func Min[T cmp.Ordered](ss []T) (min T) {
	if len(ss) == 0 {
		return
	}
	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}
	return
}

// Sum 计算数值切片中所有元素的总和
// 支持整数和浮点数类型，使用泛型实现类型安全
//
// 参数：
//   - ss: 数值切片，支持整数和浮点数类型
//
// 返回值：
//   - T: 切片中所有元素的总和
//
// 示例：
//
//	sum := Sum([]int{1, 2, 3})  // 返回 6
//	sum := Sum([]float64{1.5, 2.5})  // 返回 4.0
func Sum[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}](ss ...T) (ret T) {
	for _, s := range ss {
		ret += s
	}

	return
}

// Average 计算数值切片的平均值
func Average[T constraints.Integer | constraints.Float](ss ...T) (ret T) {
	if len(ss) == 0 {
		return
	}

	var sum float64
	for _, s := range ss {
		sum += float64(s)
	}
	return T(sum / float64(len(ss)))
}

// Abs 计算数值的绝对值
func Abs[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}](s T) T {
	if s < 0 {
		return -s
	}

	return s
}
