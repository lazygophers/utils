package candy

import "golang.org/x/exp/constraints"

// Drop 丢弃切片前 n 个元素
//
// 参数:
//   - ss: 输入切片，支持任意类型
//   - n: 要丢弃的元素个数
//
// 返回值:
//   - []T: 丢弃前 n 个元素后的新切片
//
// 特点:
//   - 支持任意类型的切片
//   - 如果 n 为负数，当作 0 处理
//   - 如果 n 大于切片长度，返回空切片
//   - 不修改原切片，返回新切片
//
// 示例:
//
//	result := Drop([]int{1, 2, 3, 4, 5}, 2)      // 返回 [3, 4, 5]
//	result := Drop([]string{"a", "b", "c"}, 0)  // 返回 ["a", "b", "c"]
//	result := Drop([]int{1, 2, 3}, 5)      // 返回 []
//	result := Drop([]int{1, 2, 3}, -1)     // 返回 [1, 2, 3]
func Drop[T any](ss []T, n int) []T {
	if n < 0 {
		n = 0
	}

	if n > len(ss) {
		n = len(ss)
	}

	return ss[n:]
}

// Filter 函数用于过滤切片中的元素，返回满足条件的所有元素
//
// 参数：
//   - ss: 要过滤的切片，可以是任意类型的切片
//   - f: 过滤函数，接收一个元素并返回布尔值，true 表示保留该元素
//
// 返回值：
//   - []T: 包含所有满足条件的元素的新切片
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	evenNumbers := Filter(numbers, func(n int) bool {
//	    return n%2 == 0
//	})
//	// evenNumbers = [2, 4]
//
// 注意事项：
//   - 使用 make 初始化返回切片，确保返回空切片而非 nil
//   - 该函数不会修改原始切片
//   - 时间复杂度为 O(n)，其中 n 为切片长度
func Filter[T any](ss []T, f func(T) bool) []T {
	if len(ss) == 0 {
		return []T{}
	}

	// 使用原始长度的1/4作为初始容量预估，减少重新分配
	ret := make([]T, 0, len(ss)/4+1)
	for _, s := range ss {
		if f(s) {
			ret = append(ret, s)
		}
	}

	return ret
}

// FilterNot 对切片进行反向过滤，保留不满足指定条件的元素
//
// 参数:
//   - ss: 输入切片，支持任意类型
//   - f: 谓词函数，接收一个元素并返回布尔值
//
// 返回值:
//   - []T: 包含所有不满足谓词函数条件的新切片
//
// 特点:
//   - 支持任意类型的切片
//   - 是 Filter 函数的补集，用于语义上更清晰的反向过滤操作
//   - 不修改原切片，返回新切片
//   - 空切片输入返回空切片
//
// 示例:
//
//	// 过滤偶数，保留奇数
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	result := FilterNot(numbers, func(n int) bool {
//	    return n % 2 == 0
//	})
//	// result = [1, 3, 5]
//
//	// 过滤空字符串，保留非空字符串
//	strings := []string{"hello", "", "world", ""}
//	result := FilterNot(strings, func(s string) bool {
//	    return s == ""
//	})
//	// result = ["hello", "world"]
//
//	// 过滤负数，保留非负数
//	nums := []int{-1, 0, 1, -2, 2}
//	result := FilterNot(nums, func(n int) bool {
//	    return n < 0
//	})
//	// result = [0, 1, 2]
func FilterNot[T any](ss []T, f func(T) bool) []T {
	// 使用 make 初始化，确保返回空切片而非 nil
	us := make([]T, 0)
	for _, s := range ss {
		if !f(s) {
			us = append(us, s)
		}
	}
	return us
}

// Contains 检查切片中是否包含指定元素
func Contains[T constraints.Ordered](ss []T, s T) bool {
	return ContainsUsing(ss, func(v T) bool {
		return s == v
	})
}

// ContainsUsing 使用自定义函数检查切片中是否包含满足条件的元素
//
// 泛型约束：T 可以是任意类型
// 参数：
//   - ss: 要搜索的切片
//   - f: 判断函数，如果函数返回 true，则表示找到匹配的元素
//
// 返回值：
//   - bool: 如果找到满足条件的元素返回 true，否则返回 false
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found := ContainsUsing(numbers, func(v int) bool {
//	    return v > 3
//	}) // 返回 true
func ContainsUsing[T any](ss []T, f func(v T) bool) bool {
	// 遍历切片中的每个元素
	for _, v := range ss {
		// 使用自定义函数判断当前元素是否满足条件
		if f(v) {
			return true
		}
	}

	// 遍历完所有元素都没有找到满足条件的，返回 false
	return false
}
