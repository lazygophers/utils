package candy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUniqueUsing 测试 UniqueUsing 函数
func TestUniqueUsing(t *testing.T) {
	t.Parallel()

	// 测试空切片场景
	t.Run("空切片场景", func(t *testing.T) {
		t.Parallel()

		// 测试整型空切片
		t.Run("空整型切片", func(t *testing.T) {
			t.Parallel()
			give := []int{}
			f := func(n int) any {
				return n
			}
			want := []int{}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "空整型切片场景下 UniqueUsing() 应返回空切片")
		})

		// 测试字符串空切片
		t.Run("空字符串切片", func(t *testing.T) {
			t.Parallel()
			give := []string{}
			f := func(s string) any {
				return s
			}
			want := []string{}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "空字符串切片场景下 UniqueUsing() 应返回空切片")
		})
	})

	// 测试单元素切片场景
	t.Run("单元素切片场景", func(t *testing.T) {
		t.Parallel()

		// 测试整型单元素切片
		t.Run("单元素整型切片", func(t *testing.T) {
			t.Parallel()
			give := []int{42}
			f := func(n int) any {
				return n
			}
			want := []int{42}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "单元素整型切片场景下 UniqueUsing() 应返回原切片")
		})

		// 测试字符串单元素切片
		t.Run("单元素字符串切片", func(t *testing.T) {
			t.Parallel()
			give := []string{"hello"}
			f := func(s string) any {
				return s
			}
			want := []string{"hello"}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "单元素字符串切片场景下 UniqueUsing() 应返回原切片")
		})

		// 测试单元素零值
		t.Run("单元素零值", func(t *testing.T) {
			t.Parallel()
			give := []int{0}
			f := func(n int) any {
				return n
			}
			want := []int{0}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "单元素零值场景下 UniqueUsing() 应返回原切片")
		})
	})

	// 测试多元素无重复场景
	t.Run("多元素无重复场景", func(t *testing.T) {
		t.Parallel()

		// 测试整型无重复切片
		t.Run("无重复整型切片", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 2, 3, 4, 5}
			f := func(n int) any {
				return n
			}
			want := []int{1, 2, 3, 4, 5}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "无重复整型切片场景下 UniqueUsing() 应返回原切片")
		})

		// 测试字符串无重复切片
		t.Run("无重复字符串切片", func(t *testing.T) {
			t.Parallel()
			give := []string{"apple", "banana", "cherry"}
			f := func(s string) any {
				return s
			}
			want := []string{"apple", "banana", "cherry"}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "无重复字符串切片场景下 UniqueUsing() 应返回原切片")
		})
	})

	// 测试多元素有重复场景
	t.Run("多元素有重复场景", func(t *testing.T) {
		t.Parallel()

		// 测试整型有重复切片-重复元素在后
		t.Run("有重复整型切片-重复元素在后", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 2, 3, 2, 1}
			f := func(n int) any {
				return n
			}
			want := []int{1, 2, 3}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "有重复整型切片场景下 UniqueUsing() 应正确去重")
		})

		// 测试整型有重复切片-重复元素在前
		t.Run("有重复整型切片-重复元素在前", func(t *testing.T) {
			t.Parallel()
			give := []int{1, 1, 2, 3, 4}
			f := func(n int) any {
				return n
			}
			want := []int{1, 2, 3, 4}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "有重复整型切片场景下 UniqueUsing() 应正确去重")
		})

		// 测试字符串有重复切片
		t.Run("有重复字符串切片", func(t *testing.T) {
			t.Parallel()
			give := []string{"a", "b", "a", "c", "b"}
			f := func(s string) any {
				return s
			}
			want := []string{"a", "b", "c"}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "有重复字符串切片场景下 UniqueUsing() 应正确去重")
		})
	})

	// 测试所有元素相同场景
	t.Run("所有元素相同场景", func(t *testing.T) {
		t.Parallel()

		// 测试整型所有元素相同
		t.Run("所有元素相同-整型", func(t *testing.T) {
			t.Parallel()
			give := []int{5, 5, 5, 5, 5}
			f := func(n int) any {
				return n
			}
			want := []int{5}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "所有元素相同整型场景下 UniqueUsing() 应返回单元素切片")
		})

		// 测试字符串所有元素相同
		t.Run("所有元素相同-字符串", func(t *testing.T) {
			t.Parallel()
			give := []string{"test", "test", "test"}
			f := func(s string) any {
				return s
			}
			want := []string{"test"}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "所有元素相同字符串场景下 UniqueUsing() 应返回单元素切片")
		})
	})

	// 测试自定义去重函数场景
	t.Run("自定义去重函数场景", func(t *testing.T) {
		t.Parallel()

		// 定义结构体类型
		type Person struct {
			Name string
			Age  int
		}

		// 测试按姓名去重
		t.Run("按姓名去重", func(t *testing.T) {
			t.Parallel()
			give := []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Alice", 35}, // 重复姓名，年龄不同
			}
			f := func(p Person) any {
				return p.Name
			}
			want := []Person{
				{"Alice", 25}, // 保留第一个出现的 Alice
				{"Bob", 30},
			}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "按姓名去重场景下 UniqueUsing() 应正确去重")
		})

		// 测试按年龄去重
		t.Run("按年龄去重", func(t *testing.T) {
			t.Parallel()
			give := []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 25}, // 重复年龄，姓名不同
			}
			f := func(p Person) any {
				return p.Age
			}
			want := []Person{
				{"Alice", 25}, // 保留第一个出现的 25 岁
				{"Bob", 30},
			}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "按年龄去重场景下 UniqueUsing() 应正确去重")
		})

		// 测试按组合键去重
		t.Run("按组合键去重", func(t *testing.T) {
			t.Parallel()
			give := []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Alice", 25}, // 完全重复
				{"Bob", 35},   // 姓名重复，年龄不同
			}
			f := func(p Person) any {
				// 使用组合键：姓名+年龄
				return fmt.Sprintf("%s_%d", p.Name, p.Age)
			}
			want := []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Bob", 35}, // 保留，因为组合键不同
			}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "按组合键去重场景下 UniqueUsing() 应正确去重")
		})
	})

	// 测试复杂数据类型场景
	t.Run("复杂数据类型场景", func(t *testing.T) {
		t.Parallel()

		// 定义复杂数据类型
		type Product struct {
			ID    int
			Name  string
			Price float64
		}

		type Order struct {
			OrderID  string
			Product  Product
			Quantity int
		}

		// 测试按订单ID去重
		t.Run("按订单ID去重", func(t *testing.T) {
			t.Parallel()
			give := []Order{
				{"ORD001", Product{1, "Laptop", 999.99}, 1},
				{"ORD002", Product{2, "Mouse", 29.99}, 2},
				{"ORD001", Product{3, "Keyboard", 49.99}, 1}, // 重复订单ID，产品不同
			}
			f := func(o Order) any {
				return o.OrderID
			}
			want := []Order{
				{"ORD001", Product{1, "Laptop", 999.99}, 1}, // 保留第一个出现的 ORD001
				{"ORD002", Product{2, "Mouse", 29.99}, 2},
			}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "按订单ID去重场景下 UniqueUsing() 应正确去重")
		})

		// 测试按产品ID去重
		t.Run("按产品ID去重", func(t *testing.T) {
			t.Parallel()
			give := []Order{
				{"ORD001", Product{1, "Laptop", 999.99}, 1},
				{"ORD002", Product{2, "Mouse", 29.99}, 2},
				{"ORD003", Product{1, "Laptop", 899.99}, 1}, // 重复产品ID，价格不同
			}
			f := func(o Order) any {
				return o.Product.ID
			}
			want := []Order{
				{"ORD001", Product{1, "Laptop", 999.99}, 1}, // 保留第一个出现的产品ID
				{"ORD002", Product{2, "Mouse", 29.99}, 2},
			}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "按产品ID去重场景下 UniqueUsing() 应正确去重")
		})

		// 测试按价格区间去重
		t.Run("按价格区间去重", func(t *testing.T) {
			t.Parallel()
			give := []Order{
				{"ORD001", Product{1, "Laptop", 999.99}, 1},
				{"ORD002", Product{2, "Mouse", 29.99}, 2},
				{"ORD003", Product{3, "Keyboard", 59.99}, 1},
				{"ORD004", Product{4, "Monitor", 299.99}, 1},
			}
			f := func(o Order) any {
				// 按价格区间分组：<50, 50-200, >200
				if o.Product.Price < 50 {
					return "low"
				} else if o.Product.Price < 200 {
					return "medium"
				} else {
					return "high"
				}
			}
			want := []Order{
				{"ORD001", Product{1, "Laptop", 999.99}, 1},  // high
				{"ORD002", Product{2, "Mouse", 29.99}, 2},    // low
				{"ORD003", Product{3, "Keyboard", 59.99}, 1}, // medium
			}
			got := UniqueUsing(give, f)
			assert.Equal(t, want, got, "按价格区间去重场景下 UniqueUsing() 应正确去重")
		})
	})

	// 测试边界情况和性能
	t.Run("边界情况和性能", func(t *testing.T) {
		t.Parallel()

		// 测试nil切片
		t.Run("nil切片", func(t *testing.T) {
			t.Parallel()
			var nilSlice []int
			f := func(n int) any {
				return n
			}
			result := UniqueUsing(nilSlice, f)
			assert.Empty(t, result, "nil切片应返回空切片")
		})

		// 测试大切片性能
		t.Run("大切片性能", func(t *testing.T) {
			t.Parallel()
			largeSlice := make([]int, 1000)
			for i := 0; i < 1000; i++ {
				largeSlice[i] = i % 100 // 创建重复数据
			}

			f := func(n int) any {
				return n
			}
			result := UniqueUsing(largeSlice, f)
			assert.Len(t, result, 100, "大切片去重后应有100个唯一元素")
		})
	})

	// 测试保留原始顺序
	t.Run("保留原始顺序", func(t *testing.T) {
		t.Parallel()
		original := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
		f := func(n int) any {
			return n
		}
		result := UniqueUsing(original, f)
		expected := []int{3, 1, 4, 5, 9, 2, 6}
		assert.Equal(t, expected, result, "去重后应保留原始出现顺序")
	})

	// 测试不修改原切片
	t.Run("不修改原切片", func(t *testing.T) {
		t.Parallel()
		original := []int{1, 2, 2, 3}
		originalCopy := make([]int, len(original))
		copy(originalCopy, original)

		f := func(n int) any {
			return n
		}
		result := UniqueUsing(original, f)

		// 确保原切片未被修改
		assert.Equal(t, originalCopy, original, "原切片应保持不变")
		// 确保返回的是新切片
		assert.NotSame(t, &original[0], &result[0], "应返回新切片")
	})
}