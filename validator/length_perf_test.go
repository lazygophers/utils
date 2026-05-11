package validator

import (
	"reflect"
	"testing"
)

// 测试数据结构
type LengthTestCase struct {
	name     string
	min      int
	max      int
	field    reflect.Value
	expected bool
}

// 生成测试数据
func generateLengthTestCases() []LengthTestCase {
	return []LengthTestCase{
		// 字符串测试
		{"valid_string_5chars", 0, 0, reflect.ValueOf("hello"), true},         // 5 in [1,10]
		{"valid_string_1char", 0, 0, reflect.ValueOf("h"), true},              // 1 in [1,10]
		{"valid_string_10chars", 0, 0, reflect.ValueOf("hellohello"), true},   // 10 in [1,10]
		{"too_short_0chars", 0, 0, reflect.ValueOf(""), false},                // 0 not in [1,10]
		{"too_long_11chars", 0, 0, reflect.ValueOf("helloworld!"), false},     // 11 not in [1,10]

		// Slice 测试
		{"valid_slice_3elems", 0, 0, reflect.ValueOf([]int{1, 2, 3}), true},
		{"valid_slice_1elem", 0, 0, reflect.ValueOf([]int{1}), true},
		{"valid_slice_10elems", 0, 0, reflect.ValueOf(make([]int, 10)), true},
		{"too_short_slice_0elems", 0, 0, reflect.ValueOf([]int{}), false},
		{"too_long_slice_11elems", 0, 0, reflect.ValueOf(make([]int, 11)), false},

		// Map 测试
		{"valid_map_2elems", 0, 0, reflect.ValueOf(map[string]int{"a": 1, "b": 2}), true},
		{"valid_map_1elem", 0, 0, reflect.ValueOf(map[string]int{"a": 1}), true},
		{"too_short_map_0elems", 0, 0, reflect.ValueOf(map[string]int{}), false},
		{"too_long_map_11elems", 0, 0, reflect.ValueOf(map[string]int{
			"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11,
		}), false},

		// Array 测试
		{"valid_array_3elems", 0, 0, reflect.ValueOf([3]int{1, 2, 3}), true},
		{"valid_array_1elem", 0, 0, reflect.ValueOf([1]int{1}), true},
		{"too_short_array_0elems", 0, 0, reflect.ValueOf([0]int{}), false},
		{"too_long_array_11elems", 0, 0, reflect.ValueOf([11]int{}), false},

		// 无效类型测试
		{"invalid_type_int", 0, 0, reflect.ValueOf(123), false},
		{"invalid_type_float", 0, 0, reflect.ValueOf(3.14), false},
	}
}

// ============================================
// 方案0: 原始实现（基准）
// ============================================
func lengthOriginal(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		length := 0
		switch field.Kind() {
		case reflect.String:
			length = len(field.String())
		case reflect.Slice, reflect.Map, reflect.Array:
			length = field.Len()
		default:
			return false
		}
		return length >= min && length <= max
	}
}

// ============================================
// 方案1: 缓存 field.Kind() 到局部变量
// ============================================
func lengthOpt1_CacheKind(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		length := 0
		switch kind {
		case reflect.String:
			length = len(field.String())
		case reflect.Slice, reflect.Map, reflect.Array:
			length = field.Len()
		default:
			return false
		}
		return length >= min && length <= max
	}
}

// ============================================
// 方案2: 字符串类型直接计算 len（避免 String() 调用）
// ============================================
func lengthOpt2_StringDirect(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		length := 0
		switch field.Kind() {
		case reflect.String:
			length = field.Len()
		case reflect.Slice, reflect.Map, reflect.Array:
			length = field.Len()
		default:
			return false
		}
		return length >= min && length <= max
	}
}

// ============================================
// 方案3: 使用 if-else 代替 switch
// ============================================
func lengthOpt3_IfElse(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		length := 0
		if kind == reflect.String {
			length = field.Len()
		} else if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
			length = field.Len()
		} else {
			return false
		}
		return length >= min && length <= max
	}
}

// ============================================
// 方案4: 消除 length 中间变量，直接返回比较结果
// ============================================
func lengthOpt4_NoIntermediate(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			l := field.Len()
			return l >= min && l <= max
		case reflect.Slice, reflect.Map, reflect.Array:
			l := field.Len()
			return l >= min && l <= max
		default:
			return false
		}
	}
}

// ============================================
// 方案5: 提前计算 min > max 的情况
// ============================================
func lengthOpt5_PreCheck(min, max int) ValidatorFunc {
	if min > max {
		min, max = max, min
	}
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			l := field.Len()
			return l >= min && l <= max
		case reflect.Slice, reflect.Map, reflect.Array:
			l := field.Len()
			return l >= min && l <= max
		default:
			return false
		}
	}
}

// ============================================
// 方案6: 分离 String 和 Container 路径
// ============================================
func lengthOpt6_SeparatePaths(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		kind := field.Kind()
		if kind == reflect.String {
			l := field.Len()
			return l >= min && l <= max
		}
		if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
			l := field.Len()
			return l >= min && l <= max
		}
		return false
	}
}

// ============================================
// 方案7: 使用 fast-path 模式（优先检查常见类型）
// ============================================
func lengthOpt7_FastPath(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		if field.Kind() == reflect.String {
			l := field.Len()
			return l >= min && l <= max
		}
		k := field.Kind()
		if k == reflect.Slice || k == reflect.Map || k == reflect.Array {
			l := field.Len()
			return l >= min && l <= max
		}
		return false
	}
}

// ============================================
// 方案8: 内联所有比较
// ============================================
func lengthOpt8_Inlined(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			l := field.Len()
			if l < min {
				return false
			}
			return l <= max
		case reflect.Slice, reflect.Map, reflect.Array:
			l := field.Len()
			if l < min {
				return false
			}
			return l <= max
		default:
			return false
		}
	}
}

// ============================================
// 方案9: 使用 goto 减少跳转（实验性）
// ============================================
func lengthOpt9_Goto(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		var l int
		switch field.Kind() {
		case reflect.String:
			l = field.Len()
			goto check
		case reflect.Slice, reflect.Map, reflect.Array:
			l = field.Len()
			goto check
		default:
			return false
		}
	check:
		if l < min {
			return false
		}
		return l <= max
	}
}

// ============================================
// 方案10: 混合优化（缓存 + if-else + 消除中间变量 + 直接Len）
// ============================================
func lengthOpt10_Hybrid(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		k := field.Kind()
		if k == reflect.String {
			l := field.Len()
			return l >= min && l <= max
		}
		if k == reflect.Slice || k == reflect.Map || k == reflect.Array {
			l := field.Len()
			return l >= min && l <= max
		}
		return false
	}
}

// ============================================
// 方案11: 极简模式（单次 Kind 检查 + 直接返回）
// ============================================
func lengthOpt11_Minimal(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		f := fl.Field()
		switch f.Kind() {
		case reflect.String:
			l := f.Len()
			return l >= min && l <= max
		case reflect.Slice, reflect.Map, reflect.Array:
			l := f.Len()
			return l >= min && l <= max
		default:
			return false
		}
	}
}

// ============================================
// 方案12: 短路优化（先检查 min）
// ============================================
func lengthOpt12_ShortCircuit(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			l := field.Len()
			return l >= min && l <= max
		case reflect.Slice, reflect.Map, reflect.Array:
			l := field.Len()
			return l >= min && l <= max
		default:
			return false
		}
	}
}

// ============================================
// 方案13: 预计算比较结果（避免重复计算）
// ============================================
func lengthOpt13_Precompute(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		var l int
		switch field.Kind() {
		case reflect.String:
			l = field.Len()
		case reflect.Slice, reflect.Map, reflect.Array:
			l = field.Len()
		default:
			return false
		}
		return l >= min && l <= max
	}
}

// ============================================
// 方案14: 使用位运算优化边界检查（实验性）
// ============================================
func lengthOpt14_BitOps(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			l := field.Len()
			return (l-min)>>31 == 0 && (max-l)>>31 == 0
		case reflect.Slice, reflect.Map, reflect.Array:
			l := field.Len()
			return (l-min)>>31 == 0 && (max-l)>>31 == 0
		default:
			return false
		}
	}
}

// ============================================
// 方案15: 完全展开 switch（避免共享代码）
// ============================================
func lengthOpt15_Expanded(min, max int) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()
		switch field.Kind() {
		case reflect.String:
			l := field.Len()
			if l < min {
				return false
			}
			if l > max {
				return false
			}
			return true
		case reflect.Slice:
			l := field.Len()
			if l < min {
				return false
			}
			if l > max {
				return false
			}
			return true
		case reflect.Map:
			l := field.Len()
			if l < min {
				return false
			}
			if l > max {
				return false
			}
			return true
		case reflect.Array:
			l := field.Len()
			if l < min {
				return false
			}
			if l > max {
				return false
			}
			return true
		default:
			return false
		}
	}
}

// ============================================
// 基准测试
// ============================================

var lengthCases = generateLengthTestCases()

func benchLengthFunc(b *testing.B, fn ValidatorFunc, cases []LengthTestCase) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			fl := mockFieldLevel{field: tc.field}
			result := fn(fl)
			if result != tc.expected {
				b.Fatalf("%s: expected %v, got %v", tc.name, tc.expected, result)
			}
		}
	}
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

// ============================================
// 正确性测试
// ============================================
func TestLengthOptimizations_Correctness(t *testing.T) {
	min, max := 1, 10
	testCases := []LengthTestCase{
		{"empty_string", 0, 0, reflect.ValueOf(""), false},
		{"valid_string", 0, 0, reflect.ValueOf("hello"), true},
		{"short_valid", 0, 0, reflect.ValueOf("hi"), true},
		{"too_short", 0, 0, reflect.ValueOf(""), false},
		{"too_long", 0, 0, reflect.ValueOf("hello world!"), false},
		{"valid_slice", 0, 0, reflect.ValueOf([]int{1, 2, 3}), true},
		{"empty_slice", 0, 0, reflect.ValueOf([]int{}), false},
		{"valid_map", 0, 0, reflect.ValueOf(map[string]int{"a": 1, "b": 2}), true},
		{"invalid_type", 0, 0, reflect.ValueOf(123), false},
	}

	functions := []struct {
		name string
		fn   ValidatorFunc
	}{
		{"Original", lengthOriginal(min, max)},
		{"Opt1_CacheKind", lengthOpt1_CacheKind(min, max)},
		{"Opt2_StringDirect", lengthOpt2_StringDirect(min, max)},
		{"Opt3_IfElse", lengthOpt3_IfElse(min, max)},
		{"Opt4_NoIntermediate", lengthOpt4_NoIntermediate(min, max)},
		{"Opt5_PreCheck", lengthOpt5_PreCheck(min, max)},
		{"Opt6_SeparatePaths", lengthOpt6_SeparatePaths(min, max)},
		{"Opt7_FastPath", lengthOpt7_FastPath(min, max)},
		{"Opt8_Inlined", lengthOpt8_Inlined(min, max)},
		{"Opt9_Goto", lengthOpt9_Goto(min, max)},
		{"Opt10_Hybrid", lengthOpt10_Hybrid(min, max)},
		{"Opt11_Minimal", lengthOpt11_Minimal(min, max)},
		{"Opt12_ShortCircuit", lengthOpt12_ShortCircuit(min, max)},
		{"Opt13_Precompute", lengthOpt13_Precompute(min, max)},
		{"Opt14_BitOps", lengthOpt14_BitOps(min, max)},
		{"Opt15_Expanded", lengthOpt15_Expanded(min, max)},
	}

	for _, tf := range functions {
		t.Run(tf.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					fl := mockFieldLevel{field: tc.field}
					result := tf.fn(fl)
					if result != tc.expected {
						t.Errorf("%s(%s): expected %v, got %v", tf.name, tc.name, tc.expected, result)
					}
				})
			}
		})
	}
}
