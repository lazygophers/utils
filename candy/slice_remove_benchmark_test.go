package candy

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// Benchmark Remove 函数的各种实现方案

type testType int

// 方案1: 当前实现（两遍扫描）
func BenchmarkRemove_Current_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Remove(data, toRemove)
	}
}

// 方案2: 单遍扫描 + append
func RemoveAppend[T constraints.Ordered](ss []T, toRemove []T) []T {
	if len(ss) == 0 {
		return make([]T, 0)
	}

	if len(toRemove) == 0 {
		result := make([]T, len(ss))
		copy(result, ss)
		return result
	}

	removeSet := make(map[T]bool, len(toRemove))
	for i := 0; i < len(toRemove); i++ {
		removeSet[toRemove[i]] = true
	}

	result := make([]T, 0, len(ss))
	for i := 0; i < len(ss); i++ {
		if !removeSet[ss[i]] {
			result = append(result, ss[i])
		}
	}
	return result
}

func BenchmarkRemove_Append_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveAppend(data, toRemove)
	}
}

// 方案3: 使用 range 循环
func RemoveRange[T constraints.Ordered](ss []T, toRemove []T) []T {
	if len(ss) == 0 {
		return make([]T, 0)
	}

	if len(toRemove) == 0 {
		result := make([]T, len(ss))
		copy(result, ss)
		return result
	}

	removeSet := make(map[T]bool, len(toRemove))
	for _, item := range toRemove {
		removeSet[item] = true
	}

	result := make([]T, 0, len(ss))
	for _, item := range ss {
		if !removeSet[item] {
			result = append(result, item)
		}
	}
	return result
}

func BenchmarkRemove_Range_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveRange(data, toRemove)
	}
}

// 方案4: 使用 struct{} 作为 map 值
func RemoveStruct[T constraints.Ordered](ss []T, toRemove []T) []T {
	if len(ss) == 0 {
		return make([]T, 0)
	}

	if len(toRemove) == 0 {
		result := make([]T, len(ss))
		copy(result, ss)
		return result
	}

	removeSet := make(map[T]struct{}, len(toRemove))
	for i := 0; i < len(toRemove); i++ {
		removeSet[toRemove[i]] = struct{}{}
	}

	result := make([]T, 0, len(ss))
	for i := 0; i < len(ss); i++ {
		if _, ok := removeSet[ss[i]]; !ok {
			result = append(result, ss[i])
		}
	}
	return result
}

func BenchmarkRemove_Struct_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveStruct(data, toRemove)
	}
}

// 方案5: 索引循环 + 两遍扫描（优化版）
func RemoveOptimized[T constraints.Ordered](ss []T, toRemove []T) []T {
	if len(ss) == 0 {
		return make([]T, 0)
	}

	if len(toRemove) == 0 {
		result := make([]T, len(ss))
		copy(result, ss)
		return result
	}

	removeSet := make(map[T]bool, len(toRemove))
	for i := 0; i < len(toRemove); i++ {
		removeSet[toRemove[i]] = true
	}

	// 第一遍扫描：计算保留元素数量
	count := 0
	for i := 0; i < len(ss); i++ {
		if !removeSet[ss[i]] {
			count++
		}
	}

	// 第二遍扫描：填充结果
	result := make([]T, count)
	j := 0
	for i := 0; i < len(ss); i++ {
		if !removeSet[ss[i]] {
			result[j] = ss[i]
			j++
		}
	}
	return result
}

func BenchmarkRemove_Optimized_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveOptimized(data, toRemove)
	}
}

// 中等数据集测试
func BenchmarkRemove_Current_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Remove(data, toRemove)
	}
}

func BenchmarkRemove_Append_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveAppend(data, toRemove)
	}
}

func BenchmarkRemove_Range_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveRange(data, toRemove)
	}
}

func BenchmarkRemove_Optimized_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveOptimized(data, toRemove)
	}
}

// 大数据集测试
func BenchmarkRemove_Current_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Remove(data, toRemove)
	}
}

func BenchmarkRemove_Append_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveAppend(data, toRemove)
	}
}

func BenchmarkRemove_Range_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveRange(data, toRemove)
	}
}

func BenchmarkRemove_Optimized_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveOptimized(data, toRemove)
	}
}

// 空移除列表测试
func BenchmarkRemove_Current_EmptyRemove(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Remove(data, toRemove)
	}
}

func BenchmarkRemove_Append_EmptyRemove(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveAppend(data, toRemove)
	}
}

func BenchmarkRemove_Optimized_EmptyRemove(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveOptimized(data, toRemove)
	}
}

// 大量移除测试
func BenchmarkRemove_Current_HighRemoval(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 500)
	for i := 0; i < 500; i++ {
		toRemove[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Remove(data, toRemove)
	}
}

func BenchmarkRemove_Append_HighRemoval(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 500)
	for i := 0; i < 500; i++ {
		toRemove[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveAppend(data, toRemove)
	}
}

func BenchmarkRemove_Optimized_HighRemoval(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 500)
	for i := 0; i < 500; i++ {
		toRemove[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveOptimized(data, toRemove)
	}
}
