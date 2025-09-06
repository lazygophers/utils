// Package max 提供查找切片最大值的功能
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMax 测试 Max 函数的正确性
func TestMax(t *testing.T) {
	// 测试整数切片
	assert.Equal(t, 3, Max([]int{1, 2, 3}))
	assert.Equal(t, 1, Max([]int{1}))
	assert.Equal(t, 0, Max([]int{0, -1, -2}))
	assert.Equal(t, 9, Max([]int{9, 2, 5, 1, 8}))

	// 测试浮点数切片
	assert.Equal(t, 3.14, Max([]float64{1.1, 2.2, 3.14}))
	assert.Equal(t, 0.5, Max([]float64{-0.1, 0.5, -0.3}))
	assert.Equal(t, 1.0, Max([]float64{1.0}))

	// 测试字符串切片
	assert.Equal(t, "c", Max([]string{"a", "b", "c"}))
	assert.Equal(t, "z", Max([]string{"z", "a", "b"}))
	assert.Equal(t, "hello", Max([]string{"hello"}))

	// 测试空切片
	require.Equal(t, 0, Max([]int{}))
	require.Equal(t, 0.0, Max([]float64{}))
	require.Equal(t, "", Max([]string{}))

	// 测试单元素切片
	require.Equal(t, 42, Max([]int{42}))
	require.Equal(t, 3.14, Max([]float64{3.14}))
	require.Equal(t, "test", Max([]string{"test"}))

	// 测试负数
	require.Equal(t, -1, Max([]int{-5, -3, -1, -10}))
	require.Equal(t, -0.5, Max([]float64{-2.5, -1.5, -0.5, -3.0}))

	// 测试重复值
	require.Equal(t, 5, Max([]int{5, 5, 5, 5}))
	require.Equal(t, "b", Max([]string{"a", "b", "b", "a"}))
}

// BenchmarkMax 测试 Max 函数的性能
func BenchmarkMax(b *testing.B) {
	// 测试小整数切片性能
	smallInts := []int{1, 2, 3, 4, 5}
	b.Run("SmallIntSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Max(smallInts)
		}
	})

	// 测试大整数切片性能
	largeInts := make([]int, 1000)
	for i := range largeInts {
		largeInts[i] = i
	}
	b.Run("LargeIntSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Max(largeInts)
		}
	})

	// 测试字符串切片性能
	strings := make([]string, 1000)
	for i := range strings {
		strings[i] = string(rune(i % 1000))
	}
	b.Run("StringSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Max(strings)
		}
	})

	// 测试浮点数切片性能
	floats := make([]float64, 1000)
	for i := range floats {
		floats[i] = float64(i) * 0.1
	}
	b.Run("FloatSlice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Max(floats)
		}
	})
}

// TestMaxEdgeCases 测试 Max 函数的边界情况
func TestMaxEdgeCases(t *testing.T) {
	// 测试极大值
	require.Equal(t, 2147483647, Max([]int{2147483647, 0, -2147483648}))

	// 测试极小值
	require.Equal(t, -2147483648, Max([]int{-2147483648, -2147483647}))

	// 测试浮点数边界
	require.Equal(t, 1.7976931348623157e+308, Max([]float64{1.7976931348623157e+308, 0.0, -1.7976931348623157e+308}))

	// 测试空字符串
	require.Equal(t, "", Max([]string{""}))

	// 测试单字符字符串
	require.Equal(t, "z", Max([]string{"a", "z", "A", "Z"}))
}

// TestMaxNilSlice 测试 Max 函数对 nil 切片的处理
func TestMaxNilSlice(t *testing.T) {
	// 测试 nil 切片
	require.Equal(t, 0, Max([]int(nil)))
	require.Equal(t, 0.0, Max([]float64(nil)))
	require.Equal(t, "", Max([]string(nil)))
}
