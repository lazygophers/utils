package candy

import "testing"

// ==================== Contains 基准测试 ====================

func BenchmarkContains_Int_Small_First_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, 0)
	}
}

func BenchmarkContains_Int_Small_Last_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, 9)
	}
}

func BenchmarkContains_Int_Medium_Middle_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, 500)
	}
}

func BenchmarkContains_Int_Large_NotFound_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, -1)
	}
}

func BenchmarkContains_String_Small_V1(b *testing.B) {
	data := make([]string, 10)
	for i := 0; i < 10; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, "test")
	}
}

func BenchmarkContains_String_Large_NotFound_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, "notfound")
	}
}

func BenchmarkContains_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(data, 42)
	}
}

// ==================== ContainsUsing 基准测试 ====================

func BenchmarkContainsUsing_Int_Small_First_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	predicate := func(v int) bool { return v == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}

func BenchmarkContainsUsing_Int_Small_Last_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	predicate := func(v int) bool { return v == 9 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}

func BenchmarkContainsUsing_Int_Medium_Middle_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(v int) bool { return v == 500 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}

func BenchmarkContainsUsing_Int_Large_NotFound_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	predicate := func(v int) bool { return v == -1 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}

func BenchmarkContainsUsing_Int_GreaterThan_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(v int) bool { return v > 500 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}

func BenchmarkContainsUsing_Int_LessThan_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(v int) bool { return v < 100 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}

func BenchmarkContainsUsing_Empty_V1(b *testing.B) {
	data := []int{}
	predicate := func(v int) bool { return v == 42 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsUsing(data, predicate)
	}
}
