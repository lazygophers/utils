package validator

import (
	"reflect"
	"testing"
)

// ============================================
// MinLength 性能优化基准测试
// 目标: 优化第899行 MinLength 函数性能
// ============================================

// 测试用例
var minlengthCases = []reflect.Value{
	reflect.ValueOf("hello"),                                // 有效字符串
	reflect.ValueOf(""),                                     // 空字符串
	reflect.ValueOf("hi"),                                   // 太短
	reflect.ValueOf("this is long enough"),                  // 长字符串
	reflect.ValueOf([]int{1, 2, 3, 4, 5}),                   // 有效切片
	reflect.ValueOf([]int{}),                                // 空切片
	reflect.ValueOf([]int{1}),                               // 单元素
	reflect.ValueOf([]string{"a", "b"}),                     // 字符串切片
	reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3}), // 有效map
	reflect.ValueOf(map[string]int{}),                       // 空map
	reflect.ValueOf([5]int{1, 2, 3, 4, 5}),                  // 数组
	reflect.ValueOf([0]int{}),                               // 空数组
}

// 辅助函数：运行 MinLength 基准测试
func benchMinLengthFunc(b *testing.B, fn ValidatorFunc, cases []reflect.Value) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			fl := &fieldLevel{field: c}
			fn(fl)
		}
	}
}

// ============================================
// 原始实现 (Baseline)
// ============================================
func minLengthOriginal(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return len(field.String()) >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案1: 缓存 Kind() 到局部变量
// ============================================
func minLengthOpt1_CacheKind(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		switch kind {
		case reflect.String:
			return len(field.String()) >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案2: 字符串使用 field.Len() 代替 String()
// ============================================
func minLengthOpt2_FieldLen(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return field.Len() >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案3: if-else 代替 switch
// ============================================
func minLengthOpt3_IfElse(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		if kind == reflect.String {
			return field.Len() >= min
		} else if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
			return field.Len() >= min
		}
		return false
	}
}

// ============================================
// 方案4: 直接返回，消除中间变量
// ============================================
func minLengthOpt4_NoIntermediate(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return field.Len() >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案5: 分离 String 和 Container 路径
// ============================================
func minLengthOpt5_SeparatePaths(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		if kind == reflect.String {
			return field.Len() >= min
		}
		if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
			return field.Len() >= min
		}
		return false
	}
}

// ============================================
// 方案6: Fast-path (优先检查常见类型 String)
// ============================================
func minLengthOpt6_FastPath(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		if field.Kind() == reflect.String {
			return field.Len() >= min
		}
		k := field.Kind()
		if k == reflect.Slice || k == reflect.Map || k == reflect.Array {
			return field.Len() >= min
		}
		return false
	}
}

// ============================================
// 方案7: 短路优化 (先检查容器类型)
// ============================================
func minLengthOpt7_ShortCircuit(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		// 容器类型通常更快
		if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
			return field.Len() >= min
		}
		if kind == reflect.String {
			return field.Len() >= min
		}
		return false
	}
}

// ============================================
// 方案8: 提前反射值检查
// ============================================
func minLengthOpt8_EarlyCheck(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		if !field.IsValid() {
			return false
		}
		switch field.Kind() {
		case reflect.String:
			return field.Len() >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案9: 内联优化 (减少函数调用)
// ============================================
func minLengthOpt9_Inlined(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		f := fl.Field()
		switch f.Kind() {
		case reflect.String:
			return f.Len() >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return f.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案10: 使用 goto 减少重复代码 (非常规优化)
// ============================================
func minLengthOpt10_Goto(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		var length int

		switch field.Kind() {
		case reflect.String:
			length = field.Len()
		case reflect.Slice, reflect.Map, reflect.Array:
			length = field.Len()
		default:
			return false
		}

		if length >= min {
			return true
		}
		return false
	}
}

// ============================================
// 方案11: 直接比较 (消除不必要的中间变量)
// ============================================
func minLengthOpt11_DirectCompare(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return field.Len() >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案12: 预计算常量
// ============================================
func minLengthOpt12_Precompute(min int) ValidatorFunc {
	// 将 min 提升为常量比较
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			return field.Len() >= min
		case reflect.Slice, reflect.Map, reflect.Array:
			return field.Len() >= min
		default:
			return false
		}
	}
}

// ============================================
// 方案13: 合并类型检查 (减少分支)
// ============================================
func minLengthOpt13_MergedTypes(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		k := field.Kind()
		// String 和容器类型都调用 Len()
		if k == reflect.String || k == reflect.Slice || k == reflect.Map || k == reflect.Array {
			return field.Len() >= min
		}
		return false
	}
}

// ============================================
// 方案14: 双重检查优化 (先 Kind 后 Len)
// ============================================
func minLengthOpt14_DoubleCheck(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		k := field.Kind()
		if k == reflect.String || k == reflect.Slice || k == reflect.Map || k == reflect.Array {
			l := field.Len()
			return l >= min
		}
		return false
	}
}

// ============================================
// 方案15: 极简版本 (最少代码行)
// ============================================
func minLengthOpt15_Minimal(min int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		f := fl.Field()
		if k := f.Kind(); k == reflect.String || k == reflect.Slice || k == reflect.Map || k == reflect.Array {
			return f.Len() >= min
		}
		return false
	}
}

// ============================================
// 基准测试函数
// ============================================
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

// ============================================
// 正确性测试
// ============================================
func TestMinLengthOptimizations_Correctness(t *testing.T) {
	min := 3
	testCases := []struct {
		name     string
		value    reflect.Value
		expected bool
	}{
		{"valid_string", reflect.ValueOf("hello"), true},
		{"empty_string", reflect.ValueOf(""), false},
		{"too_short", reflect.ValueOf("hi"), false},
		{"exact_length", reflect.ValueOf("abc"), true},
		{"valid_slice", reflect.ValueOf([]int{1, 2, 3, 4}), true},
		{"empty_slice", reflect.ValueOf([]int{}), false},
		{"short_slice", reflect.ValueOf([]int{1, 2}), false},
		{"valid_map", reflect.ValueOf(map[string]int{"a": 1, "b": 2, "c": 3}), true},
		{"empty_map", reflect.ValueOf(map[string]int{}), false},
		{"valid_array", reflect.ValueOf([5]int{1, 2, 3, 4, 5}), true},
		{"empty_array", reflect.ValueOf([0]int{}), false},
		{"invalid_type", reflect.ValueOf(123), false},
	}

	optimizations := []struct {
		name string
		fn   ValidatorFunc
	}{
		{"Original", minLengthOriginal(min)},
		{"Opt1_CacheKind", minLengthOpt1_CacheKind(min)},
		{"Opt2_FieldLen", minLengthOpt2_FieldLen(min)},
		{"Opt3_IfElse", minLengthOpt3_IfElse(min)},
		{"Opt4_NoIntermediate", minLengthOpt4_NoIntermediate(min)},
		{"Opt5_SeparatePaths", minLengthOpt5_SeparatePaths(min)},
		{"Opt6_FastPath", minLengthOpt6_FastPath(min)},
		{"Opt7_ShortCircuit", minLengthOpt7_ShortCircuit(min)},
		{"Opt8_EarlyCheck", minLengthOpt8_EarlyCheck(min)},
		{"Opt9_Inlined", minLengthOpt9_Inlined(min)},
		{"Opt10_Goto", minLengthOpt10_Goto(min)},
		{"Opt11_DirectCompare", minLengthOpt11_DirectCompare(min)},
		{"Opt12_Precompute", minLengthOpt12_Precompute(min)},
		{"Opt13_MergedTypes", minLengthOpt13_MergedTypes(min)},
		{"Opt14_DoubleCheck", minLengthOpt14_DoubleCheck(min)},
		{"Opt15_Minimal", minLengthOpt15_Minimal(min)},
	}

	for _, opt := range optimizations {
		t.Run(opt.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					fl := &fieldLevel{field: tc.value}
					result := opt.fn(fl)
					if result != tc.expected {
						t.Errorf("For value %v, expected %v, got %v",
							tc.value, tc.expected, result)
					}
				})
			}
		})
	}
}
