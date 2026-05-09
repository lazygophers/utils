package candy

import (
	"golang.org/x/exp/constraints"
	"math/rand/v2"
	"sort"
	"strconv"
	"testing"
	"time"
)

// 生成测试数据
func genInts(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = int(rand.Uint64())
	}
	return s
}

func genStrings(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = strconv.Itoa(int(rand.Uint64() % 1000))
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
	return a + b + (a*b)/2
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

// ==================== Map 函数性能优化测试 ====================

// 基准方案：当前实现
func BenchmarkMap_Current_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap_Current_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap_Current_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

// 方案1：移除空检查分支
func BenchmarkMap1_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap1_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap1_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

// 方案2：循环展开2
func BenchmarkMap2_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 2 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap2_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 2 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap2_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 2 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
		}
		_ = ret
	}
}

// 方案3：循环展开4
func BenchmarkMap3_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 4 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
			if j+2 < n {
				ret[j+2] = data[j+2] * 2
			}
			if j+3 < n {
				ret[j+3] = data[j+3] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap3_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 4 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
			if j+2 < n {
				ret[j+2] = data[j+2] * 2
			}
			if j+3 < n {
				ret[j+3] = data[j+3] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap3_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 4 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
			if j+2 < n {
				ret[j+2] = data[j+2] * 2
			}
			if j+3 < n {
				ret[j+3] = data[j+3] * 2
			}
		}
		_ = ret
	}
}

// 方案4：range循环对比
func BenchmarkMap4_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, len(data))
		for idx, v := range data {
			ret[idx] = v * 2
		}
		_ = ret
	}
}

func BenchmarkMap4_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, len(data))
		for idx, v := range data {
			ret[idx] = v * 2
		}
		_ = ret
	}
}

func BenchmarkMap4_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, len(data))
		for idx, v := range data {
			ret[idx] = v * 2
		}
		_ = ret
	}
}

// 方案5：嵌套展开4
func BenchmarkMap5_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		j := 0
		for j < n {
			if j+4 <= n {
				ret[j] = data[j] * 2
				ret[j+1] = data[j+1] * 2
				ret[j+2] = data[j+2] * 2
				ret[j+3] = data[j+3] * 2
				j += 4
			} else {
				ret[j] = data[j] * 2
				j++
			}
		}
		_ = ret
	}
}

func BenchmarkMap5_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		j := 0
		for j < n {
			if j+4 <= n {
				ret[j] = data[j] * 2
				ret[j+1] = data[j+1] * 2
				ret[j+2] = data[j+2] * 2
				ret[j+3] = data[j+3] * 2
				j += 4
			} else {
				ret[j] = data[j] * 2
				j++
			}
		}
		_ = ret
	}
}

func BenchmarkMap5_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		j := 0
		for j < n {
			if j+4 <= n {
				ret[j] = data[j] * 2
				ret[j+1] = data[j+1] * 2
				ret[j+2] = data[j+2] * 2
				ret[j+3] = data[j+3] * 2
				j += 4
			} else {
				ret[j] = data[j] * 2
				j++
			}
		}
		_ = ret
	}
}

// ==================== Reverse Benchmark Implementations ====================

// ReverseV1: 当前实现（小切片元素复制，大切片双指针）
func ReverseV1[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 32 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV2: 纯双指针交换（无分支）
func ReverseV2[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV3: 反向索引复制（无交换）
func ReverseV3[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = ss[n-1-i]
	}

	return result
}

// ReverseV4: 双指针 + 双向复制
func ReverseV4[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		result[i] = ss[j]
		result[j] = ss[i]
	}

	return result
}

// ReverseV5: 小切片阈值16
func ReverseV5[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 16 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV6: 小切片阈值64
func ReverseV6[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 64 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV7: 双指针 + 边界检查优化
func ReverseV7[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	left, right := 0, n-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}

// ReverseV8: 预分配 + 索引反转
func ReverseV8[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	for i := n/2 - 1; i >= 0; i-- {
		opp := n - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}

// ReverseV9: 小切片阈值8
func ReverseV9[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 8 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV10: 小切片阈值128
func ReverseV10[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 128 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ==================== Shuffle Benchmark Implementations ====================

// ShuffleV1: 当前实现（标准库 rand.Shuffle）
func ShuffleV1[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	rand.Shuffle(n, func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})

	return ss
}

// ShuffleV2: Fisher-Yates 手动实现
func ShuffleV2[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV3: Fisher-Yates 正向遍历
func ShuffleV3[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := 0; i < n-1; i++ {
		j := rand.IntN(n - i)
		ss[i], ss[i+j] = ss[i+j], ss[i]
	}

	return ss
}

// ShuffleV4: Fisher-Yates + 自交换检查
func ShuffleV4[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		if i != j {
			ss[i], ss[j] = ss[j], ss[i]
		}
	}

	return ss
}

// ShuffleV5: 局部变量优化
func ShuffleV5[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	s := ss
	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		s[i], s[j] = s[j], s[i]
	}

	return ss
}

// ShuffleV6: 小切片优化
func ShuffleV6[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	if n < 8 {
		for i := n - 1; i > 0; i-- {
			j := rand.IntN(i + 1)
			ss[i], ss[j] = ss[j], ss[i]
		}
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV7: 循环展开2个元素
func ShuffleV7[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV8: 分区洗牌
func ShuffleV8[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	mid := n / 2
	for i := mid - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}
	for i := n - 1; i > mid; i-- {
		j := rand.IntN(i - mid + 1)
		j += mid
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV9: 块处理
func ShuffleV9[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i -= 2 {
		if i > 0 {
			j := rand.IntN(i + 1)
			ss[i], ss[j] = ss[j], ss[i]
			i--
			if i > 0 {
				j = rand.IntN(i + 1)
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}

	return ss
}

// ShuffleV10: 简化版本
func ShuffleV10[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ==================== Sort Benchmark Implementations ====================

// SortV1: 当前实现（小切片插入排序，大切片sort.Slice）
func SortV1[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV2: 纯 sort.Slice
func SortV2[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV3: 插入排序阈值16
func SortV3[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 16 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV4: 插入排序阈值32
func SortV4[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 32 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV5: 预排序检查
func SortV5[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if sorted[i] < sorted[i-1] {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV6: 预排序 + 插入排序
func SortV6[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if sorted[i] < sorted[i-1] {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV7: 二分插入排序
func SortV7[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			left, right := 0, i
			for left < right {
				mid := (left + right) / 2
				if sorted[mid] < key {
					left = mid + 1
				} else {
					right = mid
				}
			}
			for j := i; j > left; j-- {
				sorted[j] = sorted[j-1]
			}
			sorted[left] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV8: 插入排序阈值8
func SortV8[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 8 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV9: 插入排序阈值64
func SortV9[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 64 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV10: 分层阈值（8/24/64）
func SortV10[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 8 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	if n <= 64 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// ==================== SortUsing Benchmark Implementations ====================

// SortUsingV1: 当前实现
func SortUsingV1[T any](slice []T, less func(T, T) bool) []T {
	if len(slice) < 2 {
		return slice
	}

	sorted := make([]T, len(slice))
	copy(sorted, slice)

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV2: 小切片插入排序
func SortUsingV2[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV3: 预排序检查
func SortUsingV3[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if less(sorted[i], sorted[i-1]) {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV4: 预排序 + 插入排序
func SortUsingV4[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if less(sorted[i], sorted[i-1]) {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV5: 二分插入排序
func SortUsingV5[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			left, right := 0, i
			for left < right {
				mid := (left + right) / 2
				if less(sorted[mid], key) {
					left = mid + 1
				} else {
					right = mid
				}
			}
			for j := i; j > left; j-- {
				sorted[j] = sorted[j-1]
			}
			sorted[left] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV6: 插入排序阈值16
func SortUsingV6[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 16 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV7: 插入排序阈值32
func SortUsingV7[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 32 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV8: 插入排序阈值8
func SortUsingV8[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 8 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV9: 减少闭包分配
func SortUsingV9[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	lessFunc := func(i, j int) bool {
		return less(sorted[i], sorted[j])
	}

	sort.Slice(sorted, lessFunc)

	return sorted
}

// SortUsingV10: 预检查 + 插入排序
func SortUsingV10[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n > 3 {
		if less(sorted[2], sorted[1]) || less(sorted[1], sorted[0]) {
			needsSort := false
			for i := 3; i < n; i++ {
				if less(sorted[i], sorted[i-1]) {
					needsSort = true
					break
				}
			}
			if !needsSort {
				return sorted
			}
		}
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// ==================== Reverse Benchmarks ====================

func BenchmarkReverse_Int_Small(b *testing.B) {
	data := genInts(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]int{}, data...))
		}
	})
}

func BenchmarkReverse_Int_Medium(b *testing.B) {
	data := genInts(64)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]int{}, data...))
		}
	})
}

func BenchmarkReverse_Int_Large(b *testing.B) {
	data := genInts(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]int{}, data...))
		}
	})
}

func BenchmarkReverse_String_Small(b *testing.B) {
	data := genStrings(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]string{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]string{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]string{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]string{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]string{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]string{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]string{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]string{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]string{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]string{}, data...))
		}
	})
}

func BenchmarkReverse_String_Large(b *testing.B) {
	data := genStrings(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]string{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]string{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]string{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]string{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]string{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]string{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]string{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]string{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]string{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]string{}, data...))
		}
	})
}

// ==================== Shuffle Benchmarks ====================

func BenchmarkShuffle_Int_Small(b *testing.B) {
	data := genInts(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV1(d)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV2(d)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV3(d)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV4(d)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV5(d)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV6(d)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV7(d)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV8(d)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV9(d)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV10(d)
		}
	})
}

func BenchmarkShuffle_Int_Medium(b *testing.B) {
	data := genInts(64)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV1(d)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV2(d)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV3(d)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV4(d)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV5(d)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV6(d)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV7(d)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV8(d)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV9(d)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV10(d)
		}
	})
}

func BenchmarkShuffle_Int_Large(b *testing.B) {
	data := genInts(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV1(d)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV2(d)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV3(d)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV4(d)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV5(d)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV6(d)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV7(d)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV8(d)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV9(d)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV10(d)
		}
	})
}

// ==================== Sort Benchmarks ====================

func BenchmarkSort_Int_Small(b *testing.B) {
	data := genInts(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

func BenchmarkSort_Int_Medium(b *testing.B) {
	data := genInts(64)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

func BenchmarkSort_Int_Large(b *testing.B) {
	data := genInts(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

func BenchmarkSort_Sorted_Large(b *testing.B) {
	data := make([]int, 1024)
	for i := range data {
		data[i] = i
	}
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

// ==================== SortUsing Benchmarks ====================

func BenchmarkSortUsing_Int_Small(b *testing.B) {
	data := genInts(8)
	less := func(a, b int) bool { return a < b }
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV1(append([]int{}, data...), less)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV2(append([]int{}, data...), less)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV3(append([]int{}, data...), less)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV4(append([]int{}, data...), less)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV5(append([]int{}, data...), less)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV6(append([]int{}, data...), less)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV7(append([]int{}, data...), less)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV8(append([]int{}, data...), less)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV9(append([]int{}, data...), less)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV10(append([]int{}, data...), less)
		}
	})
}

func BenchmarkSortUsing_Int_Medium(b *testing.B) {
	data := genInts(64)
	less := func(a, b int) bool { return a < b }
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV1(append([]int{}, data...), less)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV2(append([]int{}, data...), less)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV3(append([]int{}, data...), less)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV4(append([]int{}, data...), less)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV5(append([]int{}, data...), less)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV6(append([]int{}, data...), less)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV7(append([]int{}, data...), less)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV8(append([]int{}, data...), less)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV9(append([]int{}, data...), less)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV10(append([]int{}, data...), less)
		}
	})
}

func BenchmarkSortUsing_Int_Large(b *testing.B) {
	data := genInts(1024)
	less := func(a, b int) bool { return a < b }
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV1(append([]int{}, data...), less)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV2(append([]int{}, data...), less)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV3(append([]int{}, data...), less)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV4(append([]int{}, data...), less)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV5(append([]int{}, data...), less)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV6(append([]int{}, data...), less)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV7(append([]int{}, data...), less)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV8(append([]int{}, data...), less)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV9(append([]int{}, data...), less)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV10(append([]int{}, data...), less)
		}
	})
}
