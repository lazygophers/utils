package candy

import "testing"

// ==================== Index 基准测试 ====================

func BenchmarkIndex_Int_Small_First_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, 0)
	}
}

func BenchmarkIndex_Int_Small_Last_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, 9)
	}
}

func BenchmarkIndex_Int_Medium_Middle_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, 500)
	}
}

func BenchmarkIndex_Int_Large_Last_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, 9999)
	}
}

func BenchmarkIndex_Int_NotFound_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, -1)
	}
}

func BenchmarkIndex_String_Small_V1(b *testing.B) {
	data := make([]string, 10)
	for i := 0; i < 10; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, "test")
	}
}

func BenchmarkIndex_String_Large_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(data, "test")
	}
}

// ==================== Same 基准测试 ====================

func BenchmarkSame_Int_Small_NoOverlap_V1(b *testing.B) {
	a := make([]int, 10)
	b2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
		b2[i] = i + 20
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Same(a, b2)
	}
}

func BenchmarkSame_Int_Small_HalfOverlap_V1(b *testing.B) {
	a := make([]int, 10)
	b2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
		b2[i] = i + 5
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Same(a, b2)
	}
}

func BenchmarkSame_Int_Medium_FullOverlap_V1(b *testing.B) {
	a := make([]int, 1000)
	b2 := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
		b2[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Same(a, b2)
	}
}

func BenchmarkSame_Int_Large_PartialOverlap_V1(b *testing.B) {
	a := make([]int, 10000)
	b2 := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = i
		b2[i] = i + 5000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Same(a, b2)
	}
}

func BenchmarkSame_Int_OneEmpty_V1(b *testing.B) {
	a := make([]int, 1000)
	var b2 []int
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Same(a, b2)
	}
}

// ==================== Diff 基准测试 ====================

func BenchmarkDiff_Int_Small_NoDiff_V1(b *testing.B) {
	a := make([]int, 10)
	b2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
		b2[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(a, b2)
	}
}

func BenchmarkDiff_Int_Small_AllDiff_V1(b *testing.B) {
	a := make([]int, 10)
	b2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
		b2[i] = i + 20
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(a, b2)
	}
}

func BenchmarkDiff_Int_Medium_PartialDiff_V1(b *testing.B) {
	a := make([]int, 1000)
	b2 := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
		b2[i] = i + 500
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(a, b2)
	}
}

func BenchmarkDiff_Int_Large_Complex_V1(b *testing.B) {
	a := make([]int, 10000)
	b2 := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		a[i] = i
		b2[i] = i + 5000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(a, b2)
	}
}

func BenchmarkDiff_Int_OneEmpty_V1(b *testing.B) {
	a := make([]int, 1000)
	var b2 []int
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(a, b2)
	}
}

// ==================== Drop 基准测试 ====================

func BenchmarkDrop_Int_Small_1_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, 1)
	}
}

func BenchmarkDrop_Int_Small_5_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, 5)
	}
}

func BenchmarkDrop_Int_Medium_100_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, 100)
	}
}

func BenchmarkDrop_Int_Large_1000_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, 1000)
	}
}

func BenchmarkDrop_Int_Zero_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, 0)
	}
}

func BenchmarkDrop_Int_Negative_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(data, -1)
	}
}

// ==================== RemoveIndex 基准测试 ====================

func BenchmarkRemoveIndex_Int_Small_First_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndex(data, 0)
	}
}

func BenchmarkRemoveIndex_Int_Small_Last_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndex(data, 9)
	}
}

func BenchmarkRemoveIndex_Int_Small_Middle_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndex(data, 5)
	}
}

func BenchmarkRemoveIndex_Int_Medium_Middle_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndex(data, 500)
	}
}

func BenchmarkRemoveIndex_Int_Large_Middle_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndex(data, 5000)
	}
}

func BenchmarkRemoveIndex_Int_Invalid_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndex(data, -1)
	}
}
