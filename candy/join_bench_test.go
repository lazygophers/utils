package candy

import (
	"strconv"
	"testing"
)

// 生成测试数据
func genJoinInts(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = (i * 13) % 1000
	}
	return s
}

func genJoinStrings(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = strconv.Itoa((i * 13) % 1000)
	}
	return s
}

// ==================== Int 类型的基准测试 ====================

func BenchmarkJoin_Int_Small(b *testing.B) {
	data := genJoinInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_Int_Medium(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_Int_Large(b *testing.B) {
	data := genJoinInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

// ==================== String 类型的基准测试 ====================

func BenchmarkJoin_String_Small(b *testing.B) {
	data := genJoinStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_String_Medium(b *testing.B) {
	data := genJoinStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_String_Large(b *testing.B) {
	data := genJoinStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

// ==================== 不同分隔符的基准测试 ====================

func BenchmarkJoin_Int_Comma(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_Int_Dash(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "-")
	}
}

func BenchmarkJoin_Int_Pipe(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "|")
	}
}

// ==================== 空分隔符测试 ====================

func BenchmarkJoin_Int_EmptyGlue(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "")
	}
}

func BenchmarkJoin_String_EmptyGlue(b *testing.B) {
	data := genJoinStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "")
	}
}

// ==================== 空切片测试 ====================

func BenchmarkJoin_Int_EmptySlice(b *testing.B) {
	data := genJoinInts(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_String_EmptySlice(b *testing.B) {
	data := genJoinStrings(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}
