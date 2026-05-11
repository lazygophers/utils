package validator

import (
	"reflect"
	"testing"
)

// alwaysTrue returns true
func alwaysTrue(fl FieldLevel) bool {
	return true
}

// alwaysFalse returns false
func alwaysFalse(fl FieldLevel) bool {
	return false
}

// benchShortCircuitAnd creates validators that fail on first one
func benchShortCircuitAnd(count int) []ValidatorFunc {
	validators := make([]ValidatorFunc, count)
	validators[0] = alwaysFalse
	for i := 1; i < count; i++ {
		validators[i] = alwaysTrue
	}
	return validators
}

// benchNonShortCircuitAnd creates validators that execute all
func benchNonShortCircuitAnd(count int) []ValidatorFunc {
	validators := make([]ValidatorFunc, count)
	for i := 0; i < count-1; i++ {
		validators[i] = alwaysTrue
	}
	validators[count-1] = alwaysFalse
	return validators
}

// benchShortCircuitOr creates validators that succeed on first one
func benchShortCircuitOr(count int) []ValidatorFunc {
	validators := make([]ValidatorFunc, count)
	validators[0] = alwaysTrue
	for i := 1; i < count; i++ {
		validators[i] = alwaysFalse
	}
	return validators
}

// benchNonShortCircuitOr creates validators that execute all
func benchNonShortCircuitOr(count int) []ValidatorFunc {
	validators := make([]ValidatorFunc, count)
	for i := 0; i < count-1; i++ {
		validators[i] = alwaysFalse
	}
	validators[count-1] = alwaysTrue
	return validators
}

// AndIndexLoop uses index loop instead of range
func AndIndexLoop(validators ...ValidatorFunc) ValidatorFunc {
	return func(fl FieldLevel) bool {
		for i := 0; i < len(validators); i++ {
			if !validators[i](fl) {
				return false
			}
		}
		return true
	}
}

// OrIndexLoop uses index loop instead of range
func OrIndexLoop(validators ...ValidatorFunc) ValidatorFunc {
	return func(fl FieldLevel) bool {
		for i := 0; i < len(validators); i++ {
			if validators[i](fl) {
				return true
			}
		}
		return false
	}
}

// AndSwitch unrolls small number of validators
func AndSwitch(validators ...ValidatorFunc) ValidatorFunc {
	switch len(validators) {
	case 0:
		return func(fl FieldLevel) bool { return true }
	case 1:
		v0 := validators[0]
		return func(fl FieldLevel) bool { return v0(fl) }
	case 2:
		v0, v1 := validators[0], validators[1]
		return func(fl FieldLevel) bool {
			return v0(fl) && v1(fl)
		}
	case 3:
		v0, v1, v2 := validators[0], validators[1], validators[2]
		return func(fl FieldLevel) bool {
			return v0(fl) && v1(fl) && v2(fl)
		}
	default:
		return AndIndexLoop(validators...)
	}
}

// OrSwitch unrolls small number of validators
func OrSwitch(validators ...ValidatorFunc) ValidatorFunc {
	switch len(validators) {
	case 0:
		return func(fl FieldLevel) bool { return false }
	case 1:
		v0 := validators[0]
		return func(fl FieldLevel) bool { return v0(fl) }
	case 2:
		v0, v1 := validators[0], validators[1]
		return func(fl FieldLevel) bool {
			return v0(fl) || v1(fl)
		}
	case 3:
		v0, v1, v2 := validators[0], validators[1], validators[2]
		return func(fl FieldLevel) bool {
			return v0(fl) || v1(fl) || v2(fl)
		}
	default:
		return OrIndexLoop(validators...)
	}
}

// AndGoto uses goto for loop optimization
func AndGoto(validators ...ValidatorFunc) ValidatorFunc {
	if len(validators) == 0 {
		return func(fl FieldLevel) bool { return true }
	}
	return func(fl FieldLevel) bool {
		i := 0
	next:
		if i >= len(validators) {
			return true
		}
		if !validators[i](fl) {
			return false
		}
		i++
		goto next
	}
}

// OrGoto uses goto for loop optimization
func OrGoto(validators ...ValidatorFunc) ValidatorFunc {
	if len(validators) == 0 {
		return func(fl FieldLevel) bool { return false }
	}
	return func(fl FieldLevel) bool {
		i := 0
	next:
		if i >= len(validators) {
			return false
		}
		if validators[i](fl) {
			return true
		}
		i++
		goto next
	}
}

// AndStruct uses struct + method
type andValidator struct {
	validators []ValidatorFunc
}

func (a *andValidator) Validate(fl FieldLevel) bool {
	for i := 0; i < len(a.validators); i++ {
		if !a.validators[i](fl) {
			return false
		}
	}
	return true
}

func AndStruct(validators ...ValidatorFunc) ValidatorFunc {
	if len(validators) == 0 {
		return func(fl FieldLevel) bool { return true }
	}
	a := &andValidator{validators: validators}
	return a.Validate
}

// OrStruct uses struct + method
type orValidator struct {
	validators []ValidatorFunc
}

func (o *orValidator) Validate(fl FieldLevel) bool {
	for i := 0; i < len(o.validators); i++ {
		if o.validators[i](fl) {
			return true
		}
	}
	return false
}

func OrStruct(validators ...ValidatorFunc) ValidatorFunc {
	if len(validators) == 0 {
		return func(fl FieldLevel) bool { return false }
	}
	o := &orValidator{validators: validators}
	return o.Validate
}

// AndHybrid combines switch and index loop
func AndHybrid(validators ...ValidatorFunc) ValidatorFunc {
	switch len(validators) {
	case 0:
		return func(fl FieldLevel) bool { return true }
	case 1:
		v0 := validators[0]
		return func(fl FieldLevel) bool { return v0(fl) }
	case 2:
		v0, v1 := validators[0], validators[1]
		return func(fl FieldLevel) bool {
			return v0(fl) && v1(fl)
		}
	case 3:
		v0, v1, v2 := validators[0], validators[1], validators[2]
		return func(fl FieldLevel) bool {
			return v0(fl) && v1(fl) && v2(fl)
		}
	default:
		return AndIndexLoop(validators...)
	}
}

// OrHybrid combines switch and index loop
func OrHybrid(validators ...ValidatorFunc) ValidatorFunc {
	switch len(validators) {
	case 0:
		return func(fl FieldLevel) bool { return false }
	case 1:
		v0 := validators[0]
		return func(fl FieldLevel) bool { return v0(fl) }
	case 2:
		v0, v1 := validators[0], validators[1]
		return func(fl FieldLevel) bool {
			return v0(fl) || v1(fl)
		}
	case 3:
		v0, v1, v2 := validators[0], validators[1], validators[2]
		return func(fl FieldLevel) bool {
			return v0(fl) || v1(fl) || v2(fl)
		}
	default:
		return OrIndexLoop(validators...)
	}
}

// ============================================================================
// And Benchmarks - Short Circuit (2 validators)
// ============================================================================

func BenchmarkAnd_Original_Short2(b *testing.B) {
	validators := benchShortCircuitAnd(2)
	v := And(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_IndexLoop_Short2(b *testing.B) {
	validators := benchShortCircuitAnd(2)
	v := AndIndexLoop(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Switch_Short2(b *testing.B) {
	validators := benchShortCircuitAnd(2)
	v := AndSwitch(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Goto_Short2(b *testing.B) {
	validators := benchShortCircuitAnd(2)
	v := AndGoto(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Struct_Short2(b *testing.B) {
	validators := benchShortCircuitAnd(2)
	v := AndStruct(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Hybrid_Short2(b *testing.B) {
	validators := benchShortCircuitAnd(2)
	v := AndHybrid(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

// ============================================================================
// And Benchmarks - No Short Circuit (2 validators)
// ============================================================================

func BenchmarkAnd_Original_NoShort2(b *testing.B) {
	validators := benchNonShortCircuitAnd(2)
	v := And(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_IndexLoop_NoShort2(b *testing.B) {
	validators := benchNonShortCircuitAnd(2)
	v := AndIndexLoop(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Switch_NoShort2(b *testing.B) {
	validators := benchNonShortCircuitAnd(2)
	v := AndSwitch(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Goto_NoShort2(b *testing.B) {
	validators := benchNonShortCircuitAnd(2)
	v := AndGoto(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Struct_NoShort2(b *testing.B) {
	validators := benchNonShortCircuitAnd(2)
	v := AndStruct(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkAnd_Hybrid_NoShort2(b *testing.B) {
	validators := benchNonShortCircuitAnd(2)
	v := AndHybrid(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

// ============================================================================
// Or Benchmarks - Short Circuit (2 validators)
// ============================================================================

func BenchmarkOr_Original_Short2(b *testing.B) {
	validators := benchShortCircuitOr(2)
	v := Or(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkOr_IndexLoop_Short2(b *testing.B) {
	validators := benchShortCircuitOr(2)
	v := OrIndexLoop(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkOr_Switch_Short2(b *testing.B) {
	validators := benchShortCircuitOr(2)
	v := OrSwitch(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkOr_Goto_Short2(b *testing.B) {
	validators := benchShortCircuitOr(2)
	v := OrGoto(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkOr_Struct_Short2(b *testing.B) {
	validators := benchShortCircuitOr(2)
	v := OrStruct(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

func BenchmarkOr_Hybrid_Short2(b *testing.B) {
	validators := benchShortCircuitOr(2)
	v := OrHybrid(validators[0], validators[1])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}

// ============================================================================
// Not Benchmarks
// ============================================================================

func BenchmarkNot_Original(b *testing.B) {
	v := Not(alwaysTrue)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v(mockFieldLevel{field: reflect.ValueOf("")})
	}
}
