package candy

import "testing"

// ==================== First 基准测试 ====================

func BenchmarkFirst_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

func BenchmarkFirst_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		First(data)
	}
}

// ==================== FirstOr 基准测试 ====================

func BenchmarkFirstOr_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, -1)
	}
}

func BenchmarkFirstOr_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstOr(data, "default")
	}
}

// ==================== Last 基准测试 ====================

func BenchmarkLast_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

func BenchmarkLast_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Last(data)
	}
}

// ==================== LastOr 基准测试 ====================

func BenchmarkLastOr_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, -1)
	}
}

func BenchmarkLastOr_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastOr(data, "default")
	}
}

// ==================== Bottom 基准测试 ====================

func BenchmarkBottom_Int_Small_N1_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 1)
	}
}

func BenchmarkBottom_Int_Small_N5_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 5)
	}
}

func BenchmarkBottom_Int_Medium_N10_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 10)
	}
}

func BenchmarkBottom_Int_Large_N100_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 100)
	}
}

func BenchmarkBottom_Int_Large_N1000_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 1000)
	}
}

func BenchmarkBottom_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 10)
	}
}
