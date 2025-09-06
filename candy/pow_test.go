package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPow 测试 Pow 函数
func TestPow(t *testing.T) {
	tests := []struct {
		name string
		x    interface{}
		y    interface{}
		want interface{}
	}{
		// 整数类型测试
		{
			name: "整数正幂",
			x:    int(2),
			y:    int(3),
			want: int(8),
		},
		{
			name: "整数零幂",
			x:    int(5),
			y:    int(0),
			want: int(1),
		},
		{
			name: "整数负幂",
			x:    int(2),
			y:    int(-2),
			want: int(0),
		},
		{
			name: "零的整数幂",
			x:    int(0),
			y:    int(5),
			want: int(0),
		},
		{
			name: "一的整数幂",
			x:    int(1),
			y:    int(100),
			want: int(1),
		},
		// float32 类型测试
		{
			name: "float32 正幂",
			x:    float32(2.0),
			y:    float32(3.0),
			want: float32(8.0),
		},
		{
			name: "float32 小数幂",
			x:    float32(4.0),
			y:    float32(0.5),
			want: float32(2.0),
		},
		{
			name: "float32 负幂",
			x:    float32(2.0),
			y:    float32(-2.0),
			want: float32(0.25),
		},
		{
			name: "float32 零幂",
			x:    float32(3.14),
			y:    float32(0.0),
			want: float32(1.0),
		},
		// float64 类型测试
		{
			name: "float64 正幂",
			x:    float64(2.0),
			y:    float64(10.0),
			want: float64(1024.0),
		},
		{
			name: "float64 小数幂",
			x:    float64(9.0),
			y:    float64(0.5),
			want: float64(3.0),
		},
		{
			name: "float64 负幂",
			x:    float64(2.0),
			y:    float64(-3.0),
			want: float64(0.125),
		},
		{
			name: "float64 零幂",
			x:    float64(2.718),
			y:    float64(0.0),
			want: float64(1.0),
		},
		// 边界情况
		{
			name: "零的零幂",
			x:    float64(0.0),
			y:    float64(0.0),
			want: float64(1.0),
		},
		{
			name: "负数的整数幂",
			x:    float64(-2.0),
			y:    float64(3.0),
			want: float64(-8.0),
		},
		{
			name: "负数的偶数幂",
			x:    float64(-2.0),
			y:    float64(4.0),
			want: float64(16.0),
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			switch x := tt.x.(type) {
			case int:
				y := tt.y.(int)
				got := Pow(x, y)
				assert.Equal(t, tt.want, got, "Pow() 的结果应与期望值相等")
			case float32:
				y := tt.y.(float32)
				got := Pow(x, y)
				assert.InDelta(t, tt.want, got, 1e-6, "Pow() 的结果应在允许误差范围内")
			case float64:
				y := tt.y.(float64)
				got := Pow(x, y)
				assert.InDelta(t, tt.want, got, 1e-10, "Pow() 的结果应在允许误差范围内")
			}
		})
	}
}
