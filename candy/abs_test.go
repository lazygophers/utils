// Package candy 提供便捷的语法糖函数
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAbs 测试 Abs 函数
func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		give interface{}
		want interface{}
	}{
		// 整数类型测试
		{
			name: "正整数绝对值等于自身",
			give: int(42),
			want: int(42),
		},
		{
			name: "负整数绝对值等于其相反数",
			give: int(-42),
			want: int(42),
		},
		{
			name: "零的绝对值等于零",
			give: int(0),
			want: int(0),
		},
		{
			name: "最大负整数的绝对值",
			give: int(-2147483648),
			want: int(2147483648),
		},
		{
			name: "最大正整数的绝对值",
			give: int(2147483647),
			want: int(2147483647),
		},
		// int8 类型测试
		{
			name: "int8 正数",
			give: int8(127),
			want: int8(127),
		},
		{
			name: "int8 负数",
			give: int8(-128),
			want: int8(-128),
		},
		// int16 类型测试
		{
			name: "int16 正数",
			give: int16(32767),
			want: int16(32767),
		},
		{
			name: "int16 负数",
			give: int16(-32768),
			want: int16(-32768),
		},
		// int32 类型测试
		{
			name: "int32 正数",
			give: int32(2147483647),
			want: int32(2147483647),
		},
		{
			name: "int32 负数",
			give: int32(-2147483648),
			want: int32(-2147483648),
		},
		// int64 类型测试
		{
			name: "int64 正数",
			give: int64(9223372036854775807),
			want: int64(9223372036854775807),
		},
		{
			name: "int64 负数",
			give: int64(-9223372036854775808),
			want: int64(-9223372036854775808),
		},
		// uint 类型测试
		{
			name: "uint 正数",
			give: uint(42),
			want: uint(42),
		},
		{
			name: "uint 零",
			give: uint(0),
			want: uint(0),
		},
		// uint8 类型测试
		{
			name: "uint8 最大值",
			give: uint8(255),
			want: uint8(255),
		},
		// float32 类型测试
		{
			name: "float32 正数",
			give: float32(3.14),
			want: float32(3.14),
		},
		{
			name: "float32 负数",
			give: float32(-3.14),
			want: float32(3.14),
		},
		{
			name: "float32 零",
			give: float32(0.0),
			want: float32(0.0),
		},
		{
			name: "float32 极小正数",
			give: float32(1.1754944e-38),
			want: float32(1.1754944e-38),
		},
		{
			name: "float32 极大正数",
			give: float32(3.4028235e38),
			want: float32(3.4028235e38),
		},
		// float64 类型测试
		{
			name: "float64 正数",
			give: float64(3.141592653589793),
			want: float64(3.141592653589793),
		},
		{
			name: "float64 负数",
			give: float64(-3.141592653589793),
			want: float64(3.141592653589793),
		},
		{
			name: "float64 零",
			give: float64(0.0),
			want: float64(0.0),
		},
		{
			name: "float64 极小正数",
			give: float64(2.2250738585072014e-308),
			want: float64(2.2250738585072014e-308),
		},
		{
			name: "float64 极大正数",
			give: float64(1.7976931348623157e308),
			want: float64(1.7976931348623157e308),
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			switch give := tt.give.(type) {
			case int:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case int8:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case int16:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case int32:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case int64:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case uint:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case uint8:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case float32:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			case float64:
				got := Abs(give)
				assert.Equal(t, tt.want, got, "Abs() 的结果应与期望值相等")
			}
		})
	}
}
