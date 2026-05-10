package anyx

import (
	"testing"
)

func TestGetIntSimple(t *testing.T) {
	m := NewMap(map[string]interface{}{"key": 42})
	if got := m.GetInt("key"); got != 42 {
		t.Errorf("GetInt() = %v, want %v", got, 42)
	}
}

func BenchmarkGetInt_Simple(b *testing.B) {
	m := NewMap(map[string]interface{}{"key": 42})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.GetInt("key")
	}
}
