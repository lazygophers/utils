package candy

import (
	"fmt"
	"math/rand"
	"testing"
)

// ==================== ToMap Benchmark ====================

func BenchmarkToMap_JSON_Small(b *testing.B) {
	data := []byte(`{"key1":"value1","key2":"value2"}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMap(data)
	}
}

func BenchmarkToMap_JSON_Medium(b *testing.B) {
	data := []byte(`{"key1":"value1","key2":"value2","key3":"value3","key4":"value4","key5":"value5"}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMap(data)
	}
}

func BenchmarkToMap_String_Small(b *testing.B) {
	data := `{"key1":"value1","key2":"value2"}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMap(data)
	}
}

func BenchmarkToMap_MapAny(b *testing.B) {
	m := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMap(m)
	}
}

func BenchmarkToMap_Nil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMap(nil)
	}
}

// ==================== ToMapInt32String Benchmark ====================

func BenchmarkToMapInt32String_Small(b *testing.B) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt32String(m)
	}
}

func BenchmarkToMapInt32String_Medium(b *testing.B) {
	m := make(map[int]string, 100)
	for i := 0; i < 100; i++ {
		m[i] = randString(5)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt32String(m)
	}
}

func BenchmarkToMapInt32String_Large(b *testing.B) {
	m := make(map[int]string, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = randString(10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt32String(m)
	}
}

func BenchmarkToMapInt32String_FromInt64(b *testing.B) {
	m := map[int64]string{1: "a", 2: "b", 3: "c"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt32String(m)
	}
}

func BenchmarkToMapInt32String_NonMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt32String("not a map")
	}
}

// ==================== ToMapInt64String Benchmark (扩展) ====================

func BenchmarkToMapInt64String_FromInt32(b *testing.B) {
	m := map[int32]string{1: "a", 2: "b", 3: "c"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt64String(m)
	}
}

func BenchmarkToMapInt64String_XL(b *testing.B) {
	m := make(map[int]string, 10000)
	for i := 0; i < 10000; i++ {
		m[i] = randString(10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt64String(m)
	}
}

func BenchmarkToMapInt64String_NonMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapInt64String("not a map")
	}
}

// ==================== ToMapStringAny Benchmark (扩展) ====================

func BenchmarkToMapStringAny_Nil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringAny(nil)
	}
}

func BenchmarkToMapStringAny_NonMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringAny("not a map")
	}
}

func BenchmarkToMapStringAny_XL(b *testing.B) {
	m := make(map[string]interface{}, 10000)
	for i := 0; i < 10000; i++ {
		m[randString(10)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringAny(m)
	}
}

// ==================== ToMapStringArrayString Benchmark ====================

func BenchmarkToMapStringArrayString_Small(b *testing.B) {
	m := map[string][]string{"key1": {"a", "b"}, "key2": {"c"}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringArrayString(m)
	}
}

func BenchmarkToMapStringArrayString_Medium(b *testing.B) {
	m := make(map[string][]string, 50)
	for i := 0; i < 50; i++ {
		m[randString(5)] = []string{randString(3), randString(3)}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringArrayString(m)
	}
}

func BenchmarkToMapStringArrayString_Large(b *testing.B) {
	m := make(map[string][]string, 500)
	for i := 0; i < 500; i++ {
		arr := make([]string, 5)
		for j := 0; j < 5; j++ {
			arr[j] = randString(3)
		}
		m[randString(10)] = arr
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringArrayString(m)
	}
}

func BenchmarkToMapStringArrayString_Nil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringArrayString(nil)
	}
}

func BenchmarkToMapStringArrayString_XL(b *testing.B) {
	m := make(map[string][]string, 5000)
	for i := 0; i < 5000; i++ {
		arr := make([]string, 3)
		for j := 0; j < 3; j++ {
			arr[j] = randString(5)
		}
		m[randString(10)] = arr
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringArrayString(m)
	}
}

// ==================== ToMapStringInt64 Benchmark (扩展) ====================

func BenchmarkToMapStringInt64_FromInt32(b *testing.B) {
	m := map[string]int32{"key1": 1, "key2": 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64(m)
	}
}

func BenchmarkToMapStringInt64_XL(b *testing.B) {
	m := make(map[string]int, 10000)
	for i := 0; i < 10000; i++ {
		m[randString(10)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64(m)
	}
}

func BenchmarkToMapStringInt64_NonMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64("not a map")
	}
}

func BenchmarkToMapStringInt64_Huge(b *testing.B) {
	m := make(map[string]int, 100000)
	for i := 0; i < 100000; i++ {
		m[randString(10)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringInt64(m)
	}
}

// ==================== ToMapStringString Benchmark (扩展) ====================

func BenchmarkToMapStringString_NonMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringString("not a map")
	}
}

func BenchmarkToMapStringString_XL(b *testing.B) {
	m := make(map[string]string, 10000)
	for i := 0; i < 10000; i++ {
		m[randString(10)] = randString(10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringString(m)
	}
}

func BenchmarkToMapStringString_Huge(b *testing.B) {
	m := make(map[string]string, 100000)
	for i := 0; i < 100000; i++ {
		m[randString(10)] = randString(10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToMapStringString(m)
	}
}

// 辅助函数
func randString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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
