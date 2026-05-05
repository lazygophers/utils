package candy

import (
	"slices"
	"testing"
)

// ========== 实现 1: 当前实现 (baseline) ==========
func reverseImpl1[T any](ss []T) (ret []T) {
	ret = make([]T, 0, len(ss))
	for i := len(ss) - 1; i >= 0; i-- {
		ret = append(ret, ss[i])
	}
	return
}

// ========== 实现 2: 双指针交换优化 ==========
func reverseImpl2[T any](ss []T) []T {
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
func reverseImpl3[T any](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}
	result := make([]T, len(ss))
	copy(result, ss)
	slices.Reverse(result)
	return result
}

// ========== 实现 4: 直接索引赋值优化 ==========
func reverseImpl4[T any](ss []T) []T {
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
func reverseImpl5[T any](ss []T) []T {
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

// ========== 基准测试 ==========

// 小切片 (int, 10个元素)
func BenchmarkReverseImpl1_Small_Int(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl1(data)
	}
}

func BenchmarkReverseImpl2_Small_Int(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl2(data)
	}
}

func BenchmarkReverseImpl3_Small_Int(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl3(data)
	}
}

func BenchmarkReverseImpl4_Small_Int(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl4(data)
	}
}

func BenchmarkReverseImpl5_Small_Int(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl5(data)
	}
}

// 中等切片 (int, 1000个元素)
func BenchmarkReverseImpl1_Medium_Int(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl1(data)
	}
}

func BenchmarkReverseImpl2_Medium_Int(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl2(data)
	}
}

func BenchmarkReverseImpl3_Medium_Int(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl3(data)
	}
}

func BenchmarkReverseImpl4_Medium_Int(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl4(data)
	}
}

func BenchmarkReverseImpl5_Medium_Int(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl5(data)
	}
}

// 大切片 (int, 100000个元素)
func BenchmarkReverseImpl1_Large_Int(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl1(data)
	}
}

func BenchmarkReverseImpl2_Large_Int(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl2(data)
	}
}

func BenchmarkReverseImpl3_Large_Int(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl3(data)
	}
}

func BenchmarkReverseImpl4_Large_Int(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl4(data)
	}
}

func BenchmarkReverseImpl5_Large_Int(b *testing.B) {
	data := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseImpl5(data)
	}
}
