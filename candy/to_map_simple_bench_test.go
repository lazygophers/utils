package candy

import (
	"fmt"
	"testing"
)

// ========== 测试数据 ==========

func generateStringSlice(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = fmt.Sprintf("item%d", i)
	}
	return result
}

func generateIntSliceForToMap(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

func generateMapStringAny(n int) map[string]interface{} {
	result := make(map[string]interface{}, n)
	for i := 0; i < n; i++ {
		result[fmt.Sprintf("key%d", i)] = i
	}
	return result
}

func generateMapStringString(n int) map[string]string {
	result := make(map[string]string, n)
	for i := 0; i < n; i++ {
		result[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}
	return result
}

var (
	smallStringSlice  = generateStringSlice(10)
	mediumStringSlice = generateStringSlice(100)
	largeStringSlice  = generateStringSlice(1000)

	smallIntSlice  = generateIntSliceForToMap(10)
	mediumIntSlice = generateIntSliceForToMap(100)
	largeIntSlice  = generateIntSliceForToMap(1000)

	smallMapStringAny  = generateMapStringAny(10)
	mediumMapStringAny = generateMapStringAny(100)
	largeMapStringAny  = generateMapStringAny(1000)

	smallMapStringStr  = generateMapStringString(10)
	mediumMapStringStr = generateMapStringString(100)
	largeMapStringStr  = generateMapStringString(1000)
)

// ========== Slice2Map 基准测试 ==========

func BenchmarkSlice2Map_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice2Map(smallStringSlice)
	}
}

func BenchmarkSlice2Map_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice2Map(mediumStringSlice)
	}
}

func BenchmarkSlice2Map_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice2Map(largeStringSlice)
	}
}

func BenchmarkSlice2Map_Int_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice2Map(smallIntSlice)
	}
}

func BenchmarkSlice2Map_Int_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice2Map(mediumIntSlice)
	}
}

func BenchmarkSlice2Map_Int_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Slice2Map(largeIntSlice)
	}
}

// ========== ToMapStringAny 基准测试 ==========

func BenchmarkToMapStringAny_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToMapStringAny(smallMapStringAny)
	}
}

func BenchmarkToMapStringAny_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToMapStringAny(mediumMapStringAny)
	}
}

func BenchmarkToMapStringAny_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToMapStringAny(largeMapStringAny)
	}
}

// ========== ToMapStringString 基准测试 ==========

func BenchmarkToMapStringString_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToMapStringString(smallMapStringStr)
	}
}

func BenchmarkToMapStringString_Medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToMapStringString(mediumMapStringStr)
	}
}

func BenchmarkToMapStringString_Large(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToMapStringString(largeMapStringStr)
	}
}

// ========== ToMapInt64String 基准测试 ==========

func BenchmarkToMapInt64String_Small(b *testing.B) {
	testMap := make(map[int64]string, 10)
	for i := 0; i < 10; i++ {
		testMap[int64(i)] = fmt.Sprintf("value%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt64String(testMap)
	}
}

func BenchmarkToMapInt64String_Medium(b *testing.B) {
	testMap := make(map[int64]string, 100)
	for i := 0; i < 100; i++ {
		testMap[int64(i)] = fmt.Sprintf("value%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt64String(testMap)
	}
}

func BenchmarkToMapInt64String_Large(b *testing.B) {
	testMap := make(map[int64]string, 1000)
	for i := 0; i < 1000; i++ {
		testMap[int64(i)] = fmt.Sprintf("value%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt64String(testMap)
	}
}

// ========== ToMapStringInt64 基准测试 ==========

func BenchmarkToMapStringInt64_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64(smallMapStringAny)
	}
}

func BenchmarkToMapStringInt64_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64(mediumMapStringAny)
	}
}

func BenchmarkToMapStringInt64_Large(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64(largeMapStringAny)
	}
}
