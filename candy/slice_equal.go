package candy

// SliceEqual 判断两个切片是否相等，不考虑元素顺序
// 使用 map 来统计元素出现次数，确保每个元素在两个切片中出现次数相同
// 处理了 nil 切片的特殊情况：nil 和空切片视为相等
func SliceEqual[T any](a, b []T) bool {
	// 处理 nil 切片的情况：nil 和空切片视为相等
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	// 使用 map 来跟踪每个元素的出现次数
	am := make(map[any]int, len(a))
	for _, v := range a {
		am[v]++
	}

	for _, v := range b {
		if count, ok := am[v]; !ok || count == 0 {
			return false
		}
		am[v]--
	}

	// 检查所有元素的计数是否都为0
	for _, count := range am {
		if count != 0 {
			return false
		}
	}

	return true
}
