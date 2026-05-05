package candy

import (
	"math/rand"
	"slices"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

func generateIntSlice(size int, maxVal int) []int {
	slice := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		slice[i] = r.Intn(maxVal)
	}
	return slice
}

func generateSortedIntSlice(size int, maxVal int) []int {
	slice := generateIntSlice(size, maxVal)
	slices.Sort(slice)
	return slice
}

// 方案1: Baseline
func containsBaseline[T constraints.Ordered](ss []T, s T) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}

// 方案2: Int Direct
func containsIntDirect(ss []int, target int) bool {
	for _, v := range ss {
		if v == target {
			return true
		}
	}
	return false
}

// 方案3: Binary Search
func containsBinarySearch(ss []int, target int) bool {
	left, right := 0, len(ss)-1
	for left <= right {
		mid := left + (right-left)/2
		if ss[mid] == target {
			return true
		} else if ss[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 方案4: Slices Std
func containsSlicesStd[T comparable](ss []T, target T) bool {
	return slices.Contains(ss, target)
}

// 方案5: Unroll4
func containsUnroll4(ss []int, target int) bool {
	n := len(ss)
	for i := 0; i < n; i += 4 {
		if i < n && ss[i] == target {
			return true
		}
		if i+1 < n && ss[i+1] == target {
			return true
		}
		if i+2 < n && ss[i+2] == target {
			return true
		}
		if i+3 < n && ss[i+3] == target {
			return true
		}
	}
	return false
}

// 方案6: Index Access
func containsIndexAccess(ss []int, target int) bool {
	for i := 0; i < len(ss); i++ {
		if ss[i] == target {
			return true
		}
	}
	return false
}

func BenchmarkContains_Small_First(b *testing.B) {
	slice := generateIntSlice(10, 1000)
	target := slice[0]

	b.Run("Baseline", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsBaseline(slice, target)
		}
	})

	b.Run("IntDirect", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIntDirect(slice, target)
		}
	})

	b.Run("SlicesStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsSlicesStd(slice, target)
		}
	})

	b.Run("Unroll4", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsUnroll4(slice, target)
		}
	})

	b.Run("IndexAccess", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIndexAccess(slice, target)
		}
	})
}

func BenchmarkContains_Medium_Middle(b *testing.B) {
	slice := generateIntSlice(100, 1000)
	target := slice[50]

	b.Run("Baseline", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsBaseline(slice, target)
		}
	})

	b.Run("IntDirect", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIntDirect(slice, target)
		}
	})

	b.Run("SlicesStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsSlicesStd(slice, target)
		}
	})

	b.Run("Unroll4", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsUnroll4(slice, target)
		}
	})

	b.Run("IndexAccess", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIndexAccess(slice, target)
		}
	})
}

func BenchmarkContains_Large_Last(b *testing.B) {
	slice := generateIntSlice(1000, 10000)
	target := slice[len(slice)-1]

	b.Run("Baseline", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsBaseline(slice, target)
		}
	})

	b.Run("IntDirect", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIntDirect(slice, target)
		}
	})

	b.Run("SlicesStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsSlicesStd(slice, target)
		}
	})

	b.Run("Unroll4", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsUnroll4(slice, target)
		}
	})

	b.Run("IndexAccess", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIndexAccess(slice, target)
		}
	})
}

func BenchmarkContains_Sorted_Binary(b *testing.B) {
	slice := generateSortedIntSlice(1000, 10000)
	target := slice[500]

	b.Run("Baseline", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsBaseline(slice, target)
		}
	})

	b.Run("BinarySearch", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsBinarySearch(slice, target)
		}
	})

	b.Run("SlicesStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsSlicesStd(slice, target)
		}
	})
}

func BenchmarkContains_NotFound(b *testing.B) {
	slice := generateIntSlice(100, 1000)
	target := 99999

	b.Run("Baseline", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsBaseline(slice, target)
		}
	})

	b.Run("IntDirect", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIntDirect(slice, target)
		}
	})

	b.Run("SlicesStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsSlicesStd(slice, target)
		}
	})

	b.Run("Unroll4", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsUnroll4(slice, target)
		}
	})

	b.Run("IndexAccess", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			containsIndexAccess(slice, target)
		}
	})
}
