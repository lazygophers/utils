package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSqrt 测试 Sqrt 函数
func TestSqrt(t *testing.T) {
	tests := []struct {
		name string
		give interface{}
		want interface{}
	}{
		// 整数类型测试
		{
			name: "完全平方数",
			give: int(16),
			want: int(4),
		},
		{
			name: "非完全平方数",
			give: int(2),
			want: int(1),
		},
		{
			name: "零的平方根",
			give: int(0),
			want: int(0),
		},
		{
			name: "一的平方根",
			give: int(1),
			want: int(1),
		},
		// float32 类型测试
		{
			name: "float32 完全平方数",
			give: float32(16.0),
			want: float32(4.0),
		},
		{
			name: "float32 非完全平方数",
			give: float32(2.0),
			want: float32(1.4142135),
		},
		{
			name: "float32 小数平方根",
			give: float32(0.25),
			want: float32(0.5),
		},
		{
			name: "float32 零",
			give: float32(0.0),
			want: float32(0.0),
		},
		// float64 类型测试
		{
			name: "float64 完全平方数",
			give: float64(144.0),
			want: float64(12.0),
		},
		{
			name: "float64 非完全平方数",
			give: float64(3.0),
			want: float64(1.7320508075688772),
		},
		{
			name: "float64 小数平方根",
			give: float64(0.81),
			want: float64(0.9),
		},
		{
			name: "float64 零",
			give: float64(0.0),
			want: float64(0.0),
		},
		// 边界情况
		{
			name: "极大数的平方根",
			give: float64(1e20),
			want: float64(1e10),
		},
		{
			name: "极小正数的平方根",
			give: float64(1e-20),
			want: float64(1e-10),
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			switch give := tt.give.(type) {
			case int:
				got := Sqrt(give)
				assert.Equal(t, tt.want, got, "Sqrt() 的结果应与期望值相等")
			case float32:
				got := Sqrt(give)
				assert.InDelta(t, tt.want, got, 1e-6, "Sqrt() 的结果应在允许误差范围内")
			case float64:
				got := Sqrt(give)
				assert.InDelta(t, tt.want, got, 1e-12, "Sqrt() 的结果应在允许误差范围内")
			}
		})
	}
}
