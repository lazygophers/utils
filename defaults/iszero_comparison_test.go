package defaults

import (
	"fmt"
	"reflect"
	"testing"
)

// isZeroOld 原始版本（用于对比）
func isZeroOld(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

// TestIsZeroOptimization 验证优化版本的正确性和性能
func TestIsZeroOptimization(t *testing.T) {
	// 正确性验证
	testCases := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"空字符串", "", true},
		{"非空字符串", "hello", false},
		{"零整数", 0, true},
		{"非零整数", 42, false},
		{"零浮点", 0.0, true},
		{"非零浮点", 3.14, false},
		{"假布尔", false, true},
		{"真布尔", true, false},
		{"nil指针", (*int)(nil), true},
		{"非nil指针", new(int), false},
		{"nil接口", interface{}(nil), true},
		{"零值uint", uint(0), true},
		{"非零uint", uint(42), false},
		{"零值int8", int8(0), true},
		{"非零int8", int8(42), false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			v := reflect.ValueOf(tt.value)

			// 验证新版本
			newResult := isZero(v)
			if newResult != tt.expected {
				t.Errorf("isZero(%v) = %v, want %v", tt.value, newResult, tt.expected)
			}

			// 验证与旧版本结果一致
			oldResult := isZeroOld(v)
			if newResult != oldResult {
				t.Errorf("新旧版本结果不一致: 新=%v, 旧=%v", newResult, oldResult)
			}
		})
	}

	// 性能对比测试
	if testing.Short() {
		t.Skip("跳过性能测试（使用 -short 标志）")
	}

	iterations := 5000000

	fmt.Println("\nisZero 函数性能对比")
	fmt.Println("=================")

	perfCases := []struct {
		name  string
		value reflect.Value
	}{
		{"字符串", reflect.ValueOf("")},
		{"整数", reflect.ValueOf(0)},
		{"浮点数", reflect.ValueOf(0.0)},
		{"布尔值", reflect.ValueOf(false)},
		{"指针", reflect.ValueOf((*int)(nil))},
		{"接口", reflect.ValueOf(interface{}(nil))},
	}

	for _, tc := range perfCases {
		// 测试旧版本
		oldStart := testing.AllocsPerRun(iterations, func() {
			_ = isZeroOld(tc.value)
		})

		// 测试新版本
		newStart := testing.AllocsPerRun(iterations, func() {
			_ = isZero(tc.value)
		})

		fmt.Printf("%s: 旧版本=%.2f, 新版本=%.2f\n", tc.name, oldStart, newStart)
	}
}

// BenchmarkIsZeroOld 原始版本基准测试
func BenchmarkIsZeroOldString(b *testing.B) {
	v := reflect.ValueOf("")
	for i := 0; i < b.N; i++ {
		_ = isZeroOld(v)
	}
}

// BenchmarkIsZeroNew 新版本基准测试
func BenchmarkIsZeroNewString(b *testing.B) {
	v := reflect.ValueOf("")
	for i := 0; i < b.N; i++ {
		_ = isZero(v)
	}
}

func BenchmarkIsZeroOldInt(b *testing.B) {
	v := reflect.ValueOf(0)
	for i := 0; i < b.N; i++ {
		_ = isZeroOld(v)
	}
}

func BenchmarkIsZeroNewInt(b *testing.B) {
	v := reflect.ValueOf(0)
	for i := 0; i < b.N; i++ {
		_ = isZero(v)
	}
}

func BenchmarkIsZeroOldPtr(b *testing.B) {
	var p *int
	v := reflect.ValueOf(p)
	for i := 0; i < b.N; i++ {
		_ = isZeroOld(v)
	}
}

func BenchmarkIsZeroNewPtr(b *testing.B) {
	var p *int
	v := reflect.ValueOf(p)
	for i := 0; i < b.N; i++ {
		_ = isZero(v)
	}
}

func BenchmarkIsZeroOldBool(b *testing.B) {
	v := reflect.ValueOf(false)
	for i := 0; i < b.N; i++ {
		_ = isZeroOld(v)
	}
}

func BenchmarkIsZeroNewBool(b *testing.B) {
	v := reflect.ValueOf(false)
	for i := 0; i < b.N; i++ {
		_ = isZero(v)
	}
}

func BenchmarkIsZeroOldFloat(b *testing.B) {
	v := reflect.ValueOf(0.0)
	for i := 0; i < b.N; i++ {
		_ = isZeroOld(v)
	}
}

func BenchmarkIsZeroNewFloat(b *testing.B) {
	v := reflect.ValueOf(0.0)
	for i := 0; i < b.N; i++ {
		_ = isZero(v)
	}
}
