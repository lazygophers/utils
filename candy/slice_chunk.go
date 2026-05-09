package candy

// Chunk 将切片分割成指定大小的子切片
// 性能优化策略：
// - size >= len(ss): 直接返回包含整个切片的单元素结果（无分配）
// - size == 1: 使用索引循环预分配（比 range 更快）
// - 其他情况: 使用索引循环，避免 range 的开销
func Chunk[T any](ss []T, size int) (ret [][]T) {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}

	// 特殊情况：chunk 大小大于等于切片长度
	if size >= len(ss) {
		return [][]T{ss}
	}

	// 特殊情况：chunk 大小为 1，使用优化的单个元素处理
	if size == 1 {
		ret = make([][]T, len(ss))
		for i := 0; i < len(ss); i++ {
			ret[i] = []T{ss[i]}
		}
		return
	}

	// 通用情况：使用索引循环，比 range 更快
	ret = make([][]T, 0, (len(ss)+size-1)/size)
	n := len(ss)
	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		// 使用完全切片（三参数）防止容量扩展
		ret = append(ret, ss[i:end:end])
	}

	return
}
