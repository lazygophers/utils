package validator

import (
	"math/rand"
	"reflect"
	"testing"
)

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
		reflect.Int:        handleInt,
		reflect.Int8:       handleInt,
		reflect.Int16:      handleInt,
		reflect.Int32:      handleInt,
		reflect.Int64:      handleInt,
		reflect.Uint:       handleUint,
		reflect.Uint8:      handleUint,
		reflect.Uint16:     handleUint,
		reflect.Uint32:     handleUint,
		reflect.Uint64:     handleUint,
		reflect.Float32:    handleFloat,
		reflect.Float64:    handleFloat,
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
