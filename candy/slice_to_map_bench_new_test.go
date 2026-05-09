package candy

import (
	"testing"

	"golang.org/x/exp/constraints"
)

// Benchmark Slice2Map 函数的各种实现方案

// 方案1: 当前实现（使用 range）
func BenchmarkSlice2Map_Current_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

// 方案2: 使用索引循环
func Slice2MapIndex[M constraints.Ordered](list []M) map[M]bool {
	if len(list) == 0 {
		return make(map[M]bool)
	}

	result := make(map[M]bool, len(list))
	for i := 0; i < len(list); i++ {
		result[list[i]] = true
	}
	return result
}

func BenchmarkSlice2Map_Index_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

// 方案3: 不预分配容量
func Slice2MapNoPrealloc[M constraints.Ordered](list []M) map[M]bool {
	if len(list) == 0 {
		return make(map[M]bool)
	}

	result := make(map[M]bool)
	for _, item := range list {
		result[item] = true
	}
	return result
}

func BenchmarkSlice2Map_NoPrealloc_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapNoPrealloc(data)
	}
}

// 方案4: 使用 struct{} 作为值（节省内存）
func Slice2MapStruct[M constraints.Ordered](list []M) map[M]struct{} {
	if len(list) == 0 {
		return make(map[M]struct{})
	}

	result := make(map[M]struct{}, len(list))
	for i := 0; i < len(list); i++ {
		result[list[i]] = struct{}{}
	}
	return result
}

func BenchmarkSlice2Map_Struct_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

// 方案5: 检查是否已存在再设置
func Slice2MapCheck[M constraints.Ordered](list []M) map[M]bool {
	if len(list) == 0 {
		return make(map[M]bool)
	}

	result := make(map[M]bool, len(list))
	for i := 0; i < len(list); i++ {
		if !result[list[i]] {
			result[list[i]] = true
		}
	}
	return result
}

func BenchmarkSlice2Map_Check_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}

// 中等数据集测试
func BenchmarkSlice2Map_Current_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

func BenchmarkSlice2Map_Index_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

func BenchmarkSlice2Map_NoPrealloc_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapNoPrealloc(data)
	}
}

func BenchmarkSlice2Map_Struct_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

func BenchmarkSlice2Map_Check_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}

// 大数据集测试
func BenchmarkSlice2Map_Current_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

func BenchmarkSlice2Map_Index_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

func BenchmarkSlice2Map_NoPrealloc_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapNoPrealloc(data)
	}
}

func BenchmarkSlice2Map_Struct_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

func BenchmarkSlice2Map_Check_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}

// 重复数据测试
func BenchmarkSlice2Map_Current_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10 // 创建重复数据
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

func BenchmarkSlice2Map_Index_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

func BenchmarkSlice2Map_Struct_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

func BenchmarkSlice2Map_Check_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}
