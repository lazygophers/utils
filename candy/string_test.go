package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestString 测试String转换函数
// 该函数测试泛型String函数对各种有序类型的转换能力
// 包括整数、浮点数等基本类型的字符串转换
func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		give int
		want string
	}{
		{"正整数", 42, "42"},
		{"负整数", -42, "-42"},
		{"零", 0, "0"},
		{"大整数", 999999999, "999999999"},
		{"浮点零", 0.0, "0"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := String(tt.give)
			assert.Equal(t, tt.want, got, "String() 的结果应与期望值相等")
		})
	}
}
