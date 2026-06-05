package anyx

import (
	"fmt"
	"testing"
)

func BenchmarkJoinPath_2_Baseline(b *testing.B) {
	parts := []string{"root", "child"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathBaseline(parts, ".")
	}
}

func BenchmarkJoinPath_2_Optimized(b *testing.B) {
	parts := []string{"root", "child"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathOptimized(parts, ".")
	}
}

func BenchmarkJoinPath_5_Baseline(b *testing.B) {
	parts := []string{"root", "l1", "l2", "l3", "l4"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathBaseline(parts, ".")
	}
}

func BenchmarkJoinPath_5_Optimized(b *testing.B) {
	parts := []string{"root", "l1", "l2", "l3", "l4"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathOptimized(parts, ".")
	}
}

func BenchmarkJoinPath_10_Baseline(b *testing.B) {
	parts := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathBaseline(parts, ".")
	}
}

func BenchmarkJoinPath_10_Optimized(b *testing.B) {
	parts := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathOptimized(parts, ".")
	}
}

func BenchmarkJoinPath_50_Baseline(b *testing.B) {
	parts := make([]string, 50)
	for i := range parts {
		parts[i] = fmt.Sprintf("item%d", i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathBaseline(parts, ".")
	}
}

func BenchmarkJoinPath_50_Optimized(b *testing.B) {
	parts := make([]string, 50)
	for i := range parts {
		parts[i] = fmt.Sprintf("item%d", i)
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathOptimized(parts, ".")
	}
}

func BenchmarkJoinPath_100_Baseline(b *testing.B) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "x"
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathBaseline(parts, ".")
	}
}

func BenchmarkJoinPath_100_Optimized(b *testing.B) {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "x"
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = joinPathOptimized(parts, ".")
	}
}
