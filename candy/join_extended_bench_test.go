package candy

import "testing"

// ==================== Join 基准测试 ====================

func BenchmarkJoin_Int_Small_Default_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Int_Small_Comma_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data, ",")
	}
}

func BenchmarkJoin_Int_Medium_Default_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Int_Large_Default_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_String_Small_Default_V1(b *testing.B) {
	data := make([]string, 10)
	for i := 0; i < 10; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_String_Medium_Dash_V1(b *testing.B) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data, "-")
	}
}

func BenchmarkJoin_String_Large_Space_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data, " ")
	}
}

func BenchmarkJoin_Int_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Float64_Large_Default_V1(b *testing.B) {
	data := make([]float64, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Int_Single_V1(b *testing.B) {
	data := []int{42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}
