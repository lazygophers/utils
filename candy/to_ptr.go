package candy

// ToPtr 将值转换为指针
// 接受任意类型的值并返回其指针
func ToPtr[T any](v T) *T {
	return &v
}
