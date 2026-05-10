package anyx

import (
	"fmt"
	"strings"
	"testing"
)

// 原始实现
func joinPathBaseline(parts []string, sep string) string {
	switch len(parts) {
	case 0:
		return ""
	case 1:
		return parts[0]
	}

	result := parts[0]
	for _, part := range parts[1:] {
		result += sep + part
	}
	return result
}

// 优化实现
func joinPathOptimized(parts []string, sep string) string {
	switch len(parts) {
	case 0:
		return ""
	case 1:
		return parts[0]
	case 2:
		return parts[0] + sep + parts[1]
	case 3:
		return parts[0] + sep + parts[1] + sep + parts[2]
	}

	totalLen := len(sep) * (len(parts) - 1)
	for _, part := range parts {
		totalLen += len(part)
	}

	var builder strings.Builder
	builder.Grow(totalLen)
	builder.WriteString(parts[0])
	for _, part := range parts[1:] {
		builder.WriteString(sep)
		builder.WriteString(part)
	}
	return builder.String()
}

// 功能验证
func TestJoinPath_FunctionalEquivalence(t *testing.T) {
	testCases := [][]string{
		{},
		{"single"},
		{"a", "b"},
		{"a", "b", "c"},
		{"root", "level1", "level2", "level3", "level4"},
		make([]string, 100),
	}

	separators := []string{".", "/", "::", "\n"}

	for _, parts := range testCases {
		for _, sep := range separators {
			baseline := joinPathBaseline(parts, sep)
			optimized := joinPathOptimized(parts, sep)
			if baseline != optimized {
				t.Errorf("Mismatch for len=%d, sep=%q: baseline=%q, optimized=%q",
					len(parts), sep, baseline, optimized)
			}
		}
	}
}

// Benchmark 测试
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
