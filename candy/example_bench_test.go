package candy_test

import (
	"github.com/lazygophers/utils/candy"
	"testing"
)

// BenchmarkToStringSlice_StringSlice 测试 []string 类型的性能（零拷贝优化）
func BenchmarkToStringSlice_StringSlice_Small(b *testing.B) {
	b.ReportAllocs()
	data := []string{"a", "b", "c"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

func BenchmarkToStringSlice_StringSlice_Large(b *testing.B) {
	b.ReportAllocs()
	data := make([]string, 500)
	for i := range data {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

// BenchmarkToStringSlice_IntSlice 测试 []int 类型的性能（类型断言优化）
func BenchmarkToStringSlice_IntSlice_Small(b *testing.B) {
	b.ReportAllocs()
	data := []int{1, 2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

func BenchmarkToStringSlice_IntSlice_Large(b *testing.B) {
	b.ReportAllocs()
	data := make([]int, 500)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

// BenchmarkToStringSlice_AnySlice 测试 []any 类型的性能（预分配优化）
func BenchmarkToStringSlice_AnySlice_Small(b *testing.B) {
	b.ReportAllocs()
	data := []any{"a", 1, "b"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

func BenchmarkToStringSlice_AnySlice_Large(b *testing.B) {
	b.ReportAllocs()
	data := make([]any, 500)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

// BenchmarkToStringSlice_CommaString 测试逗号分隔字符串的性能
func BenchmarkToStringSlice_CommaString(b *testing.B) {
	b.ReportAllocs()
	data := "a,b,c,d,e,f,g,h,i,j"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}

// BenchmarkToStringSlice_SingleString 测试单个字符串的性能
func BenchmarkToStringSlice_SingleString(b *testing.B) {
	b.ReportAllocs()
	data := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = candy.ToStringSlice(data)
	}
}
