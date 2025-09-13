package candy

// Any 检查切片中是否存在至少一个元素满足指定条件
func Any[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}

	return false
}
