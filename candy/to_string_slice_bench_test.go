package candy

import (
	"testing"
)

// ==================== ToStringSlice Benchmark ====================

func BenchmarkToStringSlice_Empty(b *testing.B) {
	var arr []string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Nil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(nil)
	}
}

func BenchmarkToStringSlice_String_Small(b *testing.B) {
	arr := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_String_Medium(b *testing.B) {
	arr := make([]string, 100)
	for i := range arr {
		arr[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_String_Large(b *testing.B) {
	arr := make([]string, 10000)
	for i := range arr {
		arr[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Int_Small(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Int_Medium(b *testing.B) {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Int_Large(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Int64_Small(b *testing.B) {
	arr := []int64{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Int64_Medium(b *testing.B) {
	arr := make([]int64, 100)
	for i := range arr {
		arr[i] = int64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Int64_Large(b *testing.B) {
	arr := make([]int64, 10000)
	for i := range arr {
		arr[i] = int64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Float64_Small(b *testing.B) {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Float64_Medium(b *testing.B) {
	arr := make([]float64, 100)
	for i := range arr {
		arr[i] = float64(i) + 0.5
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Float64_Large(b *testing.B) {
	arr := make([]float64, 10000)
	for i := range arr {
		arr[i] = float64(i) + 0.5
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Bool_Small(b *testing.B) {
	arr := []bool{true, false, true, false, true}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Bool_Medium(b *testing.B) {
	arr := make([]bool, 100)
	for i := range arr {
		arr[i] = i%2 == 0
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Bool_Large(b *testing.B) {
	arr := make([]bool, 10000)
	for i := range arr {
		arr[i] = i%2 == 0
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Any_Small(b *testing.B) {
	arr := []any{"test", 123, true, 3.14}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Any_Medium(b *testing.B) {
	arr := make([]any, 100)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Any_Large(b *testing.B) {
	arr := make([]any, 10000)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_SingleString(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(s)
	}
}

func BenchmarkToStringSlice_SingleStringWithComma(b *testing.B) {
	s := "a,b,c,d,e"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(s)
	}
}

func BenchmarkToStringSlice_Uint64_Small(b *testing.B) {
	arr := []uint64{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Uint64_Medium(b *testing.B) {
	arr := make([]uint64, 100)
	for i := range arr {
		arr[i] = uint64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

func BenchmarkToStringSlice_Uint64_Large(b *testing.B) {
	arr := make([]uint64, 10000)
	for i := range arr {
		arr[i] = uint64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToStringSlice(arr)
	}
}

// ==================== ToArrayString Benchmark ====================
// ToArrayString 是 ToStringSlice 的别名，这里测试别名调用开销

func BenchmarkToArrayString_String_Small(b *testing.B) {
	arr := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToArrayString(arr)
	}
}

func BenchmarkToArrayString_Int_Medium(b *testing.B) {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToArrayString(arr)
	}
}

func BenchmarkToArrayString_Int64_Large(b *testing.B) {
	arr := make([]int64, 10000)
	for i := range arr {
		arr[i] = int64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToArrayString(arr)
	}
}

func BenchmarkToArrayString_SingleString(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToArrayString(s)
	}
}

func BenchmarkToArrayString_Nil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToArrayString(nil)
	}
}
