package candy

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// ============ 性能对比测试 ============

// RemoveIndex 原始实现(用于对比)
func removeIndexOld[T any](ss []T, index int) []T {
	if len(ss) == 0 || index < 0 || index >= len(ss) {
		return make([]T, 0)
	}
	if index == 0 {
		return ss[1:]
	}
	if index == len(ss)-1 {
		return ss[:len(ss)-1]
	}
	return append(ss[:index], ss[index+1:]...)
}

// Remove 原始实现(用于对比)
func removeOld[T constraints.Ordered](ss []T, toRemove []T) []T {
	result := make([]T, 0)
	removeSet := make(map[T]struct{}, len(toRemove))
	for _, item := range toRemove {
		removeSet[item] = struct{}{}
	}
	for _, item := range ss {
		if _, shouldRemove := removeSet[item]; !shouldRemove {
			result = append(result, item)
		}
	}
	return result
}

// RemoveIndex 性能对比
func BenchmarkRemoveIndex_Old_Small(b *testing.B) {
	ss := make([]int, 10)
	for i := range ss {
		ss[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = removeIndexOld(ss, 5)
	}
}

func BenchmarkRemoveIndex_New_Small(b *testing.B) {
	ss := make([]int, 10)
	for i := range ss {
		ss[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RemoveIndex(ss, 5)
	}
}

func BenchmarkRemoveIndex_Old_Large(b *testing.B) {
	ss := make([]int, 10000)
	for i := range ss {
		ss[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = removeIndexOld(ss, 5000)
	}
}

func BenchmarkRemoveIndex_New_Large(b *testing.B) {
	ss := make([]int, 10000)
	for i := range ss {
		ss[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RemoveIndex(ss, 5000)
	}
}

// Remove 性能对比
func BenchmarkRemove_Old(b *testing.B) {
	ss := make([]int, 1000)
	for i := range ss {
		ss[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = removeOld(ss, toRemove)
	}
}

func BenchmarkRemove_New(b *testing.B) {
	ss := make([]int, 1000)
	for i := range ss {
		ss[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Remove(ss, toRemove)
	}
}

// Drop 性能测试(已经是最优,仅验证)
func BenchmarkDrop_Small(b *testing.B) {
	ss := make([]int, 10)
	for i := range ss {
		ss[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Drop(ss, 5)
	}
}

func BenchmarkDrop_Large(b *testing.B) {
	ss := make([]int, 10000)
	for i := range ss {
		ss[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Drop(ss, 5000)
	}
}
