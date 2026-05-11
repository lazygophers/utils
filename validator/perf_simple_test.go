package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
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
