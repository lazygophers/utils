package candy

import (
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

	// 字符串类型测试
	t.Run("字符串类型", func(t *testing.T) {
		tests := []struct {
			name string
			give []string
			f    func(string) bool
			want []string
		}{
			{
				name: "过滤长字符串保留短字符串",
				give: []string{"apple", "banana", "cherry", "date"},
				f: func(s string) bool {
					return len(s) > 5
				},
				want: []string{"apple", "date"},
			},
			{
				name: "空字符串切片",
				give: []string{},
				f: func(s string) bool {
					return len(s) > 0
				},
				want: []string{},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				got := FilterNot(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "FilterNot() 的结果应与期望值相等")
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
			want []Person
		}{
			{
				name: "过滤年龄大于等于25的人",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 20},
				},
				f: func(p Person) bool {
					return p.Age >= 25
				},
				want: []Person{
					{"Charlie", 20},
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				got := FilterNot(tt.give, tt.f)
				assert.Equal(t, tt.want, got, "FilterNot() 的结果应与期望值相等")
			})
		}
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
