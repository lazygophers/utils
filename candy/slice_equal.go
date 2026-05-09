package candy

// SliceEqual 判断两个切片是否相等，不考虑元素顺序
// 使用 map 来统计元素出现次数，确保每个元素在两个切片中出现次数相同
// 处理了 nil 切片的特殊情况：nil 和空切片视为相等
//
// 性能优化：
//   - 快速路径：长度检查和 nil 检查
//   - 小切片使用双重循环（无内存分配）
//   - 大切片使用 map 计数（O(n) 时间复杂度）
//   - 使用 any 类型转换避免 comparable 约束限制
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

	// 小切片优化：使用双重循环避免 map 开销
	// 对于小切片，双重循环比 map 更快（无内存分配）
	const smallSliceThreshold = 32
	if len(a) < smallSliceThreshold {
		// 标记已匹配的元素
		matched := make([]bool, len(b))
		for _, va := range a {
			found := false
			for j := 0; j < len(b); j++ {
				if !matched[j] {
					// 使用 any 比较
					vaAny := any(va)
					vbAny := any(b[j])
					if vaAny == vbAny {
						matched[j] = true
						found = true
						break
					}
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	// 大切片：使用 map 计数
	am := make(map[any]int, len(a))
	for i := 0; i < len(a); i++ {
		am[a[i]]++
	}

	for i := 0; i < len(b); i++ {
		vAny := any(b[i])
		if count, ok := am[vAny]; !ok || count == 0 {
			return false
		}
		am[vAny]--
	}

	return true
}
