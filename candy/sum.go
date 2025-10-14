// Package candy 包含语法糖工具函数
package candy

// Sum 计算数值切片中所有元素的总和
// 支持整数和浮点数类型，使用泛型实现类型安全
//
// 参数：
//   - ss: 数值切片，支持整数和浮点数类型
//
// 返回值：
//   - T: 切片中所有元素的总和
//
// 示例：
//
//	sum := Sum([]int{1, 2, 3})  // 返回 6
//	sum := Sum([]float64{1.5, 2.5})  // 返回 4.0
func Sum[T interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}](ss ...T) (ret T) {
	for _, s := range ss {
		ret += s
	}

	return
}
