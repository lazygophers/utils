package candy

import (
	"slices"
	"testing"
	"unsafe"
)

// 测试数据结构
type testStruct struct {
	a int
	b string
	c float64
}

// ========== 实现 1: 当前实现 (baseline) ==========
func reverse1[T any](ss []T) (ret []T) {
	ret = make([]T, 0, len(ss))
	for i := len(ss) - 1; i >= 0; i-- {
		ret = append(ret, ss[i])
	}
	return
}

// ========== 实现 2: 双指针交换优化 (原地修改) ==========
func reverse2[T any](ss []T) []T {
	if len(ss) <= 1 {
		return ss
	}
	result := make([]T, len(ss))
	copy(result, ss)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

// ========== 实现 3: 使用 Go 1.21+ slices.Reverse ==========
func reverse3[T any](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}
	result := make([]T, len(ss))
	copy(result, ss)
	slices.Reverse(result)
	return result
}

// ========== 实现 4: 直接索引赋值优化 ==========
func reverse4[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = ss[n-1-i]
	}
	return result
}

// ========== 实现 5: 预分配长度后 append ==========
func reverse5[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}
	result := make([]T, 0, n)
	for i := n - 1; i >= 0; i-- {
		result = append(result, ss[i])
	}
	return result
}

// ========== 实现 6: 循环展开优化 (4x unroll) ==========
func reverse6[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}
	result := make([]T, n)

	// 处理前 4 的倍数部分
	i := 0
	j := n - 1
	for i < j && j-i >= 3 {
		result[i] = ss[j]
		result[i+1] = ss[j-1]
		result[i+2] = ss[j-2]
		result[i+3] = ss[j-3]
		i += 4
		j -= 4
	}

	// 处理剩余元素
	for i <= j {
		result[i] = ss[j]
		i++
		j--
	}

	return result
}

// ========== 实现 7: 针对 []byte 的特殊优化 ==========
func reverse7(ss []byte) []byte {
	n := len(ss)
	if n == 0 {
		return []byte{}
	}
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = ss[n-1-i]
	}
	return result
}

// ========== 实现 8: 使用 unsafe 优化 (仅适用于相同类型) ==========
func reverse8[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}
	result := make([]T, n)

	// 使用 unsafe.Pointer 避免边界检查
	src := (*[1 << 30]T)(unsafe.Pointer(&ss[0]))[:n:n]
	dst := (*[1 << 30]T)(unsafe.Pointer(&result[0]))[:n:n]

	for i := 0; i < n; i++ {
		dst[i] = src[n-1-i]
	}

	return result
}

// ========== 实现 9: 双指针 + 原地交换 + copy 优化 ==========
func reverse9[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		if n == 0 {
			return []T{}
		}
		result := make([]T, 1)
		result[0] = ss[0]
		return result
	}

	result := make([]T, n)
	copy(result, ss)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ========== 实现 10: 混合策略 (小切片用简单循环，大切片用双指针) ==========
func reverse10[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}

	result := make([]T, n)

	// 小切片 (< 32) 使用直接索引
	if n < 32 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
	} else {
		// 大切片使用双指针交换
		copy(result, ss)
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}

// ========== 实现 11: 循环展开优化 (8x unroll) ==========
func reverse11[T any](ss []T) []T {
	n := len(ss)
	if n == 0 {
		return []T{}
	}
	result := make([]T, n)

	// 处理前 8 的倍数部分
	i := 0
	j := n - 1
	for i < j && j-i >= 7 {
		result[i] = ss[j]
		result[i+1] = ss[j-1]
		result[i+2] = ss[j-2]
		result[i+3] = ss[j-3]
		result[i+4] = ss[j-4]
		result[i+5] = ss[j-5]
		result[i+6] = ss[j-6]
		result[i+7] = ss[j-7]
		i += 8
		j -= 8
	}

	// 处理剩余元素
	for i <= j {
		result[i] = ss[j]
		i++
		j--
	}

	return result
}

// ========== 实现 12: 原地 reverse (修改原切片) ==========
// 注意: 这个实现会修改原切片，仅用于性能对比
func reverse12[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ========== 基准测试 ==========

// 小切片 (int)
func BenchmarkReverse_Small_Int_Original(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse1(data)
	}
}

func BenchmarkReverse_Small_Int_TwoPointer(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse2(data)
	}
}

func BenchmarkReverse_Small_Int_SlicesReverse(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse3(data)
	}
}

func BenchmarkReverse_Small_Int_DirectIndex(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse4(data)
	}
}

func BenchmarkReverse_Small_Int_AppendPrealloc(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse5(data)
	}
}

func BenchmarkReverse_Small_Int_Unroll4x(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse6(data)
	}
}

func BenchmarkReverse_Small_Int_Unroll8x(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse11(data)
	}
}

func BenchmarkReverse_Small_Int_Hybrid(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse10(data)
	}
}

// 中等切片 (int)
func BenchmarkReverse_Medium_Int_Original(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse1(data)
	}
}

func BenchmarkReverse_Medium_Int_TwoPointer(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse2(data)
	}
}

func BenchmarkReverse_Medium_Int_SlicesReverse(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse3(data)
	}
}

func BenchmarkReverse_Medium_Int_DirectIndex(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse4(data)
	}
}

func BenchmarkReverse_Medium_Int_AppendPrealloc(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse5(data)
	}
}

func BenchmarkReverse_Medium_Int_Unroll4x(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse6(data)
	}
}

func BenchmarkReverse_Medium_Int_Unroll8x(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse11(data)
	}
}

func BenchmarkReverse_Medium_Int_Hybrid(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse10(data)
	}
}

// 大切片 (int)
func BenchmarkReverse_Large_Int_Original(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse1(data)
	}
}

func BenchmarkReverse_Large_Int_TwoPointer(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse2(data)
	}
}

func BenchmarkReverse_Large_Int_SlicesReverse(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse3(data)
	}
}

func BenchmarkReverse_Large_Int_DirectIndex(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse4(data)
	}
}

func BenchmarkReverse_Large_Int_AppendPrealloc(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse5(data)
	}
}

func BenchmarkReverse_Large_Int_Unroll4x(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse6(data)
	}
}

func BenchmarkReverse_Large_Int_Unroll8x(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse11(data)
	}
}

func BenchmarkReverse_Large_Int_Hybrid(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse10(data)
	}
}

// 字符串切片测试
func BenchmarkReverse_Medium_String_Original(b *testing.B) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = "test-string"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse1(data)
	}
}

func BenchmarkReverse_Medium_String_TwoPointer(b *testing.B) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = "test-string"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse2(data)
	}
}

func BenchmarkReverse_Medium_String_DirectIndex(b *testing.B) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = "test-string"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse4(data)
	}
}

// 结构体切片测试
func BenchmarkReverse_Medium_Struct_Original(b *testing.B) {
	data := make([]testStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = testStruct{a: i, b: "test", c: float64(i)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse1(data)
	}
}

func BenchmarkReverse_Medium_Struct_TwoPointer(b *testing.B) {
	data := make([]testStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = testStruct{a: i, b: "test", c: float64(i)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse2(data)
	}
}

func BenchmarkReverse_Medium_Struct_DirectIndex(b *testing.B) {
	data := make([]testStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = testStruct{a: i, b: "test", c: float64(i)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse4(data)
	}
}

// []byte 特殊优化测试
func BenchmarkReverse_Medium_Byte_Original(b *testing.B) {
	data := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse1(data)
	}
}

func BenchmarkReverse_Medium_Byte_TwoPointer(b *testing.B) {
	data := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse2(data)
	}
}

func BenchmarkReverse_Medium_Byte_ByteOptimized(b *testing.B) {
	data := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = byte(i % 256)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse7(data)
	}
}
