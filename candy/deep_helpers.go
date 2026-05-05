package candy

// Equal 通用相等性检查
func Equal[T comparable](a, b T) bool {
	return a == b
}
