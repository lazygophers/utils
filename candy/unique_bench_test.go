package candy

import (
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

// ========== 1. 当前实现（baseline map） ==========
func uniqueBaseline[T constraints.Ordered](ss []T) []T {
	ret := make([]T, 0)
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}
	return ret
}

// ========== 2. 排序后去重（不保持顺序） ==========
func uniqueSort[T constraints.Ordered](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	result := []T{sorted[0]}
	for i := 1; i < len(sorted); i++ {
		if sorted[i] != sorted[i-1] {
			result = append(result, sorted[i])
		}
	}
	return result
}

// ========== 3. 预分配 map 容量优化 ==========
func uniquePrealloc[T constraints.Ordered](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}

	estimatedSize := len(ss)
	if estimatedSize > 100 {
		estimatedSize = estimatedSize * 70 / 100
	}

	ret := make([]T, 0, estimatedSize)
	m := make(map[T]struct{}, estimatedSize)

	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}
	return ret
}

// ========== 4. 小切片优化（线性搜索） ==========
func uniqueSmallSlice[T constraints.Ordered](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}

	const threshold = 32
	if len(ss) < threshold {
		ret := make([]T, 0, len(ss))
		for _, s := range ss {
			found := false
			for _, v := range ret {
				if v == s {
					found = true
					break
				}
			}
			if !found {
				ret = append(ret, s)
			}
		}
		return ret
	}

	ret := make([]T, 0, len(ss))
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}
	return ret
}

// ========== 5. 混合策略 ==========
func uniqueHybrid[T constraints.Ordered](ss []T) []T {
	if len(ss) < 64 {
		ret := make([]T, 0, len(ss))
		for _, s := range ss {
			found := false
			for _, v := range ret {
				if v == s {
					found = true
					break
				}
			}
			if !found {
				ret = append(ret, s)
			}
		}
		return ret
	}

	ret := make([]T, 0, len(ss))
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}
	return ret
}

// ========== 6. 已排序切片优化 ==========
func uniqueSortedOptimized[T constraints.Ordered](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}

	isSorted := true
	for i := 1; i < len(ss); i++ {
		if ss[i] < ss[i-1] {
			isSorted = false
			break
		}
	}

	if isSorted {
		result := make([]T, 0, len(ss))
		result = append(result, ss[0])
		for i := 1; i < len(ss); i++ {
			if ss[i] != ss[i-1] {
				result = append(result, ss[i])
			}
		}
		return result
	}

	ret := make([]T, 0, len(ss))
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}
	return ret
}

// ========== 7. 原地重排 ==========
func uniqueInPlace[T constraints.Ordered](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}

	m := make(map[T]struct{}, len(ss))
	writeIdx := 0

	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ss[writeIdx] = s
			writeIdx++
		}
	}

	return ss[:writeIdx]
}

// ========== 8. 双指针法（需要已排序） ==========
func uniqueTwoPointer[T constraints.Ordered](ss []T) []T {
	if len(ss) == 0 {
		return []T{}
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	slow := 0
	for fast := 1; fast < len(sorted); fast++ {
		if sorted[fast] != sorted[slow] {
			slow++
			sorted[slow] = sorted[fast]
		}
	}

	return sorted[:slow+1]
}

// ========== 辅助函数：生成测试数据 ==========

func generateNoDuplicates(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i
	}
	return result
}

func generateHighDuplicates(size int) []int {
	uniqueCount := size / 10
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i % uniqueCount
	}
	return result
}

func generateMediumDuplicates(size int) []int {
	uniqueCount := size / 2
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i % uniqueCount
	}
	return result
}

func generateLowDuplicates(size int) []int {
	uniqueCount := size * 9 / 10
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i % uniqueCount
	}
	return result
}

func generateSortedDuplicates(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = i / 4
	}
	return result
}

// ========== 基准测试 ==========

func BenchmarkUnique_Small_NoDuplicates(b *testing.B) {
	data := generateNoDuplicates(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uniqueBaseline(data)
	}
}

func BenchmarkUnique_Medium_MediumDuplicates(b *testing.B) {
	data := generateMediumDuplicates(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uniqueBaseline(data)
	}
}

func BenchmarkUnique_Large_NoDuplicates(b *testing.B) {
	data := generateNoDuplicates(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uniqueBaseline(data)
	}
}

func BenchmarkUnique_Large_HighDuplicates(b *testing.B) {
	data := generateHighDuplicates(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uniqueBaseline(data)
	}
}

func BenchmarkUnique_XL_LowDuplicates(b *testing.B) {
	data := generateLowDuplicates(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = uniqueBaseline(data)
	}
}

// ========== 所有实现的性能对比 ==========

func BenchmarkAll_Medium_MediumDuplicates(b *testing.B) {
	data := generateMediumDuplicates(100)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueBaseline(data)
		}
	})

	b.Run("Sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSort(data)
		}
	})

	b.Run("Prealloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniquePrealloc(data)
		}
	})

	b.Run("SmallSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSmallSlice(data)
		}
	})

	b.Run("Hybrid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueHybrid(data)
		}
	})

	b.Run("SortedOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSortedOptimized(data)
		}
	})

	b.Run("InPlace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dataCopy := make([]int, len(data))
			copy(dataCopy, data)
			_ = uniqueInPlace(dataCopy)
		}
	})

	b.Run("TwoPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueTwoPointer(data)
		}
	})
}

func BenchmarkAll_Large_NoDuplicates(b *testing.B) {
	data := generateNoDuplicates(1000)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueBaseline(data)
		}
	})

	b.Run("Sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSort(data)
		}
	})

	b.Run("Prealloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniquePrealloc(data)
		}
	})

	b.Run("SmallSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSmallSlice(data)
		}
	})

	b.Run("Hybrid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueHybrid(data)
		}
	})

	b.Run("SortedOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSortedOptimized(data)
		}
	})

	b.Run("InPlace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			dataCopy := make([]int, len(data))
			copy(dataCopy, data)
			_ = uniqueInPlace(dataCopy)
		}
	})

	b.Run("TwoPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueTwoPointer(data)
		}
	})
}

func BenchmarkAll_XL_HighDuplicates(b *testing.B) {
	data := generateHighDuplicates(10000)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueBaseline(data)
		}
	})

	b.Run("Sort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSort(data)
		}
	})

	b.Run("Prealloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniquePrealloc(data)
		}
	})

	b.Run("SmallSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSmallSlice(data)
		}
	})

	b.Run("Hybrid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueHybrid(data)
		}
	})

	b.Run("SortedOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSortedOptimized(data)
		}
	})

	b.Run("TwoPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueTwoPointer(data)
		}
	})
}

// ========== 已排序数据的特殊测试 ==========

func BenchmarkSorted_All(b *testing.B) {
	data := generateSortedDuplicates(1000)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueBaseline(data)
		}
	})

	b.Run("TwoPointer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueTwoPointer(data)
		}
	})

	b.Run("SortedOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = uniqueSortedOptimized(data)
		}
	})
}
