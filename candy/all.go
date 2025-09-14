package candy

// All 检查切片中的所有元素是否都满足指定条件
func All[T any](ss []T, f func(T) bool) bool {
	if len(ss) == 0 {
		return true
	}

	for _, s := range ss {
		if !f(s) {
			return false
		}
	}

	return true
}
