package anyx

import (
	"testing"
)

func TestGetInt64Simple(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"test": int64(123),
	})
	if v := m.GetInt64("test"); v != 123 {
		t.Fatalf("expected 123, got %d", v)
	}
}

func BenchmarkGetInt64Original(b *testing.B) {
	m := NewMap(map[string]interface{}{
		"int64": int64(123456789),
		"int":   int(123456789),
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetInt64("int64")
		_ = m.GetInt64("int")
	}
}
