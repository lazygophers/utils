package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestToArrayString 测试 ToArrayString 函数
func TestToArrayString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{"string input", "hello", []string{"hello"}},
		{"[]string input", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"empty string", "", []string{""}},
		{"empty []string", []string{}, []string{}},
		{"string with comma", "a,b,c", []string{"a", "b", "c"}},
		{"string with separator", "a|b|c", []string{"a|b|c"}},
		{"single element []string", []string{"single"}, []string{"single"}},
		{"unicode string", "你好", []string{"你好"}},
		{"unicode []string", []string{"你好", "世界"}, []string{"你好", "世界"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ToArrayString(tt.input)
			assert.Equal(t, tt.expected, result, "ToArrayString() 的结果应与期望值相等")
		})
	}
}