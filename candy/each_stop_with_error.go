// Package candy 提供了实用的工具函数和语法糖
package candy

// EachStopWithError 遍历切片并对每个元素执行指定函数，如果函数返回错误则立即停止遍历并返回该错误
//
// 参数：
//   - ss: 要遍历的切片
//   - f: 对每个元素执行的函数，如果返回错误则停止遍历
//
// 返回值：
//   - error: 如果遍历过程中出现错误则返回该错误，否则返回 nil
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	err := EachStopWithError(numbers, func(n int) error {
//	    if n == 3 {
//	        return fmt.Errorf("数字 3 不被允许")
//	    }
//	    fmt.Println(n)
//	    return nil
//	})
func EachStopWithError[T any](ss []T, f func(T) (err error)) (err error) {
	for _, s := range ss {
		err = f(s)
		if err != nil {
			return err
		}
	}
	return nil
}
