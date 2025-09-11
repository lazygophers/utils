package candy

// PluckUint32 从结构体切片中提取指定字段的 uint32 值
func PluckUint32(list interface{}, fileName string) []uint32 {
	return pluck(list, fileName, []uint32{}).([]uint32)
}