package candy

// ToInt 将任何类型的值尽力转换为 int。
// 此函数现在使用泛型实现，提供更好的性能和类型安全。
//
// 支持的输入类型包括：
//   - bool: true 转换为 1, false 转换为 0。
//   - 所有整数类型 (int, int8, ..., uint, uint8, ...): 直接进行类型转换。
//   - 所有浮点数类型 (float32, float64): 直接进行类型转换。
//   - string, []byte: 尝试解析为整数，若解析失败则返回 0。
//
// 对于无法转换的类型(如 struct, map 等)或 nil，将返回 0。
func ToInt(val interface{}) int {
	return ConvertWithDefault[interface{}, int](val, 0)
}
