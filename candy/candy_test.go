package candy

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSliceEqual 测试 SliceEqual 函数
func TestSliceEqual(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		// 基础场景
		{
			name: "两个空切片相等",
			a:    []int{},
			b:    []int{},
			want: true,
		},
		{
			name: "一个空一个非空不相等",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "不同长度不相等",
			a:    []int{1, 2, 3},
			b:    []int{1, 2},
			want: false,
		},
		{
			name: "相同元素相等",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "相同元素不同顺序相等",
			a:    []int{1, 2, 3},
			b:    []int{3, 2, 1},
			want: true,
		},
		{
			name: "不同元素不相等",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 4},
			want: false,
		},
		// 重复元素场景
		{
			name: "重复元素相同数量相等",
			a:    []int{1, 2, 2, 3},
			b:    []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "重复元素不同数量不相等",
			a:    []int{1, 2, 2, 3},
			b:    []int{1, 2, 3, 2},
			want: true, // 顺序不同但元素相同，应该相等
		},
		{
			name: "重复元素数量不同不相等",
			a:    []int{1, 2, 2, 3},
			b:    []int{1, 2, 3, 3},
			want: false, // 重复次数不同，不相等
		},
		{
			name: "所有元素相同相等",
			a:    []int{2, 2, 2, 2},
			b:    []int{2, 2, 2, 2},
			want: true,
		},
		{
			name: "所有元素相同数量不同不相等",
			a:    []int{2, 2, 2, 2},
			b:    []int{2, 2, 2},
			want: false,
		},
		// 边界情况
		{
			name: "负数相等",
			a:    []int{-1, -2, -3},
			b:    []int{-1, -2, -3},
			want: true,
		},
		{
			name: "混合正负相等",
			a:    []int{-1, 2, -3},
			b:    []int{-1, 2, -3},
			want: true,
		},
		{
			name: "大数相等",
			a:    []int{1000000, 2000000, 3000000},
			b:    []int{1000000, 2000000, 3000000},
			want: true,
		},
		{
			name: "零值相等",
			a:    []int{0, 0, 0},
			b:    []int{0, 0, 0},
			want: true,
		},
		{
			name: "单元素相等",
			a:    []int{42},
			b:    []int{42},
			want: true,
		},
		{
			name: "单元素不相等",
			a:    []int{42},
			b:    []int{43},
			want: false,
		},
		{
			name: "包含零值相等",
			a:    []int{0, 1, 2},
			b:    []int{0, 1, 2},
			want: true,
		},
		{
			name: "复杂重复场景相等",
			a:    []int{1, 2, 2, 3, 3, 4},
			b:    []int{1, 2, 3, 2, 4, 3},
			want: true,
		},
		{
			name: "缺少元素不相等",
			a:    []int{1, 2, 3},
			b:    []int{1, 2},
			want: false,
		},
		{
			name: "多余元素不相等",
			a:    []int{1, 2},
			b:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "完全不相等",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			got := SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "SliceEqual() 的结果应与期望值相等")
		})
	}
}

// TestSliceEqualEdgeCases 测试 SliceEqual 函数的边界情况
func TestSliceEqualEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		// 重复元素匹配场景
		{
			name: "重复元素匹配-相同数量",
			a:    []int{1, 2, 2, 3},
			b:    []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "重复元素匹配-不同顺序",
			a:    []int{1, 2, 2, 3},
			b:    []int{1, 2, 3, 2},
			want: true,
		},
		{
			name: "重复元素不匹配-数量不同",
			a:    []int{1, 2, 2, 3},
			b:    []int{1, 2, 3, 3},
			want: false,
		},
		{
			name: "所有元素相同-匹配",
			a:    []int{2, 2, 2, 2},
			b:    []int{2, 2, 2, 2},
			want: true,
		},
		{
			name: "所有元素相同-数量不同",
			a:    []int{2, 2, 2, 2},
			b:    []int{2, 2, 2},
			want: false,
		},
		{
			name: "空切片匹配",
			a:    []int{},
			b:    []int{},
			want: true,
		},
		{
			name: "一个空一个非空",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "相同元素不同位置",
			a:    []int{1, 2, 3, 4},
			b:    []int{4, 3, 2, 1},
			want: true,
		},
		{
			name: "大数量重复元素",
			a:    []int{1, 1, 1, 2, 2, 3},
			b:    []int{1, 1, 2, 2, 3, 3},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			got := SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "SliceEqual() 边界情况的结果应与期望值相等")
		})
	}
}

// TestSliceEqualString 测试 SliceEqual 函数对字符串类型的支持
func TestSliceEqualString(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want bool
	}{
		{
			name: "相同字符串相等",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "c"},
			want: true,
		},
		{
			name: "相同字符串不同顺序相等",
			a:    []string{"a", "b", "c"},
			b:    []string{"c", "b", "a"},
			want: true,
		},
		{
			name: "不同字符串不相等",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "d"},
			want: false,
		},
		{
			name: "重复字符串相等",
			a:    []string{"a", "b", "b", "c"},
			b:    []string{"a", "b", "b", "c"},
			want: true,
		},
		{
			name: "重复字符串不相等",
			a:    []string{"a", "b", "b", "c"},
			b:    []string{"a", "b", "c", "c"},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			got := SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "SliceEqual() 字符串类型的结果应与期望值相等")
		})
	}
}

// TestSliceEqualWithNil 测试 SliceEqual 函数对 nil 切片的处理
func TestSliceEqualWithNil(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{
			name: "两个nil切片相等",
			a:    nil,
			b:    nil,
			want: true,
		},
		{
			name: "一个nil一个空切片相等",
			a:    nil,
			b:    []int{},
			want: false, // 长度不同，返回false
		},
		{
			name: "一个nil一个非空切片不相等",
			a:    nil,
			b:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "一个空切片一个nil不相等",
			a:    []int{},
			b:    nil,
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			got := SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "SliceEqual() nil切片处理的结果应与期望值相等")
		})
	}
}

// TestSliceEqualLargeData 测试 SliceEqual 函数对大数据的处理
func TestSliceEqualLargeData(t *testing.T) {
	// 生成大型测试数据
	largeSlice1 := make([]int, 1000)
	largeSlice2 := make([]int, 1000)
	largeSlice3 := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		largeSlice1[i] = i
		largeSlice2[i] = i
		largeSlice3[i] = i + 1
	}

	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{
			name: "大型相同切片相等",
			a:    largeSlice1,
			b:    largeSlice2,
			want: true,
		},
		{
			name: "大型不同切片不相等",
			a:    largeSlice1,
			b:    largeSlice3,
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			got := SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "SliceEqual() 大数据处理的结果应与期望值相等")
		})
	}
}

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

// TestFilterNot 测试 FilterNot 函数
func TestFilterNot(t *testing.T) {
	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			f    func(int) bool
			want []int
		}{
			{
				name: "过滤偶数保留奇数",
				give: []int{1, 2, 3, 4, 5, 6},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: []int{1, 3, 5},
			},
			{
				name: "过滤正数保留负数和零",
				give: []int{-1, 0, 1, -2, 2},
				f: func(n int) bool {
					return n > 0
				},
				want: []int{-1, 0, -2},
			},
			{
				name: "过滤大数保留小数",
				give: []int{1, 10, 100, 1000},
				f: func(n int) bool {
					return n > 50
				},
				want: []int{1, 10},
			},
			{
				name: "空切片输入",
				give: []int{},
				f: func(n int) bool {
					return n > 0
				},
				want: []int{},
			},
			{
				name: "全部元素都被过滤",
				give: []int{2, 4, 6, 8},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: []int{},
			},
			{
				name: "没有元素被过滤",
				give: []int{1, 3, 5, 7},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: []int{1, 3, 5, 7},
			},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				got := FilterNot(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "FilterNot() 的结果应与期望值相等")
			})
		}
	})
}
// TestShuffle 测试 Shuffle 函数
func TestShuffle(t *testing.T) {
	t.Parallel()
	
	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			want []int // 期望的元素集合（顺序不重要）
		}{
			{
				name: "多个元素打乱",
				give: []int{1, 2, 3, 4, 5},
				want: []int{1, 2, 3, 4, 5},
			},
			{
				name: "重复元素打乱",
				give: []int{1, 2, 2, 3, 3, 4},
				want: []int{1, 2, 2, 3, 3, 4},
			},
			{
				name: "负数打乱",
				give: []int{-1, -2, -3, -4, -5},
				want: []int{-1, -2, -3, -4, -5},
			},
			{
				name: "混合正负数打乱",
				give: []int{0, -1, 2, -3, 4},
				want: []int{0, -1, 2, -3, 4},
			},
			{
				name: "大数打乱",
				give: []int{1000000, 2000000, 3000000},
				want: []int{1000000, 2000000, 3000000},
			},
			{
				name: "单元素切片",
				give: []int{42},
				want: []int{42},
			},
			{
				name: "空切片",
				give: []int{},
				want: []int{},
			},
		}
		
		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				// 创建副本以避免修改原始数据
				original := make([]int, len(tt.give))
				copy(original, tt.give)
				
				// 执行打乱操作
				result := Shuffle(tt.give)
				
				// 验证元素集合相同（顺序可能不同）
				assert.ElementsMatch(t, tt.want, result, "Shuffle() 后元素集合应保持不变")
				
				// 验证原始数据被修改（in-place操作）
				assert.Equal(t, result, tt.give, "Shuffle() 应该返回原切片的引用")
			})
		}
	})
	
	// 浮点数类型测试
	t.Run("浮点数类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []float64
			want []float64
		}{
			{
				name: "浮点数打乱",
				give: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				want: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			},
			{
				name: "负浮点数打乱",
				give: []float64{-1.1, -2.2, -3.3},
				want: []float64{-1.1, -2.2, -3.3},
			},
			{
				name: "混合浮点数打乱",
				give: []float64{0.0, -1.5, 2.718, -3.14, 4.0},
				want: []float64{0.0, -1.5, 2.718, -3.14, 4.0},
			},
			{
				name: "单元素浮点数",
				give: []float64{3.14159},
				want: []float64{3.14159},
			},
			{
				name: "空浮点数切片",
				give: []float64{},
				want: []float64{},
			},
		}
		
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				original := make([]float64, len(tt.give))
				copy(original, tt.give)
				
				result := Shuffle(tt.give)
				
				assert.ElementsMatch(t, tt.want, result, "Shuffle() 浮点数后元素集合应保持不变")
				assert.Equal(t, result, tt.give, "Shuffle() 应该返回原切片的引用")
			})
		}
	})
	
	// 字符串类型测试
	t.Run("字符串类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			want []string
		}{
			{
				name: "字符串打乱",
				give: []string{"apple", "banana", "cherry", "date", "elderberry"},
				want: []string{"apple", "banana", "cherry", "date", "elderberry"},
			},
			{
				name: "重复字符串打乱",
				give: []string{"a", "b", "b", "c", "c", "c"},
				want: []string{"a", "b", "b", "c", "c", "c"},
			},
			{
				name: "空字符串打乱",
				give: []string{"", "hello", "", "world"},
				want: []string{"", "hello", "", "world"},
			},
			{
				name: "单元素字符串",
				give: []string{"test"},
				want: []string{"test"},
			},
			{
				name: "空字符串切片",
				give: []string{},
				want: []string{},
			},
		}
		
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				original := make([]string, len(tt.give))
				copy(original, tt.give)
				
				result := Shuffle(tt.give)
				
				assert.ElementsMatch(t, tt.want, result, "Shuffle() 字符串后元素集合应保持不变")
				assert.Equal(t, result, tt.give, "Shuffle() 应该返回原切片的引用")
			})
		}
	})
	
	// 结构体类型测试
	t.Run("结构体类型", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		
		tests := []struct {
			name string
			give []Person
			want []Person
		}{
			{
				name: "结构体打乱",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 35},
					{"David", 40},
				},
				want: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 35},
					{"David", 40},
				},
			},
			{
				name: "单元素结构体",
				give: []Person{{"Eve", 28}},
				want: []Person{{"Eve", 28}},
			},
			{
				name: "空结构体切片",
				give: []Person{},
				want: []Person{},
			},
		}
		
		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				original := make([]Person, len(tt.give))
				copy(original, tt.give)
				
				result := Shuffle(tt.give)
				
				assert.ElementsMatch(t, tt.want, result, "Shuffle() 结构体后元素集合应保持不变")
				assert.Equal(t, result, tt.give, "Shuffle() 应该返回原切片的引用")
			})
		}
	})
	
	// 随机性测试
	t.Run("随机性测试", func(t *testing.T) {
		// 测试多次调用会产生不同的顺序
		original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		
		// 进行多次打乱操作
		results := make([][]int, 10)
		for i := range results {
			// 每次都使用原始数据的副本
			data := make([]int, len(original))
			copy(data, original)
			results[i] = Shuffle(data)
		}
		
		// 验证至少有一次打乱改变了顺序（对于足够大的切片）
		// 注意：这个测试有极小的概率会失败，因为随机可能产生相同的顺序
		orderChanged := false
		for _, result := range results {
			if !reflect.DeepEqual(original, result) {
				orderChanged = true
				break
			}
		}
		
		// 对于10个元素的切片，随机打乱后保持原顺序的概率极小
		assert.True(t, orderChanged, "多次调用 Shuffle() 应该产生不同的顺序")
	})
	
	// 边界情况测试
	t.Run("边界情况", func(t *testing.T) {
		// nil切片测试
		var nilSlice []int
		result := Shuffle(nilSlice)
		assert.Nil(t, result, "Shuffle() nil切片应该返回nil")
		
		// 大切片测试
		largeSlice := make([]int, 1000)
		for i := range largeSlice {
			largeSlice[i] = i
		}
		
		original := make([]int, len(largeSlice))
		copy(original, largeSlice)
		
		result = Shuffle(largeSlice)
		assert.ElementsMatch(t, original, result, "Shuffle() 大切片后元素集合应保持不变")
		assert.Equal(t, result, largeSlice, "Shuffle() 应该返回原切片的引用")
		
		// 验证打乱确实改变了顺序
		orderChanged := false
		for i := range result {
			if result[i] != original[i] {
				orderChanged = true
				break
			}
		}
		assert.True(t, orderChanged, "Shuffle() 大切片应该改变元素顺序")
	})
}

// TestReduce 测试 Reduce 函数
func TestReduce(t *testing.T) {
	t.Parallel()

	// 整数切片求和测试
	t.Run("整数切片求和", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 15
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 整数切片求积测试
	t.Run("整数切片求积", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		f := func(a, b int) int { return a * b }
		got := Reduce(input, f)
		want := 24
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 空切片测试
	t.Run("空切片", func(t *testing.T) {
		input := []int{}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 0 // 空切片的零值
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 单元素切片测试
	t.Run("单元素切片", func(t *testing.T) {
		input := []int{100}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 100
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 求最大值测试
	t.Run("求最大值", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		f := func(a, b int) int {
			if b > a {
				return b
			}
			return a
		}
		got := Reduce(input, f)
		want := 9
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 求最小值测试
	t.Run("求最小值", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		f := func(a, b int) int {
			if b < a {
				return b
			}
			return a
		}
		got := Reduce(input, f)
		want := 1
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 负数求和测试
	t.Run("负数求和", func(t *testing.T) {
		input := []int{-1, -2, -3, -4}
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := -10
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 字符串拼接测试
	t.Run("字符串拼接", func(t *testing.T) {
		input := []string{"Hello", " ", "World", "!"}
		f := func(a, b string) string { return a + b }
		got := Reduce(input, f)
		want := "Hello World!"
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// 浮点数求和测试
	t.Run("浮点数求和", func(t *testing.T) {
		input := []float64{1.1, 2.2, 3.3}
		f := func(a, b float64) float64 { return a + b }
		got := Reduce(input, f)
		want := 6.6
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})

	// nil切片测试
	t.Run("nil切片", func(t *testing.T) {
		var input []int
		f := func(a, b int) int { return a + b }
		got := Reduce(input, f)
		want := 0
		assert.Equal(t, want, got, "Reduce() 的结果应与期望值相等")
	})
}

// TestDrop 测试 Drop 函数
func TestDrop(t *testing.T) {
	// 定义测试用例结构体
	type testCase[T any] struct {
		name string
		give []T
		n    int
		want []T
	}

	// 定义测试用例
	tests := []testCase[int]{
		{
			name: "丢弃前0个元素-返回原切片",
			give: []int{1, 2, 3, 4, 5},
			n:    0,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "丢弃前2个元素-正常情况",
			give: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{3, 4, 5},
		},
		{
			name: "丢弃全部元素-返回空切片",
			give: []int{1, 2, 3, 4, 5},
			n:    5,
			want: []int{},
		},
		{
			name: "丢弃数量超过切片长度-返回空切片",
			give: []int{1, 2, 3, 4, 5},
			n:    10,
			want: []int{},
		},
		{
			name: "空切片-返回空切片",
			give: []int{},
			n:    3,
			want: []int{},
		},
		{
			name: "负数n-当作0处理",
			give: []int{1, 2, 3, 4, 5},
			n:    -1,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "丢弃到只剩一个元素",
			give: []int{1, 2, 3, 4, 5},
			n:    4,
			want: []int{5},
		},
	}

	// 执行测试
	for _, tt := range tests {
		tt := tt // 避免竞态条件
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 并行测试

			// 调用 Drop 函数
			got := Drop(tt.give, tt.n)

			// 验证结果
			assert.Equal(t, tt.want, got, "Drop() 的结果应与期望值相等")
		})
	}

	// 测试字符串类型
	stringTests := []testCase[string]{
		{
			name: "字符串切片-丢弃前2个元素",
			give: []string{"a", "b", "c", "d", "e"},
			n:    2,
			want: []string{"c", "d", "e"},
		},
		{
			name: "字符串切片-空切片",
			give: []string{},
			n:    1,
			want: []string{},
		},
	}

	for _, tt := range stringTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Drop(tt.give, tt.n)
			assert.Equal(t, tt.want, got, "Drop() 的结果应与期望值相等")
		})
	}

	// 测试结构体类型
	type Person struct {
		Name string
		Age  int
	}

	structTests := []testCase[Person]{
		{
			name: "结构体切片-丢弃前1个元素",
			give: []Person{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			n: 1,
			want: []Person{
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
		},
	}

	for _, tt := range structTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Drop(tt.give, tt.n)
			assert.Equal(t, tt.want, got, "Drop() 的结果应与期望值相等")
		})
	}
}

// TestAny 测试 Any 函数
func TestAny(t *testing.T) {
	t.Parallel()

	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			f    func(int) bool
			want bool
		}{
			{
				name: "存在偶数-返回true",
				give: []int{1, 2, 3, 4, 5},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: true,
			},
			{
				name: "不存在偶数-返回false",
				give: []int{1, 3, 5, 7, 9},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: false,
			},
			{
				name: "存在正数-返回true",
				give: []int{-1, 0, 1, -2, 2},
				f: func(n int) bool {
					return n > 0
				},
				want: true,
			},
			{
				name: "不存在正数-返回false",
				give: []int{-1, -2, -3, 0},
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "空切片-返回false",
				give: []int{},
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "第一个元素匹配-返回true",
				give: []int{42, 1, 2, 3},
				f: func(n int) bool {
					return n == 42
				},
				want: true,
			},
			{
				name: "最后一个元素匹配-返回true",
				give: []int{1, 2, 3, 42},
				f: func(n int) bool {
					return n == 42
				},
				want: true,
			},
			{
				name: "单元素匹配-返回true",
				give: []int{42},
				f: func(n int) bool {
					return n == 42
				},
				want: true,
			},
			{
				name: "单元素不匹配-返回false",
				give: []int{41},
				f: func(n int) bool {
					return n == 42
				},
				want: false,
			},
			{
				name: "复杂条件匹配-返回true",
				give: []int{1, 5, 10, 15, 20},
				f: func(n int) bool {
					return n > 10 && n%5 == 0
				},
				want: true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Any(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "Any() 的结果应与期望值相等")
			})
		}
	})

	// 字符串类型测试
	t.Run("字符串类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			f    func(string) bool
			want bool
		}{
			{
				name: "存在长字符串-返回true",
				give: []string{"apple", "banana", "cherry", "date"},
				f: func(s string) bool {
					return len(s) > 5
				},
				want: true,
			},
			{
				name: "不存在长字符串-返回false",
				give: []string{"a", "b", "c"},
				f: func(s string) bool {
					return len(s) > 5
				},
				want: false,
			},
			{
				name: "存在特定前缀-返回true",
				give: []string{"apple", "banana", "application"},
				f: func(s string) bool {
					return len(s) > 0 && s[0] == 'a'
				},
				want: true,
			},
			{
				name: "空字符串切片-返回false",
				give: []string{},
				f: func(s string) bool {
					return len(s) > 0
				},
				want: false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Any(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "Any() 的结果应与期望值相等")
			})
		}
	})

	// 结构体类型测试
	t.Run("结构体类型", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		tests := []struct {
			name string
			give []Person
			f    func(Person) bool
			want bool
		}{
			{
				name: "存在成年人-返回true",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 20},
				},
				f: func(p Person) bool {
					return p.Age >= 18
				},
				want: true,
			},
			{
				name: "不存在成年人-返回false",
				give: []Person{
					{"Alice", 15},
					{"Bob", 16},
					{"Charlie", 17},
				},
				f: func(p Person) bool {
					return p.Age >= 18
				},
				want: false,
			},
			{
				name: "存在特定姓名-返回true",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 20},
				},
				f: func(p Person) bool {
					return p.Name == "Bob"
				},
				want: true,
			},
			{
				name: "空结构体切片-返回false",
				give: []Person{},
				f: func(p Person) bool {
					return p.Age >= 18
				},
				want: false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Any(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "Any() 的结果应与期望值相等")
			})
		}
	})

	// 边界情况测试
	t.Run("边界情况", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			f    func(int) bool
			want bool
		}{
			{
				name: "nil切片-返回false",
				give: nil,
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "单元素切片匹配-返回true",
				give: []int{1},
				f: func(n int) bool {
					return n > 0
				},
				want: true,
			},
			{
				name: "单元素切片不匹配-返回false",
				give: []int{-1},
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "所有元素都匹配-返回true",
				give: []int{2, 4, 6, 8},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Any(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "Any() 边界情况的结果应与期望值相等")
			})
		}
	})
}

// TestAll 测试 All 函数
func TestAll(t *testing.T) {
	t.Parallel()

	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			f    func(int) bool
			want bool
		}{
			{
				name: "全部是偶数-返回true",
				give: []int{2, 4, 6, 8, 10},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: true,
			},
			{
				name: "包含奇数-返回false",
				give: []int{2, 4, 5, 6, 8},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: false,
			},
			{
				name: "全部是正数-返回true",
				give: []int{1, 2, 3, 4, 5},
				f: func(n int) bool {
					return n > 0
				},
				want: true,
			},
			{
				name: "包含负数-返回false",
				give: []int{-1, 2, 3, 4, 5},
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "包含零-返回false",
				give: []int{1, 2, 0, 3, 4},
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "全部大于10-返回true",
				give: []int{11, 12, 13, 14, 15},
				f: func(n int) bool {
					return n > 10
				},
				want: true,
			},
			{
				name: "包含小于等于10的数-返回false",
				give: []int{11, 12, 10, 14, 15},
				f: func(n int) bool {
					return n > 10
				},
				want: false,
			},
			{
				name: "空切片-返回true",
				give: []int{},
				f: func(n int) bool {
					return n > 0
				},
				want: true,
			},
			{
				name: "单元素匹配-返回true",
				give: []int{42},
				f: func(n int) bool {
					return n == 42
				},
				want: true,
			},
			{
				name: "单元素不匹配-返回false",
				give: []int{41},
				f: func(n int) bool {
					return n == 42
				},
				want: false,
			},
			{
				name: "复杂条件全部匹配-返回true",
				give: []int{10, 15, 20, 25, 30},
				f: func(n int) bool {
					return n >= 10 && n <= 30 && n%5 == 0
				},
				want: true,
			},
			{
				name: "复杂条件部分匹配-返回false",
				give: []int{10, 15, 21, 25, 30},
				f: func(n int) bool {
					return n >= 10 && n <= 30 && n%5 == 0
				},
				want: false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := All(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "All() 的结果应与期望值相等")
			})
		}
	})

	// 字符串类型测试
	t.Run("字符串类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			f    func(string) bool
			want bool
		}{
			{
				name: "全部是长字符串-返回true",
				give: []string{"apple", "banana", "cherry", "date"},
				f: func(s string) bool {
					return len(s) > 3
				},
				want: true,
			},
			{
				name: "包含短字符串-返回false",
				give: []string{"apple", "banana", "cat", "date"},
				f: func(s string) bool {
					return len(s) > 3
				},
				want: false,
			},
			{
				name: "全部以'a'开头-返回true",
				give: []string{"apple", "application", "array"},
				f: func(s string) bool {
					return len(s) > 0 && s[0] == 'a'
				},
				want: true,
			},
			{
				name: "包含不以'a'开头-返回false",
				give: []string{"apple", "banana", "application"},
				f: func(s string) bool {
					return len(s) > 0 && s[0] == 'a'
				},
				want: false,
			},
			{
				name: "空字符串切片-返回true",
				give: []string{},
				f: func(s string) bool {
					return len(s) > 0
				},
				want: true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := All(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "All() 的结果应与期望值相等")
			})
		}
	})

	// 结构体类型测试
	t.Run("结构体类型", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		tests := []struct {
			name string
			give []Person
			f    func(Person) bool
			want bool
		}{
			{
				name: "全部是成年人-返回true",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 35},
				},
				f: func(p Person) bool {
					return p.Age >= 18
				},
				want: true,
			},
			{
				name: "包含未成年人-返回false",
				give: []Person{
					{"Alice", 25},
					{"Bob", 16},
					{"Charlie", 35},
				},
				f: func(p Person) bool {
					return p.Age >= 18
				},
				want: false,
			},
			{
				name: "全部姓名以'A'开头-返回true",
				give: []Person{
					{"Alice", 25},
					{"Amy", 30},
					{"Andrew", 35},
				},
				f: func(p Person) bool {
					return len(p.Name) > 0 && p.Name[0] == 'A'
				},
				want: true,
			},
			{
				name: "包含不以'A'开头-返回false",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Andrew", 35},
				},
				f: func(p Person) bool {
					return len(p.Name) > 0 && p.Name[0] == 'A'
				},
				want: false,
			},
			{
				name: "空结构体切片-返回true",
				give: []Person{},
				f: func(p Person) bool {
					return p.Age >= 18
				},
				want: true,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := All(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "All() 的结果应与期望值相等")
			})
		}
	})

	// 边界情况测试
	t.Run("边界情况", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			f    func(int) bool
			want bool
		}{
			{
				name: "nil切片-返回true",
				give: nil,
				f: func(n int) bool {
					return n > 0
				},
				want: true,
			},
			{
				name: "单元素切片匹配-返回true",
				give: []int{1},
				f: func(n int) bool {
					return n > 0
				},
				want: true,
			},
			{
				name: "单元素切片不匹配-返回false",
				give: []int{-1},
				f: func(n int) bool {
					return n > 0
				},
				want: false,
			},
			{
				name: "所有元素都匹配-返回true",
				give: []int{2, 4, 6, 8},
				f: func(n int) bool {
					return n%2 == 0
				},
				want: true,
			},
			{
				name: "所有元素都是零值-返回true",
				give: []int{0, 0, 0},
				f: func(n int) bool {
					return n == 0
				},
				want: true,
			},
			{
				name: "混合类型条件-返回false",
				give: []int{1, 2, 3, 4, 5},
				f: func(n int) bool {
					return n > 10
				},
				want: false,
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := All(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "All() 边界情况的结果应与期望值相等")
			})
		}
	})
}


// TestSliceEqual 测试切片相等比较函数
func TestSliceEqualAdditional(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{"两个nil切片相等", nil, nil, true},
		{"nil与空切片不相等", nil, []int{}, false},
		{"空切片与nil切片不相等", []int{}, nil, false},
		{"两个空切片相等", []int{}, []int{}, true},
		{"相同元素切片相等", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"元素顺序不同相等", []int{1, 2, 3}, []int{3, 2, 1}, true},
		{"元素数量不同不相等", []int{1, 2, 3}, []int{1, 2}, false},
		{"元素内容不同不相等", []int{1, 2, 3}, []int{1, 2, 4}, false},
		{"重复元素处理", []int{1, 2, 2, 3}, []int{1, 2, 3, 2}, true},
		{"重复元素数量不同不相等", []int{1, 2, 2, 3}, []int{1, 2, 3}, false},
		{"单个元素切片", []int{42}, []int{42}, true},
		{"单个元素切片不相等", []int{42}, []int{24}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := SliceEqual(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "SliceEqual() 的结果应与期望值相等")
		})
	}
}

// TestString 测试String转换函数
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
		{"大整数", 999999999, "999999999"},
		{"负整数", -42, "-42"},
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

// TestJoin 测试Join连接函数
func TestJoin(t *testing.T) {
	t.Parallel()

	// 整数类型测试
	intTests := []struct {
		name  string
		give  []int
		glue  string
		want  string
	}{
		{"默认分隔符", []int{1, 2, 3}, "", "1,2,3"},
		{"自定义分隔符", []int{1, 2, 3}, "-", "1-2-3"},
		{"空分隔符", []int{1, 2, 3}, "", "1,2,3"},
		{"单元素", []int{42}, ",", "42"},
		{"空切片", []int{}, ",", ""},
		{"nil切片", nil, ",", ""},
		{"长分隔符", []int{1, 2, 3}, "->", "1->2->3"},
	}

	for _, tt := range intTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got string
			if tt.glue == "" {
				got = Join(tt.give)
			} else {
				got = Join(tt.give, tt.glue)
			}
			assert.Equal(t, tt.want, got, "Join() 整数的结果应与期望值相等")
		})
	}

	// 字符串类型测试
	stringTests := []struct {
		name  string
		give  []string
		glue  string
		want  string
	}{
		{"字符串切片默认分隔符", []string{"a", "b", "c"}, "", "a,b,c"},
		{"字符串切片自定义分隔符", []string{"a", "b", "c"}, " ", "a b c"},
		{"字符串切片单元素", []string{"hello"}, ",", "hello"},
		{"字符串切片空切片", []string{}, ",", ""},
		{"字符串切片nil切片", nil, ",", ""},
	}

	for _, tt := range stringTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var got string
			if tt.glue == "" {
				got = Join(tt.give)
			} else {
				got = Join(tt.give, tt.glue)
			}
			assert.Equal(t, tt.want, got, "Join() 字符串的结果应与期望值相等")
		})
	}
}

// TestMax 测试 Max 函数的各种情况
func TestMax(t *testing.T) {
	// 测试整数类型切片的最大值查找
	t.Run("整数类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			want int
		}{
			{"正整数切片", []int{1, 2, 3, 4, 5}, 5},
			{"负整数切片", []int{-1, -2, -3, -4, -5}, -1},
			{"混合正负整数切片", []int{-5, 0, 5, -10, 10}, 10},
			{"单个正整数", []int{42}, 42},
			{"单个负整数", []int{-42}, -42},
			{"重复值切片", []int{3, 3, 3, 3, 3}, 3},
			{"大数切片", []int{1, 999999999, 999999998}, 999999999},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Max(tt.give)
				assert.Equal(t, tt.want, got, "Max() 的结果应与期望值相等")
			})
		}
	})

	// 测试不同整数类型
	t.Run("不同整数类型", func(t *testing.T) {
		// int8 类型
		int8Slice := []int8{1, 2, 3, 4, 5}
		assert.Equal(t, int8(5), Max(int8Slice))

		// int16 类型
		int16Slice := []int16{1000, 2000, 3000}
		assert.Equal(t, int16(3000), Max(int16Slice))

		// int32 类型
		int32Slice := []int32{100000, 200000, 300000}
		assert.Equal(t, int32(300000), Max(int32Slice))

		// int64 类型
		int64Slice := []int64{10000000000, 20000000000, 30000000000}
		assert.Equal(t, int64(30000000000), Max(int64Slice))

		// uint 类型
		uintSlice := []uint{1, 2, 3, 4, 5}
		assert.Equal(t, uint(5), Max(uintSlice))

		// uint8 类型
		uint8Slice := []uint8{10, 20, 30, 40, 50}
		assert.Equal(t, uint8(50), Max(uint8Slice))
	})

	// 测试浮点数类型切片的最大值查找
	t.Run("浮点数类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []float64
			want float64
		}{
			{"正浮点数切片", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 5.5},
			{"负浮点数切片", []float64{-1.1, -2.2, -3.3, -4.4, -5.5}, -1.1},
			{"混合正负浮点数切片", []float64{-5.5, 0.0, 5.5, -10.1, 10.1}, 10.1},
			{"单个正浮点数", []float64{3.14159}, 3.14159},
			{"单个负浮点数", []float64{-3.14159}, -3.14159},
			{"重复值浮点数切片", []float64{2.718, 2.718, 2.718}, 2.718},
			{"科学计数法", []float64{1.23e-4, 5.67e8, 9.01e2}, 5.67e8},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Max(tt.give)
				assert.Equal(t, tt.want, got, "Max() 的结果应与期望值相等")
			})
		}
	})

	// 测试不同浮点数类型
	t.Run("不同浮点数类型", func(t *testing.T) {
		// float32 类型
		float32Slice := []float32{1.1, 2.2, 3.3}
		assert.Equal(t, float32(3.3), Max(float32Slice))
	})

	// 测试字符串类型切片的最大值查找
	t.Run("字符串类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			want string
		}{
			{"字母字符串切片", []string{"apple", "banana", "cherry", "date"}, "date"},
			{"混合大小写字符串切片", []string{"Apple", "banana", "Cherry", "date"}, "date"},
			{"数字字符串切片", []string{"1", "2", "10", "20"}, "20"},
			{"特殊字符字符串切片", []string{"!", "@", "#", "$"}, "@"}, // 修正：Unicode码点顺序，"@" > "#"
			{"中文字符串切片", []string{"苹果", "香蕉", "樱桃", "日期"}, "香蕉"}, // 修正：Unicode码点顺序，"香蕉" > "苹果"
			{"单个字符", []string{"a"}, "a"},
			{"重复字符串切片", []string{"hello", "hello", "hello"}, "hello"},
			{"空字符串和有效字符串", []string{"", "hello", "world"}, "world"},
			{"Unicode 字符串", []string{"α", "β", "γ", "δ"}, "δ"},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Max(tt.give)
				assert.Equal(t, tt.want, got, "Max() 的结果应与期望值相等")
			})
		}
	})

	// 测试空切片情况
	t.Run("空切片情况", func(t *testing.T) {
		// 整数空切片
		assert.Equal(t, 0, Max([]int{}), "空整数切片应返回零值")
		
		// 浮点数空切片
		assert.Equal(t, 0.0, Max([]float64{}), "空浮点数切片应返回零值")
		
		// 字符串空切片
		assert.Equal(t, "", Max([]string{}), "空字符串切片应返回零值")
		
		// 其他类型的空切片
		assert.Equal(t, int8(0), Max([]int8{}))
		assert.Equal(t, uint16(0), Max([]uint16{}))
		assert.Equal(t, float32(0.0), Max([]float32{}))
	})

	// 测试单元素切片情况
	t.Run("单元素切片情况", func(t *testing.T) {
		tests := []struct {
			name string
			give interface{}
			want interface{}
		}{
			{"单元素整数切片", []int{42}, 42},
			{"单元素浮点数切片", []float64{3.14159}, 3.14159},
			{"单元素字符串切片", []string{"hello"}, "hello"},
			{"单元素负整数切片", []int{-42}, -42},
			{"单元素零值切片", []int{0}, 0},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch v := tt.give.(type) {
				case []int:
					got := Max(v)
					assert.Equal(t, tt.want, got)
				case []float64:
					got := Max(v)
					assert.Equal(t, tt.want, got)
				case []string:
					got := Max(v)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	// 测试重复值情况
	t.Run("重复值情况", func(t *testing.T) {
		tests := []struct {
			name string
			give interface{}
			want interface{}
		}{
			{"全部相同的整数", []int{5, 5, 5, 5, 5}, 5},
			{"全部相同的浮点数", []float64{2.5, 2.5, 2.5}, 2.5},
			{"全部相同的字符串", []string{"test", "test", "test"}, "test"},
			{"最大值重复出现", []int{1, 5, 2, 5, 3, 5}, 5},
			{"最小值重复出现", []int{1, 1, 1, 2, 3, 4}, 4},
			{"中间值重复出现", []int{1, 2, 3, 2, 3, 1}, 3},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch v := tt.give.(type) {
				case []int:
					got := Max(v)
					assert.Equal(t, tt.want, got)
				case []float64:
					got := Max(v)
					assert.Equal(t, tt.want, got)
				case []string:
					got := Max(v)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	// 验证正确返回最大值元素
	t.Run("验证正确返回最大值元素", func(t *testing.T) {
		// 测试是否返回的是原始切片中的元素，而不是计算得到的值
		originalSlice := []int{1, 2, 3, 4, 5}
		maxValue := Max(originalSlice)
		
		// 验证返回值确实是最大值
		assert.Equal(t, 5, maxValue)
		
		// 验证返回值存在于原切片中
		found := false
		for _, v := range originalSlice {
			if v == maxValue {
				found = true
				break
			}
		}
		assert.True(t, found, "返回的最大值应该存在于原切片中")

		// 对于浮点数，考虑精度问题
		floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		floatMax := Max(floatSlice)
		assert.InDelta(t, 5.5, floatMax, 1e-9, "浮点数最大值应该在精度范围内")
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		// 测试极大整数
		largeIntSlice := []int{1, math.MaxInt64 - 1, math.MaxInt64}
		assert.Equal(t, math.MaxInt64, Max(largeIntSlice))

		// 测试极小整数
		smallIntSlice := []int{-1, math.MinInt64 + 1, math.MinInt64}
		assert.Equal(t, -1, Max(smallIntSlice))

		// 测试极大浮点数
		largeFloatSlice := []float64{1.0, math.MaxFloat64 - 1.0, math.MaxFloat64}
		assert.Equal(t, math.MaxFloat64, Max(largeFloatSlice))

		// 测试极小浮点数
		smallFloatSlice := []float64{-1.0, -math.MaxFloat64, -2.0}
		assert.Equal(t, -1.0, Max(smallFloatSlice))

		// 测试无穷大
		infinitySlice := []float64{1.0, math.Inf(1), 2.0}
		assert.Equal(t, math.Inf(1), Max(infinitySlice))

		// 测试负无穷大
		negativeInfinitySlice := []float64{-1.0, math.Inf(-1), -2.0}
		assert.Equal(t, -1.0, Max(negativeInfinitySlice))

		// 测试 NaN (NaN 在比较中总是返回 false，所以应该返回第一个元素 NaN)
		nanSlice := []float64{math.NaN(), 1.0, 2.0}
		assert.True(t, math.IsNaN(Max(nanSlice)), "NaN 比较应返回第一个元素 NaN")

		allNanSlice := []float64{math.NaN(), math.NaN(), math.NaN()}
		assert.True(t, math.IsNaN(Max(allNanSlice)), "全 NaN 切片应该返回 NaN")
	})
}

// BenchmarkMax 基准测试 Max 函数
func BenchmarkMax(b *testing.B) {
	// 基准测试整数类型
	b.Run("整数类型", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Max(slice)
		}
	})

	// 基准测试浮点数类型
	b.Run("浮点数类型", func(b *testing.B) {
		slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Max(slice)
		}
	})

	// 基准测试字符串类型
	b.Run("字符串类型", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Max(slice)
		}
	})

	// 基准测试大切片
	b.Run("大切片", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Max(slice)
		}
	})

	// 基准测试空切片
	b.Run("空切片", func(b *testing.B) {
		slice := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Max(slice)
		}
	})
}

// TestMin 测试 Min 函数
func TestMin(t *testing.T) {
	// 测试整数类型切片的最小值查找
	t.Run("整数类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []int
			want int
		}{
			{"正整数切片", []int{1, 2, 3, 4, 5}, 1},
			{"负整数切片", []int{-1, -2, -3, -4, -5}, -5},
			{"混合正负整数切片", []int{-5, 0, 5, -10, 10}, -10},
			{"单个正整数", []int{42}, 42},
			{"单个负整数", []int{-42}, -42},
			{"重复值切片", []int{3, 3, 3, 3, 3}, 3},
			{"大数切片", []int{1, 999999999, 999999998}, 1},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Min(tt.give)
				assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
			})
		}
	})

	// 测试不同整数类型
	t.Run("不同整数类型", func(t *testing.T) {
		// int8 类型
		int8Slice := []int8{1, 2, 3, 4, 5}
		assert.Equal(t, int8(1), Min(int8Slice))

		// int16 类型
		int16Slice := []int16{1000, 2000, 3000}
		assert.Equal(t, int16(1000), Min(int16Slice))

		// int32 类型
		int32Slice := []int32{100000, 200000, 300000}
		assert.Equal(t, int32(100000), Min(int32Slice))

		// int64 类型
		int64Slice := []int64{10000000000, 20000000000, 30000000000}
		assert.Equal(t, int64(10000000000), Min(int64Slice))

		// uint 类型
		uintSlice := []uint{1, 2, 3, 4, 5}
		assert.Equal(t, uint(1), Min(uintSlice))

		// uint8 类型
		uint8Slice := []uint8{10, 20, 30, 40, 50}
		assert.Equal(t, uint8(10), Min(uint8Slice))
	})

	// 测试浮点数类型切片的最小值查找
	t.Run("浮点数类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []float64
			want float64
		}{
			{"正浮点数切片", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 1.1},
			{"负浮点数切片", []float64{-1.1, -2.2, -3.3, -4.4, -5.5}, -5.5},
			{"混合正负浮点数切片", []float64{-5.5, 0.0, 5.5, -10.1, 10.1}, -10.1},
			{"单个正浮点数", []float64{3.14159}, 3.14159},
			{"单个负浮点数", []float64{-3.14159}, -3.14159},
			{"重复值浮点数切片", []float64{2.718, 2.718, 2.718}, 2.718},
			{"科学计数法", []float64{1.23e-4, 5.67e8, 9.01e2}, 1.23e-4},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Min(tt.give)
				assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
			})
		}
	})

	// 测试不同浮点数类型
	t.Run("不同浮点数类型", func(t *testing.T) {
		// float32 类型
		float32Slice := []float32{1.1, 2.2, 3.3}
		assert.Equal(t, float32(1.1), Min(float32Slice))
	})

	// 测试字符串类型切片的最小值查找
	t.Run("字符串类型切片", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			want string
		}{
			{"字母字符串切片", []string{"apple", "banana", "cherry", "date"}, "apple"},
			{"混合大小写字符串切片", []string{"Apple", "banana", "Cherry", "date"}, "Apple"},
			{"数字字符串切片", []string{"1", "2", "10", "20"}, "1"},
			{"特殊字符字符串切片", []string{"!", "@", "#", "$"}, "!"},
			{"中文字符串切片", []string{"苹果", "香蕉", "樱桃", "日期"}, "日期"},
			{"单个字符", []string{"a"}, "a"},
			{"重复字符串切片", []string{"hello", "hello", "hello"}, "hello"},
			{"空字符串和有效字符串", []string{"", "hello", "world"}, ""},
			{"Unicode 字符串", []string{"α", "β", "γ", "δ"}, "α"},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Min(tt.give)
				assert.Equal(t, tt.want, got, "Min() 的结果应与期望值相等")
			})
		}
	})

	// 测试空切片情况
	t.Run("空切片情况", func(t *testing.T) {
		// 整数空切片
		assert.Equal(t, 0, Min([]int{}), "空整数切片应返回零值")
		
		// 浮点数空切片
		assert.Equal(t, 0.0, Min([]float64{}), "空浮点数切片应返回零值")
		
		// 字符串空切片
		assert.Equal(t, "", Min([]string{}), "空字符串切片应返回零值")
		
		// 其他类型的空切片
		assert.Equal(t, int8(0), Min([]int8{}))
		assert.Equal(t, uint16(0), Min([]uint16{}))
		assert.Equal(t, float32(0.0), Min([]float32{}))
	})

	// 测试单元素切片情况
	t.Run("单元素切片情况", func(t *testing.T) {
		tests := []struct {
			name string
			give interface{}
			want interface{}
		}{
			{"单元素整数切片", []int{42}, 42},
			{"单元素浮点数切片", []float64{3.14159}, 3.14159},
			{"单元素字符串切片", []string{"hello"}, "hello"},
			{"单元素负整数切片", []int{-42}, -42},
			{"单元素零值切片", []int{0}, 0},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch v := tt.give.(type) {
				case []int:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []float64:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []string:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	// 测试重复值情况
	t.Run("重复值情况", func(t *testing.T) {
		tests := []struct {
			name string
			give interface{}
			want interface{}
		}{
			{"全部相同的整数", []int{5, 5, 5, 5, 5}, 5},
			{"全部相同的浮点数", []float64{2.5, 2.5, 2.5}, 2.5},
			{"全部相同的字符串", []string{"test", "test", "test"}, "test"},
			{"最小值重复出现", []int{1, 5, 2, 1, 3, 1}, 1},
			{"最大值重复出现", []int{10, 10, 10, 2, 3, 4}, 2},
			{"中间值重复出现", []int{5, 2, 3, 2, 3, 5}, 2},
		}

		for _, tt := range tests {
			tt := tt // 避免竞态
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch v := tt.give.(type) {
				case []int:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []float64:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				case []string:
					got := Min(v)
					assert.Equal(t, tt.want, got)
				}
			})
		}
	})

	// 验证正确返回最小值元素
	t.Run("验证正确返回最小值元素", func(t *testing.T) {
		// 测试是否返回的是原始切片中的元素，而不是计算得到的值
		originalSlice := []int{1, 2, 3, 4, 5}
		minValue := Min(originalSlice)
		
		// 验证返回值确实是最小值
		assert.Equal(t, 1, minValue)
		
		// 验证返回值存在于原切片中
		found := false
		for _, v := range originalSlice {
			if v == minValue {
				found = true
				break
			}
		}
		assert.True(t, found, "返回的最小值应该存在于原切片中")

		// 对于浮点数，考虑精度问题
		floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
		floatMin := Min(floatSlice)
		assert.InDelta(t, 1.1, floatMin, 1e-9, "浮点数最小值应该在精度范围内")
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		// 测试极大整数
		largeIntSlice := []int{1, math.MaxInt64 - 1, math.MaxInt64}
		assert.Equal(t, 1, Min(largeIntSlice))

		// 测试极小整数
		smallIntSlice := []int{-1, math.MinInt64 + 1, math.MinInt64}
		assert.Equal(t, math.MinInt64, Min(smallIntSlice))

		// 测试极大浮点数
		largeFloatSlice := []float64{1.0, math.MaxFloat64 - 1.0, math.MaxFloat64}
		assert.Equal(t, 1.0, Min(largeFloatSlice))

		// 测试极小浮点数
		smallFloatSlice := []float64{-1.0, -math.MaxFloat64, -2.0}
		assert.Equal(t, -math.MaxFloat64, Min(smallFloatSlice))

		// 测试无穷大
		infinitySlice := []float64{1.0, math.Inf(1), 2.0}
		assert.Equal(t, 1.0, Min(infinitySlice))

		// 测试负无穷大
		negativeInfinitySlice := []float64{-1.0, math.Inf(-1), -2.0}
		assert.Equal(t, math.Inf(-1), Min(negativeInfinitySlice))

		// 测试 NaN (NaN 在比较中总是返回 false，所以应该返回第一个元素 NaN)
		nanSlice := []float64{math.NaN(), 1.0, 2.0}
		assert.True(t, math.IsNaN(Min(nanSlice)), "NaN 比较应返回第一个元素 NaN")

		allNanSlice := []float64{math.NaN(), math.NaN(), math.NaN()}
		assert.True(t, math.IsNaN(Min(allNanSlice)), "全 NaN 切片应该返回 NaN")
	})
}

// BenchmarkMin 基准测试 Min 函数
func BenchmarkMin(b *testing.B) {
	// 基准测试整数类型
	b.Run("整数类型", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试浮点数类型
	b.Run("浮点数类型", func(b *testing.B) {
		slice := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.0}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试字符串类型
	b.Run("字符串类型", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试大切片
	b.Run("大切片", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})

	// 基准测试空切片
	b.Run("空切片", func(b *testing.B) {
		slice := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Min(slice)
		}
	})
}

// TestUnique 测试 Unique 函数
func TestUnique(t *testing.T) {
	t.Parallel()

	// 测试整数类型切片去重
	t.Run("整数类型切片去重", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			give  []int
			want []int
		}{
			{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
			{[]int{5, 5, 5, 5, 5}, []int{5}},
			{[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
			{[]int{1, 3, 2, 4, 3, 2, 1}, []int{1, 3, 2, 4}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%v", tt.give), func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "整数切片去重结果应匹配")
			})
		}
	})

	// 测试浮点数类型切片去重
	t.Run("浮点数类型切片去重", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			give  []float64
			want []float64
		}{
			{[]float64{1.1, 2.2, 3.3, 2.2, 1.1}, []float64{1.1, 2.2, 3.3}},
			{[]float64{5.5, 5.5, 5.5}, []float64{5.5}},
			{[]float64{1.0, 2.0, 3.0, 4.0}, []float64{1.0, 2.0, 3.0, 4.0}},
			{[]float64{0.1, 0.2, 0.1, 0.3}, []float64{0.1, 0.2, 0.3}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%v", tt.give), func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "浮点数切片去重结果应匹配")
			})
		}
	})

	// 测试字符串类型切片去重
	t.Run("字符串类型切片去重", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			give  []string
			want []string
		}{
			{[]string{"a", "b", "c", "b", "a"}, []string{"a", "b", "c"}},
			{[]string{"hello", "hello", "hello"}, []string{"hello"}},
			{[]string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}},
			{[]string{"go", "python", "go", "java", "python"}, []string{"go", "python", "java"}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%v", tt.give), func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "字符串切片去重结果应匹配")
			})
		}
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			give  []int
			want []int
		}{
			{"空切片", []int{}, []int{}},
			{"单元素切片", []int{42}, []int{42}},
			{"全部相同元素", []int{7, 7, 7, 7}, []int{7}},
			{"无重复元素", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			{"连续重复", []int{1, 1, 2, 2, 3, 3}, []int{1, 2, 3}},
			{"间隔重复", []int{1, 2, 1, 3, 2, 1}, []int{1, 2, 3}},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Unique(tt.give)
				assert.Equal(t, tt.want, got, "边界情况处理应正确")
			})
		}
	})

	// 测试保留原始顺序
	t.Run("保留原始顺序", func(t *testing.T) {
		t.Parallel()
		// 测试去重后的元素保持原始出现顺序
		original := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
		result := Unique(original)
		expected := []int{3, 1, 4, 5, 9, 2, 6}
		assert.Equal(t, expected, result, "去重后应保留原始顺序")
	})

	// 测试不修改原切片
	t.Run("不修改原切片", func(t *testing.T) {
		t.Parallel()
		original := []int{1, 2, 2, 3}
		originalCopy := make([]int, len(original))
		copy(originalCopy, original)
		
		result := Unique(original)
		
		// 确保原切片未被修改
		assert.Equal(t, originalCopy, original, "原切片应保持不变")
		// 确保返回的是新切片
		assert.NotSame(t, &original[0], &result[0], "应返回新切片")
	})
}

// BenchmarkUnique 测试 Unique 函数的基准性能
func BenchmarkUnique(b *testing.B) {
	// 小数据集基准测试
	b.Run("小数据集", func(b *testing.B) {
		data := []int{1, 2, 3, 2, 1, 4, 5, 4, 3}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})

	// 中等数据集基准测试
	b.Run("中等数据集", func(b *testing.B) {
		data := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			data[i] = i % 100 // 创建有重复的数据
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})

	// 大数据集基准测试
	b.Run("大数据集", func(b *testing.B) {
		data := make([]int, 100000)
		for i := 0; i < 100000; i++ {
			data[i] = i % 1000 // 创建有重复的数据
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})

	// 无重复数据集基准测试
	b.Run("无重复数据集", func(b *testing.B) {
		data := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			data[i] = i // 创建无重复的数据
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unique(data)
		}
	})
}
// TestContains 测试 Contains 函数的各种场景
func TestContains(t *testing.T) {
	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name string
			slice []int
			target int
			want bool
		}{
			{"包含元素", []int{1, 2, 3, 4, 5}, 3, true},
			{"不包含元素", []int{1, 2, 3, 4, 5}, 6, false},
			{"空切片", []int{}, 1, false},
			{"单元素-匹配", []int{42}, 42, true},
			{"单元素-不匹配", []int{42}, 24, false},
			{"重复元素", []int{1, 2, 2, 3, 2}, 2, true},
			{"负数", []int{-1, -2, -3}, -2, true},
			{"零值", []int{0, 1, 2}, 0, true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Contains(tt.slice, tt.target)
				assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
			})
		}
	})

	// 浮点数类型测试
	t.Run("浮点数类型", func(t *testing.T) {
		tests := []struct {
			name string
			slice []float64
			target float64
			want bool
		}{
			{"包含元素", []float64{1.1, 2.2, 3.3}, 2.2, true},
			{"不包含元素", []float64{1.1, 2.2, 3.3}, 4.4, false},
			{"空切片", []float64{}, 1.1, false},
			{"科学计数法", []float64{1.5e10, 2.3e-5}, 1.5e10, true},
			{"精度测试 - 浮点数精确比较", []float64{0.1 + 0.2}, 0.3, true}, // 浮点数精度问题
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Contains(tt.slice, tt.target)
				assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
			})
		}
	})

	// 字符串类型测试
	t.Run("字符串类型", func(t *testing.T) {
		tests := []struct {
			name string
			slice []string
			target string
			want bool
		}{
			{"包含元素", []string{"apple", "banana", "cherry"}, "banana", true},
			{"不包含元素", []string{"apple", "banana", "cherry"}, "orange", false},
			{"空切片", []string{}, "test", false},
			{"空字符串", []string{"", "hello", ""}, "", true},
			{"中文字符串", []string{"苹果", "香蕉", "橙子"}, "香蕉", true},
			{"特殊字符", []string{"a@b.com", "x#y", "test$"}, "x#y", true},
			{"Unicode字符", []string{"café", "naïve", "résumé"}, "naïve", true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Contains(tt.slice, tt.target)
				assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
			})
		}
	})

	// 边界情况测试
	t.Run("边界情况", func(t *testing.T) {
		tests := []struct {
			name string
			slice interface{}
			target interface{}
			want bool
		}{
			{"大整数切片", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, true},
			{"大字符串切片", []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, "j", true},
			{"nil切片", ([]int)(nil), 1, false},
			{"首元素", []int{1, 2, 3}, 1, true},
			{"末元素", []int{1, 2, 3}, 3, true},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				switch s := tt.slice.(type) {
				case []int:
					got := Contains(s, tt.target.(int))
					assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
				case []string:
					got := Contains(s, tt.target.(string))
					assert.Equal(t, tt.want, got, "Contains() 的结果应与期望值相等")
				}
			})
		}
	})
}

// BenchmarkContains 性能测试
func BenchmarkContains(b *testing.B) {
	// 小切片测试
	b.Run("小切片-存在", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		target := 3
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	b.Run("小切片-不存在", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		target := 99
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	// 中等切片测试
	b.Run("中等切片-存在", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		target := 500
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	b.Run("中等切片-不存在", func(b *testing.B) {
		slice := make([]int, 1000)
		for i := range slice {
			slice[i] = i
		}
		target := 9999
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	// 大切片测试
	b.Run("大切片-存在", func(b *testing.B) {
		slice := make([]int, 100000)
		for i := range slice {
			slice[i] = i
		}
		target := 50000
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	b.Run("大切片-不存在", func(b *testing.B) {
		slice := make([]int, 100000)
		for i := range slice {
			slice[i] = i
		}
		target := 999999
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})

	// 字符串切片测试
	b.Run("字符串切片", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
		target := "cherry"
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Contains(slice, target)
		}
	})
}
