// Package candy 提供通用的语法糖工具函数
package candy

// Bottom 返回切片的最后 n 个元素
// 如果 n 大于切片长度，则返回整个切片
// 使用泛型支持任意类型的切片
func Bottom[T any](ss []T, n int) (ret []T) {
	if n <= 0 {
		return []T{}
	}
	if n > len(ss) {
		n = len(ss)
	}

	ret = make([]T, n)
	copy(ret, ss[len(ss)-n:])
	return ret
}
