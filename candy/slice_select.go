package candy

// First 返回切片中的第一个元素
// 如果切片为空，返回类型的零值
//
// 泛型参数 T 可以是任意类型
//
// 示例：
//
//	nums := []int{1, 2, 3}
//	first := First(nums) // 返回 1
//
//	empty := []string{}
//	first := First(empty) // 返回 "" (string 的零值)
func First[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[0]
}

// FirstOr 返回切片的第一个元素，如果切片为空则返回指定的默认值
//
// 该函数使用泛型支持任意类型的切片，提供了安全的空切片处理机制。
// 在访问切片第一个元素之前，会先检查切片长度，避免 panic。
//
// 参数:
//   - ss: 任意类型的切片
//   - or: 当切片为空时返回的默认值
//
// 返回:
//   - 切片的第一个元素，如果切片为空则返回默认值
//
// 示例:
//
//	numbers := []int{1, 2, 3}
//	first := FirstOr(numbers, 0)     // 返回 1
//
//	empty := []int{}
//	defaultVal := FirstOr(empty, 0) // 返回 0
func FirstOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[0]
}

// Last 返回切片中的最后一个元素
// 如果切片为空，返回类型的零值
// 该函数使用泛型实现，支持任意类型的切片
//
// 参数:
//   - ss: 任意类型的切片
//
// 返回:
//   - T: 切片中的最后一个元素，如果切片为空则返回类型零值
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last := Last(numbers) // 返回 5
//
//	empty := []string{}
//	result := Last(empty) // 返回 ""
func Last[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[len(ss)-1]
}

// LastOr 返回切片中的最后一个元素
// 如果切片为空，返回指定的默认值
// 该函数使用泛型实现，支持任意类型的切片
//
// 参数:
//   - ss: 任意类型的切片
//   - or: 当切片为空时返回的默认值
//
// 返回:
//   - T: 切片中的最后一个元素，如果切片为空则返回指定的默认值
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last := LastOr(numbers, 0) // 返回 5
//
//	empty := []string{}
//	result := LastOr(empty, "default") // 返回 "default"
func LastOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[len(ss)-1]
}

// Bottom 返回切片的最后 n 个元素
func Bottom[T any](ss []T, n int) (ret []T) {
	if n <= 0 {
		return []T{}
	}
	if n > len(ss) {
		n = len(ss)
	}

	ret = make([]T, n)
	copy(ret, ss[len(ss)-n:])
	return ret
}
