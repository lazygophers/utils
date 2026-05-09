package candy

import (
	"time"

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
func Max[T constraints.Ordered](ss ...T) (max T) {
	n := len(ss)
	if n == 0 {
		return
	}
	// 优化：使用索引循环避免 range 的额外开销
	max = ss[0]
	for i := 1; i < n; i++ {
		// 优化：减少边界检查，先存入局部变量
		v := ss[i]
		if v > max {
			max = v
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
func Min[T constraints.Ordered](ss ...T) (min T) {
	n := len(ss)
	if n == 0 {
		return
	}
	// 优化：使用索引循环避免 range 的额外开销
	min = ss[0]
	for i := 1; i < n; i++ {
		// 优化：减少边界检查，先存入局部变量
		v := ss[i]
		if v < min {
			min = v
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
func Sum[T constraints.Ordered](ss ...T) (ret T) {
	// 优化：SIMD友好的4路累加，减少依赖链
	n := len(ss)
	var sum1, sum2, sum3, sum4 T
	i := 0
	// 每次处理4个元素，避免数据依赖
	for i = 0; i+4 <= n; i += 4 {
		sum1 += ss[i]
		sum2 += ss[i+1]
		sum3 += ss[i+2]
		sum4 += ss[i+3]
	}
	// 处理剩余元素
	for j := i; j < n; j++ {
		ret += ss[j]
	}
	ret += sum1 + sum2 + sum3 + sum4
	return
}

// Average 计算数值切片的平均值
func Average[T constraints.Integer | constraints.Float | time.Duration](ss ...T) (ret T) {
	n := len(ss)
	if n == 0 {
		return 0
	}
	// 优化：单次遍历，SIMD友好的4路累加
	var sum1, sum2, sum3, sum4 T
	i := 0
	for i = 0; i+4 <= n; i += 4 {
		sum1 += ss[i]
		sum2 += ss[i+1]
		sum3 += ss[i+2]
		sum4 += ss[i+3]
	}
	for j := i; j < n; j++ {
		sum1 += ss[j]
	}
	ret = (sum1 + sum2 + sum3 + sum4) / T(n)
	return
}

// Abs 计算数值的绝对值
func Abs[T constraints.Integer | constraints.Float](s T) T {
	if s < 0 {
		return -s
	}

	return s
}
