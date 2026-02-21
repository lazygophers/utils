package candy

// Clone 通用克隆函数，根据类型选择最佳策略
func Clone[T any](src T) T {
	var dst T
	DeepCopy(src, &dst)
	return dst
}

// CloneSlice 切片克隆的便捷函数
func CloneSlice[T any](src []T) []T {
	return TypedSliceCopy(src)
}

// CloneMap map 克隆的便捷函数
func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	return TypedMapCopy(src)
}
