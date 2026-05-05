package candy

import (
	"reflect"
	"testing"
)

// 性能对比基准测试

func BenchmarkDeepEqual_Int(b *testing.B) {
	x, y := 42, 42
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, y)
	}
}

func BenchmarkReflectDeepEqual_Int(b *testing.B) {
	x, y := 42, 42
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, y)
	}
}

func BenchmarkDeepEqual_Slice(b *testing.B) {
	x, y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, y)
	}
}

func BenchmarkReflectDeepEqual_Slice(b *testing.B) {
	x, y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, y)
	}
}

func BenchmarkDeepEqual_Map(b *testing.B) {
	x, y := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, y)
	}
}

func BenchmarkReflectDeepEqual_Map(b *testing.B) {
	x, y := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, y)
	}
}

func BenchmarkDeepEqual_LargeSlice(b *testing.B) {
	x, y := make([]int, 1000), make([]int, 1000)
	for i := range x {
		x[i], y[i] = i, i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, y)
	}
}

func BenchmarkReflectDeepEqual_LargeSlice(b *testing.B) {
	x, y := make([]int, 1000), make([]int, 1000)
	for i := range x {
		x[i], y[i] = i, i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, y)
	}
}

func BenchmarkDeepEqual_LargeMap(b *testing.B) {
	x, y := make(map[string]int, 100), make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := string(rune('a' + i))
		x[key], y[key] = i, i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, y)
	}
}

func BenchmarkReflectDeepEqual_LargeMap(b *testing.B) {
	x, y := make(map[string]int, 100), make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := string(rune('a' + i))
		x[key], y[key] = i, i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, y)
	}
}

func BenchmarkDeepEqual_SameSliceRef(b *testing.B) {
	x := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, x)
	}
}

func BenchmarkReflectDeepEqual_SameSliceRef(b *testing.B) {
	x := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, x)
	}
}

func BenchmarkDeepEqual_SameMapRef(b *testing.B) {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepEqual(x, x)
	}
}

func BenchmarkReflectDeepEqual_SameMapRef(b *testing.B) {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, x)
	}
}
