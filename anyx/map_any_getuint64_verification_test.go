package anyx

import (
	"strconv"
	"testing"
)

// BenchmarkVerifyGetUint64Performance 验证 GetUint64 性能优化的基准测试
func BenchmarkVerifyGetUint64Performance(b *testing.B) {
	// 准备测试数据
	m := make(map[string]interface{}, 100)
	for i := 0; i < 100; i++ {
		m["key_"+strconv.Itoa(i)] = uint64(i)
	}
	mm := NewMap(m)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mm.GetUint64("key_50")
	}
}
