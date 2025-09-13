// Package candy 提供了实用的工具函数和语法糖
package candy

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestEachStopWithError 测试 EachStopWithError 函数的各种场景
func TestEachStopWithError(t *testing.T) {
	// 定义测试用的错误
	testError := errors.New("测试错误")
	anotherError := errors.New("另一个错误")

	tests := []struct {
		name     string
		input    []int
		fn       func(int) error
		wantErr  error
		wantCall int // 期望函数被调用的次数
	}{
		{
			name:     "空切片",
			input:    []int{},
			fn:       func(n int) error { return nil },
			wantErr:  nil,
			wantCall: 0,
		},
		{
			name:     "正常遍历无错误",
			input:    []int{1, 2, 3, 4, 5},
			fn:       func(n int) error { return nil },
			wantErr:  nil,
			wantCall: 5,
		},
		{
			name:  "第一个元素就返回错误",
			input: []int{1, 2, 3, 4, 5},
			fn: func(n int) error {
				if n == 1 {
					return testError
				}
				return nil
			},
			wantErr:  testError,
			wantCall: 1,
		},
		{
			name:  "中间元素返回错误",
			input: []int{1, 2, 3, 4, 5},
			fn: func(n int) error {
				if n == 3 {
					return testError
				}
				return nil
			},
			wantErr:  testError,
			wantCall: 3,
		},
		{
			name:  "最后一个元素返回错误",
			input: []int{1, 2, 3, 4, 5},
			fn: func(n int) error {
				if n == 5 {
					return testError
				}
				return nil
			},
			wantErr:  testError,
			wantCall: 5,
		},
		{
			name:     "单元素切片无错误",
			input:    []int{42},
			fn:       func(n int) error { return nil },
			wantErr:  nil,
			wantCall: 1,
		},
		{
			name:     "单元素切片有错误",
			input:    []int{42},
			fn:       func(n int) error { return testError },
			wantErr:  testError,
			wantCall: 1,
		},
		{
			name:  "验证错误传递",
			input: []int{1, 2, 3},
			fn: func(n int) error {
				if n == 2 {
					return anotherError
				}
				return nil
			},
			wantErr:  anotherError,
			wantCall: 2,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态条件
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // 启用并行测试

			// 记录函数调用次数
			callCount := 0
			wrappedFn := func(n int) error {
				callCount++
				return tt.fn(n)
			}

			// 执行测试
			err := EachStopWithError(tt.input, wrappedFn)

			// 验证结果
			if tt.wantErr != nil {
				require.Error(t, err, "期望返回错误")
				assert.Equal(t, tt.wantErr, err, "返回的错误应该与期望一致")
			} else {
				require.NoError(t, err, "期望返回 nil")
			}

			// 验证函数调用次数
			assert.Equal(t, tt.wantCall, callCount, "函数调用次数应该与期望一致")
		})
	}
}

// TestEachStopWithErrorString 测试 EachStopWithError 函数处理字符串切片
func TestEachStopWithErrorString(t *testing.T) {
	testError := errors.New("字符串处理错误")

	tests := []struct {
		name     string
		input    []string
		fn       func(string) error
		wantErr  error
		wantCall int
	}{
		{
			name:     "字符串切片处理",
			input:    []string{"hello", "world", "go"},
			fn:       func(s string) error { return nil },
			wantErr:  nil,
			wantCall: 3,
		},
		{
			name:  "字符串处理遇到错误",
			input: []string{"hello", "world", "go"},
			fn: func(s string) error {
				if s == "world" {
					return testError
				}
				return nil
			},
			wantErr:  testError,
			wantCall: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			callCount := 0
			wrappedFn := func(s string) error {
				callCount++
				return tt.fn(s)
			}

			err := EachStopWithError(tt.input, wrappedFn)

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.wantCall, callCount)
		})
	}
}

// TestEachStopWithErrorStruct 测试 EachStopWithError 函数处理结构体切片
func TestEachStopWithErrorStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testError := errors.New("年龄验证错误")

	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 17}, // 未成年
		{Name: "Charlie", Age: 30},
	}

	callCount := 0
	fn := func(p Person) error {
		callCount++
		if p.Age < 18 {
			return testError
		}
		return nil
	}

	err := EachStopWithError(people, fn)

	require.Error(t, err, "未成年应该返回错误")
	assert.Equal(t, testError, err, "错误应该与期望一致")
	assert.Equal(t, 2, callCount, "应该只处理前两个人")
}

// TestEachStopWithErrorNilSlice 测试 EachStopWithError 函数处理 nil 切片
func TestEachStopWithErrorNilSlice(t *testing.T) {
	var nilSlice []int
	callCount := 0

	err := EachStopWithError(nilSlice, func(n int) error {
		callCount++
		return nil
	})

	require.NoError(t, err, "nil 切片应该不返回错误")
	assert.Equal(t, 0, callCount, "nil 切片不应该调用函数")
}

// BenchmarkEachStopWithError 基准测试 EachStopWithError 函数的性能
func BenchmarkEachStopWithError(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	// 无错误情况的基准测试
	b.Run("NoError", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			err := EachStopWithError(data, func(n int) error {
				return nil
			})
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	// 中间有错误情况的基准测试
	b.Run("ErrorInMiddle", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			err := EachStopWithError(data, func(n int) error {
				if n == 500 {
					return errors.New("基准测试错误")
				}
				return nil
			})
			if err == nil {
				b.Fatal("期望错误但没有得到")
			}
		}
	})

	// 第一个元素就错误的基准测试
	b.Run("ErrorFirst", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			err := EachStopWithError(data, func(n int) error {
				return errors.New("第一个错误")
			})
			if err == nil {
				b.Fatal("期望错误但没有得到")
			}
		}
	})
}

// BenchmarkEachStopWithErrorDifferentSizes 基准测试不同数据大小的性能
func BenchmarkEachStopWithErrorDifferentSizes(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		data := make([]int, size)
		for i := range data {
			data[i] = i
		}

		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := EachStopWithError(data, func(n int) error {
					return nil
				})
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
