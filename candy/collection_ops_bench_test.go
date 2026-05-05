package candy

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 生成测试数据
func genInts(n int) []int {
	r := rand.New(rand.NewSource(42))
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = r.Int()
	}
	return s
}

func genStrings(n int) []string {
	r := rand.New(rand.NewSource(42))
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = strconv.Itoa(r.Intn(1000))
	}
	return s
}

// 简单累加函数
func simpleAdd(a, b int) int {
	return a + b
}

// 复杂函数（模拟真实场景）
func complexOp(a, b int) int {
	time.Sleep(1 * time.Microsecond)
	return a + b + (a * b) / 2
}

// 字符串连接
func concatStr(a, b string) string {
	return a + b
}

// ==================== 基准测试：不同大小的切片 ====================

// Baseline: 当前实现
func BenchmarkReduce_Current_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

func BenchmarkReduce_Current_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

func BenchmarkReduce_Current_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

// ==================== 10种优化方案 ====================

// 方案1: 避免切片重新分配，直接索引访问
func Reduce1[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}
	result := ss[0]
	for i := 1; i < len(ss); i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce1_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, simpleAdd)
	}
}

func BenchmarkReduce1_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, simpleAdd)
	}
}

// 方案2: 预检查并针对小切片优化
func Reduce2[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	if n == 2 {
		return f(ss[0], ss[1])
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce2_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce2(data, simpleAdd)
	}
}

func BenchmarkReduce2_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce2(data, simpleAdd)
	}
}

// 方案3: 循环展开（每4个元素）
func Reduce3[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}

	result := ss[0]
	i := 1
	// 处理剩余元素，确保有4的倍数个元素可展开
	for i < n && ((n-i)%4) != 0 {
		result = f(result, ss[i])
		i++
	}

	// 4路展开
	for i+3 < n {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
		i += 4
	}

	// 处理剩余元素
	for i < n {
		result = f(result, ss[i])
		i++
	}

	return result
}

func BenchmarkReduce3_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce3(data, simpleAdd)
	}
}

func BenchmarkReduce3_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce3(data, simpleAdd)
	}
}

// 方案4: 使用指针传递（仅适用于指针类型）
func Reduce4[T any](ss []T, f func(*T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(&result, ss[i])
	}
	return result
}

func ptrAdd(a *int, b int) int {
	return *a + b
}

func BenchmarkReduce4_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce4(data, ptrAdd)
	}
}

func BenchmarkReduce4_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce4(data, ptrAdd)
	}
}

// 方案5: int类型特化版本
func ReduceInt(ss []int, f func(int, int) int) int {
	n := len(ss)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return ss[0]
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduceInt_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceInt(data, simpleAdd)
	}
}

func BenchmarkReduceInt_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceInt(data, simpleAdd)
	}
}

// 方案6: string类型特化版本
func ReduceString(ss []string, f func(string, string) string) string {
	n := len(ss)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return ss[0]
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduceString_Small(b *testing.B) {
	data := genStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceString(data, concatStr)
	}
}

func BenchmarkReduceString_Large(b *testing.B) {
	data := genStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceString(data, concatStr)
	}
}

// 方案7: 混合策略 - 小切片用简单循环，大切片用展开
func Reduce7[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n <= 4 {
		// 小切片：直接展开
		if n == 1 {
			return ss[0]
		}
		if n == 2 {
			return f(ss[0], ss[1])
		}
		if n == 3 {
			return f(f(ss[0], ss[1]), ss[2])
		}
		return f(f(f(ss[0], ss[1]), ss[2]), ss[3])
	}

	// 大切片：循环展开
	result := ss[0]
	i := 1
	for i+7 < n {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
		result = f(result, ss[i+4])
		result = f(result, ss[i+5])
		result = f(result, ss[i+6])
		result = f(result, ss[i+7])
		i += 8
	}

	for i < n {
		result = f(result, ss[i])
		i++
	}

	return result
}

func BenchmarkReduce7_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, simpleAdd)
	}
}

func BenchmarkReduce7_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, simpleAdd)
	}
}

// 方案8: 针对复杂函数的优化（减少函数调用开销）
func Reduce8[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	// 对小切片直接返回
	if n <= 4 {
		switch n {
		case 1:
			return ss[0]
		case 2:
			return f(ss[0], ss[1])
		case 3:
			return f(f(ss[0], ss[1]), ss[2])
		case 4:
			return f(f(f(ss[0], ss[1]), ss[2]), ss[3])
		}
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce8_Complex_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce8(data, simpleAdd)
	}
}

func BenchmarkReduce8_Complex_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce8(data, simpleAdd)
	}
}

// 方案9: 使用临时变量减少边界检查
func Reduce9[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	result := ss[0]
	length := n
	for i := 1; i < length; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce9_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce9(data, simpleAdd)
	}
}

func BenchmarkReduce9_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce9(data, simpleAdd)
	}
}

// 方案10: 无边界检查的循环（使用goto）
func Reduce10[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	result := ss[0]
	i := 1
	if i >= n {
		return result
	}

loop:
	result = f(result, ss[i])
	i++
	if i < n {
		goto loop
	}

	return result
}

func BenchmarkReduce10_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce10(data, simpleAdd)
	}
}

func BenchmarkReduce10_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce10(data, simpleAdd)
	}
}

// ==================== 测试复杂函数场景 ====================

func BenchmarkReduce_Current_Complex(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

func BenchmarkReduce1_Complex(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, simpleAdd)
	}
}

func BenchmarkReduce7_Complex(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, simpleAdd)
	}
}

// ==================== 字符串测试 ====================

func BenchmarkReduce_Current_String(b *testing.B) {
	data := genStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, concatStr)
	}
}

func BenchmarkReduce1_String(b *testing.B) {
	data := genStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, concatStr)
	}
}

func BenchmarkReduce7_String(b *testing.B) {
	data := genStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, concatStr)
	}
}
