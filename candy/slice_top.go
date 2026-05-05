package candy

// Top 返回切片中的前 n 个元素
// 如果 n 大于切片长度，则返回整个切片的副本
// 使用 copy 确保返回的是新切片，避免修改原切片
func Top[T any](ss []T, n int) (ret []T) {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}

	ret = make([]T, n)
	copy(ret, ss[:n])
	return ret
}
