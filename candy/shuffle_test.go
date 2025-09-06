package candy

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

// BenchmarkShuffle 基准测试 Shuffle 函数
func BenchmarkShuffle(b *testing.B) {
	// 基准测试小切片
	b.Run("小切片", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Shuffle(slice)
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
			Shuffle(slice)
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
			Shuffle(slice)
		}
	})

	// 基准测试空切片
	b.Run("空切片", func(b *testing.B) {
		slice := []int{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Shuffle(slice)
		}
	})

	// 基准测试单元素切片
	b.Run("单元素切片", func(b *testing.B) {
		slice := []int{42}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Shuffle(slice)
		}
	})
}
