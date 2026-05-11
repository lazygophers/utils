package validator

import (
	"reflect"
	"testing"
)

// 测试数据生成
type testStruct struct {
	Name     string `validate:"required"`
	Email    string `validate:"email"`
	Age      int    `validate:"min=18,max=100"`
	Username string `validate:"alpha"`
	Code     string `validate:"alphanum"`
}

func genFieldLevel() FieldLevel {
	v := testStruct{
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      25,
		Username: "johndoe",
		Code:     "ABC123",
	}
	rv := reflect.ValueOf(v)
	fl := &fieldLevel{
		top:             rv,
		parent:          rv,
		field:           rv.FieldByName("Email"),
		fieldName:       "Email",
		structFieldName: "Email",
		param:           "",
		structField:     reflect.TypeOf(v).Field(1),
	}
	return fl
}

// 常用标签列表（按使用频率排序）
var commonTags = []string{
	"required", "email", "min", "max", "len",
	"alpha", "alphanum", "url", "numeric", "eq",
	"ne", "eqfield", "nefield", "required_if",
	"required_with", "required_without",
}

// ===== 当前实现 (Baseline) =====

func BenchmarkValidateField_Current(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField(fl, tag)
		}
	}
}

func BenchmarkValidateField_Current_Single(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.validateField(fl, "email")
	}
}

// ===== 方案2: 内联 map 查找 =====

func BenchmarkValidateField_Opt2_InlineMap(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField_Opt2_InlineMap(fl, tag)
		}
	}
}

// ===== 方案3: 单次查找 =====

func BenchmarkValidateField_Opt3_SingleLookup(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField_Opt3_SingleLookup(fl, tag)
		}
	}
}

// ===== 方案5: 热路径 switch =====

func BenchmarkValidateField_Opt5_HotPathSwitch(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField_Opt5_HotPathSwitch(fl, tag)
		}
	}
}

// ===== 方案6: 完整 switch =====

func BenchmarkValidateField_Opt6_FullSwitch(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField_Opt6_FullSwitch(fl, tag)
		}
	}
}

// ===== 方案11: 内联验证器函数 =====

func BenchmarkValidateField_Opt11_InlinedValidators(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField_Opt11_InlinedValidators(fl, tag)
		}
	}
}

// ===== 方案13: goto 优化 =====

func BenchmarkValidateField_Opt13_GotoOptimized(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tag := range commonTags {
			_ = e.validateField_Opt13_GotoOptimized(fl, tag)
		}
	}
}

// ===== 内存分配基准 =====

func BenchmarkValidateField_Alloc_Current(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.validateField(fl, "email")
	}
}

func BenchmarkValidateField_Alloc_Opt5_Switch(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.validateField_Opt5_HotPathSwitch(fl, "email")
	}
}

func BenchmarkValidateField_Alloc_Opt11_Inlined(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.validateField_Opt11_InlinedValidators(fl, "email")
	}
}

// ===== 并行基准 =====

func BenchmarkValidateField_Parallel_Current(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = e.validateField(fl, "email")
		}
	})
}

func BenchmarkValidateField_Parallel_Opt5_Switch(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = e.validateField_Opt5_HotPathSwitch(fl, "email")
		}
	})
}

func BenchmarkValidateField_Parallel_Opt11_Inlined(b *testing.B) {
	e := NewEngine()
	fl := genFieldLevel()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = e.validateField_Opt11_InlinedValidators(fl, "email")
		}
	})
}
