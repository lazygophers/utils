package candy

import (
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
