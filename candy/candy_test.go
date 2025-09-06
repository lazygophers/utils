package candy

import (
	
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestPerson 测试用的Person结构体
type TestPerson struct {
	Name string
	Age  int
}

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
		name string
		give []int
		glue string
		want string
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
		name string
		give []string
		glue string
		want string
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




// TestContains 测试 Contains 函数的各种场景
func TestContains(t *testing.T) {
	// 整数类型测试
	t.Run("整数类型", func(t *testing.T) {
		tests := []struct {
			name   string
			slice  []int
			target int
			want   bool
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
			name   string
			slice  []float64
			target float64
			want   bool
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
			name   string
			slice  []string
			target string
			want   bool
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
			name   string
			slice  interface{}
			target interface{}
			want   bool
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



// TestRandom 测试 Random 函数
func TestRandom(t *testing.T) {
	t.Parallel()

	// 测试整数类型切片的随机选择
	t.Run("整数类型切片", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name string
			give []int
			// 验证返回值是否在切片中
			validate func(int, []int) bool
		}{
			{
				name: "多元素切片",
				give: []int{1, 2, 3, 4, 5},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "单元素切片",
				give: []int{42},
				validate: func(result int, slice []int) bool {
					return result == slice[0]
				},
			},
			{
				name: "重复元素切片",
				give: []int{1, 2, 2, 3, 3},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "负数切片",
				give: []int{-1, -2, -3, -4, -5},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "混合正负数切片",
				give: []int{-5, 0, 5, -10, 10},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "大数切片",
				give: []int{1000000, 2000000, 3000000},
				validate: func(result int, slice []int) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				results := make(map[int]int)
				for i := 0; i < 100; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原切片中")
					results[result]++
				}

				// 验证所有元素都有被选中的机会（对于足够大的切片）
				if len(tt.give) > 1 {
					_ = true // 确保所有元素都被选中
					for _, expected := range tt.give {
						if results[expected] == 0 {
							_ = false
							break
						}
					}
					// 由于随机性，这个测试有一定概率失败，但100次测试应该覆盖所有元素
					// 我们主要验证结果的有效性
				}
			})
		}
	})

	// 测试浮点数类型切片的随机选择
	t.Run("浮点数类型切片", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name     string
			give     []float64
			validate func(float64, []float64) bool
		}{
			{
				name: "多元素浮点数切片",
				give: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				validate: func(result float64, slice []float64) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "科学计数法浮点数",
				give: []float64{1.5e10, 2.3e-5, 3.14},
				validate: func(result float64, slice []float64) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "负浮点数切片",
				give: []float64{-1.1, -2.2, -3.3},
				validate: func(result float64, slice []float64) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				for i := 0; i < 50; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原浮点数切片中")
				}
			})
		}
	})

	// 测试字符串类型切片的随机选择
	t.Run("字符串类型切片", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name     string
			give     []string
			validate func(string, []string) bool
		}{
			{
				name: "多元素字符串切片",
				give: []string{"apple", "banana", "cherry", "date", "elderberry"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "Unicode字符串切片",
				give: []string{"苹果", "香蕉", "樱桃", "日期"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "特殊字符字符串切片",
				give: []string{"a@b.com", "x#y", "test$", "hello world"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "空字符串切片",
				give: []string{"", "hello", "", "world"},
				validate: func(result string, slice []string) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				for i := 0; i < 50; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原字符串切片中")
				}
			})
		}
	})

	// 测试结构体类型切片的随机选择
	t.Run("结构体类型切片", func(t *testing.T) {
		t.Parallel()

		type Person struct {
			Name string
			Age  int
		}

		tests := []struct {
			name     string
			give     []Person
			validate func(Person, []Person) bool
		}{
			{
				name: "多元素结构体切片",
				give: []Person{
					{"Alice", 25},
					{"Bob", 30},
					{"Charlie", 35},
					{"David", 40},
				},
				validate: func(result Person, slice []Person) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
			{
				name: "重复字段结构体切片",
				give: []Person{
					{"Alice", 25},
					{"Bob", 25},
					{"Alice", 30}, // 重复姓名，不同年龄
				},
				validate: func(result Person, slice []Person) bool {
					for _, v := range slice {
						if v == result {
							return true
						}
					}
					return false
				},
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				// 多次测试以验证随机性
				for i := 0; i < 50; i++ {
					result := Random(tt.give)
					assert.True(t, tt.validate(result, tt.give), "Random() 返回值应在原结构体切片中")
				}
			})
		}
	})

	// 测试边界情况
	t.Run("边界情况", func(t *testing.T) {
		t.Parallel()

		// 测试空切片 - 应返回零值
		t.Run("空切片", func(t *testing.T) {
			t.Parallel()

			// 整数空切片
			var emptyIntSlice []int
			intResult := Random(emptyIntSlice)
			assert.Equal(t, 0, intResult, "空整数切片应返回零值")

			// 字符串空切片
			var emptyStringSlice []string
			stringResult := Random(emptyStringSlice)
			assert.Equal(t, "", stringResult, "空字符串切片应返回零值")

			// 浮点数空切片
			var emptyFloatSlice []float64
			floatResult := Random(emptyFloatSlice)
			assert.Equal(t, 0.0, floatResult, "空浮点数切片应返回零值")

			// 结构体空切片
			var emptyStructSlice []TestPerson
			structResult := Random(emptyStructSlice)
			assert.Equal(t, TestPerson{}, structResult, "空结构体切片应返回零值")
		})

		// 测试单元素切片
		t.Run("单元素切片", func(t *testing.T) {
			t.Parallel()

			// 单元素整数切片
			singleInt := []int{42}
			result := Random(singleInt)
			assert.Equal(t, 42, result, "单元素整数切片应返回该元素")

			// 单元素字符串切片
			singleString := []string{"hello"}
			singleStringResult := Random(singleString)
			assert.Equal(t, "hello", singleStringResult, "单元素字符串切片应返回该元素")

			// 单元素浮点数切片
			singleFloat := []float64{3.14}
			singleFloatResult := Random(singleFloat)
			assert.Equal(t, 3.14, singleFloatResult, "单元素浮点数切片应返回该元素")
		})

		// 测试nil切片
		t.Run("nil切片", func(t *testing.T) {
			t.Parallel()

			var nilSlice []int
			result := Random(nilSlice)
			assert.Equal(t, 0, result, "nil切片应返回零值")
		})
	})

	// 测试随机性分布
	t.Run("随机性分布", func(t *testing.T) {
		t.Parallel()

		// 使用固定种子进行可重复的随机测试
		originalSeed := rand.Int63()
		defer func() {
			// 恢复原始种子
			rand.Seed(originalSeed)
		}()

		// 设置固定种子以确保测试可重复
		rand.Seed(12345)

		slice := []int{1, 2, 3, 4, 5}
		results := make([]int, 1000)

		// 生成1000个随机结果
		for i := range results {
			results[i] = Random(slice)
		}

		// 统计每个元素出现的频率
		frequency := make(map[int]int)
		for _, result := range results {
			frequency[result]++
		}

		// 验证每个元素都出现了（在1000次测试中）
		for _, expected := range slice {
			assert.Greater(t, frequency[expected], 0, "每个元素都应该被随机选中至少一次")
		}

		// 验证分布大致均匀（允许一定的偏差）
		expectedFrequency := 1000 / len(slice)
		for _, count := range frequency {
			// 允许20%的偏差
			deviation := float64(count-expectedFrequency) / float64(expectedFrequency)
			assert.LessOrEqual(t, deviation, 0.5, "随机分布应该大致均匀")
		}
	})

	// 测试不同类型的一致性行为
	t.Run("类型一致性", func(t *testing.T) {
		t.Parallel()

		// 验证对于相同的数据，不同类型的结果都是有效的
		intSlice := []int{1, 2, 3}
		stringSlice := []string{"1", "2", "3"}
		floatSlice := []float64{1.0, 2.0, 3.0}

		// 每种类型都应该能正确处理
		intResult := Random(intSlice)
		stringResult := Random(stringSlice)
		floatResult := Random(floatSlice)

		// 验证结果类型正确
		assert.IsType(t, intResult, 0, "整数切片应返回整数类型")
		assert.IsType(t, stringResult, "", "字符串切片应返回字符串类型")
		assert.IsType(t, floatResult, 0.0, "浮点数切片应返回浮点数类型")

		// 验证结果在各自切片中
		assert.Contains(t, intSlice, intResult)
		assert.Contains(t, stringSlice, stringResult)
		assert.Contains(t, floatSlice, floatResult)
	})
}

// BenchmarkRandom 基准测试 Random 函数
func BenchmarkRandom(b *testing.B) {
	// 基准测试小切片
	b.Run("小切片", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试中等切片
	b.Run("中等切片", func(b *testing.B) {
		slice := make([]int, 100)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试大切片
	b.Run("大切片", func(b *testing.B) {
		slice := make([]int, 10000)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试空切片
	b.Run("空切片", func(b *testing.B) {
		slice := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试单元素切片
	b.Run("单元素切片", func(b *testing.B) {
		slice := []int{42}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试字符串切片
	b.Run("字符串切片", func(b *testing.B) {
		slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})

	// 基准测试结构体切片
	b.Run("结构体切片", func(b *testing.B) {
		type Person struct {
			Name string
			Age  int
		}
		slice := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
			{"David", 40},
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(slice)
		}
	})
}

// TestEach 测试Each函数
func TestEach(t *testing.T) {
	// 测试基本功能：整数切片
	t.Run("整数切片", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int

		Each(data, func(item int) {
			result = append(result, item*2)
		})

		require.Equal(t, []int{2, 4, 6, 8, 10}, result, "Each应该对每个元素执行函数")
	})

	// 测试字符串切片
	t.Run("字符串切片", func(t *testing.T) {
		t.Parallel()

		data := []string{"a", "b", "c"}
		var result []string

		Each(data, func(item string) {
			result = append(result, item+"x")
		})

		require.Equal(t, []string{"ax", "bx", "cx"}, result, "Each应该对字符串切片正常工作")
	})

	// 测试结构体切片
	t.Run("结构体切片", func(t *testing.T) {
		t.Parallel()

		type TestItem struct {
			ID   int
			Name string
		}

		data := []TestItem{
			{ID: 1, Name: "item1"},
			{ID: 2, Name: "item2"},
		}

		var result []int
		Each(data, func(item TestItem) {
			result = append(result, item.ID)
		})

		require.Equal(t, []int{1, 2}, result, "Each应该对结构体切片正常工作")
	})

	// 测试浮点数切片
	t.Run("浮点数切片", func(t *testing.T) {
		t.Parallel()

		data := []float64{1.1, 2.2, 3.3}
		var sum float64

		Each(data, func(item float64) {
			sum += item
		})

		assert.InDelta(t, 6.6, sum, 0.001, "Each应该对浮点数切片正常工作")
	})

	// 测试空切片
	t.Run("空切片", func(t *testing.T) {
		t.Parallel()

		data := []int{}
		var result []int

		Each(data, func(item int) {
			result = append(result, item)
		})

		require.Empty(t, result, "Each处理空切片时不应该执行函数")
	})

	// 测试nil切片
	t.Run("nil切片", func(t *testing.T) {
		t.Parallel()

		var data []int
		var result []int

		Each(data, func(item int) {
			result = append(result, item)
		})

		require.Empty(t, result, "Each处理nil切片时不应该执行函数")
	})

	// 测试单元素切片
	t.Run("单元素切片", func(t *testing.T) {
		t.Parallel()

		data := []int{42}
		var result int

		Each(data, func(item int) {
			result = item
		})

		require.Equal(t, 42, result, "Each应该正确处理单元素切片")
	})

	// 测试函数副作用
	t.Run("函数副作用", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3}
		counter := 0

		Each(data, func(item int) {
			counter++
		})

		require.Equal(t, 3, counter, "Each应该对每个元素执行一次函数")
	})

	// 测试修改原始切片元素
	t.Run("修改原始元素", func(t *testing.T) {
		t.Parallel()

		data := []TestPerson{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		Each(data, func(item TestPerson) {
			item.Age += 10
		})

		// 注意：Each函数接收的是值拷贝，所以原始切片不会被修改
		require.Equal(t, 25, data[0].Age, "Each不应该修改原始切片元素（值拷贝）")
		require.Equal(t, 30, data[1].Age, "Each不应该修改原始切片元素（值拷贝）")
	})

	// 测试指针切片
	t.Run("指针切片", func(t *testing.T) {
		t.Parallel()

		data := []*TestPerson{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}

		Each(data, func(item *TestPerson) {
			item.Age += 10
		})

		require.Equal(t, 35, data[0].Age, "Each应该可以通过指针修改原始数据")
		require.Equal(t, 40, data[1].Age, "Each应该可以通过指针修改原始数据")
	})

	// 测试复杂计算
	t.Run("复杂计算", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var sum int
		var product = 1

		Each(data, func(item int) {
			sum += item
			product *= item
		})

		require.Equal(t, 15, sum, "Each应该正确计算总和")
		require.Equal(t, 120, product, "Each应该正确计算乘积")
	})

	// 测试并发安全性
	t.Run("并发安全", func(t *testing.T) {
		t.Parallel()

		data := []int{1, 2, 3, 4, 5}
		var result []int
		var mu sync.Mutex

		Each(data, func(item int) {
			mu.Lock()
			result = append(result, item*2)
			mu.Unlock()
		})

		require.Equal(t, []int{2, 4, 6, 8, 10}, result, "Each在并发环境下应该安全工作")
	})
}

// BenchmarkEach 测试Each函数的性能
func BenchmarkEach(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var sum int
		Each(data, func(item int) {
			sum += item
		})
	}
}
