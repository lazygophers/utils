package candy

import (
	"testing"
	"time"
)

// ==================== Max 基准测试 ====================

func BenchmarkMax_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(data...)
	}
}

func BenchmarkMax_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(data...)
	}
}

func BenchmarkMax_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(data...)
	}
}

func BenchmarkMax_Float64_Small_V1(b *testing.B) {
	data := make([]float64, 10)
	for i := 0; i < 10; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(data...)
	}
}

func BenchmarkMax_Float64_Large_V1(b *testing.B) {
	data := make([]float64, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(data...)
	}
}

// ==================== Min 基准测试 ====================

func BenchmarkMin_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(data...)
	}
}

func BenchmarkMin_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(data...)
	}
}

func BenchmarkMin_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(data...)
	}
}

func BenchmarkMin_Float64_Small_V1(b *testing.B) {
	data := make([]float64, 10)
	for i := 0; i < 10; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(data...)
	}
}

func BenchmarkMin_Float64_Large_V1(b *testing.B) {
	data := make([]float64, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(data...)
	}
}

// ==================== Sum 基准测试 ====================

func BenchmarkSum_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

func BenchmarkSum_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

func BenchmarkSum_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

func BenchmarkSum_Float64_Small_V1(b *testing.B) {
	data := make([]float64, 10)
	for i := 0; i < 10; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

func BenchmarkSum_Float64_Large_V1(b *testing.B) {
	data := make([]float64, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

// ==================== Average 基准测试 ====================

func BenchmarkAverage_Int_Small_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}

func BenchmarkAverage_Int_Medium_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}

func BenchmarkAverage_Int_Large_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}

func BenchmarkAverage_Float64_Small_V1(b *testing.B) {
	data := make([]float64, 10)
	for i := 0; i < 10; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}

func BenchmarkAverage_Float64_Large_V1(b *testing.B) {
	data := make([]float64, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}

// ==================== Abs 基准测试 ====================

func BenchmarkAbs_Int_Positive_V1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Abs(42)
	}
}

func BenchmarkAbs_Int_Negative_V1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Abs(-42)
	}
}

func BenchmarkAbs_Float64_Positive_V1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Abs(42.5)
	}
}

func BenchmarkAbs_Float64_Negative_V1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Abs(-42.5)
	}
}

func BenchmarkAbs_Int_Loop_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = -i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range data {
			Abs(v)
		}
	}
}

// ==================== Duration 基准测试 ====================

func BenchmarkSum_Duration_Small_V1(b *testing.B) {
	data := make([]time.Duration, 10)
	for i := 0; i < 10; i++ {
		data[i] = time.Duration(i) * time.Millisecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

func BenchmarkSum_Duration_Large_V1(b *testing.B) {
	data := make([]time.Duration, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = time.Duration(i) * time.Millisecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(data...)
	}
}

func BenchmarkAverage_Duration_Small_V1(b *testing.B) {
	data := make([]time.Duration, 10)
	for i := 0; i < 10; i++ {
		data[i] = time.Duration(i) * time.Millisecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}

func BenchmarkAverage_Duration_Large_V1(b *testing.B) {
	data := make([]time.Duration, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = time.Duration(i) * time.Millisecond
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(data...)
	}
}
