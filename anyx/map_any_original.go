package anyx

import (
	"testing"

	"github.com/lazygophers/utils/candy"
)

// GetUint16Original 保留原始实现用于性能对比
func (p *MapAny) GetUint16Original(key string) uint16 {
	val, ok := p.get(key)
	if !ok {
		return 0
	}

	return candy.ToUint16(val)
}

// BenchmarkGetUint16_OriginalFunc 原始函数实现的基准测试
func BenchmarkGetUint16_OriginalFunc(b *testing.B) {
	m := NewMap(map[string]interface{}{
		"value": uint16(12345),
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetUint16Original("value")
	}
}
