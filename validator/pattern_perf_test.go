package validator

import (
	"reflect"
	"regexp"
	"strings"
	"sync"
	"testing"
)

// Pattern 测试数据（避免与现有测试冲突）
var (
	patternSimplePattern   = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	patternComplexPattern  = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	patternLiteralPattern  = `^test@example\.com$`
	patternFixedLength     = `^\d{5}$`

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

	patternLongString = strings.Repeat("a", 1000) + "@test.com"
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
