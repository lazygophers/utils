package candy

// PluckInt64 从结构体切片中提取指定字段的 int64 值
func PluckInt64(list interface{}, fieldName string) []int64 {
	return pluck(list, fieldName, []int64{}).([]int64)
}
