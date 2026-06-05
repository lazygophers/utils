package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ========== isFieldNotEmpty 优化方案 ==========

// 原始实现
func isFieldNotEmptyOriginal(field reflect.Value) bool {
	if !field.IsValid() {
		return false
	}

	switch field.Kind() {
	case reflect.String:
		return field.String() != ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface:
		return !field.IsNil() && isFieldNotEmptyOriginal(field.Elem())
	case reflect.Bool:
		return field.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return field.Float() != 0
	default:
		return !field.IsZero()
	}
}

// 方案1: 快速路径优化
func isFieldNotEmptyFastPath(field reflect.Value) bool {
	if !field.IsValid() {
		return false
	}

	kind := field.Kind()

	// 快速路径: 最常见类型
	if kind == reflect.String {
		return field.String() != ""
	}
	if kind == reflect.Int {
		return field.Int() != 0
	}
	if kind == reflect.Ptr {
		return !field.IsNil() && isFieldNotEmptyFastPath(field.Elem())
	}

	// 其他类型
	switch kind {
	case reflect.Slice, reflect.Map, reflect.Array:
		return field.Len() > 0
	case reflect.Interface:
		return !field.IsNil() && isFieldNotEmptyFastPath(field.Elem())
	case reflect.Bool:
		return field.Bool()
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return field.Float() != 0
	default:
		return !field.IsZero()
	}
}

// 方案2: 减少分支
func isFieldNotEmptyReducedBranches(field reflect.Value) bool {
	if !field.IsValid() {
		return false
	}

	kind := field.Kind()

	// 字符串
	if kind == reflect.String {
		return field.String() != ""
	}

	// 整数类型（有符号和无符号分开处理）
	if kind >= reflect.Int && kind <= reflect.Int64 {
		return field.Int() != 0
	}
	if kind >= reflect.Uint && kind <= reflect.Uint64 {
		return field.Uint() != 0
	}

	// 浮点
	if kind >= reflect.Float32 && kind <= reflect.Float64 {
		return field.Float() != 0
	}

	// 集合类型
	if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
		return field.Len() > 0
	}

	// 指针和接口
	if kind == reflect.Ptr || kind == reflect.Interface {
		return !field.IsNil() && isFieldNotEmptyReducedBranches(field.Elem())
	}

	// 布尔
	if kind == reflect.Bool {
		return field.Bool()
	}

	return !field.IsZero()
}

// 方案3: 混合优化（最优方案）
func isFieldNotEmptyHybrid(field reflect.Value) bool {
	if !field.IsValid() {
		return false
	}

	kind := field.Kind()

	// 超快速路径: 最常见类型
	switch kind {
	case reflect.String:
		return field.String() != ""
	case reflect.Int:
		return field.Int() != 0
	case reflect.Ptr:
		if field.IsNil() {
			return false
		}
		return isFieldNotEmptyHybrid(field.Elem())
	}

	// 范围检查优化
	if kind >= reflect.Int8 && kind <= reflect.Int64 {
		return field.Int() != 0
	}
	if kind >= reflect.Uint8 && kind <= reflect.Uint64 {
		return field.Uint() != 0
	}
	if kind == reflect.Float32 || kind == reflect.Float64 {
		return field.Float() != 0
	}
	if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
		return field.Len() > 0
	}
	if kind == reflect.Interface {
		return !field.IsNil() && isFieldNotEmptyHybrid(field.Elem())
	}
	if kind == reflect.Bool {
		return field.Bool()
	}

	return !field.IsZero()
}

// ========== getFieldValueAsString 优化方案 ==========

// 原始实现
func getFieldValueAsStringOriginal(field reflect.Value) string {
	if !field.IsValid() {
		return ""
	}

	switch field.Kind() {
	case reflect.String:
		return field.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(field.Bool())
	default:
		return fmt.Sprintf("%v", field.Interface())
	}
}

// 方案1: 快速路径
func getFieldValueAsStringFastPath(field reflect.Value) string {
	if !field.IsValid() {
		return ""
	}

	kind := field.Kind()

	// 快速路径: 字符串直接返回
	if kind == reflect.String {
		return field.String()
	}

	// 快速路径: 布尔值
	if kind == reflect.Bool {
		if field.Bool() {
			return "true"
		}
		return "false"
	}

	// 其他类型
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(field.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(field.Float(), 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", field.Interface())
	}
}

// 方案2: 混合优化（最优方案）
func getFieldValueAsStringHybrid(field reflect.Value) string {
	if !field.IsValid() {
		return ""
	}

	kind := field.Kind()

	// 超快速路径
	switch kind {
	case reflect.String:
		return field.String()
	case reflect.Bool:
		if field.Bool() {
			return "true"
		}
		return "false"
	}

	// 整数优化
	if kind >= reflect.Int8 && kind <= reflect.Int64 {
		i := field.Int()
		if i >= 0 && i < 100 {
			return smallIntStrings[i]
		}
		return strconv.FormatInt(i, 10)
	}
	if kind >= reflect.Uint8 && kind <= reflect.Uint64 {
		u := field.Uint()
		if u < 100 {
			return smallIntStrings[u]
		}
		return strconv.FormatUint(u, 10)
	}
	if kind == reflect.Float32 || kind == reflect.Float64 {
		return strconv.FormatFloat(field.Float(), 'f', -1, 64)
	}

	return fmt.Sprintf("%v", field.Interface())
}

// 预分配小整数字符串表
var smallIntStrings = [100]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
	"20", "21", "22", "23", "24", "25", "26", "27", "28", "29",
	"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
	"40", "41", "42", "43", "44", "45", "46", "47", "48", "49",
	"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
	"60", "61", "62", "63", "64", "65", "66", "67", "68", "69",
	"70", "71", "72", "73", "74", "75", "76", "77", "78", "79",
	"80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
	"90", "91", "92", "93", "94", "95", "96", "97", "98", "99",
}

// ========== compareFields 优化方案 ==========

// 原始实现
func compareFieldsOriginal(current, target reflect.Value) int {
	if !current.IsValid() || !target.IsValid() {
		return 0
	}

	switch current.Kind() {
	case reflect.String:
		currentStr := current.String()
		targetStr := target.String()
		if currentStr == targetStr {
			return 0
		}
		if currentStr < targetStr {
			return -1
		}
		return 1
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		currentInt := current.Int()
		targetInt := target.Int()
		if currentInt == targetInt {
			return 0
		}
		if currentInt < targetInt {
			return -1
		}
		return 1
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		currentUint := current.Uint()
		targetUint := target.Uint()
		if currentUint == targetUint {
			return 0
		}
		if currentUint < targetUint {
			return -1
		}
		return 1
	case reflect.Float32, reflect.Float64:
		currentFloat := current.Float()
		targetFloat := target.Float()
		if currentFloat == targetFloat {
			return 0
		}
		if currentFloat < targetFloat {
			return -1
		}
		return 1
	case reflect.Bool:
		currentBool := current.Bool()
		targetBool := target.Bool()
		if currentBool == targetBool {
			return 0
		}
		if !currentBool && targetBool {
			return -1
		}
		return 1
	case reflect.Ptr, reflect.Interface:
		if current.IsNil() && target.IsNil() {
			return 0
		}
		if current.IsNil() {
			return -1
		}
		if target.IsNil() {
			return 1
		}
		return compareFieldsOriginal(current.Elem(), target.Elem())
	default:
		currentStr := getFieldValueAsStringOriginal(current)
		targetStr := getFieldValueAsStringOriginal(target)
		if currentStr == targetStr {
			return 0
		}
		if currentStr < targetStr {
			return -1
		}
		return 1
	}
}

// 方案1: 快速路径
func compareFieldsFastPath(current, target reflect.Value) int {
	if !current.IsValid() || !target.IsValid() {
		return 0
	}

	kind := current.Kind()

	// 快速路径: 字符串
	if kind == reflect.String {
		cs := current.String()
		ts := target.String()
		if cs == ts {
			return 0
		}
		if cs < ts {
			return -1
		}
		return 1
	}

	// 快速路径: 整数
	if kind == reflect.Int {
		ci := current.Int()
		ti := target.Int()
		if ci == ti {
			return 0
		}
		if ci < ti {
			return -1
		}
		return 1
	}

	// 其他类型
	switch kind {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ci := current.Int()
		ti := target.Int()
		if ci == ti {
			return 0
		}
		if ci < ti {
			return -1
		}
		return 1
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		cu := current.Uint()
		tu := target.Uint()
		if cu == tu {
			return 0
		}
		if cu < tu {
			return -1
		}
		return 1
	case reflect.Float32, reflect.Float64:
		cf := current.Float()
		tf := target.Float()
		if cf == tf {
			return 0
		}
		if cf < tf {
			return -1
		}
		return 1
	case reflect.Bool:
		cb := current.Bool()
		tb := target.Bool()
		if cb == tb {
			return 0
		}
		if !cb && tb {
			return -1
		}
		return 1
	case reflect.Ptr, reflect.Interface:
		if current.IsNil() && target.IsNil() {
			return 0
		}
		if current.IsNil() {
			return -1
		}
		if target.IsNil() {
			return 1
		}
		return compareFieldsFastPath(current.Elem(), target.Elem())
	default:
		cs := getFieldValueAsStringFastPath(current)
		ts := getFieldValueAsStringFastPath(target)
		if cs == ts {
			return 0
		}
		if cs < ts {
			return -1
		}
		return 1
	}
}

// 方案2: 混合优化（最优方案）
func compareFieldsHybrid(current, target reflect.Value) int {
	if !current.IsValid() || !target.IsValid() {
		return 0
	}

	kind := current.Kind()

	// 超快速路径: 最常见类型
	switch kind {
	case reflect.String:
		cs, ts := current.String(), target.String()
		if cs == ts {
			return 0
		}
		if cs < ts {
			return -1
		}
		return 1
	case reflect.Int:
		ci, ti := current.Int(), target.Int()
		if ci == ti {
			return 0
		}
		if ci < ti {
			return -1
		}
		return 1
	case reflect.Ptr:
		if current.IsNil() && target.IsNil() {
			return 0
		}
		if current.IsNil() {
			return -1
		}
		if target.IsNil() {
			return 1
		}
		return compareFieldsHybrid(current.Elem(), target.Elem())
	}

	// 范围检查优化
	if kind >= reflect.Int8 && kind <= reflect.Int64 {
		ci, ti := current.Int(), target.Int()
		if ci == ti {
			return 0
		}
		if ci < ti {
			return -1
		}
		return 1
	}
	if kind >= reflect.Uint8 && kind <= reflect.Uint64 {
		cu, tu := current.Uint(), target.Uint()
		if cu == tu {
			return 0
		}
		if cu < tu {
			return -1
		}
		return 1
	}
	if kind == reflect.Float32 || kind == reflect.Float64 {
		cf, tf := current.Float(), target.Float()
		if cf == tf {
			return 0
		}
		if cf < tf {
			return -1
		}
		return 1
	}
	if kind == reflect.Bool {
		cb, tb := current.Bool(), target.Bool()
		if cb == tb {
			return 0
		}
		if !cb && tb {
			return -1
		}
		return 1
	}
	if kind == reflect.Interface {
		if current.IsNil() && target.IsNil() {
			return 0
		}
		if current.IsNil() {
			return -1
		}
		if target.IsNil() {
			return 1
		}
		return compareFieldsHybrid(current.Elem(), target.Elem())
	}

	// 默认
	cs, ts := getFieldValueAsStringFastPath(current), getFieldValueAsStringFastPath(target)
	if cs == ts {
		return 0
	}
	if cs < ts {
		return -1
	}
	return 1
}

// ========== 测试数据 ==========

func generateBenchCases() []reflect.Value {
	str := "hello"
	num := 42
	flt := 3.14
	bl := true
	slice := []int{1, 2, 3}
	m := map[string]int{"a": 1}

	return []reflect.Value{
		reflect.ValueOf(str),
		reflect.ValueOf(""),
		reflect.ValueOf(num),
		reflect.ValueOf(int8(8)),
		reflect.ValueOf(int16(16)),
		reflect.ValueOf(int32(32)),
		reflect.ValueOf(int64(64)),
		reflect.ValueOf(uint(42)),
		reflect.ValueOf(uint8(8)),
		reflect.ValueOf(uint16(16)),
		reflect.ValueOf(uint32(32)),
		reflect.ValueOf(uint64(64)),
		reflect.ValueOf(float32(3.14)),
		reflect.ValueOf(flt),
		reflect.ValueOf(bl),
		reflect.ValueOf(slice),
		reflect.ValueOf([]int{}),
		reflect.ValueOf(m),
		reflect.ValueOf(map[string]int{}),
		reflect.ValueOf(&str),
		reflect.ValueOf((*string)(nil)),
	}
}

// ========== Benchmark 测试 ==========

// ========== 正确性测试 ==========

func TestOptimizationCorrectness(t *testing.T) {
	cases := generateBenchCases()

	t.Run("isFieldNotEmpty", func(t *testing.T) {
		for _, v := range cases {
			original := isFieldNotEmptyOriginal(v)
			if got := isFieldNotEmptyHybrid(v); got != original {
				t.Errorf("Hybrid mismatch: got %v, want %v", got, original)
			}
		}
	})

	t.Run("getFieldValueAsString", func(t *testing.T) {
		for _, v := range cases {
			original := getFieldValueAsStringOriginal(v)
			if got := getFieldValueAsStringHybrid(v); got != original {
				t.Errorf("Hybrid mismatch: got %q, want %q", got, original)
			}
		}
	})

	t.Run("compareFields", func(t *testing.T) {
		for _, v := range cases {
			original := compareFieldsOriginal(v, v)
			if got := compareFieldsHybrid(v, v); got != original {
				t.Errorf("Hybrid mismatch: got %d, want %d", got, original)
			}
		}
	})
}

// ========== 场景化 Benchmarks ==========

// 验证优化后的 Pattern 函数功能正确性
func TestPatternOptimization(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		value    string
		expected bool
	}{
		{
			name:     "有效邮箱",
			pattern:  `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
			value:    "test@example.com",
			expected: true,
		},
		{
			name:     "无效邮箱-缺少@",
			pattern:  `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
			value:    "invalid",
			expected: false,
		},
		{
			name:     "固定长度-5位数字",
			pattern:  `^\d{5}$`,
			value:    "12345",
			expected: true,
		},
		{
			name:     "固定长度-不足5位",
			pattern:  `^\d{5}$`,
			value:    "123",
			expected: false,
		},
		{
			name:     "字面量匹配",
			pattern:  `^hello$`,
			value:    "hello",
			expected: true,
		},
		{
			name:     "字面量不匹配",
			pattern:  `^hello$`,
			value:    "world",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := Pattern(tt.pattern)
			fl := &testFieldLevel{value: reflect.ValueOf(tt.value)}
			result := validator(fl)
			if result != tt.expected {
				t.Errorf("Pattern() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 测试用 FieldLevel 实现
type testFieldLevel struct {
	value reflect.Value
}

func (t *testFieldLevel) Top() reflect.Value {
	return t.value
}

func (t *testFieldLevel) Parent() reflect.Value {
	return t.value
}

func (t *testFieldLevel) Field() reflect.Value {
	return t.value
}

func (t *testFieldLevel) FieldName() string {
	return "test"
}

func (t *testFieldLevel) StructFieldName() string {
	return "Test"
}

func (t *testFieldLevel) Param() string {
	return ""
}

func (t *testFieldLevel) GetTag(key string) string {
	return ""
}

func (t *testFieldLevel) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// 性能对比基准测试

// 测试数据
var (
	testCases = []struct {
		name  string
		field FieldLevel
		want  bool
	}{
		{"empty_string", testFL{reflect.ValueOf("")}, false},
		{"nonempty_string", testFL{reflect.ValueOf("hello")}, true},
		{"empty_slice", testFL{reflect.ValueOf([]int{})}, false},
		{"nonempty_slice", testFL{reflect.ValueOf([]int{1})}, true},
		{"empty_map", testFL{reflect.ValueOf(map[string]int{})}, false},
		{"nonempty_map", testFL{reflect.ValueOf(map[string]int{"a": 1})}, true},
		{"nil_ptr", testFL{reflect.ValueOf((*int)(nil))}, false},
		{"nonnil_ptr", testFL{reflect.ValueOf(ptr(42))}, true},
		{"zero_int", testFL{reflect.ValueOf(0)}, false},
		{"nonzero_int", testFL{reflect.ValueOf(42)}, true},
		{"nil_interface", testFL{reflect.ValueOf(nil)}, false},
		{"nonnil_interface", testFL{reflect.ValueOf(42)}, true},
	}
)

type testFL struct{ field reflect.Value }

func (t testFL) Field() reflect.Value                { return t.field }
func (t testFL) Top() reflect.Value                  { return reflect.Value{} }
func (t testFL) Parent() reflect.Value               { return reflect.Value{} }
func (t testFL) FieldName() string                   { return "" }
func (t testFL) StructFieldName() string             { return "" }
func (t testFL) Param() string                       { return "" }
func (t testFL) GetTag(string) string                { return "" }
func (t testFL) GetFieldByName(string) reflect.Value { return reflect.Value{} }

func ptr(v int) *int { return &v }

// 原始实现
func reqOrig(fl FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.String:
		return field.String() != ""
	case reflect.Slice, reflect.Map, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface:
		return !field.IsNil()
	default:
		return field.IsValid() && !field.IsZero()
	}
}

// FastPath 优化
func reqFast(fl FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	if kind == reflect.String {
		return field.String() != ""
	}
	if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
		return field.Len() > 0
	}
	if kind == reflect.Ptr || kind == reflect.Interface {
		return !field.IsNil()
	}
	return field.IsValid() && !field.IsZero()
}

// 分离变量优化
func reqSep(fl FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	if kind == reflect.String {
		s := field.String()
		return s != ""
	}
	if kind == reflect.Slice || kind == reflect.Map || kind == reflect.Array {
		l := field.Len()
		return l > 0
	}
	if kind == reflect.Ptr || kind == reflect.Interface {
		return !field.IsNil()
	}
	return field.IsValid() && !field.IsZero()
}

func TestRequiredCorrectness(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := reqOrig(tc.field); got != tc.want {
				t.Errorf("reqOrig() = %v, want %v", got, tc.want)
			}
			if got := reqFast(tc.field); got != tc.want {
				t.Errorf("reqFast() = %v, want %v", got, tc.want)
			}
			if got := reqSep(tc.field); got != tc.want {
				t.Errorf("reqSep() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestParseTag 测试 parseTag 功能正确性
func TestParseTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		expected []validationRule
	}{
		{
			name: "简单标签",
			tag:  "required,email,max=100",
			expected: []validationRule{
				{tag: "required", param: ""},
				{tag: "email", param: ""},
				{tag: "max", param: "100"},
			},
		},
		{
			name: "带空格标签",
			tag:  "required , email , max = 100",
			expected: []validationRule{
				{tag: "required", param: ""},
				{tag: "email", param: ""},
				{tag: "max", param: "100"},
			},
		},
		{
			name: "复杂标签",
			tag:  "required,email,min=18,max=100,len=6-20",
			expected: []validationRule{
				{tag: "required", param: ""},
				{tag: "email", param: ""},
				{tag: "min", param: "18"},
				{tag: "max", param: "100"},
				{tag: "len", param: "6-20"},
			},
		},
		{
			name:     "空标签",
			tag:      "",
			expected: []validationRule(nil),
		},
		{
			name:     "只有空格",
			tag:      "   ,  ,   ",
			expected: []validationRule(nil),
		},
		{
			name: "带参数值有空格",
			tag:  "regex=^[a-z]+$ , url , max = 100",
			expected: []validationRule{
				{tag: "regex", param: "^[a-z]+$"},
				{tag: "url", param: ""},
				{tag: "max", param: "100"},
			},
		},
		{
			name: "单规则",
			tag:  "required",
			expected: []validationRule{
				{tag: "required", param: ""},
			},
		},
		{
			name: "多参数规则",
			tag:  "in=1,2,3,notin=4,5,6",
			expected: []validationRule{
				{tag: "in", param: "1"},
				{tag: "2", param: ""},
				{tag: "3", param: ""},
				{tag: "notin", param: "4"},
				{tag: "5", param: ""},
				{tag: "6", param: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Engine{}
			result := e.parseTag(tt.tag)

			// 验证结果长度
			assert.Equal(t, len(tt.expected), len(result), "结果长度不匹配")

			// 验证每个规则
			for i := range tt.expected {
				assert.Equal(t, tt.expected[i].tag, result[i].tag, "规则 %d: tag 不匹配", i)
				assert.Equal(t, tt.expected[i].param, result[i].param, "规则 %d: param 不匹配", i)
			}
		})
	}
}

// TestParseTagPerformance 测试 parseTag 性能
func TestParseTagPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过性能测试")
	}

	e := &Engine{}
	tag := "required,email,min=18,max=100,len=6-20,in=1,2,3,notin=4,5,6,regex=^[a-z]+$"

	// 预热
	for i := 0; i < 1000; i++ {
		_ = e.parseTag(tag)
	}

	// 测试
	start := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = e.parseTag(tag)
		}
	})

	t.Logf("性能测试结果: %v", start)
}

// TestAndComposer 测试 And 组合验证器
func TestAndComposer(t *testing.T) {
	v, _ := New()

	// 注册组合验证器：密码长度 8-20 且包含特殊字符
	v.RegisterValidation("strong_pwd", And(
		Length(8, 20),
		ContainsSpecial(),
	))

	type User struct {
		Password string `validate:"strong_pwd"`
	}

	// 测试有效密码
	valid := User{Password: "abc123!@"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试太短
	tooShort := User{Password: "ab1!"}
	err = v.Struct(tooShort)
	if err == nil {
		t.Error("Expected error for too short password")
	}

	// 测试缺少特殊字符
	noSpecial := User{Password: "abc12345"}
	err = v.Struct(noSpecial)
	if err == nil {
		t.Error("Expected error for password without special char")
	}
}

// TestOrComposer 测试 Or 组合验证器
func TestOrComposer(t *testing.T) {
	v, _ := New()

	// 注册组合验证器：可以是手机号或邮箱
	v.RegisterValidation("phone_or_email", Or(
		Pattern(`^1[3-9]\d{9}$`),
		Email(),
	))

	type Contact struct {
		Contact string `validate:"phone_or_email"`
	}

	// 测试手机号
	phone := Contact{Contact: "13812345678"}
	err := v.Struct(phone)
	if err != nil {
		t.Errorf("Expected valid phone, got error: %v", err)
	}

	// 测试邮箱
	email := Contact{Contact: "test@example.com"}
	err = v.Struct(email)
	if err != nil {
		t.Errorf("Expected valid email, got error: %v", err)
	}

	// 测试无效
	invalid := Contact{Contact: "invalid"}
	err = v.Struct(invalid)
	if err == nil {
		t.Error("Expected error for invalid contact")
	}
}

// TestNotComposer 测试 Not 组合验证器
func TestNotComposer(t *testing.T) {
	v, _ := New()

	// 注册组合验证器：不能是 admin
	v.RegisterValidation("not_admin", Not(
		In("admin", "root", "system"),
	))

	type User struct {
		Username string `validate:"not_admin"`
	}

	// 测试有效用户
	valid := User{Username: "user123"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效用户名
	invalid := User{Username: "admin"}
	err = v.Struct(invalid)
	if err == nil {
		t.Error("Expected error for admin username")
	}
}

// TestRangeComposer 测试范围验证器
func TestRangeComposer(t *testing.T) {
	v, _ := New()

	type Product struct {
		Price float64 `validate:"range=0,10000"`
	}

	// 注册范围验证器
	v.RegisterValidation("range", Range(0, 10000))

	// 测试有效价格
	valid := Product{Price: 99.99}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试超出范围
	tooHigh := Product{Price: 15000}
	err = v.Struct(tooHigh)
	if err == nil {
		t.Error("Expected error for price too high")
	}
}

// TestLengthComposer 测试长度验证器
func TestLengthComposer(t *testing.T) {
	v, _ := New()

	type User struct {
		Name string `validate:"length=2,10"`
	}

	// 注册长度验证器
	v.RegisterValidation("length", Length(2, 10))

	// 测试有效长度
	valid := User{Name: "Alice"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试太短
	tooShort := User{Name: "A"}
	err = v.Struct(tooShort)
	if err == nil {
		t.Error("Expected error for name too short")
	}

	// 测试太长
	tooLong := User{Name: "ThisNameIsWayTooLong"}
	err = v.Struct(tooLong)
	if err == nil {
		t.Error("Expected error for name too long")
	}
}

// TestChainedComposition 测试链式组合
func TestChainedComposition(t *testing.T) {
	v, _ := New()

	// 复杂组合：用户名必须 3-20 字符，字母开头，只含字母数字下划线
	v.RegisterValidation("username", And(
		Length(3, 20),
		Pattern(`^[a-zA-Z][a-zA-Z0-9_]*$`),
		Not(In("admin", "root", "system")),
	))

	type User struct {
		Username string `validate:"username"`
	}

	// 测试有效用户名
	validCases := []string{"user123", "Alice", "bob_the_builder"}
	for _, username := range validCases {
		user := User{Username: username}
		err := v.Struct(user)
		if err != nil {
			t.Errorf("Expected valid for %s, got error: %v", username, err)
		}
	}

	// 测试无效用户名
	invalidCases := []struct {
		username string
		reason   string
	}{
		{"ab", "too short"},
		{"1invalid", "must start with letter"},
		{"admin", "reserved name"},
		{"user@name", "invalid chars"},
	}

	for _, tc := range invalidCases {
		user := User{Username: tc.username}
		err := v.Struct(user)
		if err == nil {
			t.Errorf("Expected error for %s (%s)", tc.username, tc.reason)
		}
	}
}

// TestRequiredWithMinLength 测试必填+最小长度组合
func TestRequiredWithMinLength(t *testing.T) {
	v, _ := New()

	type Form struct {
		Password string `validate:"required_min"`
	}

	// 注册组合验证器
	v.RegisterValidation("required_min", And(
		Required(),
		MinLength(8),
	))

	// 测试有效密码
	valid := Form{Password: "password123"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试空密码
	empty := Form{Password: ""}
	err = v.Struct(empty)
	if err == nil {
		t.Error("Expected error for empty password")
	}

	// 测试太短
	tooShort := Form{Password: "short"}
	err = v.Struct(tooShort)
	if err == nil {
		t.Error("Expected error for short password")
	}
}

// TestInNotInComposer 测试 In/NotIn 验证器
func TestInNotInComposer(t *testing.T) {
	v, _ := New()

	type Product struct {
		Category string `validate:"valid_category"`
		Size     string `validate:"not_reserved_size"`
	}

	// 注册验证器
	v.RegisterValidation("valid_category", In("electronics", "clothing", "food"))
	v.RegisterValidation("not_reserved_size", NotIn("xs", "xl"))

	// 测试有效产品
	valid := Product{Category: "electronics", Size: "m"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效分类
	invalidCat := Product{Category: "invalid", Size: "m"}
	err = v.Struct(invalidCat)
	if err == nil {
		t.Error("Expected error for invalid category")
	}

	// 测试保留尺寸
	reserved := Product{Category: "electronics", Size: "xs"}
	err = v.Struct(reserved)
	if err == nil {
		t.Error("Expected error for reserved size")
	}
}

// TestEmailURLComposer 测试 Email/URL 构造函数
func TestEmailURLComposer(t *testing.T) {
	v, _ := New()

	type Contact struct {
		EmailAddr string `validate:"custom_email"`
		Website   string `validate:"custom_url"`
	}

	// 注册验证器
	v.RegisterValidation("custom_email", Email())
	v.RegisterValidation("custom_url", URL())

	// 测试有效数据
	valid := Contact{EmailAddr: "test@example.com", Website: "https://example.com"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效邮箱
	invalidEmail := Contact{EmailAddr: "invalid", Website: "https://example.com"}
	err = v.Struct(invalidEmail)
	if err == nil {
		t.Error("Expected error for invalid email")
	}

	// 测试无效 URL
	invalidURL := Contact{EmailAddr: "test@example.com", Website: "invalid"}
	err = v.Struct(invalidURL)
	if err == nil {
		t.Error("Expected error for invalid URL")
	}
}

// TestPatternComposer 测试正则表达式构造函数
func TestPatternComposer(t *testing.T) {
	v, _ := New()

	type User struct {
		ZipCode string `validate:"zip_code"`
	}

	// 注册验证器
	v.RegisterValidation("zip_code", Pattern(`^\d{5}$`))

	// 测试有效邮编
	valid := User{ZipCode: "12345"}
	err := v.Struct(valid)
	if err != nil {
		t.Errorf("Expected valid, got error: %v", err)
	}

	// 测试无效邮编
	invalid := User{ZipCode: "abc"}
	err = v.Struct(invalid)
	if err == nil {
		t.Error("Expected error for invalid zip code")
	}
}

// 全面性能测试 - 测试不同场景
func TestComprehensivePerf(t *testing.T) {
	fmt.Println("\n================== 全面性能测试 ==================")

	// 场景1: 非短路 And（所有验证器都执行）
	fmt.Println("\n--- 场景1: 非短路 And（所有验证器都执行） ---")
	testAndNoShortCircuit(t)

	// 场景2: 大量验证器（20个）
	fmt.Println("\n--- 场景2: 大量验证器 And（20个） ---")
	testAndManyValidators(t)

	// 场景3: 非短路 Or
	fmt.Println("\n--- 场景3: 非短路 Or（所有验证器都执行） ---")
	testOrNoShortCircuit(t)

	// 场景4: 大量验证器 Or（20个）
	fmt.Println("\n--- 场景4: 大量验证器 Or（20个） ---")
	testOrManyValidators(t)

	fmt.Println("\n================== 测试完成 ==================")
}

func testAndNoShortCircuit(t *testing.T) {
	validators := make([]ValidatorFunc, 5)
	for i := 0; i < 4; i++ {
		validators[i] = func(fl FieldLevel) bool { return true }
	}
	validators[4] = func(fl FieldLevel) bool { return false }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := And(validators...)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := AndIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

func testAndManyValidators(t *testing.T) {
	validators := make([]ValidatorFunc, 20)
	validators[0] = func(fl FieldLevel) bool { return false }
	for i := 1; i < 20; i++ {
		validators[i] = func(fl FieldLevel) bool { return true }
	}

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := And(validators...)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := AndIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

func testOrNoShortCircuit(t *testing.T) {
	validators := make([]ValidatorFunc, 5)
	for i := 0; i < 4; i++ {
		validators[i] = func(fl FieldLevel) bool { return false }
	}
	validators[4] = func(fl FieldLevel) bool { return true }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := Or(validators...)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := OrIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

func testOrManyValidators(t *testing.T) {
	validators := make([]ValidatorFunc, 20)
	validators[0] = func(fl FieldLevel) bool { return true }
	for i := 1; i < 20; i++ {
		validators[i] = func(fl FieldLevel) bool { return false }
	}

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 原始版本
	vOriginal := Or(validators...)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 索引循环版本
	vIndex := OrIndexLoop(validators...)
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
}

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

func TestAlphaOptimized(t *testing.T) {
	validator := Alpha()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"纯字母", "HelloWorld", true},
		{"混合大小写", "HeLLoWoRLd", true},
		{"长字符串", "TheQuickBrownFoxJumpsOverTheLazyDog", true},
		{"单字符", "H", true},
		{"空字符串", "", false},
		{"包含数字", "Hello123", false},
		{"包含特殊字符", "Test@123", false},
		{"纯数字", "123", false},
		{"包含空格", "Hello World", false},
		{"包含下划线", "Hello_World", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &fieldLevel{field: reflect.ValueOf(tt.input)}
			result := validator(fl)
			if result != tt.expected {
				t.Errorf("Alpha(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAlphanumOptimized(t *testing.T) {
	validator := Alphanum()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"字母数字", "User123", true},
		{"混合", "AbC123xYz", true},
		{"长字符串", "UserID1234567890ABCDEFghijklmnopQRSTUVWXYZ9876543210", true},
		{"单字符字母", "A", true},
		{"单字符数字", "1", true},
		{"空字符串", "", false},
		{"包含特殊字符", "Test@123", false},
		{"包含空格", "Hello World", false},
		{"包含短横线", "Test-123", false},
		{"包含下划线", "Test_123", false},
		{"纯字母", "HelloWorld", true},
		{"纯数字", "123456", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fl := &fieldLevel{field: reflect.ValueOf(tt.input)}
			result := validator(fl)
			if result != tt.expected {
				t.Errorf("Alphanum(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAlphaPerformance(t *testing.T) {
	validator := Alpha()
	fl := &fieldLevel{field: reflect.ValueOf("HelloWorldTest")}

	// 运行多次确保没有性能退化
	iterations := 100000
	for i := 0; i < iterations; i++ {
		if !validator(fl) {
			t.Error("Alpha validation failed")
		}
	}
}

func TestAlphanumPerformance(t *testing.T) {
	validator := Alphanum()
	fl := &fieldLevel{field: reflect.ValueOf("User123456789")}

	// 运行多次确保没有性能退化
	iterations := 100000
	for i := 0; i < iterations; i++ {
		if !validator(fl) {
			t.Error("Alphanum validation failed")
		}
	}
}

// 手动性能测试 - And 函数对比
func TestManualAndPerf(t *testing.T) {
	// 创建测试验证器
	v1 := func(fl FieldLevel) bool { return false }
	v2 := func(fl FieldLevel) bool { return true }
	v3 := func(fl FieldLevel) bool { return true }
	v4 := func(fl FieldLevel) bool { return true }
	v5 := func(fl FieldLevel) bool { return true }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 测试原始版本（range）
	vOriginal := And(v1, v2, v3, v4, v5)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 测试索引循环版本
	vIndex := AndIndexLoop(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	// 测试 switch 版本
	vSwitch := AndSwitch(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vSwitch(fl)
	}
	switchTime := time.Since(start)

	// 测试 goto 版本
	vGoto := AndGoto(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vGoto(fl)
	}
	gotoTime := time.Since(start)

	// 测试结构体版本
	vStruct := AndStruct(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vStruct(fl)
	}
	structTime := time.Since(start)

	// 输出结果
	fmt.Println("\n========================================")
	fmt.Println("And 函数性能测试结果 (短路场景，5个验证器，1000万次迭代):")
	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
	fmt.Printf("Switch展开:       %v (%.1f%%)\n", switchTime, float64(switchTime)*100/float64(originalTime))
	fmt.Printf("Goto优化:         %v (%.1f%%)\n", gotoTime, float64(gotoTime)*100/float64(originalTime))
	fmt.Printf("结构体方法:       %v (%.1f%%)\n", structTime, float64(structTime)*100/float64(originalTime))
	fmt.Println("========================================")
}

// 手动性能测试 - Or 函数对比
func TestManualOrPerf(t *testing.T) {
	// 创建测试验证器
	v1 := func(fl FieldLevel) bool { return true }
	v2 := func(fl FieldLevel) bool { return false }
	v3 := func(fl FieldLevel) bool { return false }
	v4 := func(fl FieldLevel) bool { return false }
	v5 := func(fl FieldLevel) bool { return false }

	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 测试原始版本（range）
	vOriginal := Or(v1, v2, v3, v4, v5)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 测试索引循环版本
	vIndex := OrIndexLoop(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vIndex(fl)
	}
	indexTime := time.Since(start)

	// 测试 switch 版本
	vSwitch := OrSwitch(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vSwitch(fl)
	}
	switchTime := time.Since(start)

	// 测试 goto 版本
	vGoto := OrGoto(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vGoto(fl)
	}
	gotoTime := time.Since(start)

	// 测试结构体版本
	vStruct := OrStruct(v1, v2, v3, v4, v5)
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		vStruct(fl)
	}
	structTime := time.Since(start)

	// 输出结果
	fmt.Println("\n========================================")
	fmt.Println("Or 函数性能测试结果 (短路场景，5个验证器，1000万次迭代):")
	fmt.Printf("原始版本(range):  %v\n", originalTime)
	fmt.Printf("索引循环:         %v (%.1f%%)\n", indexTime, float64(indexTime)*100/float64(originalTime))
	fmt.Printf("Switch展开:       %v (%.1f%%)\n", switchTime, float64(switchTime)*100/float64(originalTime))
	fmt.Printf("Goto优化:         %v (%.1f%%)\n", gotoTime, float64(gotoTime)*100/float64(originalTime))
	fmt.Printf("结构体方法:       %v (%.1f%%)\n", structTime, float64(structTime)*100/float64(originalTime))
	fmt.Println("========================================")
}

// 手动性能测试 - Not 函数对比
func TestManualNotPerf(t *testing.T) {
	// 创建测试验证器
	v1 := func(fl FieldLevel) bool { return true }
	fl := mockFieldLevel{field: reflect.ValueOf("")}

	// 测试原始版本
	vOriginal := Not(v1)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		vOriginal(fl)
	}
	originalTime := time.Since(start)

	// 输出结果
	fmt.Println("\n========================================")
	fmt.Println("Not 函数性能测试结果 (1000万次迭代):")
	fmt.Printf("原始版本:  %v\n", originalTime)
	fmt.Println("结论: Not 函数已经是最优实现（仅一个取反操作）")
	fmt.Println("========================================")
}

func TestFieldLevelMethods(t *testing.T) {
	t.Run("Top", func(t *testing.T) {
		type TestStruct struct {
			Name string
		}
		s := TestStruct{Name: "test"}
		fl := &fieldLevel{
			top: reflect.ValueOf(s),
		}
		result := fl.Top()
		assert.True(t, result.IsValid())
	})

	t.Run("Parent", func(t *testing.T) {
		type ParentStruct struct {
			Child string
		}
		p := ParentStruct{Child: "test"}
		fl := &fieldLevel{
			parent: reflect.ValueOf(p),
		}
		result := fl.Parent()
		assert.True(t, result.IsValid())
	})

	t.Run("Field", func(t *testing.T) {
		fl := &fieldLevel{
			field: reflect.ValueOf("test"),
		}
		result := fl.Field()
		assert.True(t, result.IsValid())
		assert.Equal(t, "test", result.String())
	})

	t.Run("FieldName", func(t *testing.T) {
		fl := &fieldLevel{
			fieldName: "test_field",
		}
		result := fl.FieldName()
		assert.Equal(t, "test_field", result)
	})

	t.Run("StructFieldName", func(t *testing.T) {
		fl := &fieldLevel{
			structFieldName: "TestField",
		}
		result := fl.StructFieldName()
		assert.Equal(t, "TestField", result)
	})

	t.Run("Param", func(t *testing.T) {
		fl := &fieldLevel{
			param: "10",
		}
		result := fl.Param()
		assert.Equal(t, "10", result)
	})

	t.Run("GetTag", func(t *testing.T) {
		type TestStruct struct {
			Field string `json:"field_name" validate:"required"`
		}
		s := TestStruct{Field: "test"}
		rt := reflect.TypeOf(s)
		fieldType := rt.Field(0)
		fl := &fieldLevel{
			structField: fieldType,
		}

		jsonTag := fl.GetTag("json")
		assert.Equal(t, "field_name", jsonTag)

		validateTag := fl.GetTag("validate")
		assert.Equal(t, "required", validateTag)

		emptyTag := fl.GetTag("nonexistent")
		assert.Equal(t, "", emptyTag)
	})
}

func TestEngineSetTagName(t *testing.T) {
	e := NewEngine()
	e.SetTagName("custom_tag")
	assert.Equal(t, "custom_tag", e.tagName)
}

func TestEngineSetFieldNameFunc(t *testing.T) {
	e := NewEngine()

	customFunc := func(field reflect.StructField) string {
		return "custom_" + field.Name
	}
	e.SetFieldNameFunc(customFunc)
	assert.NotNil(t, e.fieldNameFunc)

	e.SetFieldNameFunc(nil)
	assert.NotNil(t, e.fieldNameFunc)
}

func TestDefaultFieldNameFunc(t *testing.T) {
	t.Run("with_json_tag", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:"field_name"`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "field_name", result)
	})

	t.Run("with_json_tag_omitempty", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:"field_name,omitempty"`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "field_name", result)
	})

	t.Run("with_json_tag_dash", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:"-"`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "FieldName", result)
	})

	t.Run("without_json_tag", func(t *testing.T) {
		type TestStruct struct {
			FieldName string
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "FieldName", result)
	})

	t.Run("with_empty_json_tag", func(t *testing.T) {
		type TestStruct struct {
			FieldName string `json:""`
		}
		rt := reflect.TypeOf(TestStruct{})
		field := rt.Field(0)
		result := defaultFieldNameFunc(field)
		assert.Equal(t, "FieldName", result)
	})
}

func TestStructFieldNameFunc(t *testing.T) {
	type TestStruct struct {
		FieldName string
	}
	rt := reflect.TypeOf(TestStruct{})
	field := rt.Field(0)
	result := structFieldNameFunc(field)
	assert.Equal(t, "FieldName", result)
}

func TestEngineWithCustomTagName(t *testing.T) {
	e := NewEngine()
	e.SetTagName("check")

	type TestStruct struct {
		Email string `check:"email"`
	}
	s := TestStruct{Email: "invalid-email"}
	err := e.Struct(s)
	assert.Error(t, err)
}

func TestEngineWithCustomFieldNameFunc(t *testing.T) {
	e := NewEngine()

	e.SetFieldNameFunc(func(field reflect.StructField) string {
		return "custom_" + field.Name
	})

	type TestStruct struct {
		Email string `validate:"email"`
	}
	s := TestStruct{Email: "invalid-email"}
	err := e.Struct(s)
	assert.Error(t, err)

	valErr, ok := err.(ValidationErrors)
	require.True(t, ok)
	assert.True(t, valErr.HasField("custom_Email"))
}

func TestEngineVarWithComplexTags(t *testing.T) {
	e := NewEngine()

	t.Run("multiple_tags_with_params", func(t *testing.T) {
		err := e.Var("test@example.com", "required,email")
		assert.NoError(t, err)
	})

	t.Run("tag_with_spaces", func(t *testing.T) {
		err := e.Var("test", "required , alpha")
		assert.NoError(t, err)
	})

	t.Run("empty_tag", func(t *testing.T) {
		err := e.Var("test", "")
		assert.NoError(t, err)
	})

	t.Run("tag_only_commas", func(t *testing.T) {
		err := e.Var("test", ",,")
		assert.NoError(t, err)
	})
}

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
		{"valid_string_5chars", 0, 0, reflect.ValueOf("hello"), true},       // 5 in [1,10]
		{"valid_string_1char", 0, 0, reflect.ValueOf("h"), true},            // 1 in [1,10]
		{"valid_string_10chars", 0, 0, reflect.ValueOf("hellohello"), true}, // 10 in [1,10]
		{"too_short_0chars", 0, 0, reflect.ValueOf(""), false},              // 0 not in [1,10]
		{"too_long_11chars", 0, 0, reflect.ValueOf("helloworld!"), false},   // 11 not in [1,10]

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
