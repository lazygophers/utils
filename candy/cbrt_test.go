package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCbrt 测试 Cbrt 函数
func TestCbrt(t *testing.T) {
	tests := []struct {
		name string
		give interface{}
		want interface{}
	}{
		// 整数类型测试
		{
			name: "完全立方数",
			give: int(27),
			want: int(3),
		},
		{
			name: "负的完全立方数",
			give: int(-8),
			want: int(-2),
		},
		{
			name: "零的立方根",
			give: int(0),
			want: int(0),
		},
		{
			name: "一的立方根",
			give: int(1),
			want: int(1),
		},
		// float32 类型测试
		{
			name: "float32 完全立方数",
			give: float32(8.0),
			want: float32(2.0),
		},
		{
			name: "float32 负的完全立方数",
			give: float32(-27.0),
			want: float32(-3.0),
		},
		{
			name: "float32 非完全立方数",
			give: float32(10.0),
			want: float32(2.1544347),
		},
		{
			name: "float32 零",
			give: float32(0.0),
			want: float32(0.0),
		},
		// float64 类型测试
		{
			name: "float64 完全立方数",
			give: float64(125.0),
			want: float64(5.0),
		},
		{
			name: "float64 负的完全立方数",
			give: float64(-64.0),
			want: float64(-4.0),
		},
		{
			name: "float64 非完全立方数",
			give: float64(20.0),
			want: float64(2.7144176),
		},
		{
			name: "float64 零",
			give: float64(0.0),
			want: float64(0.0),
		},
		// 边界情况
		{
			name: "极大数的立方根",
			give: float64(1e30),
			want: float64(1e10),
		},
		{
			name: "极小正数的立方根",
			give: float64(1e-30),
			want: float64(1e-10),
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			switch give := tt.give.(type) {
			case int:
				got := Cbrt(give)
				assert.Equal(t, tt.want, got, "Cbrt() 的结果应与期望值相等")
			case float32:
				got := Cbrt(give)
				assert.InDelta(t, tt.want, got, 1e-6, "Cbrt() 的结果应在允许误差范围内")
			case float64:
				got := Cbrt(give)
				assert.InDelta(t, tt.want, got, 1e-7, "Cbrt() 的结果应在允许误差范围内")
			}
		})
	}
}