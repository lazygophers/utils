package candy

import (
	"testing"
)

// 简单的性能对比测试
func BenchmarkMapKeys_EmptyMap_Optimized(b *testing.B) {
	var m map[int]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeys(m)
	}
}

func BenchmarkMapKeysInt_TypeAssertFastPath(b *testing.B) {
	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		m[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeysInt(m)
	}
}

func BenchmarkMapKeysInt_ReflectPath(b *testing.B) {
	m := make(map[string]int)
	for i := 0; i < 100; i++ {
		m[string(rune('a'+i))] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapKeysInt(m)
	}
}

func BenchmarkMapValues_EmptyMap_Optimized(b *testing.B) {
	var m map[int]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapValues(m)
	}
}
