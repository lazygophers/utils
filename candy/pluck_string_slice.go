package candy

// PluckStringSlice 从结构体切片中提取指定字段的 []string 值
func PluckStringSlice(list interface{}, fieldName string) [][]string {
	return pluck(list, fieldName, [][]string{}).([][]string)
}