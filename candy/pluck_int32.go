package candy

// PluckInt32 从结构体切片中提取指定字段的 int32 值
func PluckInt32(list interface{}, fieldName string) []int32 {
	return pluck(list, fieldName, []int32{}).([]int32)
}
