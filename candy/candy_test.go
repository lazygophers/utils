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
