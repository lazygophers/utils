// Package candy 提供了 Go 语言中常用的集合操作函数和工具方法
package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCollections 集合操作功能的综合测试
func TestCollections(t *testing.T) {
	// 元素访问组测试
	t.Run("ElementAccess", func(t *testing.T) {
		// First 函数测试
		t.Run("First", func(t *testing.T) {
			t.Run("non-empty int slice", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := First(input)
				assert.Equal(t, 1, result)
			})

			t.Run("empty int slice", func(t *testing.T) {
				input := []int{}
				result := First(input)
				assert.Equal(t, 0, result) // int 的零值
			})

			t.Run("single element", func(t *testing.T) {
				input := []string{"hello"}
				result := First(input)
				assert.Equal(t, "hello", result)
			})

			t.Run("empty string slice", func(t *testing.T) {
				input := []string{}
				result := First(input)
				assert.Equal(t, "", result) // string 的零值
			})

			t.Run("float64 slice", func(t *testing.T) {
				input := []float64{3.14, 2.71}
				result := First(input)
				assert.Equal(t, 3.14, result)
			})

			t.Run("bool slice", func(t *testing.T) {
				input := []bool{true, false}
				result := First(input)
				assert.Equal(t, true, result)
			})

			t.Run("empty bool slice", func(t *testing.T) {
				input := []bool{}
				result := First(input)
				assert.Equal(t, false, result) // bool 的零值
			})

			// FirstOr 函数测试
			t.Run("FirstOr", func(t *testing.T) {
				t.Run("non-empty int slice", func(t *testing.T) {
					input := []int{1, 2, 3}
					result := FirstOr(input, 99)
					assert.Equal(t, 1, result)
				})

				t.Run("empty int slice with default", func(t *testing.T) {
					input := []int{}
					result := FirstOr(input, 99)
					assert.Equal(t, 99, result)
				})

				t.Run("single element", func(t *testing.T) {
					input := []string{"hello"}
					result := FirstOr(input, "default")
					assert.Equal(t, "hello", result)
				})

				t.Run("empty string slice with default", func(t *testing.T) {
					input := []string{}
					result := FirstOr(input, "default")
					assert.Equal(t, "default", result)
				})

				t.Run("float64 slice", func(t *testing.T) {
					input := []float64{3.14, 2.71}
					result := FirstOr(input, 1.0)
					assert.Equal(t, 3.14, result)
				})

				t.Run("empty float64 slice", func(t *testing.T) {
					input := []float64{}
					result := FirstOr(input, 1.0)
					assert.Equal(t, 1.0, result)
				})

				t.Run("bool slice", func(t *testing.T) {
					input := []bool{true, false}
					result := FirstOr(input, false)
					assert.Equal(t, true, result)
				})

				t.Run("empty bool slice with default true", func(t *testing.T) {
					input := []bool{}
					result := FirstOr(input, true)
					assert.Equal(t, true, result)
				})
			})
		})

		// Last 函数测试
		t.Run("Last", func(t *testing.T) {
			t.Run("non-empty int slice", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Last(input)
				assert.Equal(t, 5, result)
			})

			t.Run("empty int slice", func(t *testing.T) {
				input := []int{}
				result := Last(input)
				assert.Equal(t, 0, result) // int 的零值
			})

			t.Run("single element", func(t *testing.T) {
				input := []string{"hello"}
				result := Last(input)
				assert.Equal(t, "hello", result)
			})

			t.Run("empty string slice", func(t *testing.T) {
				input := []string{}
				result := Last(input)
				assert.Equal(t, "", result) // string 的零值
			})

			t.Run("float64 slice", func(t *testing.T) {
				input := []float64{3.14, 2.71, 1.41}
				result := Last(input)
				assert.Equal(t, 1.41, result)
			})

			t.Run("bool slice", func(t *testing.T) {
				input := []bool{true, false, true}
				result := Last(input)
				assert.Equal(t, true, result)
			})

			t.Run("empty bool slice", func(t *testing.T) {
				input := []bool{}
				result := Last(input)
				assert.Equal(t, false, result) // bool 的零值
			})

			// LastOr 函数测试
			t.Run("LastOr", func(t *testing.T) {
				t.Run("non-empty int slice", func(t *testing.T) {
					input := []int{1, 2, 3, 4, 5}
					result := LastOr(input, 99)
					assert.Equal(t, 5, result)
				})

				t.Run("empty int slice with default", func(t *testing.T) {
					input := []int{}
					result := LastOr(input, 99)
					assert.Equal(t, 99, result)
				})

				t.Run("single element", func(t *testing.T) {
					input := []string{"hello"}
					result := LastOr(input, "default")
					assert.Equal(t, "hello", result)
				})

				t.Run("empty string slice with default", func(t *testing.T) {
					input := []string{}
					result := LastOr(input, "default")
					assert.Equal(t, "default", result)
				})

				t.Run("float64 slice", func(t *testing.T) {
					input := []float64{3.14, 2.71, 1.41}
					result := LastOr(input, 1.0)
					assert.Equal(t, 1.41, result)
				})

				t.Run("empty float64 slice", func(t *testing.T) {
					input := []float64{}
					result := LastOr(input, 1.0)
					assert.Equal(t, 1.0, result)
				})

				t.Run("bool slice", func(t *testing.T) {
					input := []bool{true, false, true}
					result := LastOr(input, false)
					assert.Equal(t, true, result)
				})

				t.Run("empty bool slice with default true", func(t *testing.T) {
					input := []bool{}
					result := LastOr(input, true)
					assert.Equal(t, true, result)
				})
			})
		})

		// Bottom 函数测试
		t.Run("Bottom", func(t *testing.T) {
			t.Run("basic int slice", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Bottom(input, 3)
				expected := []int{3, 4, 5}
				assert.Equal(t, expected, result)
			})

			t.Run("n equals slice length", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := Bottom(input, 3)
				expected := []int{1, 2, 3}
				assert.Equal(t, expected, result)
			})

			t.Run("n greater than slice length", func(t *testing.T) {
				input := []int{1, 2}
				result := Bottom(input, 5)
				expected := []int{1, 2}
				assert.Equal(t, expected, result)
			})

			t.Run("n is zero", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := Bottom(input, 0)
				expected := []int{}
				assert.Equal(t, expected, result)
			})

			t.Run("empty slice", func(t *testing.T) {
				input := []int{}
				result := Bottom(input, 3)
				expected := []int{}
				assert.Equal(t, expected, result)
			})

			t.Run("string slice", func(t *testing.T) {
				input := []string{"a", "b", "c", "d"}
				result := Bottom(input, 2)
				expected := []string{"c", "d"}
				assert.Equal(t, expected, result)
			})

			t.Run("single element", func(t *testing.T) {
				input := []int{42}
				result := Bottom(input, 1)
				expected := []int{42}
				assert.Equal(t, expected, result)
			})

			t.Run("negative n", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := Bottom(input, -1)
				expected := []int{}
				assert.Equal(t, expected, result)
			})
		})

		// Top 函数测试
		t.Run("Top", func(t *testing.T) {
			t.Run("basic int slice", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Top(input, 3)
				expected := []int{1, 2, 3}
				assert.Equal(t, expected, result)
			})

			t.Run("n equals slice length", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := Top(input, 3)
				expected := []int{1, 2, 3}
				assert.Equal(t, expected, result)
			})

			t.Run("n greater than slice length", func(t *testing.T) {
				input := []int{1, 2}
				result := Top(input, 5)
				expected := []int{1, 2}
				assert.Equal(t, expected, result)
			})

			t.Run("n is zero", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := Top(input, 0)
				expected := []int{}
				assert.Equal(t, expected, result)
			})

			t.Run("empty slice", func(t *testing.T) {
				input := []int{}
				result := Top(input, 3)
				expected := []int{}
				assert.Equal(t, expected, result)
			})

			t.Run("string slice", func(t *testing.T) {
				input := []string{"a", "b", "c", "d"}
				result := Top(input, 2)
				expected := []string{"a", "b"}
				assert.Equal(t, expected, result)
			})

			t.Run("single element", func(t *testing.T) {
				input := []int{42}
				result := Top(input, 1)
				expected := []int{42}
				assert.Equal(t, expected, result)
			})

			t.Run("negative n", func(t *testing.T) {
				input := []int{1, 2, 3}
				result := Top(input, -1)
				expected := []int{}
				assert.Equal(t, expected, result)
			})

			t.Run("modifying result doesn't affect original", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Top(input, 3)
				result[0] = 999
				assert.Equal(t, 1, input[0])    // 原切片不受影响
				assert.Equal(t, 999, result[0]) // 结果切片被修改
			})
		})

		// Index 函数测试
		t.Run("Index", func(t *testing.T) {
			t.Run("found at beginning", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Index(input, 1)
				assert.Equal(t, 0, result)
			})

			t.Run("found in middle", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Index(input, 3)
				assert.Equal(t, 2, result)
			})

			t.Run("found at end", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Index(input, 5)
				assert.Equal(t, 4, result)
			})

			t.Run("not found", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5}
				result := Index(input, 6)
				assert.Equal(t, -1, result)
			})

			t.Run("empty slice", func(t *testing.T) {
				input := []int{}
				result := Index(input, 1)
				assert.Equal(t, -1, result)
			})

			t.Run("single element found", func(t *testing.T) {
				input := []int{42}
				result := Index(input, 42)
				assert.Equal(t, 0, result)
			})

			t.Run("single element not found", func(t *testing.T) {
				input := []int{42}
				result := Index(input, 1)
				assert.Equal(t, -1, result)
			})

			t.Run("string slice found", func(t *testing.T) {
				input := []string{"apple", "banana", "cherry"}
				result := Index(input, "banana")
				assert.Equal(t, 1, result)
			})

			t.Run("string slice not found", func(t *testing.T) {
				input := []string{"apple", "banana", "cherry"}
				result := Index(input, "grape")
				assert.Equal(t, -1, result)
			})

			t.Run("duplicate elements - returns first", func(t *testing.T) {
				input := []int{1, 2, 3, 2, 5}
				result := Index(input, 2)
				assert.Equal(t, 1, result) // 返回第一个匹配的索引
			})

			t.Run("float64 slice", func(t *testing.T) {
				input := []float64{1.1, 2.2, 3.3}
				result := Index(input, 2.2)
				assert.Equal(t, 1, result)
			})
		})
	})

	// 过滤和搜索组测试
	t.Run("FilterAndSearch", func(t *testing.T) {
		// Filter 函数测试
		t.Run("Filter", func(t *testing.T) {
			// 定义测试结构体
			type Person struct {
				Name string
				Age  int
			}

			// 基本过滤测试
			t.Run("整数切片过滤偶数", func(t *testing.T) {
				input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
				result := Filter(input, func(n int) bool { return n%2 == 0 })
				expected := []int{2, 4, 6, 8, 10}
				assert.Equal(t, expected, result, "整数切片过滤结果不匹配")
			})

			t.Run("字符串切片过滤长度", func(t *testing.T) {
				input := []string{"apple", "banana", "cherry", "date", "elderberry"}
				result := Filter(input, func(s string) bool { return len(s) > 5 })
				expected := []string{"banana", "cherry", "elderberry"}
				assert.Equal(t, expected, result, "字符串切片过滤结果不匹配")
			})

			t.Run("结构体切片过滤年龄", func(t *testing.T) {
				input := []Person{
					{Name: "Alice", Age: 25},
					{Name: "Bob", Age: 17},
					{Name: "Charlie", Age: 30},
					{Name: "David", Age: 16},
				}
				result := Filter(input, func(p Person) bool { return p.Age > 18 })
				expected := []Person{
					{Name: "Alice", Age: 25},
					{Name: "Charlie", Age: 30},
				}
				assert.Equal(t, expected, result, "结构体切片过滤结果不匹配")
			})

			// 边界情况
			t.Run("空切片", func(t *testing.T) {
				result := Filter([]int{}, func(n int) bool { return n%2 == 0 })
				assert.Empty(t, result, "空切片过滤结果应该为空")
				assert.NotNil(t, result, "空切片过滤结果不应该为nil")
			})

			t.Run("所有元素都不满足条件", func(t *testing.T) {
				result := Filter([]int{1, 3, 5, 7, 9}, func(n int) bool { return n%2 == 0 })
				assert.Empty(t, result, "所有元素都不满足条件时结果应该为空")
				assert.NotNil(t, result, "结果不应该为nil")
			})

			t.Run("所有元素都满足条件", func(t *testing.T) {
				input := []int{2, 4, 6, 8, 10}
				result := Filter(input, func(n int) bool { return n%2 == 0 })
				assert.Equal(t, input, result, "所有元素都满足条件时结果应该等于输入")
			})
		})

		// FilterNot 函数测试
		t.Run("FilterNot", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Contains 函数测试
		t.Run("Contains", func(t *testing.T) {
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
		})

		// All 函数测试
		t.Run("All", func(t *testing.T) {
			t.Run("all elements satisfy condition", func(t *testing.T) {
				input := []int{2, 4, 6, 8, 10}
				result := All(input, func(n int) bool { return n%2 == 0 })
				assert.True(t, result)
			})

			t.Run("not all elements satisfy condition", func(t *testing.T) {
				input := []int{2, 4, 5, 8, 10}
				result := All(input, func(n int) bool { return n%2 == 0 })
				assert.False(t, result)
			})

			t.Run("empty slice", func(t *testing.T) {
				input := []int{}
				result := All(input, func(n int) bool { return n > 0 })
				assert.True(t, result) // vacuous truth
			})

			t.Run("single element true", func(t *testing.T) {
				input := []int{2}
				result := All(input, func(n int) bool { return n%2 == 0 })
				assert.True(t, result)
			})

			t.Run("single element false", func(t *testing.T) {
				input := []int{3}
				result := All(input, func(n int) bool { return n%2 == 0 })
				assert.False(t, result)
			})

			t.Run("strings all non-empty", func(t *testing.T) {
				input := []string{"hello", "world", "test"}
				result := All(input, func(s string) bool { return len(s) > 0 })
				assert.True(t, result)
			})

			t.Run("strings with empty", func(t *testing.T) {
				input := []string{"hello", "", "test"}
				result := All(input, func(s string) bool { return len(s) > 0 })
				assert.False(t, result)
			})
		})

		// Any 函数测试
		t.Run("Any", func(t *testing.T) {
			t.Run("some elements satisfy condition", func(t *testing.T) {
				input := []int{1, 3, 4, 7, 9}
				result := Any(input, func(n int) bool { return n%2 == 0 })
				assert.True(t, result)
			})

			t.Run("no elements satisfy condition", func(t *testing.T) {
				input := []int{1, 3, 5, 7, 9}
				result := Any(input, func(n int) bool { return n%2 == 0 })
				assert.False(t, result)
			})

			t.Run("empty slice", func(t *testing.T) {
				input := []int{}
				result := Any(input, func(n int) bool { return n > 0 })
				assert.False(t, result)
			})

			t.Run("single element true", func(t *testing.T) {
				input := []int{2}
				result := Any(input, func(n int) bool { return n%2 == 0 })
				assert.True(t, result)
			})

			t.Run("single element false", func(t *testing.T) {
				input := []int{3}
				result := Any(input, func(n int) bool { return n%2 == 0 })
				assert.False(t, result)
			})

			t.Run("strings any empty", func(t *testing.T) {
				input := []string{"hello", "world", ""}
				result := Any(input, func(s string) bool { return len(s) == 0 })
				assert.True(t, result)
			})

			t.Run("strings none empty", func(t *testing.T) {
				input := []string{"hello", "world", "test"}
				result := Any(input, func(s string) bool { return len(s) == 0 })
				assert.False(t, result)
			})
		})

		// Same 函数测试
		t.Run("Same", func(t *testing.T) {
			// 测试内容将在这里添加
		})
	})

	// 变换操作组测试
	t.Run("Transformations", func(t *testing.T) {
		// Sort 函数测试
		t.Run("Sort", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// SortUsing 函数测试
		t.Run("SortUsing", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Reverse 函数测试
		t.Run("Reverse", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Shuffle 函数测试
		t.Run("Shuffle", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Unique 函数测试
		t.Run("Unique", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// UniqueUsing 函数测试
		t.Run("UniqueUsing", func(t *testing.T) {
			// 测试内容将在这里添加
		})
	})

	// 切片操作组测试
	t.Run("SliceOperations", func(t *testing.T) {
		// Chunk 函数测试
		t.Run("Chunk", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Drop 函数测试
		t.Run("Drop", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// RemoveIndex 函数测试
		t.Run("RemoveIndex", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Remove 函数测试
		t.Run("Remove", func(t *testing.T) {
			t.Run("basic int slices", func(t *testing.T) {
				ss := []int{1, 2, 3, 4, 5}
				toRemove := []int{2, 4, 6}
				result := Remove(ss, toRemove)
				expected := []int{1, 3, 5}
				assert.Equal(t, expected, result)
			})

			t.Run("remove all elements", func(t *testing.T) {
				ss := []int{1, 2, 3}
				toRemove := []int{1, 2, 3}
				result := Remove(ss, toRemove)
				assert.Empty(t, result)
			})

			t.Run("remove no elements", func(t *testing.T) {
				ss := []int{1, 2, 3}
				toRemove := []int{4, 5, 6}
				result := Remove(ss, toRemove)
				assert.Equal(t, ss, result)
			})

			t.Run("empty source slice", func(t *testing.T) {
				ss := []int{}
				toRemove := []int{1, 2, 3}
				result := Remove(ss, toRemove)
				assert.Empty(t, result)
			})

			t.Run("empty remove slice", func(t *testing.T) {
				ss := []int{1, 2, 3}
				toRemove := []int{}
				result := Remove(ss, toRemove)
				assert.Equal(t, ss, result)
			})

			t.Run("both empty slices", func(t *testing.T) {
				ss := []int{}
				toRemove := []int{}
				result := Remove(ss, toRemove)
				assert.Empty(t, result)
			})

			t.Run("string slices", func(t *testing.T) {
				ss := []string{"apple", "banana", "cherry", "date"}
				toRemove := []string{"banana", "date", "elderberry"}
				result := Remove(ss, toRemove)
				expected := []string{"apple", "cherry"}
				assert.Equal(t, expected, result)
			})

			t.Run("float64 slices", func(t *testing.T) {
				ss := []float64{1.1, 2.2, 3.3, 4.4}
				toRemove := []float64{2.2, 4.4, 5.5}
				result := Remove(ss, toRemove)
				expected := []float64{1.1, 3.3}
				assert.Equal(t, expected, result)
			})

			t.Run("duplicates in source", func(t *testing.T) {
				ss := []int{1, 2, 2, 3, 3, 3}
				toRemove := []int{2, 3}
				result := Remove(ss, toRemove)
				expected := []int{1}
				assert.Equal(t, expected, result)
			})

			t.Run("duplicates in remove list", func(t *testing.T) {
				ss := []int{1, 2, 3, 4, 5}
				toRemove := []int{2, 2, 4, 4}
				result := Remove(ss, toRemove)
				expected := []int{1, 3, 5}
				assert.Equal(t, expected, result)
			})

			t.Run("single element removal", func(t *testing.T) {
				ss := []int{1}
				toRemove := []int{1}
				result := Remove(ss, toRemove)
				assert.Empty(t, result)
			})

			t.Run("single element no removal", func(t *testing.T) {
				ss := []int{1}
				toRemove := []int{2}
				result := Remove(ss, toRemove)
				assert.Equal(t, []int{1}, result)
			})
		})

		// RemoveSlice 函数测试
		t.Run("RemoveSlice", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Diff 函数测试
		t.Run("Diff", func(t *testing.T) {
			t.Run("basic int slices", func(t *testing.T) {
				ss := []int{1, 2, 3}
				against := []int{2, 3, 4}
				added, removed := Diff(ss, against)
				assert.Equal(t, []int{4}, added)
				assert.Equal(t, []int{1}, removed)
			})

			t.Run("identical slices", func(t *testing.T) {
				ss := []int{1, 2, 3}
				against := []int{1, 2, 3}
				added, removed := Diff(ss, against)
				assert.Empty(t, added)
				assert.Empty(t, removed)
			})

			t.Run("completely different slices", func(t *testing.T) {
				ss := []int{1, 2, 3}
				against := []int{4, 5, 6}
				added, removed := Diff(ss, against)
				assert.Equal(t, []int{4, 5, 6}, added)
				assert.Equal(t, []int{1, 2, 3}, removed)
			})

			t.Run("empty slices", func(t *testing.T) {
				ss := []int{}
				against := []int{}
				added, removed := Diff(ss, against)
				assert.Empty(t, added)
				assert.Empty(t, removed)
			})

			t.Run("one empty slice", func(t *testing.T) {
				ss := []int{1, 2, 3}
				against := []int{}
				added, removed := Diff(ss, against)
				assert.Empty(t, added)
				assert.Equal(t, []int{1, 2, 3}, removed)
			})

			t.Run("against empty slice", func(t *testing.T) {
				ss := []int{}
				against := []int{1, 2, 3}
				added, removed := Diff(ss, against)
				assert.Equal(t, []int{1, 2, 3}, added)
				assert.Empty(t, removed)
			})

			t.Run("string slices", func(t *testing.T) {
				ss := []string{"apple", "banana", "cherry"}
				against := []string{"banana", "cherry", "date"}
				added, removed := Diff(ss, against)
				assert.Equal(t, []string{"date"}, added)
				assert.Equal(t, []string{"apple"}, removed)
			})

			t.Run("float64 slices", func(t *testing.T) {
				ss := []float64{1.1, 2.2, 3.3}
				against := []float64{2.2, 3.3, 4.4}
				added, removed := Diff(ss, against)
				assert.Equal(t, []float64{4.4}, added)
				assert.Equal(t, []float64{1.1}, removed)
			})

			t.Run("duplicates in slices", func(t *testing.T) {
				ss := []int{1, 2, 2, 3}
				against := []int{2, 3, 3, 4}
				added, removed := Diff(ss, against)
				assert.Equal(t, []int{4}, added)
				assert.Equal(t, []int{1}, removed)
			})

			t.Run("subset relationship", func(t *testing.T) {
				ss := []int{1, 2}
				against := []int{1, 2, 3, 4}
				added, removed := Diff(ss, against)
				assert.Equal(t, []int{3, 4}, added)
				assert.Empty(t, removed)
			})

			t.Run("superset relationship", func(t *testing.T) {
				ss := []int{1, 2, 3, 4}
				against := []int{1, 2}
				added, removed := Diff(ss, against)
				assert.Empty(t, added)
				assert.Equal(t, []int{3, 4}, removed)
			})

			t.Run("single element difference", func(t *testing.T) {
				ss := []int{1}
				against := []int{2}
				added, removed := Diff(ss, against)
				assert.Equal(t, []int{2}, added)
				assert.Equal(t, []int{1}, removed)
			})

			t.Run("byte slices", func(t *testing.T) {
				ss := []byte{65, 66, 67} // A, B, C
				against := []byte{66, 67, 68} // B, C, D
				added, removed := Diff(ss, against)
				assert.Equal(t, []byte{68}, added)
				assert.Equal(t, []byte{65}, removed)
			})
		})

		// DiffSlice 函数测试
		t.Run("DiffSlice", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// SliceEqual 函数测试
		t.Run("SliceEqual", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Spare 函数测试
		t.Run("Spare", func(t *testing.T) {
			// 测试内容将在这里添加
		})
	})

	// 数据提取组测试
	t.Run("DataExtraction", func(t *testing.T) {
		// Pluck 相关函数测试
		t.Run("Pluck", func(t *testing.T) {
			// 测试内容将在这里添加
		})
	})

	// 遍历操作组测试
	t.Run("Iteration", func(t *testing.T) {
		// Each 函数测试
		t.Run("Each", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// EachReverse 函数测试
		t.Run("EachReverse", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// EachStopWithError 函数测试
		t.Run("EachStopWithError", func(t *testing.T) {
			// 测试内容将在这里添加
		})
	})

	// 其他工具组测试
	t.Run("Utilities", func(t *testing.T) {
		// Join 函数测试
		t.Run("Join", func(t *testing.T) {
			// 测试内容将在这里添加
		})

		// Reduce 函数测试
		t.Run("Reduce", func(t *testing.T) {
			// 测试内容将在这里添加
		})
	})
}

// BenchmarkCollections 集合操作的基准测试
func BenchmarkCollections(b *testing.B) {
	// 基准测试内容将在这里添加
}

