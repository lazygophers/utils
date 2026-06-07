package validator

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
	"testing"
)

// stringFieldHolder benchmark 中模拟单 string 字段的 struct
type stringFieldHolder struct {
	Field string
}

var testEmails = []string{
	"test@example.com",
	"user.name@example.com",
	"user+tag@example.com",
	"invalid",
	"@example.com",
	"user@",
	"a@b.c",
	"very.long.email.address@very.long.domain.name.com",
	"", // 空字符串
	"admin@mail.net",
	"hello@world.io",
	"info@company.co",
	"user123@test-domain.com",
	"first.last@sub.domain.org",
}

// 测试原始正则表达式方案（保留用于对比）
func BenchmarkEmail_OriginalRegex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, email := range testEmails {
			emailRegex.MatchString(email)
		}
	}
}

// 测试新的优化实现
func BenchmarkEmail_Optimized(b *testing.B) {
	validator := Email()
	fl := &fieldLevel{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, email := range testEmails {
			fl.field = reflect.ValueOf(email)
			validator(fl)
		}
	}
}

// 简单的内联版本用于基准测试
func BenchmarkEmail_InlineOptimized(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, email := range testEmails {
			if email == "" {
				continue
			}
			n := len(email)
			if n < 6 || n > 254 {
				continue
			}

			atIndex := strings.IndexByte(email, '@')
			if atIndex == -1 || atIndex == 0 || atIndex == n-1 {
				continue
			}

			localPart := email[:atIndex]
			domainPart := email[atIndex+1:]

			if len(localPart) == 0 || len(localPart) > 64 {
				continue
			}

			if len(domainPart) == 0 {
				continue
			}

			lastDot := strings.LastIndexByte(domainPart, '.')
			if lastDot == -1 || lastDot == 0 {
				continue
			}

			if len(domainPart)-lastDot-1 < 2 {
				continue
			}
		}
	}
}

// mockFieldLevel is a mock implementation of FieldLevel for testing
type mockFieldLevel struct {
	top       reflect.Value
	parent    reflect.Value
	field     reflect.Value
	fieldName string
	param     string
}

func (m mockFieldLevel) Top() reflect.Value {
	return m.top
}

func (m mockFieldLevel) Parent() reflect.Value {
	return m.parent
}

func (m mockFieldLevel) Field() reflect.Value {
	return m.field
}

func (m mockFieldLevel) FieldName() string {
	return m.fieldName
}

func (m mockFieldLevel) StructFieldName() string {
	return m.fieldName
}

func (m mockFieldLevel) Param() string {
	return m.param
}

func (m mockFieldLevel) GetTag(key string) string {
	return ""
}

func (m mockFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

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

// 生成测试数据（固定种子保证可重复）
func genFloats(n int) []float64 {
	r := rand.New(rand.NewSource(42))
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = r.Float64() * 200 // 0-200 范围
	}
	return s
}

func genInts(n int) []int {
	r := rand.New(rand.NewSource(42))
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = r.Intn(200) // 0-200 范围
	}
	return s
}

// 创建测试用的 FieldLevel
type rangeTestFieldLevel struct {
	val reflect.Value
}

func (fl rangeTestFieldLevel) Field() reflect.Value {
	return fl.val
}

func (fl rangeTestFieldLevel) Top() reflect.Value {
	return fl.val
}

func (fl rangeTestFieldLevel) FieldName() string {
	return "test"
}

func (fl rangeTestFieldLevel) Parent() reflect.Value {
	return fl.val
}

func (fl rangeTestFieldLevel) StructFieldName() string {
	return "test"
}

func (fl rangeTestFieldLevel) Param() string {
	return ""
}

func (fl rangeTestFieldLevel) GetTag(key string) string {
	return ""
}

func (fl rangeTestFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// ========== 方案1：当前实现（Baseline）==========
func Range_Original(min, max float64) ValidatorFunc {
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

// ========== 方案2：缓存 Kind ==========
func Range_CacheKind(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
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

// ========== 方案3：快速失败（先检查类型）==========
func Range_FastFail(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		k := field.Kind()

		// 快速失败：非数值类型
		if k < reflect.Int || k > reflect.Float64 {
			return false
		}

		// 处理有符号整数
		if k >= reflect.Int && k <= reflect.Int64 {
			val := float64(field.Int())
			return val >= min && val <= max
		}

		// 处理无符号整数
		if k >= reflect.Uint && k <= reflect.Uint64 {
			val := float64(field.Uint())
			return val >= min && val <= max
		}

		// 处理浮点数
		val := field.Float()
		return val >= min && val <= max
	}
}

// ========== 方案4：提前计算边界（避免重复比较）==========
func Range_PrecomputeBounds(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := float64(field.Int())
			gte := val >= min
			lte := val <= max
			return gte && lte
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val := float64(field.Uint())
			gte := val >= min
			lte := val <= max
			return gte && lte
		case reflect.Float32, reflect.Float64:
			val := field.Float()
			gte := val >= min
			lte := val <= max
			return gte && lte
		default:
			return false
		}
	}
}

// ========== 方案5：使用 Interface() 进行类型断言 ==========
func Range_InterfaceAssertion(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		i := field.Interface()

		switch v := i.(type) {
		case int:
			val := float64(v)
			return val >= min && val <= max
		case int8:
			val := float64(v)
			return val >= min && val <= max
		case int16:
			val := float64(v)
			return val >= min && val <= max
		case int32:
			val := float64(v)
			return val >= min && val <= max
		case int64:
			val := float64(v)
			return val >= min && val <= max
		case uint:
			val := float64(v)
			return val >= min && val <= max
		case uint8:
			val := float64(v)
			return val >= min && val <= max
		case uint16:
			val := float64(v)
			return val >= min && val <= max
		case uint32:
			val := float64(v)
			return val >= min && val <= max
		case uint64:
			val := float64(v)
			return val >= min && val <= max
		case float32:
			val := float64(v)
			return val >= min && val <= max
		case float64:
			return v >= min && v <= max
		default:
			return false
		}
	}
}

// ========== 方案6：整数优化路径（避免 float64 转换）==========
func Range_IntegerOptimized(min, max float64) ValidatorFunc {
	minInt := int64(min)
	maxInt := int64(max)
	isIntRange := min == float64(minInt) && max == float64(maxInt)

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
			if isIntRange && minInt >= 0 {
				val := int64(field.Uint())
				return val >= minInt && val <= maxInt
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

// ========== 方案7：分支预测优化（热门路径优先）==========
func Range_BranchPrediction(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		k := field.Kind()

		// 假设 float64 是最常见的情况
		if k == reflect.Float64 {
			val := field.Float()
			return val >= min && val <= max
		}

		// 其次是 int
		if k == reflect.Int {
			val := float64(field.Int())
			return val >= min && val <= max
		}

		// 其他情况
		switch k {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Float32:
			val := field.Float()
			return val >= min && val <= max
		default:
			return false
		}
	}
}

// ========== 方案8：无 switch（使用 if-else 链）==========
func Range_NoSwitch(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		k := field.Kind()

		if k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
			val := float64(field.Int())
			return val >= min && val <= max
		}

		if k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 {
			val := float64(field.Uint())
			return val >= min && val <= max
		}

		if k == reflect.Float32 || k == reflect.Float64 {
			val := field.Float()
			return val >= min && val <= max
		}

		return false
	}
}

// ========== 方案9：直接方法调用（避免 Kind()）==========
func Range_DirectMethod(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 尝试直接调用，成功则返回
		if field.CanInt() {
			val := float64(field.Int())
			return val >= min && val <= max
		}

		if field.CanUint() {
			val := float64(field.Uint())
			return val >= min && val <= max
		}

		if field.CanFloat() {
			val := field.Float()
			return val >= min && val <= max
		}

		return false
	}
}

// ========== 方案10：联合比较（单次比较）==========
func Range_CombineCompare(min, max float64) ValidatorFunc {
	diff := max - min
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := float64(field.Int())
			return val-min >= 0 && val-min <= diff
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val := float64(field.Uint())
			return val-min >= 0 && val-min <= diff
		case reflect.Float32, reflect.Float64:
			val := field.Float()
			return val-min >= 0 && val-min <= diff
		default:
			return false
		}
	}
}

// ========== 方案11：查表法（类型到处理函数映射）==========
type rangeHandler func(reflect.Value, float64, float64) bool

func handleInt(v reflect.Value, min, max float64) bool {
	val := float64(v.Int())
	return val >= min && val <= max
}

func handleUint(v reflect.Value, min, max float64) bool {
	val := float64(v.Uint())
	return val >= min && val <= max
}

func handleFloat(v reflect.Value, min, max float64) bool {
	val := v.Float()
	return val >= min && val <= max
}

func Range_LookupTable(min, max float64) ValidatorFunc {
	// 预构建类型处理表
	handlers := map[reflect.Kind]rangeHandler{
		reflect.Int:     handleInt,
		reflect.Int8:    handleInt,
		reflect.Int16:   handleInt,
		reflect.Int32:   handleInt,
		reflect.Int64:   handleInt,
		reflect.Uint:    handleUint,
		reflect.Uint8:   handleUint,
		reflect.Uint16:  handleUint,
		reflect.Uint32:  handleUint,
		reflect.Uint64:  handleUint,
		reflect.Float32: handleFloat,
		reflect.Float64: handleFloat,
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		handler, ok := handlers[field.Kind()]
		if !ok {
			return false
		}
		return handler(field, min, max)
	}
}

// ========== 方案12：内联优化（展开所有分支）==========
func Range_Inlined(min, max float64) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 完全展开，避免任何函数调用
		switch field.Kind() {
		case reflect.Int:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Int8:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Int16:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Int32:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Int64:
			val := float64(field.Int())
			return val >= min && val <= max
		case reflect.Uint:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Uint8:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Uint16:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Uint32:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Uint64:
			val := float64(field.Uint())
			return val >= min && val <= max
		case reflect.Float32:
			val := field.Float()
			return val >= min && val <= max
		case reflect.Float64:
			val := field.Float()
			return val >= min && val <= max
		default:
			return false
		}
	}
}

// ========== 基准测试 ==========

func BenchmarkRange_Original_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_Original(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_CacheKind_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_CacheKind(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_FastFail_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_FastFail(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_PrecomputeBounds_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_PrecomputeBounds(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_InterfaceAssertion_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_InterfaceAssertion(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_IntegerOptimized_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_IntegerOptimized(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_BranchPrediction_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_BranchPrediction(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_NoSwitch_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_NoSwitch(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_DirectMethod_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_DirectMethod(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_CombineCompare_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_CombineCompare(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_LookupTable_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_LookupTable(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_Inlined_Float64(b *testing.B) {
	data := genFloats(100)
	validator := Range_Inlined(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

// ========== Int 类型基准测试 ==========

func BenchmarkRange_Original_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_Original(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_CacheKind_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_CacheKind(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_FastFail_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_FastFail(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_PrecomputeBounds_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_PrecomputeBounds(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_InterfaceAssertion_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_InterfaceAssertion(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_IntegerOptimized_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_IntegerOptimized(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_BranchPrediction_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_BranchPrediction(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_NoSwitch_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_NoSwitch(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_DirectMethod_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_DirectMethod(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_CombineCompare_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_CombineCompare(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_LookupTable_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_LookupTable(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_Inlined_Int(b *testing.B) {
	data := genInts(100)
	validator := Range_Inlined(10, 100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

// ========== 内存分配基准测试 ==========

func BenchmarkRange_Original_Float64_Alloc(b *testing.B) {
	data := genFloats(100)
	validator := Range_Original(10, 100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_CacheKind_Float64_Alloc(b *testing.B) {
	data := genFloats(100)
	validator := Range_CacheKind(10, 100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_FastFail_Float64_Alloc(b *testing.B) {
	data := genFloats(100)
	validator := Range_FastFail(10, 100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_DirectMethod_Float64_Alloc(b *testing.B) {
	data := genFloats(100)
	validator := Range_DirectMethod(10, 100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

func BenchmarkRange_LookupTable_Float64_Alloc(b *testing.B) {
	data := genFloats(100)
	validator := Range_LookupTable(10, 100)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, v := range data {
			fl := rangeTestFieldLevel{val: reflect.ValueOf(v)}
			validator(fl)
		}
	}
}

// 测试用的复杂嵌套结构体
type Address struct {
	Street  string `validate:"required"`
	City    string `validate:"required"`
	ZipCode string `validate:"len=5"`
}

type Person struct {
	Name    string   `validate:"required"`
	Age     int      `validate:"gte=0,lte=150"`
	Email   string   `validate:"email"`
	Address Address  `validate:""`
	Tags    []string `validate:"dive,omitempty"`
}

type Company struct {
	Name      string   `validate:"required"`
	CEO       Person   `validate:"required"`
	Employees []Person `validate:"dive"`
}

// 生成测试数据
func generatePersonData(n int) []Person {
	people := make([]Person, n)
	for i := 0; i < n; i++ {
		people[i] = Person{
			Name:  fmt.Sprintf("Person%d", i),
			Age:   20 + i%50,
			Email: fmt.Sprintf("person%d@example.com", i),
			Address: Address{
				Street:  fmt.Sprintf("%d Street", i),
				City:    "City",
				ZipCode: "12345",
			},
			Tags: []string{"tag1", "tag2", "tag3"},
		}
	}
	return people
}

// 基准测试：当前实现
func BenchmarkValidateStruct_Current_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(person)
	}
}

func BenchmarkValidateStruct_Current_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(company)
	}
}

func BenchmarkValidateStruct_Current_Large(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(100),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(company)
	}
}

// 方案1: 缓存 field.Kind() 结果
func (e *Engine) validateStruct_Opt1_KindCache(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt1_KindCache(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt1_KindCache(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt1_KindCache(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt1_KindCache(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt1_KindCache(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt1_KindCache(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt1_KindCache_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt1_KindCache(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt1_KindCache_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt1_KindCache(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

// 方案2: 减少字段名称字符串拼接（只在需要时拼接）
func (e *Engine) validateStruct_Opt2_LazyNamespace(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		var fieldName string
		if namespace == "" {
			fieldName = fieldType.Name
		} else {
			fieldName = namespace + "." + fieldType.Name
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			if field.Kind() == reflect.Struct {
				e.validateStruct_Opt2_LazyNamespace(top, field, fieldName, errors)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt2_LazyNamespace(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fieldName + "[" + fmt.Sprintf("%d", j) + "]"

						if elem.Kind() == reflect.Struct {
							e.validateStruct_Opt2_LazyNamespace(top, elem, elemFieldName, errors)
						} else if elem.Kind() == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt2_LazyNamespace(top, elem.Elem(), elemFieldName, errors)
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		if field.Kind() == reflect.Struct {
			e.validateStruct_Opt2_LazyNamespace(top, field, fieldName, errors)
		} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt2_LazyNamespace(top, field.Elem(), fieldName, errors)
		}
	}
}

func BenchmarkValidateStruct_Opt2_LazyNamespace_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt2_LazyNamespace(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案3: 减少反射调用（内联字段访问）
func (e *Engine) validateStruct_Opt3_InlineAccess(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt3_InlineAccess(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt3_InlineAccess(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt3_InlineAccess(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt3_InlineAccess(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt3_InlineAccess(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt3_InlineAccess(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt3_InlineAccess_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt3_InlineAccess(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案4: 使用 strings.Builder 替代字符串拼接
func (e *Engine) validateStruct_Opt4_StringBuilder(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		var fieldName string
		if namespace == "" {
			fieldName = fieldType.Name
		} else {
			var builder strings.Builder
			builder.Grow(len(namespace) + len(fieldType.Name) + 1)
			builder.WriteString(namespace)
			builder.WriteByte('.')
			builder.WriteString(fieldType.Name)
			fieldName = builder.String()
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt4_StringBuilder(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt4_StringBuilder(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						var builder strings.Builder
						builder.Grow(len(fieldName) + 12)
						builder.WriteString(fieldName)
						builder.WriteByte('[')
						builder.WriteString(fmt.Sprint(j))
						builder.WriteByte(']')
						elemFieldName := builder.String()

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt4_StringBuilder(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt4_StringBuilder(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt4_StringBuilder(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt4_StringBuilder(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt4_StringBuilder_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt4_StringBuilder(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案5: 使用 sync.Pool 对象池复用 fieldLevel
// 注意：对象池已在 engine.go 中定义，此处使用全局对象池

func (e *Engine) validateStruct_Opt5_ObjectPool(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt5_ObjectPool(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt5_ObjectPool(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fieldName + "[" + fmt.Sprint(j) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt5_ObjectPool(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt5_ObjectPool(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := fieldLevelPool.Get().(*fieldLevel)
							elemFl.top = top
							elemFl.parent = field
							elemFl.field = elem
							elemFl.fieldName = elemFieldName
							elemFl.structFieldName = elemFieldName
							elemFl.structField = fieldType

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}

							fieldLevelPool.Put(elemFl)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt5_ObjectPool(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt5_ObjectPool(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt5_ObjectPool_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt5_ObjectPool(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案6: 组合优化（Kind缓存 + 内联访问）
func (e *Engine) validateStruct_Opt6_Combined(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt6_Combined(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt6_Combined(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt6_Combined(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt6_Combined(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt6_Combined(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt6_Combined(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt6_Combined_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt6_Combined(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt6_Combined_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt6_Combined(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

// 方案7: 组合优化 + 对象池
func (e *Engine) validateStruct_Opt7_FullCombined(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt7_FullCombined(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt7_FullCombined(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt7_FullCombined(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt7_FullCombined(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := fieldLevelPool.Get().(*fieldLevel)
							elemFl.top = top
							elemFl.parent = field
							elemFl.field = elem
							elemFl.fieldName = elemFieldName
							elemFl.structFieldName = elemFieldName
							elemFl.structField = fieldType

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}

							fieldLevelPool.Put(elemFl)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt7_FullCombined(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt7_FullCombined(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt7_FullCombined_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt7_FullCombined(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt7_FullCombined_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt7_FullCombined(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

// 方案8: 预先提取常用值到局部变量
func (e *Engine) validateStruct_Opt8_LocalVars(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()
	tagName := e.tagName
	fieldNameFunc := e.fieldNameFunc

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt8_LocalVars(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt8_LocalVars(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fieldName + "[" + fmt.Sprint(j) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt8_LocalVars(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt8_LocalVars(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							for _, elemRule := range elemRules {
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt8_LocalVars(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt8_LocalVars(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt8_LocalVars_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt8_LocalVars(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案9: 减少函数调用（内联简单检查）
func (e *Engine) validateStruct_Opt9_InlinedChecks(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		// 内联 IsExported 检查
		if fieldType.PkgPath != "" {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt9_InlinedChecks(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt9_InlinedChecks(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt9_InlinedChecks(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt9_InlinedChecks(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := &fieldLevel{
								top:             top,
								parent:          field,
								field:           elem,
								fieldName:       elemFieldName,
								structFieldName: elemFieldName,
								structField:     fieldType,
							}

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt9_InlinedChecks(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt9_InlinedChecks(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt9_InlinedChecks_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt9_InlinedChecks(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

// 方案10: 完整优化组合
func (e *Engine) validateStruct_Opt10_AllOptimizations(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()
	tagName := e.tagName
	fieldNameFunc := e.fieldNameFunc

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if fieldType.PkgPath != "" {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt10_AllOptimizations(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() {
				elem := field.Elem()
				if elem.Kind() == reflect.Struct {
					e.validateStruct_Opt10_AllOptimizations(top, elem, fieldName, errors)
				}
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fieldName + "[" + fmt.Sprint(k) + "]"

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt10_AllOptimizations(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() {
							elemElem := elem.Elem()
							if elemElem.Kind() == reflect.Struct {
								e.validateStruct_Opt10_AllOptimizations(top, elemElem, elemFieldName, errors)
							}
						} else if rule.param != "" {
							elemRules := e.parseTag(rule.param)
							elemFl := fieldLevelPool.Get().(*fieldLevel)
							elemFl.top = top
							elemFl.parent = field
							elemFl.field = elem
							elemFl.fieldName = elemFieldName
							elemFl.structFieldName = elemFieldName
							elemFl.structField = fieldType

							numElemRules := len(elemRules)
							for l := 0; l < numElemRules; l++ {
								elemRule := elemRules[l]
								elemFl.param = elemRule.param
								if !e.validateField(elemFl, elemRule.tag) {
									*errors = append(*errors, &FieldError{
										Field:       elemFieldName,
										Tag:         elemRule.tag,
										Value:       elem.Interface(),
										Param:       elemRule.param,
										ActualTag:   elemRule.tag,
										Namespace:   elemFieldName,
										StructField: elemFieldName,
										Message:     fmt.Sprintf("validation failed for tag '%s'", elemRule.tag),
									})
								}
							}

							fieldLevelPool.Put(elemFl)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt10_AllOptimizations(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() {
			elem := field.Elem()
			if elem.Kind() == reflect.Struct {
				e.validateStruct_Opt10_AllOptimizations(top, elem, fieldName, errors)
			}
		}
	}
}

func BenchmarkValidateStruct_Opt10_AllOptimizations_Simple(b *testing.B) {
	v, _ := New()
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt10_AllOptimizations(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt10_AllOptimizations_Nested(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(10),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt10_AllOptimizations(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

func BenchmarkValidateStruct_Opt10_AllOptimizations_Large(b *testing.B) {
	v, _ := New()
	company := Company{
		Name: "Tech Corp",
		CEO: Person{
			Name:  "CEO",
			Age:   50,
			Email: "ceo@example.com",
			Address: Address{
				Street:  "CEO Street",
				City:    "CEO City",
				ZipCode: "12345",
			},
		},
		Employees: generatePersonData(100),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var errs ValidationErrors
		v.engine.validateStruct_Opt10_AllOptimizations(reflect.ValueOf(company), reflect.ValueOf(company), "", &errs)
	}
}

type mockURLFieldLevel struct {
	s string
}

func (m *mockURLFieldLevel) Field() reflect.Value {
	return reflect.ValueOf(m.s)
}

func (m *mockURLFieldLevel) FieldName() string {
	return ""
}

func (m *mockURLFieldLevel) StructFieldName() string {
	return ""
}

func (m *mockURLFieldLevel) Param() string {
	return ""
}

func (m *mockURLFieldLevel) GetTag(key string) string {
	return ""
}

func (m *mockURLFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

func (m *mockURLFieldLevel) Top() reflect.Value {
	return reflect.Value{}
}

func (m *mockURLFieldLevel) Parent() reflect.Value {
	return reflect.Value{}
}

func BenchmarkURL(b *testing.B) {
	urls := []string{
		"http://example.com",
		"https://example.com",
		"ftp://example.com",
		"ws://example.com",
		"wss://example.com",
		"invalid",
		"",
	}

	fn := URL()
	for i := 0; i < b.N; i++ {
		for _, u := range urls {
			_ = fn(&mockURLFieldLevel{s: u})
		}
	}
}

func BenchmarkIsFieldNotEmptyOriginal(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			isFieldNotEmptyOriginal(v)
		}
	}
}

func BenchmarkIsFieldNotEmptyFastPath(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			isFieldNotEmptyFastPath(v)
		}
	}
}

func BenchmarkIsFieldNotEmptyReducedBranches(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			isFieldNotEmptyReducedBranches(v)
		}
	}
}

func BenchmarkIsFieldNotEmptyHybrid(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			isFieldNotEmptyHybrid(v)
		}
	}
}

func BenchmarkGetFieldValueAsStringOriginal(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			getFieldValueAsStringOriginal(v)
		}
	}
}

func BenchmarkGetFieldValueAsStringFastPath(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			getFieldValueAsStringFastPath(v)
		}
	}
}

func BenchmarkGetFieldValueAsStringHybrid(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			getFieldValueAsStringHybrid(v)
		}
	}
}

func BenchmarkCompareFieldsOriginal(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			compareFieldsOriginal(v, v)
		}
	}
}

func BenchmarkCompareFieldsFastPath(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			compareFieldsFastPath(v, v)
		}
	}
}

func BenchmarkCompareFieldsHybrid(b *testing.B) {
	cases := generateBenchCases()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			compareFieldsHybrid(v, v)
		}
	}
}

func BenchmarkStringHeavy(b *testing.B) {
	s := "test"
	v := reflect.ValueOf(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isFieldNotEmptyHybrid(v)
	}
}

func BenchmarkIntHeavy(b *testing.B) {
	n := 42
	v := reflect.ValueOf(n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isFieldNotEmptyHybrid(v)
	}
}

func BenchmarkPtrHeavy(b *testing.B) {
	s := "test"
	v := reflect.ValueOf(&s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isFieldNotEmptyHybrid(v)
	}
}

func BenchmarkPatternOptimized(b *testing.B) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	validator := Pattern(pattern)
	fl := &testFieldLevel{value: reflect.ValueOf("test@example.com")}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// ========== Alpha 优化方案 ==========

func alphaManualRange(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return false
		}
	}
	return true
}

func alphaSinglePass(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i] | 0x20
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func alphaBitOps(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		isLower := (c|0x20) == 'a' && c-'a' < 26
		if !isLower {
			return false
		}
	}
	return true
}

// ========== Alphanum 优化方案 ==========

func alphanumManualRange(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
}

func alphanumSinglePass(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		cl := c | 0x20
		if (cl < 'a' || cl > 'z') && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}

func alphanumBitOps(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		isLetter := (c|0x20) == 'a' && c-'a' < 26
		isDigit := c-'0' < 10
		if !isLetter && !isDigit {
			return false
		}
	}
	return true
}

// ========== 基准测试 ==========

var testAlphaMedium = "HelloWorldTest"
var testAlphanumMedium = "User123456789"

func BenchmarkAlphaRegex(b *testing.B) {
	fl := &fieldLevel{field: reflect.ValueOf(testAlphaMedium)}
	validator := Alpha()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkAlphaManualRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphaManualRange(testAlphaMedium)
	}
}

func BenchmarkAlphaSinglePass(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphaSinglePass(testAlphaMedium)
	}
}

func BenchmarkAlphaBitOps(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphaBitOps(testAlphaMedium)
	}
}

func BenchmarkAlphanumRegex(b *testing.B) {
	fl := &fieldLevel{field: reflect.ValueOf(testAlphanumMedium)}
	validator := Alphanum()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkAlphanumManualRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphanumManualRange(testAlphanumMedium)
	}
}

func BenchmarkAlphanumSinglePass(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphanumSinglePass(testAlphanumMedium)
	}
}

func BenchmarkAlphanumBitOps(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphanumBitOps(testAlphanumMedium)
	}
}

func BenchmarkRequired(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				reqOrig(tc.field)
			}
		}
	})

	b.Run("FastPath", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				reqFast(tc.field)
			}
		}
	})

	b.Run("SeparatedVars", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				reqSep(tc.field)
			}
		}
	})
}

// ========== 测试数据生成 ==========

// 生成固定种子的测试标签（保证可重复性）
func genTags(tagType string, count int) string {
	switch tagType {
	case "simple":
		// 简单标签：required,email,max=100
		return "required,email,max=100"
	case "medium":
		// 中等复杂度：required,email,min=18,max=100,len=6-20
		return "required,email,min=18,max=100,len=6-20"
	case "complex":
		// 复杂标签：多个规则
		return "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$,url,alpha,alphanum,alpha_dash"
	case "whitespace":
		// 带空格：required , email , max = 100
		return "required , email , max = 100"
	case "realistic":
		// 真实场景：用户验证规则
		return "required,email,min=3,max=50,len=6-20,regex=^[a-zA-Z0-9]+$"
	case "many":
		// 大量规则：20个规则
		rules := make([]string, count)
		for i := 0; i < count; i++ {
			if i%3 == 0 {
				rules[i] = "rule" + string(rune('A'+i))
			} else if i%3 == 1 {
				rules[i] = "rule" + string(rune('A'+i)) + "=param"
			} else {
				rules[i] = "rule" + string(rune('A'+i)) + " = value "
			}
		}
		return strings.Join(rules, ",")
	default:
		return "required"
	}
}

// ========== 当前实现（Baseline） ==========

func parseTag_Current(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案1：预分配切片容量 ==========

func parseTag_Prealloc(tag string) []validationRule {
	// 估算规则数：按逗号分隔，假设平均每个规则1个字符
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案2：使用 IndexByte 代替 Index ==========

func parseTag_IndexByte(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案3：预分配 + IndexByte ==========

func parseTag_Prealloc_IndexByte(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案4：减少 TrimSpace 调用（只在需要时trim） ==========

func parseTag_ReduceTrim(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		// 只在开始时trim一次
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			// 提取子串时不再trim，假设用户格式正确
			ruleName := part[:idx]
			param := part[idx+1:]
			// 手动trim前导空格
			for len(ruleName) > 0 && ruleName[0] == ' ' {
				ruleName = ruleName[1:]
			}
			for len(ruleName) > 0 && ruleName[len(ruleName)-1] == ' ' {
				ruleName = ruleName[:len(ruleName)-1]
			}
			for len(param) > 0 && param[0] == ' ' {
				param = param[1:]
			}
			for len(param) > 0 && param[len(param)-1] == ' ' {
				param = param[:len(param)-1]
			}
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案5：手动字节级解析（零分配） ==========

func parseTag_Bytes(tag string) []validationRule {
	// 预估规则数
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	start := 0
	inTag := true
	var ruleStart, ruleEnd int
	var paramStart, paramEnd int

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		if inTag {
			// 跳过前导空格
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == '=' || ch == ',' || i == len(tag) {
				ruleStart = start
				ruleEnd = i

				// trim尾部空格
				for ruleEnd > ruleStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				if ch == '=' {
					inTag = false
					start = i + 1
				} else {
					// 完成规则解析
					if ruleStart < ruleEnd {
						rules = append(rules, validationRule{
							tag:   tag[ruleStart:ruleEnd],
							param: "",
						})
					}
					start = i + 1
				}
			}
		} else {
			// 跳过前导空格
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == ',' || i == len(tag) {
				paramStart = start
				paramEnd = i

				// trim尾部空格
				for paramEnd > paramStart && tag[paramEnd-1] == ' ' {
					paramEnd--
				}

				// 完成规则解析
				rules = append(rules, validationRule{
					tag:   tag[ruleStart:ruleEnd],
					param: tag[paramStart:paramEnd],
				})

				inTag = true
				start = i + 1
			}
		}
	}

	return rules
}

// ========== 优化方案6：strings.Builder 重用 ==========

func parseTag_Builder(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	var builder strings.Builder
	builder.Grow(50)

	start := 0
	inTag := true

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		if inTag {
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == '=' || ch == ',' || i == len(tag) {
				end := i
				for end > start && tag[end-1] == ' ' {
					end--
				}

				if ch == '=' {
					builder.Reset()
					builder.WriteString(tag[start:end])
					inTag = false
					start = i + 1
				} else {
					if start < end {
						rules = append(rules, validationRule{
							tag:   tag[start:end],
							param: "",
						})
					}
					start = i + 1
				}
			}
		} else {
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == ',' || i == len(tag) {
				end := i
				for end > start && tag[end-1] == ' ' {
					end--
				}

				ruleName := builder.String()
				rules = append(rules, validationRule{
					tag:   ruleName,
					param: tag[start:end],
				})

				inTag = true
				start = i + 1
			}
		}
	}

	return rules
}

// ========== 优化方案7：单次遍历 + 字符串切片重用 ==========

func parseTag_SinglePass(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	var currentRule strings.Builder
	var currentParam strings.Builder
	currentRule.Grow(20)
	currentParam.Grow(20)

	inParam := false
	trimmed := false

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		// 跳过前导空格
		if !trimmed && ch == ' ' {
			continue
		}
		trimmed = true

		if ch == '=' {
			inParam = true
			trimmed = false
			continue
		}

		if ch == ',' || i == len(tag) {
			if inParam {
				// trim尾部空格
				paramStr := currentParam.String()
				end := len(paramStr)
				for end > 0 && paramStr[end-1] == ' ' {
					end--
				}
				rules = append(rules, validationRule{
					tag:   currentRule.String(),
					param: paramStr[:end],
				})
				currentParam.Reset()
			} else {
				ruleStr := currentRule.String()
				rules = append(rules, validationRule{
					tag:   ruleStr,
					param: "",
				})
			}
			currentRule.Reset()
			inParam = false
			trimmed = false
			continue
		}

		if inParam {
			currentParam.WriteByte(ch)
		} else {
			currentRule.WriteByte(ch)
		}
	}

	return rules
}

// ========== 优化方案8：使用 strings.Split 后批量处理 ==========

func parseTag_Batch(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案9：手动解析 + 索引优化 ==========

func parseTag_ManualIndex(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	start := 0
	eqIndex := -1
	trimStart := 0

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		// 跟踪等号位置
		if ch == '=' {
			eqIndex = i
			continue
		}

		if ch == ',' || i == len(tag) {
			// 处理前导空格
			for trimStart < i && tag[trimStart] == ' ' {
				trimStart++
			}
			if trimStart >= i {
				start = i + 1
				trimStart = start
				eqIndex = -1
				continue
			}

			// 处理尾部空格
			end := i
			for end > trimStart && tag[end-1] == ' ' {
				end--
			}

			if eqIndex != -1 && eqIndex < end {
				// 有参数
				ruleEnd := eqIndex
				for ruleEnd > trimStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				paramStart := eqIndex + 1
				for paramStart < end && tag[paramStart] == ' ' {
					paramStart++
				}

				rules = append(rules, validationRule{
					tag:   tag[trimStart:ruleEnd],
					param: tag[paramStart:end],
				})
			} else {
				// 无参数
				rules = append(rules, validationRule{
					tag:   tag[trimStart:end],
					param: "",
				})
			}

			start = i + 1
			trimStart = start
			eqIndex = -1
		}
	}

	return rules
}

// ========== 优化方案10：混合方案（Split + IndexByte + 预分配） ==========

func parseTag_Hybrid(tag string) []validationRule {
	// 使用 Split 简化逻辑，但优化其他部分
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		// 快速跳过空字符串
		if part == "" {
			continue
		}

		// 跳过前导空格
		start := 0
		for start < len(part) && part[start] == ' ' {
			start++
		}
		if start >= len(part) {
			continue
		}

		// 跳过尾部空格
		end := len(part)
		for end > start && part[end-1] == ' ' {
			end--
		}

		trimmed := part[start:end]

		// 使用 IndexByte
		if idx := strings.IndexByte(trimmed, '='); idx != -1 {
			ruleName := trimmed[:idx]
			param := trimmed[idx+1:]

			// 只在必要时trim参数
			paramStart := 0
			for paramStart < len(param) && param[paramStart] == ' ' {
				paramStart++
			}
			paramEnd := len(param)
			for paramEnd > paramStart && param[paramEnd-1] == ' ' {
				paramEnd--
			}

			rules = append(rules, validationRule{
				tag:   ruleName,
				param: param[paramStart:paramEnd],
			})
		} else {
			rules = append(rules, validationRule{tag: trimmed, param: ""})
		}
	}

	return rules
}

// ========== 基准测试 ==========

func BenchmarkParseTag_Current_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Prealloc_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc(tag)
	}
}

func BenchmarkParseTag_IndexByte_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_IndexByte(tag)
	}
}

func BenchmarkParseTag_Prealloc_IndexByte_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc_IndexByte(tag)
	}
}

func BenchmarkParseTag_ReduceTrim_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ReduceTrim(tag)
	}
}

func BenchmarkParseTag_Bytes_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Bytes(tag)
	}
}

func BenchmarkParseTag_Builder_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Builder(tag)
	}
}

func BenchmarkParseTag_SinglePass_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_SinglePass(tag)
	}
}

func BenchmarkParseTag_Batch_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Batch(tag)
	}
}

func BenchmarkParseTag_ManualIndex_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ManualIndex(tag)
	}
}

func BenchmarkParseTag_Hybrid_Simple(b *testing.B) {
	tag := genTags("simple", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Hybrid(tag)
	}
}

// ========== 复杂场景基准 ==========

func BenchmarkParseTag_Current_Medium(b *testing.B) {
	tag := genTags("medium", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Complex(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Whitespace(b *testing.B) {
	tag := genTags("whitespace", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Realistic(b *testing.B) {
	tag := genTags("realistic", 0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Current_Many(b *testing.B) {
	tag := genTags("many", 20)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

// ========== 内存分配测试 ==========

func BenchmarkParseTag_Current_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Current(tag)
	}
}

func BenchmarkParseTag_Prealloc_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc(tag)
	}
}

func BenchmarkParseTag_IndexByte_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_IndexByte(tag)
	}
}

func BenchmarkParseTag_Prealloc_IndexByte_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Prealloc_IndexByte(tag)
	}
}

func BenchmarkParseTag_ReduceTrim_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ReduceTrim(tag)
	}
}

func BenchmarkParseTag_Bytes_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Bytes(tag)
	}
}

func BenchmarkParseTag_Builder_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Builder(tag)
	}
}

func BenchmarkParseTag_SinglePass_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_SinglePass(tag)
	}
}

func BenchmarkParseTag_Batch_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Batch(tag)
	}
}

func BenchmarkParseTag_ManualIndex_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_ManualIndex(tag)
	}
}

func BenchmarkParseTag_Hybrid_Alloc(b *testing.B) {
	tag := genTags("complex", 0)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTag_Hybrid(tag)
	}
}

// Pattern 测试数据（避免与现有测试冲突）
var (
	patternSimplePattern  = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	patternComplexPattern = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	patternLiteralPattern = `^test@example\.com$`
	patternFixedLength    = `^\d{5}$`

	// Pattern 测试用例
	patternTestEmails = []string{
		"test@example.com",
		"user.name@domain.co.uk",
		"admin@test.org",
	}

	patternInvalidEmails = []string{
		"invalid",
		"@example.com",
		"test@",
		"test@.com",
	}

	patternLongString  = strings.Repeat("a", 1000) + "@test.com"
	patternEmptyString = ""
)

// ============ 方案 0: 当前实现（基线） ============
func patternBaseline(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	return func(fl FieldLevel) bool {
		return regex.MatchString(fl.Field().String())
	}
}

// ============ 方案 1: 缓存 Field().String() 结果 ============
func pattern1_CacheString(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	return func(fl FieldLevel) bool {
		str := fl.Field().String()
		return regex.MatchString(str)
	}
}

// ============ 方案 2: 使用 regex.Match() 字节级匹配 ============
func pattern2_ByteMatch(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			s := field.String()
			return regex.MatchString(s)
		default:
			return regex.Match([]byte(field.String()))
		}
	}
}

// ============ 方案 3: 快速路径 - 检测纯字符串模式 ============
func pattern3_FastLiteral(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)

	// 检测是否为纯字符串字面量模式（无元字符）
	isLiteral := true
	for _, c := range pattern {
		if c == '\\' || c == '^' || c == '$' || c == '.' || c == '*' ||
			c == '+' || c == '?' || c == '|' || c == '[' || c == '(' ||
			c == ')' || c == '{' || c == '}' {
			isLiteral = false
			break
		}
	}

	if isLiteral {
		return func(fl FieldLevel) bool {
			return strings.Contains(fl.Field().String(), pattern)
		}
	}

	return func(fl FieldLevel) bool {
		return regex.MatchString(fl.Field().String())
	}
}

// ============ 方案 4: 使用 field.Interface() + 类型断言 ============
func pattern4_InterfaceTypeAssert(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	return func(fl FieldLevel) bool {
		field := fl.Field()
		if str, ok := field.Interface().(string); ok {
			return regex.MatchString(str)
		}
		return regex.MatchString(field.String())
	}
}

// ============ 方案 5: 预检查字符串长度（适用于固定长度模式） ============
func pattern5_LengthCheck(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)

	// 尝试解析固定长度要求，如 ^\d{5}$
	expectedLen := -1
	if strings.HasPrefix(pattern, `^`) && strings.HasSuffix(pattern, `$`) {
		// 简单检测 {n} 模式
		start := strings.Index(pattern, `{`)
		end := strings.Index(pattern, `}`)
		if start != -1 && end != -1 && end > start {
			lenStr := pattern[start+1 : end]
			if lenStr == "5" {
				expectedLen = 5
			}
		}
	}

	if expectedLen > 0 {
		return func(fl FieldLevel) bool {
			field := fl.Field()
			str := field.String()
			if len(str) != expectedLen {
				return false
			}
			return regex.MatchString(str)
		}
	}

	return func(fl FieldLevel) bool {
		return regex.MatchString(fl.Field().String())
	}
}

// ============ 方案 6: 使用 sync.Pool 复用 byte slice ============
func pattern6_BytePool(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 128)
		},
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		str := field.String()
		buf := pool.Get().([]byte)
		buf = append(buf[:0], str...)
		matched := regex.Match(buf)
		pool.Put(buf)
		return matched
	}
}

// ============ 方案 7: 避免反射 - 直接使用反射缓存 ============
func pattern7_ReflectCache(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	return func(fl FieldLevel) bool {
		field := fl.Field()
		// 直接调用 String() 方法，减少反射层次
		return regex.MatchString(field.String())
	}
}

// ============ 方案 8: 闭包变量捕获优化 - 内联正则 ============
func pattern8_InlineRegex(pattern string) ValidatorFunc {
	return func(fl FieldLevel) bool {
		// 每次都编译正则（用于对比，预期性能较差）
		regex := regexp.MustCompile(pattern)
		return regex.MatchString(fl.Field().String())
	}
}

// ============ 方案 9: 混合优化 - 字符串类型特殊处理 ============
func pattern9_HybridOptimization(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)

	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 快速路径：字符串类型
		if field.Kind() == reflect.String {
			return regex.MatchString(field.String())
		}

		// 其他类型转换后匹配
		return regex.MatchString(field.String())
	}
}

// ============ 方案 10: 完全优化版本 ============
func pattern10_FullyOptimized(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)

	// 检测模式特征
	isLiteral := true
	expectedLen := -1

	// 检查是否为字面量
	if strings.HasPrefix(pattern, `^`) && strings.HasSuffix(pattern, `$`) {
		inner := pattern[1 : len(pattern)-1]
		for _, c := range inner {
			if c == '\\' || c == '.' || c == '*' || c == '+' ||
				c == '?' || c == '|' || c == '[' || c == '(' ||
				c == ')' || c == '{' || c == '}' {
				isLiteral = false
				break
			}
		}

		// 检查固定长度
		if strings.Contains(inner, `{`) && strings.Contains(inner, `}`) {
			start := strings.Index(inner, `{`)
			end := strings.Index(inner, `}`)
			if end > start {
				lenStr := inner[start+1 : end]
				if lenStr == "5" || lenStr == "3" {

					if lenStr == "5" {
						expectedLen = 5
					} else if lenStr == "3" {
						expectedLen = 3
					}
				}
			}
		}
	}

	// 根据模式特征返回不同的实现
	if isLiteral && expectedLen > 0 {
		literal := pattern[1 : len(pattern)-1]
		return func(fl FieldLevel) bool {
			str := fl.Field().String()
			return len(str) == expectedLen && str == literal
		}
	}

	if isLiteral {
		literal := pattern[1 : len(pattern)-1]
		return func(fl FieldLevel) bool {
			return fl.Field().String() == literal
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		return regex.MatchString(field.String())
	}
}

// ============ 方案 11: 预编译正则 + 直接字节访问 ============
func pattern11_PrecompiledBytes(pattern string) ValidatorFunc {
	regex := regexp.MustCompile(pattern)
	return func(fl FieldLevel) bool {
		field := fl.Field()
		if field.Kind() == reflect.String {
			s := field.String()
			// 直接使用字节切片
			return regex.Match([]byte(s))
		}
		return regex.MatchString(field.String())
	}
}

// ============ 方案 12: 使用 regexp.Compile 而非 MustCompile ============
func pattern12_Compile(pattern string) ValidatorFunc {
	regex, _ := regexp.Compile(pattern)
	return func(fl FieldLevel) bool {
		return regex.MatchString(fl.Field().String())
	}
}

// ============ 测试辅助函数 ============

type patternMockFieldLevel struct {
	field reflect.Value
}

func (m patternMockFieldLevel) Top() reflect.Value {
	return m.field
}

func (m patternMockFieldLevel) Parent() reflect.Value {
	return m.field
}

func (m patternMockFieldLevel) Field() reflect.Value {
	return m.field
}

func (m patternMockFieldLevel) FieldName() string {
	return "test"
}

func (m patternMockFieldLevel) StructFieldName() string {
	return "Test"
}

func (m patternMockFieldLevel) Param() string {
	return ""
}

func (m patternMockFieldLevel) GetTag(key string) string {
	return ""
}

func (m patternMockFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// ============ 基准测试 ============

// 简单正则 - 有效邮箱
func BenchmarkPattern_Baseline_Simple_Valid(b *testing.B) {
	validator := patternBaseline(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_CacheString_Simple_Valid(b *testing.B) {
	validator := pattern1_CacheString(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_ByteMatch_Simple_Valid(b *testing.B) {
	validator := pattern2_ByteMatch(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FastLiteral_Simple_Valid(b *testing.B) {
	validator := pattern3_FastLiteral(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_InterfaceTypeAssert_Simple_Valid(b *testing.B) {
	validator := pattern4_InterfaceTypeAssert(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_LengthCheck_Simple_Valid(b *testing.B) {
	validator := pattern5_LengthCheck(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_BytePool_Simple_Valid(b *testing.B) {
	validator := pattern6_BytePool(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_ReflectCache_Simple_Valid(b *testing.B) {
	validator := pattern7_ReflectCache(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_HybridOptimization_Simple_Valid(b *testing.B) {
	validator := pattern9_HybridOptimization(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_Simple_Valid(b *testing.B) {
	validator := pattern10_FullyOptimized(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_PrecompiledBytes_Simple_Valid(b *testing.B) {
	validator := pattern11_PrecompiledBytes(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 复杂正则 - 有效邮箱
func BenchmarkPattern_Baseline_Complex_Valid(b *testing.B) {
	validator := patternBaseline(patternComplexPattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_CacheString_Complex_Valid(b *testing.B) {
	validator := pattern1_CacheString(patternComplexPattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_Complex_Valid(b *testing.B) {
	validator := pattern10_FullyOptimized(patternComplexPattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_HybridOptimization_Complex_Valid(b *testing.B) {
	validator := pattern9_HybridOptimization(patternComplexPattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternTestEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 无效邮箱 - 快速失败
func BenchmarkPattern_Baseline_Simple_Invalid(b *testing.B) {
	validator := patternBaseline(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternInvalidEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_Simple_Invalid(b *testing.B) {
	validator := pattern10_FullyOptimized(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternInvalidEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_HybridOptimization_Simple_Invalid(b *testing.B) {
	validator := pattern9_HybridOptimization(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternInvalidEmails[0])}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 固定长度模式
func BenchmarkPattern_Baseline_FixedLength(b *testing.B) {
	validator := patternBaseline(patternFixedLength)
	fl := patternMockFieldLevel{field: reflect.ValueOf("12345")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_LengthCheck_FixedLength(b *testing.B) {
	validator := pattern5_LengthCheck(patternFixedLength)
	fl := patternMockFieldLevel{field: reflect.ValueOf("12345")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_FixedLength(b *testing.B) {
	validator := pattern10_FullyOptimized(patternFixedLength)
	fl := patternMockFieldLevel{field: reflect.ValueOf("12345")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 长字符串
func BenchmarkPattern_Baseline_LongString(b *testing.B) {
	validator := patternBaseline(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternLongString)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_LongString(b *testing.B) {
	validator := pattern10_FullyOptimized(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternLongString)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_HybridOptimization_LongString(b *testing.B) {
	validator := pattern9_HybridOptimization(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternLongString)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 空字符串
func BenchmarkPattern_Baseline_EmptyString(b *testing.B) {
	validator := patternBaseline(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternEmptyString)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_EmptyString(b *testing.B) {
	validator := pattern10_FullyOptimized(patternSimplePattern)
	fl := patternMockFieldLevel{field: reflect.ValueOf(patternEmptyString)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 字面量模式
func BenchmarkPattern_Baseline_Literal(b *testing.B) {
	validator := patternBaseline(`^hello$`)
	fl := patternMockFieldLevel{field: reflect.ValueOf("hello")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkPattern_FullyOptimized_Literal(b *testing.B) {
	validator := pattern10_FullyOptimized(`^hello$`)
	fl := patternMockFieldLevel{field: reflect.ValueOf("hello")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// ============== 方案1: 预转换 reflect.Value ==============
func InV1(values ...interface{}) ValidatorFunc {
	// 预先转换所有值为 reflect.Value，避免重复转换
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV1(values ...interface{}) ValidatorFunc {
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 方案2: int 类型专用快速路径 ==============
func InV2(values ...interface{}) ValidatorFunc {
	// 检测是否全是 int
	allInt := true
	intValues := make([]int, 0, len(values))

	for _, v := range values {
		if i, ok := v.(int); ok {
			intValues = append(intValues, i)
		} else {
			allInt = false
			break
		}
	}

	if allInt {
		// 使用 map 优化
		intMap := make(map[int]bool, len(intValues))
		for _, v := range intValues {
			intMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.Int {
				return intMap[int(field.Int())]
			}
			// 降级到原始方法
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return true
				}
			}
			return false
		}
	}

	// 非纯 int，使用原始方法
	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV2(values ...interface{}) ValidatorFunc {
	allInt := true
	intValues := make([]int, 0, len(values))

	for _, v := range values {
		if i, ok := v.(int); ok {
			intValues = append(intValues, i)
		} else {
			allInt = false
			break
		}
	}

	if allInt {
		intMap := make(map[int]bool, len(intValues))
		for _, v := range intValues {
			intMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.Int {
				return !intMap[int(field.Int())]
			}
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return false
				}
			}
			return true
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 方案3: string 类型专用快速路径 ==============
func InV3(values ...interface{}) ValidatorFunc {
	// 检测是否全是 string
	allString := true
	stringValues := make([]string, 0, len(values))

	for _, v := range values {
		if s, ok := v.(string); ok {
			stringValues = append(stringValues, s)
		} else {
			allString = false
			break
		}
	}

	if allString {
		// 使用 map 优化
		stringMap := make(map[string]bool, len(stringValues))
		for _, v := range stringValues {
			stringMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.String {
				return stringMap[field.String()]
			}
			// 降级
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return true
				}
			}
			return false
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV3(values ...interface{}) ValidatorFunc {
	allString := true
	stringValues := make([]string, 0, len(values))

	for _, v := range values {
		if s, ok := v.(string); ok {
			stringValues = append(stringValues, s)
		} else {
			allString = false
			break
		}
	}

	if allString {
		stringMap := make(map[string]bool, len(stringValues))
		for _, v := range stringValues {
			stringMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.String {
				return !stringMap[field.String()]
			}
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return false
				}
			}
			return true
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 方案4: 统一类型用 map ==============
func InV4(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 检查类型是否统一
	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return intMap[int(field.Int())]
				}
				return false
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return stringMap[field.String()]
				}
				return false
			}
		}
	}

	// 混合类型或未处理类型，使用线性查找
	return InV1(values...)
}

func NotInV4(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return !intMap[int(field.Int())]
				}
				return true
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return !stringMap[field.String()]
				}
				return true
			}
		}
	}

	return NotInV1(values...)
}

// ============== 方案5: 少量枚举用展开 switch (硬编码 3 个) ==============
func InV5(values ...interface{}) ValidatorFunc {
	if len(values) == 3 {
		v0, v1, v2 := values[0], values[1], values[2]
		rv0, rv1, rv2 := reflect.ValueOf(v0), reflect.ValueOf(v1), reflect.ValueOf(v2)

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if compareFields(field, rv0) == 0 {
				return true
			}
			if compareFields(field, rv1) == 0 {
				return true
			}
			if compareFields(field, rv2) == 0 {
				return true
			}
			return false
		}
	}

	// 降级
	return InV1(values...)
}

func NotInV5(values ...interface{}) ValidatorFunc {
	if len(values) == 3 {
		v0, v1, v2 := values[0], values[1], values[2]
		rv0, rv1, rv2 := reflect.ValueOf(v0), reflect.ValueOf(v1), reflect.ValueOf(v2)

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if compareFields(field, rv0) == 0 {
				return false
			}
			if compareFields(field, rv1) == 0 {
				return false
			}
			if compareFields(field, rv2) == 0 {
				return false
			}
			return true
		}
	}

	return NotInV1(values...)
}

// ============== 方案6: 直接 interface{} 比较（避免反射） ==============
func InV6(values ...interface{}) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 尝试直接获取 interface{} 值
		var fieldInterface interface{}
		if field.CanInterface() {
			fieldInterface = field.Interface()
		} else {
			// 无法直接获取，使用反射
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return true
				}
			}
			return false
		}

		// 直接比较
		for _, v := range values {
			if fieldInterface == v {
				return true
			}
		}
		return false
	}
}

func NotInV6(values ...interface{}) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()

		var fieldInterface interface{}
		if field.CanInterface() {
			fieldInterface = field.Interface()
		} else {
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return false
				}
			}
			return true
		}

		for _, v := range values {
			if fieldInterface == v {
				return false
			}
		}
		return true
	}
}

// ============== 方案7: sort + binary search（有序枚举） ==============
func InV7(values ...interface{}) ValidatorFunc {
	if len(values) <= 10 {
		// 少量值不排序，直接线性查找
		return InV1(values...)
	}

	// 复制并排序
	sortedValues := make([]interface{}, len(values))
	copy(sortedValues, values)

	sort.Slice(sortedValues, func(i, j int) bool {
		vi := reflect.ValueOf(sortedValues[i])
		vj := reflect.ValueOf(sortedValues[j])
		return compareFields(vi, vj) < 0
	})

	reflectValues := make([]reflect.Value, len(sortedValues))
	for i, v := range sortedValues {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 二分查找
		low, high := 0, len(reflectValues)-1
		for low <= high {
			mid := (low + high) / 2
			cmp := compareFields(field, reflectValues[mid])
			if cmp == 0 {
				return true
			} else if cmp < 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return false
	}
}

func NotInV7(values ...interface{}) ValidatorFunc {
	if len(values) <= 10 {
		return NotInV1(values...)
	}

	sortedValues := make([]interface{}, len(values))
	copy(sortedValues, values)

	sort.Slice(sortedValues, func(i, j int) bool {
		vi := reflect.ValueOf(sortedValues[i])
		vj := reflect.ValueOf(sortedValues[j])
		return compareFields(vi, vj) < 0
	})

	reflectValues := make([]reflect.Value, len(sortedValues))
	for i, v := range sortedValues {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		low, high := 0, len(reflectValues)-1
		for low <= high {
			mid := (low + high) / 2
			cmp := compareFields(field, reflectValues[mid])
			if cmp == 0 {
				return false
			} else if cmp < 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return true
	}
}

// ============== 方案8: 混合优化（智能选择策略） ==============
func InV8(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 少量枚举：直接线性
	if len(values) <= 5 {
		return InV1(values...)
	}

	// 检查是否统一类型
	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return intMap[int(field.Int())]
				}
				return false
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return stringMap[field.String()]
				}
				return false
			}
		}
	}

	// 大量混合类型：二分查找
	return InV7(values...)
}

func NotInV8(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	if len(values) <= 5 {
		return NotInV1(values...)
	}

	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return !intMap[int(field.Int())]
				}
				return true
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return !stringMap[field.String()]
				}
				return true
			}
		}
	}

	return NotInV7(values...)
}

// ============== 方案9: 使用 sync.Map 缓存验证结果 ==============
func InV9(values ...interface{}) ValidatorFunc {
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	cache := &sync.Map{}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 尝试从缓存获取
		fieldKey := field.Interface()
		if cached, ok := cache.Load(fieldKey); ok {
			return cached.(bool)
		}

		// 计算并缓存
		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				cache.Store(fieldKey, true)
				return true
			}
		}
		cache.Store(fieldKey, false)
		return false
	}
}

func NotInV9(values ...interface{}) ValidatorFunc {
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	cache := &sync.Map{}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		fieldKey := field.Interface()
		if cached, ok := cache.Load(fieldKey); ok {
			return !cached.(bool)
		}

		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				cache.Store(fieldKey, true)
				return false
			}
		}
		cache.Store(fieldKey, false)
		return true
	}
}

// ============== 方案10: 分组优化（按类型分组） ==============
func InV10(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 按类型分组
	typeGroups := make(map[reflect.Type][]reflect.Value)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		typeGroups[rt] = append(typeGroups[rt], rv)
	}

	// 为每种类型创建优化的查找器
	typeCheckers := make(map[reflect.Kind]func(reflect.Value) bool)

	for rt, rvs := range typeGroups {
		kind := rt.Kind()

		switch kind {
		case reflect.Int:
			intMap := make(map[int64]bool, len(rvs))
			for _, rv := range rvs {
				intMap[rv.Int()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return intMap[field.Int()]
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(rvs))
			for _, rv := range rvs {
				stringMap[rv.String()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return stringMap[field.String()]
			}

		default:
			// 其他类型使用线性查找
			typeCheckers[kind] = func(field reflect.Value) bool {
				for _, rv := range rvs {
					if compareFields(field, rv) == 0 {
						return true
					}
				}
				return false
			}
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		if checker, ok := typeCheckers[field.Kind()]; ok {
			return checker(field)
		}
		return false
	}
}

func NotInV10(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	typeGroups := make(map[reflect.Type][]reflect.Value)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		typeGroups[rt] = append(typeGroups[rt], rv)
	}

	typeCheckers := make(map[reflect.Kind]func(reflect.Value) bool)

	for rt, rvs := range typeGroups {
		kind := rt.Kind()

		switch kind {
		case reflect.Int:
			intMap := make(map[int64]bool, len(rvs))
			for _, rv := range rvs {
				intMap[rv.Int()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return !intMap[field.Int()]
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(rvs))
			for _, rv := range rvs {
				stringMap[rv.String()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return !stringMap[field.String()]
			}

		default:
			typeCheckers[kind] = func(field reflect.Value) bool {
				for _, rv := range rvs {
					if compareFields(field, rv) == 0 {
						return false
					}
				}
				return true
			}
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		if checker, ok := typeCheckers[field.Kind()]; ok {
			return checker(field)
		}
		return true
	}
}

// ============== 方案11: 组合优化（预转换 + 类型检测 + map） ==============
func InV11(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 预转换并分析
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	// 检查是否统一类型
	unifiedType := reflectValues[0].Type()
	allSameType := true
	for _, rv := range reflectValues {
		if rv.Type() != unifiedType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch unifiedType.Kind() {
		case reflect.Int:
			intMap := make(map[int64]bool, len(values))
			for _, rv := range reflectValues {
				intMap[rv.Int()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return intMap[field.Int()]
				}
				return false
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, rv := range reflectValues {
				stringMap[rv.String()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return stringMap[field.String()]
				}
				return false
			}

		case reflect.Float64, reflect.Float32:
			floatMap := make(map[float64]bool, len(values))
			for _, rv := range reflectValues {
				floatMap[rv.Float()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
					return floatMap[field.Float()]
				}
				return false
			}
		}
	}

	// 混合类型，使用预转换的线性查找
	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, rv := range reflectValues {
			if compareFields(field, rv) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV11(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	unifiedType := reflectValues[0].Type()
	allSameType := true
	for _, rv := range reflectValues {
		if rv.Type() != unifiedType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch unifiedType.Kind() {
		case reflect.Int:
			intMap := make(map[int64]bool, len(values))
			for _, rv := range reflectValues {
				intMap[rv.Int()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return !intMap[field.Int()]
				}
				return true
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, rv := range reflectValues {
				stringMap[rv.String()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return !stringMap[field.String()]
				}
				return true
			}

		case reflect.Float64, reflect.Float32:
			floatMap := make(map[float64]bool, len(values))
			for _, rv := range reflectValues {
				floatMap[rv.Float()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
					return !floatMap[field.Float()]
				}
				return true
			}
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, rv := range reflectValues {
			if compareFields(field, rv) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 基准测试 ==============

// 测试数据准备
type TestStruct struct {
	Value int
}

func (t TestStruct) Field() reflect.Value {
	return reflect.ValueOf(t.Value)
}

func (t TestStruct) FieldName() string {
	return "Field"
}

func (t TestStruct) StructFieldName() string {
	return "TestStruct.Field"
}

func (t TestStruct) Param() string {
	return ""
}

func (t TestStruct) GetTag(key string) string {
	return ""
}

func (t TestStruct) Top() reflect.Value {
	return reflect.ValueOf(t)
}

func (t TestStruct) Parent() reflect.Value {
	return reflect.Value{}
}

func (t TestStruct) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// ============== In 函数基准测试 ==============

// 少量 int 枚举（3个）
func BenchmarkIn_Original_3Int(b *testing.B) {
	validator := In(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V1_3Int(b *testing.B) {
	validator := InV1(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V2_3Int(b *testing.B) {
	validator := InV2(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V4_3Int(b *testing.B) {
	validator := InV4(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V5_3Int(b *testing.B) {
	validator := InV5(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V6_3Int(b *testing.B) {
	validator := InV6(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V8_3Int(b *testing.B) {
	validator := InV8(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V10_3Int(b *testing.B) {
	validator := InV10(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V11_3Int(b *testing.B) {
	validator := InV11(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 中等 int 枚举（15个）
func BenchmarkIn_Original_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := In(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V1_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV1(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V2_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV2(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V4_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV4(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V7_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV7(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V8_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV8(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V10_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV10(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V11_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV11(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 大量 int 枚举（100个）
func BenchmarkIn_Original_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := In(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V1_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV1(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V2_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV2(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V4_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV4(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V7_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV7(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V8_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV8(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V10_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV10(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V11_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV11(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// String 类型测试（中等枚举）
func BenchmarkIn_Original_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := In(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := stringFieldHolder{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V1_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV1(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := stringFieldHolder{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V3_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV3(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := stringFieldHolder{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V4_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV4(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := stringFieldHolder{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V11_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV11(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := stringFieldHolder{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

// ============== NotIn 函数基准测试 ==============

// 少量 int 枚举
func BenchmarkNotIn_Original_3Int(b *testing.B) {
	validator := NotIn(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V1_3Int(b *testing.B) {
	validator := NotInV1(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V2_3Int(b *testing.B) {
	validator := NotInV2(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V11_3Int(b *testing.B) {
	validator := NotInV11(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 中等 int 枚举
func BenchmarkNotIn_Original_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotIn(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V1_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotInV1(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V2_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotInV2(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V11_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotInV11(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 大量 int 枚举
func BenchmarkNotIn_Original_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotIn(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V1_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotInV1(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V2_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotInV2(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V11_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotInV11(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// ========== 当前实现（Baseline） ==========

func parseTagBaseline(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案1：预分配切片 ==========

func parseTagOpt1(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.Index(part, "="); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案2：IndexByte ==========

func parseTagOpt2(tag string) []validationRule {
	var rules []validationRule
	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案3：预分配 + IndexByte ==========

func parseTagOpt3(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案4：批量处理 ==========

func parseTagOpt4(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if idx := strings.IndexByte(part, '='); idx != -1 {
			ruleName := strings.TrimSpace(part[:idx])
			param := strings.TrimSpace(part[idx+1:])
			rules = append(rules, validationRule{tag: ruleName, param: param})
		} else {
			rules = append(rules, validationRule{tag: part, param: ""})
		}
	}

	return rules
}

// ========== 优化方案5：手动 trim ==========

func parseTagOpt5(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		start := 0
		for start < len(part) && part[start] == ' ' {
			start++
		}
		if start >= len(part) {
			continue
		}

		end := len(part)
		for end > start && part[end-1] == ' ' {
			end--
		}

		trimmed := part[start:end]

		if idx := strings.IndexByte(trimmed, '='); idx != -1 {
			ruleName := trimmed[:idx]
			param := trimmed[idx+1:]

			paramStart := 0
			for paramStart < len(param) && param[paramStart] == ' ' {
				paramStart++
			}
			paramEnd := len(param)
			for paramEnd > paramStart && param[paramEnd-1] == ' ' {
				paramEnd--
			}

			rules = append(rules, validationRule{
				tag:   ruleName,
				param: param[paramStart:paramEnd],
			})
		} else {
			rules = append(rules, validationRule{tag: trimmed, param: ""})
		}
	}

	return rules
}

// ========== 优化方案6：单次遍历 ==========

func parseTagOpt6(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	start := 0
	inTag := true
	var ruleStart, ruleEnd int
	var paramStart, paramEnd int

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		if inTag {
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == '=' || ch == ',' || i == len(tag) {
				ruleStart = start
				ruleEnd = i

				for ruleEnd > ruleStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				if ch == '=' {
					inTag = false
					start = i + 1
				} else {
					if ruleStart < ruleEnd {
						rules = append(rules, validationRule{
							tag:   tag[ruleStart:ruleEnd],
							param: "",
						})
					}
					start = i + 1
				}
			}
		} else {
			if start <= i && ch == ' ' {
				start = i + 1
				continue
			}

			if ch == ',' || i == len(tag) {
				paramStart = start
				paramEnd = i

				for paramEnd > paramStart && tag[paramEnd-1] == ' ' {
					paramEnd--
				}

				rules = append(rules, validationRule{
					tag:   tag[ruleStart:ruleEnd],
					param: tag[paramStart:paramEnd],
				})

				inTag = true
				start = i + 1
			}
		}
	}

	return rules
}

// ========== 优化方案7：优化 Index 追踪 ==========

func parseTagOpt7(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	start := 0
	eqIndex := -1
	trimStart := 0

	for i := 0; i <= len(tag); i++ {
		ch := byte(0)
		if i < len(tag) {
			ch = tag[i]
		}

		if ch == '=' {
			eqIndex = i
			continue
		}

		if ch == ',' || i == len(tag) {
			for trimStart < i && tag[trimStart] == ' ' {
				trimStart++
			}
			if trimStart >= i {
				start = i + 1
				trimStart = start
				eqIndex = -1
				continue
			}

			end := i
			for end > trimStart && tag[end-1] == ' ' {
				end--
			}

			if eqIndex != -1 && eqIndex < end {
				ruleEnd := eqIndex
				for ruleEnd > trimStart && tag[ruleEnd-1] == ' ' {
					ruleEnd--
				}

				paramStart := eqIndex + 1
				for paramStart < end && tag[paramStart] == ' ' {
					paramStart++
				}

				rules = append(rules, validationRule{
					tag:   tag[trimStart:ruleEnd],
					param: tag[paramStart:end],
				})
			} else {
				rules = append(rules, validationRule{
					tag:   tag[trimStart:end],
					param: "",
				})
			}

			start = i + 1
			trimStart = start
			eqIndex = -1
		}
	}

	return rules
}

// ========== 优化方案8：混合方案 ==========

func parseTagOpt8(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		if part == "" {
			continue
		}

		start := 0
		for start < len(part) && part[start] == ' ' {
			start++
		}
		if start >= len(part) {
			continue
		}

		end := len(part)
		for end > start && part[end-1] == ' ' {
			end--
		}

		trimmed := part[start:end]

		if idx := strings.IndexByte(trimmed, '='); idx != -1 {
			ruleName := trimmed[:idx]
			param := trimmed[idx+1:]

			paramStart := 0
			for paramStart < len(param) && param[paramStart] == ' ' {
				paramStart++
			}
			paramEnd := len(param)
			for paramEnd > paramStart && param[paramEnd-1] == ' ' {
				paramEnd--
			}

			rules = append(rules, validationRule{
				tag:   ruleName,
				param: param[paramStart:paramEnd],
			})
		} else {
			rules = append(rules, validationRule{tag: trimmed, param: ""})
		}
	}

	return rules
}

// ========== 优化方案9：最简优化 ==========

func parseTagOpt9(tag string) []validationRule {
	parts := strings.Split(tag, ",")
	rules := make([]validationRule, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		idx := strings.IndexByte(part, '=')
		if idx == -1 {
			rules = append(rules, validationRule{tag: part, param: ""})
			continue
		}

		ruleName := part[:idx]
		param := part[idx+1:]

		trimRuleEnd := len(ruleName)
		for trimRuleEnd > 0 && ruleName[trimRuleEnd-1] == ' ' {
			trimRuleEnd--
		}

		trimParamStart := 0
		for trimParamStart < len(param) && param[trimParamStart] == ' ' {
			trimParamStart++
		}

		trimParamEnd := len(param)
		for trimParamEnd > trimParamStart && param[trimParamEnd-1] == ' ' {
			trimParamEnd--
		}

		rules = append(rules, validationRule{
			tag:   ruleName[:trimRuleEnd],
			param: param[trimParamStart:trimParamEnd],
		})
	}

	return rules
}

// ========== 优化方案10：完全手动解析 ==========

func parseTagOpt10(tag string) []validationRule {
	estimatedCount := (len(tag) + 1) / 2
	rules := make([]validationRule, 0, estimatedCount)

	i := 0
	n := len(tag)

	for i < n {
		// 跳过前导空格和逗号
		for i < n && (tag[i] == ' ' || tag[i] == ',') {
			i++
		}
		if i >= n {
			break
		}

		start := i
		eqPos := -1

		// 查找规则结束位置（逗号）和等号位置
		for i < n && tag[i] != ',' {
			if tag[i] == '=' && eqPos == -1 {
				eqPos = i
			}
			i++
		}

		end := i
		// trim尾部空格
		for end > start && tag[end-1] == ' ' {
			end--
		}

		if eqPos != -1 && eqPos < end {
			// 有参数
			ruleEnd := eqPos
			for ruleEnd > start && tag[ruleEnd-1] == ' ' {
				ruleEnd--
			}

			paramStart := eqPos + 1
			for paramStart < end && tag[paramStart] == ' ' {
				paramStart++
			}

			rules = append(rules, validationRule{
				tag:   tag[start:ruleEnd],
				param: tag[paramStart:end],
			})
		} else {
			// 无参数
			rules = append(rules, validationRule{
				tag:   tag[start:end],
				param: "",
			})
		}

		i++
	}

	return rules
}

// ========== 基准测试 ==========

func BenchmarkParseTag_Baseline_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagBaseline(tag)
	}
}

func BenchmarkParseTag_Opt1_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt1(tag)
	}
}

func BenchmarkParseTag_Opt2_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt2(tag)
	}
}

func BenchmarkParseTag_Opt3_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt3(tag)
	}
}

func BenchmarkParseTag_Opt4_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt4(tag)
	}
}

func BenchmarkParseTag_Opt5_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt5(tag)
	}
}

func BenchmarkParseTag_Opt6_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt6(tag)
	}
}

func BenchmarkParseTag_Opt7_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt7(tag)
	}
}

func BenchmarkParseTag_Opt8_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt8(tag)
	}
}

func BenchmarkParseTag_Opt9_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt9(tag)
	}
}

func BenchmarkParseTag_Opt10_Simple(b *testing.B) {
	tag := "required,email,max=100"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt10(tag)
	}
}

// ========== 复杂场景 ==========

func BenchmarkParseTag_Baseline_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagBaseline(tag)
	}
}

func BenchmarkParseTag_Opt3_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt3(tag)
	}
}

func BenchmarkParseTag_Opt4_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt4(tag)
	}
}

func BenchmarkParseTag_Opt10_Complex(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt10(tag)
	}
}

// ========== 内存分配测试 ==========

func BenchmarkParseTag_Baseline_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagBaseline(tag)
	}
}

func BenchmarkParseTag_Opt3_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt3(tag)
	}
}

func BenchmarkParseTag_Opt4_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt4(tag)
	}
}

func BenchmarkParseTag_Opt10_Alloc(b *testing.B) {
	tag := "required,email,min=18,max=100,len=6-20"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parseTagOpt10(tag)
	}
}

func BenchmarkMinLength_Original(b *testing.B) {
	fn := minLengthOriginal(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt1_CacheKind(b *testing.B) {
	fn := minLengthOpt1_CacheKind(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt2_FieldLen(b *testing.B) {
	fn := minLengthOpt2_FieldLen(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt3_IfElse(b *testing.B) {
	fn := minLengthOpt3_IfElse(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt4_NoIntermediate(b *testing.B) {
	fn := minLengthOpt4_NoIntermediate(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt5_SeparatePaths(b *testing.B) {
	fn := minLengthOpt5_SeparatePaths(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt6_FastPath(b *testing.B) {
	fn := minLengthOpt6_FastPath(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt7_ShortCircuit(b *testing.B) {
	fn := minLengthOpt7_ShortCircuit(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt8_EarlyCheck(b *testing.B) {
	fn := minLengthOpt8_EarlyCheck(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt9_Inlined(b *testing.B) {
	fn := minLengthOpt9_Inlined(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt10_Goto(b *testing.B) {
	fn := minLengthOpt10_Goto(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt11_DirectCompare(b *testing.B) {
	fn := minLengthOpt11_DirectCompare(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt12_Precompute(b *testing.B) {
	fn := minLengthOpt12_Precompute(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt13_MergedTypes(b *testing.B) {
	fn := minLengthOpt13_MergedTypes(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt14_DoubleCheck(b *testing.B) {
	fn := minLengthOpt14_DoubleCheck(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkMinLength_Opt15_Minimal(b *testing.B) {
	fn := minLengthOpt15_Minimal(3)
	benchMinLengthFunc(b, fn, minlengthCases)
}

func BenchmarkLength_Original(b *testing.B) {
	fn := lengthOriginal(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt1_CacheKind(b *testing.B) {
	fn := lengthOpt1_CacheKind(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt2_StringDirect(b *testing.B) {
	fn := lengthOpt2_StringDirect(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt3_IfElse(b *testing.B) {
	fn := lengthOpt3_IfElse(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt4_NoIntermediate(b *testing.B) {
	fn := lengthOpt4_NoIntermediate(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt5_PreCheck(b *testing.B) {
	fn := lengthOpt5_PreCheck(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt6_SeparatePaths(b *testing.B) {
	fn := lengthOpt6_SeparatePaths(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt7_FastPath(b *testing.B) {
	fn := lengthOpt7_FastPath(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt8_Inlined(b *testing.B) {
	fn := lengthOpt8_Inlined(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt9_Goto(b *testing.B) {
	fn := lengthOpt9_Goto(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt10_Hybrid(b *testing.B) {
	fn := lengthOpt10_Hybrid(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt11_Minimal(b *testing.B) {
	fn := lengthOpt11_Minimal(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt12_ShortCircuit(b *testing.B) {
	fn := lengthOpt12_ShortCircuit(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt13_Precompute(b *testing.B) {
	fn := lengthOpt13_Precompute(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt14_BitOps(b *testing.B) {
	fn := lengthOpt14_BitOps(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

func BenchmarkLength_Opt15_Expanded(b *testing.B) {
	fn := lengthOpt15_Expanded(1, 10)
	benchLengthFunc(b, fn, lengthCases)
}

// 快速测试数据结构
type QuickPerson struct {
	Name    string   `validate:"required"`
	Age     int      `validate:"gte=0,lte=150"`
	Email   string   `validate:"email"`
	Address Address  `validate:""`
	Tags    []string `validate:"dive,omitempty"`
}

func generateQuickPersonData(n int) []QuickPerson {
	people := make([]QuickPerson, n)
	for i := 0; i < n; i++ {
		people[i] = QuickPerson{
			Name:  fmt.Sprintf("Person%d", i),
			Age:   20 + i%50,
			Email: fmt.Sprintf("person%d@example.com", i),
			Address: Address{
				Street:  fmt.Sprintf("%d Street", i),
				City:    "City",
				ZipCode: "12345",
			},
			Tags: []string{"tag1", "tag2", "tag3"},
		}
	}
	return people
}

// Baseline: 当前实现
func (e *Engine) validateStruct_Baseline(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			if field.Kind() == reflect.Struct {
				e.validateStruct_Baseline(top, field, fieldName, errors)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Baseline(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						if elem.Kind() == reflect.Struct {
							e.validateStruct_Baseline(top, elem, elemFieldName, errors)
						} else if elem.Kind() == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Baseline(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		if field.Kind() == reflect.Struct {
			e.validateStruct_Baseline(top, field, fieldName, errors)
		} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Baseline(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案1: 缓存 Kind
func (e *Engine) validateStruct_Opt1(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt1(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt1(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt1(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt1(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt1(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt1(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案2: 内联访问
func (e *Engine) validateStruct_Opt2(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt2(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt2(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, k)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt2(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt2(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt2(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt2(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案3: 对象池
func (e *Engine) validateStruct_Opt3(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()

	for i := 0; i < current.NumField(); i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt3(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt3(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		for _, rule := range rules {
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					for j := 0; j < field.Len(); j++ {
						elem := field.Index(j)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, j)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt3(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt3(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt3(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt3(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案4: 组合优化 (Kind缓存 + 内联访问)
func (e *Engine) validateStruct_Opt4(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt4(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt4(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := &fieldLevel{
			top:             top,
			parent:          current,
			field:           field,
			fieldName:       displayName,
			structFieldName: fieldType.Name,
			structField:     fieldType,
		}

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, k)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt4(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt4(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt4(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt4(top, field.Elem(), fieldName, errors)
		}
	}
}

// 方案5: 完整优化 (Kind缓存 + 内联访问 + 对象池)
func (e *Engine) validateStruct_Opt5(top, current reflect.Value, namespace string, errors *ValidationErrors) {
	rt := current.Type()
	numField := current.NumField()

	for i := 0; i < numField; i++ {
		field := current.Field(i)
		fieldType := rt.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		fieldName := fieldType.Name
		if namespace != "" {
			fieldName = namespace + "." + fieldName
		}

		tag := fieldType.Tag.Get(e.tagName)
		if tag == "" || tag == "-" {
			fieldKind := field.Kind()
			if fieldKind == reflect.Struct {
				e.validateStruct_Opt5(top, field, fieldName, errors)
			} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				e.validateStruct_Opt5(top, field.Elem(), fieldName, errors)
			}
			continue
		}

		rules := e.parseTag(tag)
		displayName := e.fieldNameFunc(fieldType)

		fl := fieldLevelPool.Get().(*fieldLevel)
		fl.top = top
		fl.parent = current
		fl.field = field
		fl.fieldName = displayName
		fl.structFieldName = fieldType.Name
		fl.structField = fieldType

		numRules := len(rules)
		for j := 0; j < numRules; j++ {
			rule := rules[j]
			fl.param = rule.param

			if rule.tag == "dive" {
				fieldKind := field.Kind()
				if fieldKind == reflect.Slice || fieldKind == reflect.Array {
					fieldLen := field.Len()
					for k := 0; k < fieldLen; k++ {
						elem := field.Index(k)
						elemFieldName := fmt.Sprintf("%s[%d]", fieldName, k)

						elemKind := elem.Kind()
						if elemKind == reflect.Struct {
							e.validateStruct_Opt5(top, elem, elemFieldName, errors)
						} else if elemKind == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
							e.validateStruct_Opt5(top, elem.Elem(), elemFieldName, errors)
						}
					}
				}
				continue
			}

			if !e.validateField(fl, rule.tag) {
				*errors = append(*errors, &FieldError{
					Field:       displayName,
					Tag:         rule.tag,
					Value:       field.Interface(),
					Param:       rule.param,
					ActualTag:   rule.tag,
					Namespace:   fieldName,
					StructField: fieldType.Name,
					Message:     formatMessage(getDefaultMessage(rule.tag), "var", rule.tag, rule.param),
				})
			}
		}

		fieldLevelPool.Put(fl)

		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			e.validateStruct_Opt5(top, field, fieldName, errors)
		} else if fieldKind == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
			e.validateStruct_Opt5(top, field.Elem(), fieldName, errors)
		}
	}
}

// 基准测试
func BenchmarkValidateStruct_Comparison_Simple(b *testing.B) {
	v, _ := New()
	person := QuickPerson{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			ZipCode: "12345",
		},
	}

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Baseline(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt1-KindCache", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt1(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt2-InlineAccess", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt2(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt3-ObjectPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt3(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt4-Combined", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt4(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})

	b.Run("Opt5-FullOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errs ValidationErrors
			v.engine.validateStruct_Opt5(reflect.ValueOf(person), reflect.ValueOf(person), "", &errs)
		}
	})
}

func BenchmarkValidateStruct_Comparison_Nested(b *testing.B) {
	v, _ := New()
	people := generateQuickPersonData(10)

	b.Run("Baseline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Baseline(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt1-KindCache", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt1(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt2-InlineAccess", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt2(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt3-ObjectPool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt3(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt4-Combined", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt4(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})

	b.Run("Opt5-FullOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := range people {
				var errs ValidationErrors
				v.engine.validateStruct_Opt5(reflect.ValueOf(people[j]), reflect.ValueOf(people[j]), "", &errs)
			}
		}
	})
}
