package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAccessGenericSlice_FullCoverage 全面测试 accessGenericSlice 的覆盖率
func TestAccessGenericSlice_FullCoverage(t *testing.T) {
	tests := []struct {
		name        string
		slice       any
		index       int
		wantErr     bool
		errType     error
		description string
	}{
		// 当前实现测试（始终返回错误）
		{"nil 切片", nil, 0, true, ErrInvalidSlice, "nil 值应返回错误"},
		{"非切片类型", "string", 0, true, ErrInvalidSlice, "字符串类型应返回错误"},
		{"整数", 42, 0, true, ErrInvalidSlice, "整数类型应返回错误"},
		{"map", map[string]any{}, 0, true, ErrInvalidSlice, "map 类型应返回错误"},

		// 未支持的切片类型
		{"[]uint", []uint{1, 2, 3}, 1, true, ErrInvalidSlice, "uint 切片应返回错误"},
		{"[]float32", []float32{1.1, 2.2}, 0, true, ErrInvalidSlice, "float32 切片应返回错误"},
		{"[]int32", []int32{1, 2}, 0, true, ErrInvalidSlice, "int32 切片应返回错误"},
		{"[]uint8", []uint8{1, 2, 3}, 2, true, ErrInvalidSlice, "uint8 切片应返回错误"},
		{"[]int16", []int16{10, 20}, 1, true, ErrInvalidSlice, "int16 切片应返回错误"},
		{"[]uint16", []uint16{100, 200}, 0, true, ErrInvalidSlice, "uint16 切片应返回错误"},
		{"[]uint64", []uint64{1, 2, 3}, 2, true, ErrInvalidSlice, "uint64 切片应返回错误"},

		// 自定义类型
		{"自定义切片", []struct{ X int }{{1}, {2}}, 0, true, ErrInvalidSlice, "自定义结构体切片应返回错误"},

		// 边界情况
		{"[]uint 负索引", []uint{1, 2, 3}, -1, true, ErrInvalidSlice, "负索引应返回错误"},
		{"[]uint 超大索引", []uint{1, 2, 3}, 100, true, ErrInvalidSlice, "超大索引应返回错误"},
		{"空切片", []uint{}, 0, true, ErrInvalidSlice, "空切片应返回错误"},
		{"单元素切片", []uint{42}, 0, true, ErrInvalidSlice, "单元素切片应返回错误"},

		// 指针类型
		{"指针类型", &[]struct{}{}, 0, true, ErrInvalidSlice, "指针类型应返回错误"},

		// 接口切片（显式）
		{"[]interface{} 显式", []interface{}{"a", "b"}, 0, true, ErrInvalidSlice, "显式接口切片应返回错误"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessGenericSlice(tt.slice, tt.index)

			// 验证错误
			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
				assert.Nil(t, result, tt.description+" 结果应为 nil")
			} else {
				assert.NoError(t, err, tt.description)
				assert.NotNil(t, result, tt.description+" 结果不应为 nil")
			}
		})
	}
}

// TestAccessGenericSlice_ErrorMessages 测试错误消息格式
func TestAccessGenericSlice_ErrorMessages(t *testing.T) {
	tests := []struct {
		name     string
		slice    any
		index    int
		contains string
	}{
		{"错误消息包含类型信息", []uint{1, 2}, 0, "[]uint"},
		{"错误消息包含索引", []float32{1.1}, 1, "1"},
		{"nil 错误消息", nil, 0, "nil"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := accessGenericSlice(tt.slice, tt.index)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.contains)
		})
	}
}

// TestAccessGenericSlice_ReflectImpl 测试 Reflect 实现（如果替换）
func TestAccessGenericSlice_ReflectImpl(t *testing.T) {
	// 这些测试用于验证 reflect 实现的正确性
	// 当实现改为 reflect 时，取消注释这些测试

	tests := []struct {
		name        string
		slice       any
		index       int
		want        any
		wantErr     bool
		errType     error
		description string
	}{
		{"[]uint 有效访问", []uint{10, 20, 30}, 1, uint(20), false, nil, "应返回第二个元素"},
		{"[]float32 有效访问", []float32{1.1, 2.2, 3.3}, 2, float32(3.3), false, nil, "应返回第三个元素"},
		{"[]int32 有效访问", []int32{-5, 0, 5}, 0, int32(-5), false, nil, "应返回第一个元素"},
		{"[]uint64 有效访问", []uint64{100, 200}, 1, uint64(200), false, nil, "应返回第二个元素"},

		{"[]uint 越界", []uint{1, 2, 3}, 5, nil, true, ErrOutOfRange, "应返回越界错误"},
		{"[]uint 负索引", []uint{1, 2, 3}, -1, nil, true, ErrOutOfRange, "负索引应返回错误"},
		{"[]float32 空切片", []float32{}, 0, nil, true, ErrOutOfRange, "空切片应返回错误"},

		{"非切片类型", "string", 0, nil, true, ErrInvalidSlice, "字符串应返回错误"},
		{"nil 值", nil, 0, nil, true, ErrInvalidSlice, "nil 应返回错误"},
		{"整数类型", 42, 0, nil, true, ErrInvalidSlice, "整数应返回错误"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用 reflect 实现测试
			result, err := accessGenericSlice_ReflectValue(tt.slice, tt.index)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
				assert.Nil(t, result, tt.description+" 结果应为 nil")
			} else {
				assert.NoError(t, err, tt.description)
				assert.Equal(t, tt.want, result, tt.description+" 结果应匹配")
			}
		})
	}
}

// TestAccessGenericSlice_TypeAssertFirst 测试类型断言优先实现
func TestAccessGenericSlice_TypeAssertFirst(t *testing.T) {
	tests := []struct {
		name        string
		slice       any
		index       int
		want        any
		wantErr     bool
		errType     error
		description string
	}{
		{"[]uint fast path", []uint{1, 2, 3}, 1, uint(2), false, nil, "应命中 fast path"},
		{"[]float32 fast path", []float32{1.1, 2.2}, 0, float32(1.1), false, nil, "应命中 fast path"},
		{"[]int32 fast path", []int32{10, 20}, 1, int32(20), false, nil, "应命中 fast path"},

		{"[]uint 越界", []uint{1, 2, 3}, 10, nil, true, ErrOutOfRange, "应返回越界错误"},
		{"[]float32 空切片", []float32{}, 0, nil, true, ErrOutOfRange, "空切片应返回错误"},

		{"自定义类型 fallback", []int16{1, 2, 3}, 1, int16(2), false, nil, "应 fallback 到 reflect"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessGenericSlice_TypeAssertFirst(tt.slice, tt.index)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
			} else {
				assert.NoError(t, err, tt.description)
				assert.Equal(t, tt.want, result, tt.description+" 结果应匹配")
			}
		})
	}
}

// TestAccessGenericSlice_SimpleError 测试简化错误实现
func TestAccessGenericSlice_SimpleError(t *testing.T) {
	tests := []struct {
		name        string
		slice       any
		index       int
		want        any
		wantErr     bool
		errType     error
		description string
	}{
		{"[]uint 有效访问", []uint{10, 20}, 0, uint(10), false, nil, "应返回第一个元素"},
		{"[]float32 有效访问", []float32{1.1, 2.2}, 1, float32(2.2), false, nil, "应返回第二个元素"},

		{"非切片类型", "string", 0, nil, true, ErrInvalidSlice, "应返回错误（无详细信息）"},
		{"越界访问", []uint{1, 2}, 5, nil, true, ErrOutOfRange, "应返回越界错误"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessGenericSlice_SimpleError(tt.slice, tt.index)

			if tt.wantErr {
				assert.Error(t, err, tt.description)
				if tt.errType != nil {
					assert.ErrorIs(t, err, tt.errType, tt.description)
				}
			} else {
				assert.NoError(t, err, tt.description)
				assert.Equal(t, tt.want, result, tt.description)
			}
		})
	}
}

// TestAccessGenericSlice_EdgeCases 边界情况测试
func TestAccessGenericSlice_EdgeCases(t *testing.T) {
	t.Run("nil 切片", func(t *testing.T) {
		_, err := accessGenericSlice_ReflectValue(nil, 0)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidSlice)
	})

	t.Run("零长度切片", func(t *testing.T) {
		slice := []int{}
		_, err := accessGenericSlice_ReflectValue(slice, 0)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrOutOfRange)
	})

	t.Run("最大有效索引", func(t *testing.T) {
		slice := []uint{1, 2, 3, 4, 5}
		result, err := accessGenericSlice_ReflectValue(slice, 4)
		assert.NoError(t, err)
		assert.Equal(t, uint(5), result)
	})

	t.Run("索引刚好越界", func(t *testing.T) {
		slice := []uint{1, 2, 3}
		_, err := accessGenericSlice_ReflectValue(slice, 3)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrOutOfRange)
	})

	t.Run("大索引", func(t *testing.T) {
		slice := []uint{1}
		_, err := accessGenericSlice_ReflectValue(slice, 1000000)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrOutOfRange)
	})
}

// TestAccessGenericSlice_CustomTypes 自定义类型测试
func TestAccessGenericSlice_CustomTypes(t *testing.T) {
	type Point struct {
		X, Y int
	}

	type PointSlice []Point

	tests := []struct {
		name        string
		slice       any
		index       int
		want        any
		wantErr     bool
		description string
	}{
		{"自定义结构体切片", []Point{{1, 2}, {3, 4}}, 0, Point{1, 2}, false, "应返回结构体"},
		{"自定义类型切片", PointSlice{{5, 6}}, 0, Point{5, 6}, false, "应返回元素"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessGenericSlice_ReflectValue(tt.slice, tt.index)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result, tt.description)
			}
		})
	}
}

// TestAccessGenericSlice_SliceElementTypes 切片元素类型覆盖测试
func TestAccessGenericSlice_SliceElementTypes(t *testing.T) {
	tests := []struct {
		name        string
		slice       any
		index       int
		want        any
		description string
	}{
		{"[]int8", []int8{1, 2, 3}, 1, int8(2), "int8 元素"},
		{"[]int16", []int16{100, 200}, 0, int16(100), "int16 元素"},
		{"[]int32", []int32{1000, 2000}, 1, int32(2000), "int32 元素"},
		{"[]int64", []int64{10000, 20000}, 0, int64(10000), "int64 元素"},
		{"[]uint8", []uint8{10, 20}, 1, uint8(20), "uint8 元素"},
		{"[]uint16", []uint16{1000, 2000}, 0, uint16(1000), "uint16 元素"},
		{"[]uint32", []uint32{10000, 20000}, 1, uint32(20000), "uint32 元素"},
		{"[]uint64", []uint64{100000, 200000}, 0, uint64(100000), "uint64 元素"},
		{"[]float32", []float32{1.1, 2.2}, 0, float32(1.1), "float32 元素"},
		{"[]float64", []float64{1.11, 2.22}, 1, float64(2.22), "float64 元素"},
		{"[]bool", []bool{true, false}, 0, true, "bool 元素"},
		{"[]string", []string{"a", "b"}, 1, "b", "string 元素"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := accessGenericSlice_ReflectValue(tt.slice, tt.index)
			assert.NoError(t, err, tt.description)
			assert.Equal(t, tt.want, result, tt.description+" 值应匹配")
		})
	}
}
