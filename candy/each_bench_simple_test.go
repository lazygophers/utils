package candy

import (
	"math/rand"
	"testing"
)

// ==================== 测试数据生成 ====================

// 生成整数切片（避免与 collection_ops_bench_test.go 冲突）
func genIntsEach(n int) []int {
	r := rand.New(rand.NewSource(42))
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = r.Int()
	}
	return s
}

// ==================== 回调函数 ====================

// 简单回调（累加）
var sumResult int

func simpleSum(v int) {
	sumResult += v
}

// ==================== Each 函数的 10 种优化方案 ====================

// 方案1: 当前实现（baseline）- 使用 range
func Each1[T any](values []T, fn func(value T)) {
	for _, value := range values {
		fn(value)
	}
}

// 方案2: 使用索引循环
func Each2[T any](values []T, fn func(value T)) {
	n := len(values)
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案3: 预先检查空切片
func Each3[T any](values []T, fn func(value T)) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		fn(value)
	}
}

// 方案4: 使用索引循环 + 空切片检查
func Each4[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案5: 循环展开（4路）
func Each5[T any](values []T, fn func(value T)) {
	n := len(values)
	i := 0

	// 处理不能被4整除的部分
	for i < n && (n-i)%4 != 0 {
		fn(values[i])
		i++
	}

	// 4路展开
	for i+3 < n {
		fn(values[i])
		fn(values[i+1])
		fn(values[i+2])
		fn(values[i+3])
		i += 4
	}

	// 处理剩余元素
	for i < n {
		fn(values[i])
		i++
	}
}

// 方案6: 循环展开（8路）
func Each6[T any](values []T, fn func(value T)) {
	n := len(values)
	i := 0

	// 处理不能被8整除的部分
	for i < n && (n-i)%8 != 0 {
		fn(values[i])
		i++
	}

	// 8路展开
	for i+7 < n {
		fn(values[i])
		fn(values[i+1])
		fn(values[i+2])
		fn(values[i+3])
		fn(values[i+4])
		fn(values[i+5])
		fn(values[i+6])
		fn(values[i+7])
		i += 8
	}

	// 处理剩余元素
	for i < n {
		fn(values[i])
		i++
	}
}

// 方案7: 小切片特殊处理
func Each7[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}

	// 小切片直接展开
	if n <= 4 {
		if n >= 1 {
			fn(values[0])
		}
		if n >= 2 {
			fn(values[1])
		}
		if n >= 3 {
			fn(values[2])
		}
		if n >= 4 {
			fn(values[3])
		}
		return
	}

	// 大切片使用循环展开
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案8: 使用临时变量减少边界检查
func Each8[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}

	length := n
	for i := 0; i < length; i++ {
		fn(values[i])
	}
}

// 方案9: 混合策略 - 小切片用range，大切片用索引
func Each9[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}

	// 小切片使用range（代码简洁）
	if n < 32 {
		for _, value := range values {
			fn(value)
		}
		return
	}

	// 大切片使用索引循环（性能更好）
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案10: 使用指针访问（适用于特定类型）
func Each10(values []int, fn func(*int)) {
	n := len(values)
	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		fn(&values[i])
	}
}

func ptrSum(v *int) {
	sumResult += *v
}

// ==================== EachReverse 函数的 10 种优化方案 ====================

// 方案1: 当前实现（baseline）
func EachReverse1[T any](ss []T, f func(T)) {
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案2: 预先计算长度
func EachReverse2[T any](ss []T, f func(T)) {
	n := len(ss)
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案3: 空切片检查
func EachReverse3[T any](ss []T, f func(T)) {
	if len(ss) == 0 {
		return
	}
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案4: 预先计算长度 + 空切片检查
func EachReverse4[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案5: 循环展开（4路）
func EachReverse5[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	i := n - 1

	// 处理不能被4整除的部分
	for i >= 0 && ((n-1-i)%4) != 0 {
		f(ss[i])
		i--
	}

	// 4路展开
	for i-3 >= 0 {
		f(ss[i])
		f(ss[i-1])
		f(ss[i-2])
		f(ss[i-3])
		i -= 4
	}

	// 处理剩余元素
	for i >= 0 {
		f(ss[i])
		i--
	}
}

// 方案6: 循环展开（8路）
func EachReverse6[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	i := n - 1

	// 处理不能被8整除的部分
	for i >= 0 && ((n-1-i)%8) != 0 {
		f(ss[i])
		i--
	}

	// 8路展开
	for i-7 >= 0 {
		f(ss[i])
		f(ss[i-1])
		f(ss[i-2])
		f(ss[i-3])
		f(ss[i-4])
		f(ss[i-5])
		f(ss[i-6])
		f(ss[i-7])
		i -= 8
	}

	// 处理剩余元素
	for i >= 0 {
		f(ss[i])
		i--
	}
}

// 方案7: 小切片特殊处理
func EachReverse7[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	// 小切片直接展开
	if n <= 4 {
		if n >= 1 {
			f(ss[n-1])
		}
		if n >= 2 {
			f(ss[n-2])
		}
		if n >= 3 {
			f(ss[n-3])
		}
		if n >= 4 {
			f(ss[n-4])
		}
		return
	}

	// 大切片使用常规循环
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案8: 使用临时变量减少边界检查
func EachReverse8[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	length := n
	for i := length - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案9: 混合策略 - 小和大切片不同处理
func EachReverse9[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	// 小切片直接展开
	if n <= 4 {
		for i := n - 1; i >= 0; i-- {
			f(ss[i])
		}
		return
	}

	// 大切片使用预先计算长度
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案10: 使用 while 循环风格（goto）
func EachReverse10[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	i := n - 1
loop:
	f(ss[i])
	i--
	if i >= 0 {
		goto loop
	}
}

// ==================== 基准测试 ====================

// Each 测试 - 小切片
func BenchmarkEach1_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each1(data, simpleSum)
	}
}

func BenchmarkEach2_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each2(data, simpleSum)
	}
}

func BenchmarkEach3_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each3(data, simpleSum)
	}
}

func BenchmarkEach4_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each4(data, simpleSum)
	}
}

func BenchmarkEach5_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each5(data, simpleSum)
	}
}

func BenchmarkEach6_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each6(data, simpleSum)
	}
}

func BenchmarkEach7_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each7(data, simpleSum)
	}
}

func BenchmarkEach8_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each8(data, simpleSum)
	}
}

func BenchmarkEach9_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each9(data, simpleSum)
	}
}

func BenchmarkEach10_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each10(data, ptrSum)
	}
}

// Each 测试 - 大切片
func BenchmarkEach1_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each1(data, simpleSum)
	}
}

func BenchmarkEach2_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each2(data, simpleSum)
	}
}

func BenchmarkEach4_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each4(data, simpleSum)
	}
}

func BenchmarkEach5_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each5(data, simpleSum)
	}
}

func BenchmarkEach6_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each6(data, simpleSum)
	}
}

func BenchmarkEach9_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each9(data, simpleSum)
	}
}

// EachReverse 测试 - 小切片
func BenchmarkEachReverse1_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse1(data, simpleSum)
	}
}

func BenchmarkEachReverse2_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse2(data, simpleSum)
	}
}

func BenchmarkEachReverse4_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse4(data, simpleSum)
	}
}

func BenchmarkEachReverse9_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse9(data, simpleSum)
	}
}

// EachReverse 测试 - 大切片
func BenchmarkEachReverse1_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse1(data, simpleSum)
	}
}

func BenchmarkEachReverse2_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse2(data, simpleSum)
	}
}

func BenchmarkEachReverse4_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse4(data, simpleSum)
	}
}

func BenchmarkEachReverse9_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse9(data, simpleSum)
	}
}
