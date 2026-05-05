package candy

// RemoveIndex 移除指定索引的元素
// 该函数从切片中移除指定索引位置的元素，并返回新的切片
// 如果索引无效（超出范围或为负数），则返回空切片
//
// 性能优化：
//   - 边界情况（首/尾）使用零拷贝切片操作
//   - 小切片（<64元素）使用 append（避免预分配开销）
//   - 大切片使用预分配 + copy（减少内存重分配）
func RemoveIndex[T any](ss []T, index int) []T {
	n := len(ss)
	// 边界检查：如果切片为空或索引无效，返回空切片
	if n == 0 || index < 0 || index >= n {
		return make([]T, 0)
	}

	// 处理移除第一个元素的特殊情况（零拷贝）
	if index == 0 {
		return ss[1:]
	}

	// 处理移除最后一个元素的特殊情况（零拷贝）
	if index == n-1 {
		return ss[:n-1]
	}

	// 混合策略：小切片使用 append，大切片使用预分配 copy
	// 阈值64是性能测试得出的最佳平衡点
	if n < 64 {
		return append(ss[:index], ss[index+1:]...)
	}

	// 大切片：预分配并使用 copy（性能最优）
	result := make([]T, n-1)
	copy(result, ss[:index])
	copy(result[index:], ss[index+1:])
	return result
}
