package anyx

// mapGetWithSeparatorOptimized 已弃用，请直接使用 mapGetWithSeparator
// 保留此函数仅为向后兼容
// Deprecated: 使用 mapGetWithSeparator 代替
func mapGetWithSeparatorOptimized(m map[string]any, key string, sep string) (any, error) {
	return mapGetWithSeparator(m, key, sep)
}
