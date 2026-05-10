package cryptox

import (
	"testing"
)

// 测试数据
var testInput = "The quick brown fox jumps over the lazy dog. The quick brown fox jumps over the lazy dog."

// Benchmark 函数
func BenchmarkSha256Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256Original(testInput)
	}
}

func BenchmarkSha256V1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V1(testInput)
	}
}

func BenchmarkSha256V2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V2(testInput)
	}
}

func BenchmarkSha256V3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V3(testInput)
	}
}

func BenchmarkSha256V4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V4(testInput)
	}
}

func BenchmarkSha256V5(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V5(testInput)
	}
}

func BenchmarkSha256V6(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V6(testInput)
	}
}

func BenchmarkSha256V7(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V7(testInput)
	}
}

func BenchmarkSha256V8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V8(testInput)
	}
}

func BenchmarkSha256V9(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V9(testInput)
	}
}

func BenchmarkSha256V10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V10(testInput)
	}
}

func BenchmarkSha256V11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V11(testInput)
	}
}

func BenchmarkSha256V12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Sha256V12(testInput)
	}
}

// 正确性验证测试
func TestSha256Variants(t *testing.T) {
	input := "Hello, World!"
	expected := Sha256Original[string](input)

	variants := []struct {
		name string
		fn   func(string) string
	}{
		{"Original", func(s string) string { return Sha256Original[string](s) }},
		{"V1", func(s string) string { return Sha256V1[string](s) }},
		{"V2", func(s string) string { return Sha256V2[string](s) }},
		{"V3", func(s string) string { return Sha256V3[string](s) }},
		{"V4", func(s string) string { return Sha256V4[string](s) }},
		{"V5", func(s string) string { return Sha256V5[string](s) }},
		{"V6", func(s string) string { return Sha256V6[string](s) }},
		{"V7", func(s string) string { return Sha256V7[string](s) }},
		{"V8", func(s string) string { return Sha256V8[string](s) }},
		{"V9", func(s string) string { return Sha256V9[string](s) }},
		{"V10", func(s string) string { return Sha256V10[string](s) }},
		{"V11", func(s string) string { return Sha256V11[string](s) }},
		{"V12", func(s string) string { return Sha256V12[string](s) }},
	}

	for _, tc := range variants {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.fn(input)
			if result != expected {
				t.Errorf("%s: expected %q, got %q", tc.name, expected, result)
			}
		})
	}
}
