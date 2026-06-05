package candy

import (
	"testing"
)

// Benchmark ToPtr 函数的各种实现方案
// ToPtr 非常简单，主要是测试不同的语法变体是否有性能差异

// 版本1: 当前实现
func toPtrV1[T any](v T) *T {
	return &v
}

// 版本2: 使用临时变量（理论上无区别，但测试一下）
func toPtrV2[T any](v T) *T {
	result := v
	return &result
}

// 版本3: 使用 new + 赋值
func toPtrV3[T any](v T) *T {
	p := new(T)
	*p = v
	return p
}

// Benchmark 测试
func BenchmarkToPtr_Int(b *testing.B) {
	b.Run("V1_Current", func(b *testing.B) {
		val := 42
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV1(val)
		}
	})

	b.Run("V2_TempVar", func(b *testing.B) {
		val := 42
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV2(val)
		}
	})

	b.Run("V3_NewAssign", func(b *testing.B) {
		val := 42
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV3(val)
		}
	})
}

func BenchmarkToPtr_String(b *testing.B) {
	b.Run("V1_Current", func(b *testing.B) {
		val := "test string"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV1(val)
		}
	})

	b.Run("V2_TempVar", func(b *testing.B) {
		val := "test string"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV2(val)
		}
	})

	b.Run("V3_NewAssign", func(b *testing.B) {
		val := "test string"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV3(val)
		}
	})
}

func BenchmarkToPtr_Struct(b *testing.B) {
	type testStruct struct {
		a, b, c int
		d, e, f string
	}

	val := testStruct{a: 1, b: 2, c: 3, d: "test", e: "data", f: "value"}

	b.Run("V1_Current", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV1(val)
		}
	})

	b.Run("V2_TempVar", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV2(val)
		}
	})

	b.Run("V3_NewAssign", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = toPtrV3(val)
		}
	})
}
