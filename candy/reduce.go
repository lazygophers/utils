package candy

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
