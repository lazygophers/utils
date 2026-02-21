package candy

// RemoveIndex 移除指定索引的元素
// 该函数从切片中移除指定索引位置的元素，并返回新的切片
// 如果索引无效（超出范围或为负数），则返回空切片
func RemoveIndex[T any](ss []T, index int) []T {
	// 边界检查：如果切片为空或索引无效，返回空切片
	if len(ss) == 0 || index < 0 || index >= len(ss) {
		return make([]T, 0)
	}

	// 处理移除第一个元素的特殊情况
	if index == 0 {
		return ss[1:]
	}

	// 处理移除最后一个元素的特殊情况
	if index == len(ss)-1 {
		return ss[:len(ss)-1]
	}

	// 一般情况：使用 append 将索引前后的元素拼接起来
	return append(ss[:index], ss[index+1:]...)
}
