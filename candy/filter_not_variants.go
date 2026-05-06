package candy

// 方案1: 预分配一半容量
// 适用于保留约一半元素的场景
func filterNotV1_PreHalf[T any](ss []T, f func(T) bool) []T {
	capacity := len(ss) / 2
	if capacity == 0 {
		capacity = 1
	}
	ret := make([]T, 0, capacity)
	for _, s := range ss {
		if !f(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

// 方案2: 预分配全容量
// 适用于保留大部分元素的场景
func filterNotV2_PreFull[T any](ss []T, f func(T) bool) []T {
	ret := make([]T, 0, len(ss))
	for _, s := range ss {
		if !f(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

// 方案3: 索引循环替代 range
// 避免range的值拷贝开销
func filterNotV3_IndexLoop[T any](ss []T, f func(T) bool) []T {
	ret := make([]T, 0)
	n := len(ss)
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			ret = append(ret, ss[i])
		}
	}
	return ret
}

// 方案4: 预分配一半容量 + 索引循环
// 结合方案1和方案3的优点
func filterNotV4_PreHalfIndex[T any](ss []T, f func(T) bool) []T {
	capacity := len(ss) / 2
	if capacity == 0 {
		capacity = 1
	}
	ret := make([]T, 0, capacity)
	n := len(ss)
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			ret = append(ret, ss[i])
		}
	}
	return ret
}

// 方案5: 预分配全容量 + 索引循环
// 结合方案2和方案3的优点
func filterNotV5_PreFullIndex[T any](ss []T, f func(T) bool) []T {
	ret := make([]T, 0, len(ss))
	n := len(ss)
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			ret = append(ret, ss[i])
		}
	}
	return ret
}

// 方案6: 两遍扫描优化
// 第一遍计数，第二遍精确分配
func filterNotV6_TwoPass[T any](ss []T, f func(T) bool) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}

	// 第一遍：计数
	count := 0
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			count++
		}
	}

	if count == 0 {
		return []T{}
	}

	// 第二遍：填充
	ret := make([]T, count)
	j := 0
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			ret[j] = ss[i]
			j++
		}
	}
	return ret
}

// 方案7: 复制后切片
// 适用于需要保留原始切片的场景
func filterNotV7_CopySlice[T any](ss []T, f func(T) bool) []T {
	if len(ss) == 0 {
		return []T{}
	}

	// 复制整个切片
	copied := make([]T, len(ss))
	copy(copied, ss)

	// 原地过滤
	n := 0
	for _, s := range copied {
		if !f(s) {
			copied[n] = s
			n++
		}
	}
	return copied[:n]
}

// 方案8: 原地修改（非零拷贝）
// 直接在原切片上操作，会修改原切片
// 注意：此实现会修改原切片，仅用于性能对比
func filterNotV8_InPlace[T any](ss []T, f func(T) bool) []T {
	// 先创建副本避免修改原切片
	copied := make([]T, len(ss))
	copy(copied, ss)

	n := 0
	for _, s := range copied {
		if !f(s) {
			copied[n] = s
			n++
		}
	}
	return copied[:n]
}

// 方案9: 动态容量调整
// 根据已保留的比例动态调整容量
func filterNotV9_Dynamic[T any](ss []T, f func(T) bool) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}

	ret := make([]T, 0, 1) // 初始容量为1
	kept := 0
	total := 0

	for i := 0; i < n; i++ {
		total++
		if !f(ss[i]) {
			ret = append(ret, ss[i])
			kept++

			// 动态调整容量
			if kept == cap(ret) && total < n {
				// 根据当前保留比例预测
				ratio := float64(kept) / float64(total)
				predicted := int(float64(n-total) * ratio)
				newCap := kept + predicted
				if newCap > cap(ret) {
					newRet := make([]T, kept, newCap)
					copy(newRet, ret)
					ret = newRet
				}
			}
		}
	}
	return ret
}

// 方案10: 空切片快速路径 + 预分配一半
// 对空切片快速返回，对其他情况使用半容量预分配
func filterNotV10_FastPath[T any](ss []T, f func(T) bool) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}

	capacity := n / 2
	if capacity == 0 {
		capacity = 1
	}
	ret := make([]T, 0, capacity)

	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			ret = append(ret, ss[i])
		}
	}
	return ret
}

// 方案11: 使用 Filter + 取反逻辑
// 通过组合Filter函数实现
func filterNotV11_Composed[T any](ss []T, f func(T) bool) []T {
	return Filter(ss, func(t T) bool {
		return !f(t)
	})
}
