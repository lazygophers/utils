package validator

import (
	"reflect"
	"testing"
)

// 方案1: 预计算整数比较
func Range1(min, max float64) ValidatorFunc {
	minInt := int64(min)
	maxInt := int64(max)
	minUint := uint64(min)
	maxUint := uint64(max)
	isIntRange := float64(minInt) == min && float64(maxInt) == max
	isUintRange := float64(minUint) == min && float64(maxUint) == max

	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if isIntRange {
				val := field.Int()
				return val >= minInt && val <= maxInt
			}
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if isUintRange {
				val := field.Uint()
				return val >= minUint && val <= maxUint
			}
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Float32, reflect.Float64:
			val := field.Float()
			return val >= min && val <= max
		default:
			return false
		}
	}
}

// 方案6: 预计算带条件
func Range6(min, max float64) ValidatorFunc {
	minInt := int64(min)
	maxInt := int64(max)
	minUint := uint64(min)
	maxUint := uint64(max)
	isIntRange := min == float64(minInt) && max == float64(maxInt) && minInt >= 0
	isUintRange := min == float64(minUint) && max == float64(maxUint)

	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := field.Int()
			if isIntRange {
				return val >= minInt && val <= maxInt
			}
			fval := float64(val)
			return fval >= min && fval <= max
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val := field.Uint()
			if isUintRange {
				return val >= minUint && val <= maxUint
			}
			fval := float64(val)
			return fval >= min && fval <= max
		case reflect.Float32, reflect.Float64:
			val := field.Float()
			return val >= min && val <= max
		default:
			return false
		}
	}
}

// 方案12: 优化快速路径
func Range12(min, max float64) ValidatorFunc {
	minInt := int64(min)
	maxInt := int64(max)
	isIntRange := min == float64(minInt) && max == float64(maxInt)

	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()

		if isIntRange && kind >= reflect.Int && kind <= reflect.Int64 {
			val := field.Int()
			return val >= minInt && val <= maxInt
		}

		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Float32, reflect.Float64:
			val := field.Float()
			return val >= min && val <= max
		default:
			return false
		}
	}
}

// 原始实现
func RangeOriginal(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Float32, reflect.Float64:
			val := field.Float()
			return val >= min && val <= max
		default:
			return false
		}
	}
}

// Mock 实现
type mockFieldLevel struct {
	field reflect.Value
}

func (m mockFieldLevel) Top() reflect.Value {
	return reflect.Value{}
}

func (m mockFieldLevel) Parent() reflect.Value {
	return reflect.Value{}
}

func (m mockFieldLevel) Field() reflect.Value {
	return m.field
}

func (m mockFieldLevel) FieldName() string {
	return ""
}

func (m mockFieldLevel) StructFieldName() string {
	return ""
}

func (m mockFieldLevel) Param() string {
	return ""
}

func (m mockFieldLevel) GetTag(key string) string {
	return ""
}

func (m mockFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// ============ Int8 基准测试 ============
func BenchmarkRange_Original_Int8_Valid(b *testing.B) {
	validator := RangeOriginal(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange1_Int8_Valid(b *testing.B) {
	validator := Range1(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange6_Int8_Valid(b *testing.B) {
	validator := Range6(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange12_Int8_Valid(b *testing.B) {
	validator := Range12(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

// ============ Int64 基准测试 ============
func BenchmarkRange_Original_Int64_Valid(b *testing.B) {
	validator := RangeOriginal(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int64(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange1_Int64_Valid(b *testing.B) {
	validator := Range1(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int64(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange6_Int64_Valid(b *testing.B) {
	validator := Range6(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int64(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange12_Int64_Valid(b *testing.B) {
	validator := Range12(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(int64(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

// ============ Float64 基准测试 ============
func BenchmarkRange_Original_Float64_Valid(b *testing.B) {
	validator := RangeOriginal(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(float64(50.0))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange1_Float64_Valid(b *testing.B) {
	validator := Range1(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(float64(50.0))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange6_Float64_Valid(b *testing.B) {
	validator := Range6(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(float64(50.0))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange12_Float64_Valid(b *testing.B) {
	validator := Range12(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(float64(50.0))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

// ============ Uint8 基准测试 ============
func BenchmarkRange_Original_Uint8_Valid(b *testing.B) {
	validator := RangeOriginal(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(uint8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange1_Uint8_Valid(b *testing.B) {
	validator := Range1(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(uint8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange6_Uint8_Valid(b *testing.B) {
	validator := Range6(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(uint8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}

func BenchmarkRange12_Uint8_Valid(b *testing.B) {
	validator := Range12(0, 100)
	field := mockFieldLevel{field: reflect.ValueOf(uint8(50))}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(field)
	}
}
