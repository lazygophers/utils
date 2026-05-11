package defaults

import (
	"testing"
	"time"
)

// 测试用的结构体类型

type BenchmarkSimple struct {
	Field string `default:"test"`
}

type BenchmarkNested struct {
	Value string `default:"nested"`
}

type BenchmarkComplex struct {
	StringField    string            `default:"test_string"`
	IntField       int               `default:"42"`
	UintField      uint              `default:"100"`
	FloatField     float64           `default:"3.14"`
	BoolField      bool              `default:"true"`
	SliceField     []int             `default:"[1,2,3,4,5]"`
	MapField       map[string]string `default:"{\"key\":\"value\"}"`
	TimeField      time.Time         `default:"2024-01-01"`
	PtrField       *string           `default:"ptr_value"`
	NestedField    BenchmarkNested
	InterfaceField interface{}
}

type BenchmarkVeryComplex struct {
	Field1  string   `default:"f1"`
	Field2  int      `default:"1"`
	Field3  float64  `default:"1.0"`
	Field4  bool     `default:"true"`
	Field5  []string `default:"[a,b,c]"`
	Field6  BenchmarkNested
	Field7  *BenchmarkNested
	Field8  map[string]int `default:"{\"x\":1}"`
	Field9  time.Time      `default:"2024-01-01"`
	Field10 []int          `default:"[1,2,3]"`
}

// Benchmark 原始实现

func BenchmarkOriginalSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s BenchmarkSimple
		SetDefaults(&s)
	}
}

func BenchmarkOriginalComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var c BenchmarkComplex
		SetDefaults(&c)
	}
}

func BenchmarkOriginalVeryComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v BenchmarkVeryComplex
		SetDefaults(&v)
	}
}

// Benchmark 优化实现

func BenchmarkOptimizedSimple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s BenchmarkSimple
		SetDefaultsOptimized(&s)
	}
}

func BenchmarkOptimizedComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var c BenchmarkComplex
		SetDefaultsOptimized(&c)
	}
}

func BenchmarkOptimizedVeryComplex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v BenchmarkVeryComplex
		SetDefaultsOptimized(&v)
	}
}

// 正确性验证

func TestOptimizedCorrectness(t *testing.T) {
	type TestStruct struct {
		StringField string  `default:"test"`
		IntField    int     `default:"42"`
		FloatField  float64 `default:"3.14"`
		BoolField   bool    `default:"true"`
		Nested      struct {
			Value string `default:"nested"`
		}
	}

	// 使用原始实现
	var original TestStruct
	SetDefaults(&original)

	// 使用优化实现
	var optimized TestStruct
	SetDefaultsOptimized(&optimized)

	// 验证结果一致
	if original.StringField != optimized.StringField {
		t.Errorf("StringField mismatch: original=%v, optimized=%v", original.StringField, optimized.StringField)
	}
	if original.IntField != optimized.IntField {
		t.Errorf("IntField mismatch: original=%v, optimized=%v", original.IntField, optimized.IntField)
	}
	if original.FloatField != optimized.FloatField {
		t.Errorf("FloatField mismatch: original=%v, optimized=%v", original.FloatField, optimized.FloatField)
	}
	if original.BoolField != optimized.BoolField {
		t.Errorf("BoolField mismatch: original=%v, optimized=%v", original.BoolField, optimized.BoolField)
	}
	if original.Nested.Value != optimized.Nested.Value {
		t.Errorf("Nested.Value mismatch: original=%v, optimized=%v", original.Nested.Value, optimized.Nested.Value)
	}
}

// 对比 Benchmark

func BenchmarkCompareSimple(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		var s BenchmarkSimple
		for i := 0; i < b.N; i++ {
			SetDefaults(&s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		var s BenchmarkSimple
		for i := 0; i < b.N; i++ {
			SetDefaultsOptimized(&s)
		}
	})
}

func BenchmarkCompareComplex(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		var c BenchmarkComplex
		for i := 0; i < b.N; i++ {
			SetDefaults(&c)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		var c BenchmarkComplex
		for i := 0; i < b.N; i++ {
			SetDefaultsOptimized(&c)
		}
	})
}

func BenchmarkCompareVeryComplex(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		var v BenchmarkVeryComplex
		for i := 0; i < b.N; i++ {
			SetDefaults(&v)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		var v BenchmarkVeryComplex
		for i := 0; i < b.N; i++ {
			SetDefaultsOptimized(&v)
		}
	})
}
