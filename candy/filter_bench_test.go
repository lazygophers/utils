package candy

import "testing"

func BenchmarkFilter_Small_50Percent(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Medium_50Percent(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Large_50Percent(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Medium_10Percent(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n < 100 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Medium_90Percent(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n < 900 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}
