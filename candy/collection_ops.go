package candy

import (
	"math/rand/v2"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

// All 检查切片中的所有元素是否都满足指定条件
func All[T any](ss []T, f func(T) bool) bool {
	if len(ss) == 0 {
		return true
	}

	for _, s := range ss {
		if !f(s) {
			return false
		}
	}

	return true
}

// Any 检查切片中是否存在至少一个元素满足指定条件
func Any[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}

	return false
}

// Each 遍历切片并对每个元素执行指定函数
func Each[T any](values []T, fn func(value T)) {
	for _, value := range values {
		fn(value)
	}
}

// EachReverse 反向遍历切片并对每个元素执行指定函数
// 从切片的最后一个元素开始，向前遍历到第一个元素
// 对于每个元素，都会调用传入的函数 f 进行处理
//
// 参数:
//   - ss: 要遍历的切片
//   - f: 对每个元素执行的函数，接收一个类型为 T 的参数
//
// 泛型参数:
//   - T: 切片中元素的类型，可以是任意类型
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	EachReverse(numbers, func(n int) {
//	    fmt.Println(n) // 输出: 5, 4, 3, 2, 1
//	})
func EachReverse[T any](ss []T, f func(T)) {
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// Map 对切片中的每个元素应用指定的函数，返回新的切片
// 该函数使用泛型支持任意类型的输入和输出
//
// 参数:
//   - ss: 输入切片，类型为 []T
//   - f: 映射函数，接收类型 T 的参数，返回类型 U 的结果
//
// 返回:
//   - []U: 应用映射函数后的新切片
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	doubled := Map(numbers, func(n int) int {
//	    return n * 2
//	})
//	// doubled 为 []int{2, 4, 6, 8, 10}
func Map[T, U any](ss []T, f func(T) U) []U {
	n := len(ss)
	if n == 0 {
		return []U{}
	}

	// 直接分配最终长度的切片，避免 append 操作
	ret := make([]U, n)
	// 使用传统 for 循环避免 range 的值拷贝开销
	for i := 0; i < n; i++ {
		ret[i] = f(ss[i])
	}

	return ret
}

// Reduce 对切片进行归约操作，使用指定的二元函数将切片元素合并为单个值
func Reduce[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}

	result := ss[0]
	for _, s := range ss[1:] {
		result = f(result, s)
	}
	return result
}

// Reverse 返回一个反转后的切片，原切片保持不变
// 该函数使用泛型支持任意类型的切片，返回一个新的反转后的切片
// 采用混合策略优化性能：小切片使用直接索引，大切片使用双指针交换
func Reverse[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}

	result := make([]T, n)

	// 小切片使用直接索引赋值，性能最优
	if n < 32 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	// 大切片使用双指针交换优化
	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Shuffle 随机打乱切片中的元素顺序
//
// 类型参数：
//   - T: 任意类型
//
// 参数：
//   - ss: 待打乱的切片，可以是任意类型
//
// 返回值：
//   - []T: 打乱后的切片（原地修改，返回原切片的引用）
//
// 特点：
//   - 使用 Fisher-Yates 洗牌算法，确保均匀随机分布
//   - 原地修改，不创建新切片，内存效率高
//   - 支持任意类型的切片
//   - 对于空切片或单元素切片，直接返回原切片
//   - 高性能优化：使用 randx 包的高性能随机数生成器
//
// 示例：
//
//	// 打乱整数切片
//	data := []int{1, 2, 3, 4, 5}
//	result := Shuffle(data)
//	// result 是打乱后的切片，与 data 是同一个切片
//
//	// 打乱字符串切片
//	names := []string{"Alice", "Bob", "Charlie", "David"}
//	shuffled := Shuffle(names)
//	// shuffled 包含随机顺序的名字
func Shuffle[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	// 使用标准库 rand.Shuffle 获得最佳性能
	// 基准测试显示在大多数情况下，特别是大切片上性能最优
	rand.Shuffle(n, func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})

	return ss
}

// Sort 对有序类型的切片进行排序
// 接受一个实现了 constraints.Ordered 接口的切片，返回一个新的已排序切片
// 原始切片不会被修改，返回的是排序后的副本
//
// 性能优化：
//   - 小切片（≤24元素）：使用插入排序，避免快速排序的递归开销
//   - 大切片（>24元素）：使用标准 sort.Slice
//   - 该阈值在常见用例中提供最佳性能平衡
func Sort[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	// 创建新切片用于排序
	sorted := make([]T, n)
	copy(sorted, ss)

	// 小切片使用插入排序（性能更优）
	// 阈值24是性能测试得出的最佳平衡点
	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	// 大切片使用标准排序
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortUsing 使用自定义比较函数对切片进行排序
func SortUsing[T any](slice []T, less func(T, T) bool) []T {
	if len(slice) < 2 {
		return slice
	}

	sorted := make([]T, len(slice))
	copy(sorted, slice)

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// Join 将有序类型的切片按指定分隔符连接成字符串
// 该函数提供了通用的切片连接功能，支持所有实现了 constraints.Ordered 接口的类型
// 包括整数、浮点数和字符串等基本类型
//
// 参数:
//   - ss: 输入切片，类型为 []T，其中 T 必须实现 constraints.Ordered 接口
//   - glue: 可选参数，指定连接分隔符，默认为 ","
//
// 返回:
//   - string: 连接后的字符串
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	result := Join(numbers, "-")
//	// result 为 "1-2-3-4-5"
//
//	words := []string{"Hello", "World", "Go"}
//	result := Join(words, " ")
//	// result 为 "Hello World Go"
func Join[T constraints.Ordered](ss []T, glue ...string) string {
	// 设置默认分隔符
	seq := ","
	if len(glue) > 0 {
		seq = glue[0]
	}

	// 使用 Map 函数将切片元素转换为字符串，然后用 strings.Join 连接
	return strings.Join(Map(ss, func(s T) string {
		return ToString(s)
	}), seq)
}
