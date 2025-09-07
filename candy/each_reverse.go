package candy

// EachReverse 反向遍历切片并对每个元素执行指定函数
// 从切片的最后一个元素开始，向前遍历到第一个元素
// 对于每个元素，都会调用传入的函数 f 进行处理
//
// 参数:
//   - ss: 要遍历的切片
//   - f: 对每个元素执行的函数，接收一个类型为 T 的参数
//
// 泛型参数:
//   - T: 切片中元素的类型，可以是任意类型
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	EachReverse(numbers, func(n int) {
//	    fmt.Println(n) // 输出: 5, 4, 3, 2, 1
//	})
func EachReverse[T any](ss []T, f func(T)) {
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}
