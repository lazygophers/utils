// Package candy 包含常用的语法糖函数，提供简洁的编程接口
package candy

// First 返回切片中的第一个元素
// 如果切片为空，返回类型的零值
//
// 泛型参数 T 可以是任意类型
//
// 示例：
//
//	nums := []int{1, 2, 3}
//	first := First(nums) // 返回 1
//
//	empty := []string{}
//	first := First(empty) // 返回 "" (string 的零值)
func First[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[0]
}
