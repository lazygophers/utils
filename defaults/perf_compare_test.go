package defaults

import (
	"reflect"
	"testing"
	"time"
)

type PerfTestSimple struct {
	Field string `default:"test"`
}

type PerfTestComplex struct {
	StringField string            `default:"test_string"`
	IntField    int               `default:"42"`
	UintField   uint              `default:"100"`
	FloatField  float64           `default:"3.14"`
	BoolField   bool              `default:"true"`
	SliceField  []int             `default:"[1,2,3,4,5]"`
	MapField    map[string]string `default:"{\"key\":\"value\"}"`
	TimeField   time.Time         `default:"2024-01-01"`
	PtrField    *string           `default:"ptr_value"`
	NestedField struct {
		Value string `default:"nested"`
	}
	InterfaceField interface{}
}

// 性能对比测试
func TestPerformanceCompare(t *testing.T) {
	iterations := 10000

	// 测试原始实现性能
	t.Run("Original_Simple", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			var s PerfTestSimple
			setDefaultWithOptions(reflect.ValueOf(&s), "", defaultOptions)
		}
		duration := time.Since(start)
		t.Logf("Original implementation (Simple): %v for %d iterations (%.2f ns/op)",
			duration, iterations, float64(duration.Nanoseconds())/float64(iterations))
	})

	t.Run("Optimized_Simple", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			var s PerfTestSimple
			setDefaultOptimized(reflect.ValueOf(&s), "", defaultOptions)
		}
		duration := time.Since(start)
		t.Logf("Optimized implementation (Simple): %v for %d iterations (%.2f ns/op)",
			duration, iterations, float64(duration.Nanoseconds())/float64(iterations))
	})

	// 测试复杂结构体性能
	t.Run("Original_Complex", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			var c PerfTestComplex
			setDefaultWithOptions(reflect.ValueOf(&c), "", defaultOptions)
		}
		duration := time.Since(start)
		t.Logf("Original implementation (Complex): %v for %d iterations (%.2f ns/op)",
			duration, iterations, float64(duration.Nanoseconds())/float64(iterations))
	})

	t.Run("Optimized_Complex", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			var c PerfTestComplex
			setDefaultOptimized(reflect.ValueOf(&c), "", defaultOptions)
		}
		duration := time.Since(start)
		t.Logf("Optimized implementation (Complex): %v for %d iterations (%.2f ns/op)",
			duration, iterations, float64(duration.Nanoseconds())/float64(iterations))
	})
}

// 正确性验证
func TestOptimizedCorrectnessDetailed(t *testing.T) {
	type ComplexTest struct {
		StringField string            `default:"test"`
		IntField    int               `default:"42"`
		UintField   uint              `default:"100"`
		FloatField  float64           `default:"3.14"`
		BoolField   bool              `default:"true"`
		SliceField  []int             `default:"[1,2,3]"`
		MapField    map[string]string `default:"{\"key\":\"value\"}"`
		Nested      struct {
			Value string `default:"nested"`
		}
	}

	// 测试原始实现
	var original ComplexTest
	err := setDefaultWithOptions(reflect.ValueOf(&original), "", defaultOptions)
	if err != nil {
		t.Fatalf("Original implementation failed: %v", err)
	}

	// 测试优化实现
	var optimized ComplexTest
	err = setDefaultOptimized(reflect.ValueOf(&optimized), "", defaultOptions)
	if err != nil {
		t.Fatalf("Optimized implementation failed: %v", err)
	}

	// 验证字符串字段
	if original.StringField != optimized.StringField {
		t.Errorf("StringField mismatch: original=%v, optimized=%v", original.StringField, optimized.StringField)
	}

	// 验证整数字段
	if original.IntField != optimized.IntField {
		t.Errorf("IntField mismatch: original=%v, optimized=%v", original.IntField, optimized.IntField)
	}

	// 验证无符号整数字段
	if original.UintField != optimized.UintField {
		t.Errorf("UintField mismatch: original=%v, optimized=%v", original.UintField, optimized.UintField)
	}

	// 验证浮点数字段
	if original.FloatField != optimized.FloatField {
		t.Errorf("FloatField mismatch: original=%v, optimized=%v", original.FloatField, optimized.FloatField)
	}

	// 验证布尔字段
	if original.BoolField != optimized.BoolField {
		t.Errorf("BoolField mismatch: original=%v, optimized=%v", original.BoolField, optimized.BoolField)
	}

	// 验证切片字段
	if len(original.SliceField) != len(optimized.SliceField) {
		t.Errorf("SliceField length mismatch: original=%v, optimized=%v",
			len(original.SliceField), len(optimized.SliceField))
	} else {
		for i := range original.SliceField {
			if original.SliceField[i] != optimized.SliceField[i] {
				t.Errorf("SliceField[%d] mismatch: original=%v, optimized=%v",
					i, original.SliceField[i], optimized.SliceField[i])
			}
		}
	}

	// 验证映射字段
	if len(original.MapField) != len(optimized.MapField) {
		t.Errorf("MapField length mismatch: original=%v, optimized=%v",
			len(original.MapField), len(optimized.MapField))
	} else {
		for k, v := range original.MapField {
			if optimized.MapField[k] != v {
				t.Errorf("MapField[%s] mismatch: original=%v, optimized=%v",
					k, v, optimized.MapField[k])
			}
		}
	}

	// 验证嵌套结构
	if original.Nested.Value != optimized.Nested.Value {
		t.Errorf("Nested.Value mismatch: original=%v, optimized=%v",
			original.Nested.Value, optimized.Nested.Value)
	}

	t.Log("All correctness checks passed!")
}

// 基准测试 - 简单结构体
func BenchmarkPerfSimple_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s PerfTestSimple
		_ = setDefaultWithOptions(reflect.ValueOf(&s), "", defaultOptions)
	}
}

func BenchmarkPerfSimple_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s PerfTestSimple
		_ = setDefaultOptimized(reflect.ValueOf(&s), "", defaultOptions)
	}
}

// 基准测试 - 复杂结构体
func BenchmarkPerfComplex_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c PerfTestComplex
		_ = setDefaultWithOptions(reflect.ValueOf(&c), "", defaultOptions)
	}
}

func BenchmarkPerfComplex_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var c PerfTestComplex
		_ = setDefaultOptimized(reflect.ValueOf(&c), "", defaultOptions)
	}
}
