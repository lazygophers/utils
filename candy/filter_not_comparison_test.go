package candy

import "testing"

// 性能对比测试：新实现 vs 原始实现

// 原始实现（保留用于对比）
func filterNotOriginal[T any](ss []T, f func(T) bool) []T {
	us := make([]T, 0)
	for _, s := range ss {
		if !f(s) {
			us = append(us, s)
		}
	}
	return us
}

func BenchmarkFilterNot_Comparison_Original_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotOriginal(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_Comparison_New_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNot(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_Comparison_Original_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotOriginal(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_Comparison_New_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNot(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_Comparison_Original_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotOriginal(largeDataInt, predicate)
	}
}

func BenchmarkFilterNot_Comparison_New_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNot(largeDataInt, predicate)
	}
}
