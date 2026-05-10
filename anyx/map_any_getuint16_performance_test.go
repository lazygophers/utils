package anyx

import (
	"testing"
)

// BenchmarkGetUint16_Optimized 优化后的实现性能测试
func BenchmarkGetUint16_Optimized(b *testing.B) {
	m := NewMap(map[string]interface{}{
		"value": uint16(12345),
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetUint16("value")
	}
}

// BenchmarkGetUint16_Optimized_Miss 优化后未命中测试
func BenchmarkGetUint16_Optimized_Miss(b *testing.B) {
	m := NewMap(map[string]interface{}{
		"value": uint16(12345),
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetUint16("notexist")
	}
}

// BenchmarkGetUint16_Optimized_Nested 优化后嵌套路径测试
func BenchmarkGetUint16_Optimized_Nested(b *testing.B) {
	m := NewMap(map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": uint16(12345),
		},
	})
	m.EnableCut(".")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetUint16("level1.level2")
	}
}

// BenchmarkGetUint16_Optimized_TypeMismatch 优化后类型不匹配测试
func BenchmarkGetUint16_Optimized_TypeMismatch(b *testing.B) {
	m := NewMap(map[string]interface{}{
		"value": "12345",
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetUint16("value")
	}
}
