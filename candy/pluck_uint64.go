package candy

// PluckUint64 从结构体切片中提取指定字段的 uint64 值
func PluckUint64(list interface{}, fieldName string) []uint64 {
	return pluck(list, fieldName, []uint64{}).([]uint64)
}
