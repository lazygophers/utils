package candy

import (
	"reflect"
	"testing"
)

// ============================================================================
// 方案对比基准测试
// ============================================================================

// 基准测试数据
type DCBenchStruct struct {
	ID   int
	Name string
	Tags []string
	Meta map[string]interface{}
}

var (
	benchSmallSlice = []int{1, 2, 3, 4, 5}
	benchLargeSlice = make([]int, 1000)
	benchSmallMap   = map[string]int{"a": 1, "b": 2, "c": 3}
	benchLargeMap   = make(map[string]int, 100)
	benchDCStruct   = DCBenchStruct{
		ID:   1,
		Name: "test",
		Tags: []string{"a", "b", "c"},
		Meta: map[string]interface{}{"key": "value"},
	}
)

func init() {
	for i := range benchLargeSlice {
		benchLargeSlice[i] = i
	}
	for i := 0; i < 100; i++ {
		benchLargeMap[testKey(i)] = i
	}
}

func testKey(i int) string {
	return "key" + string(rune('0'+i))
}

// ============================================================================
// 方案1: 原始实现
// ============================================================================

func Benchmark_v1_Original_SmallSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopy(benchSmallSlice, &dst)
	}
}

func Benchmark_v1_Original_LargeSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopy(benchLargeSlice, &dst)
	}
}

func Benchmark_v1_Original_SmallMap(b *testing.B) {
	var dst map[string]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopy(benchSmallMap, &dst)
	}
}

func Benchmark_v1_Original_Struct(b *testing.B) {
	var dst DCBenchStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopy(benchDCStruct, &dst)
	}
}

// ============================================================================
// 方案2: 优化版本 - 预分配 + 基本类型快速路径
// ============================================================================

func DeepCopyOptimized(src, dst any) {
	v1, v2 := reflect.ValueOf(src), reflect.ValueOf(dst)
	deepCopyOptimizedValue(v1, v2)
}

func deepCopyOptimizedValue(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() {
		return
	}

	// 解引用指针
	for v1.Kind() == reflect.Ptr {
		if v1.IsNil() {
			return
		}
		v1 = v1.Elem()
	}
	for v2.Kind() == reflect.Ptr {
		if v2.IsNil() {
			v2.Set(reflect.New(v2.Type().Elem()))
		}
		v2 = v2.Elem()
	}

	if v1.Type() != v2.Type() {
		return
	}

	switch v1.Kind() {
	case reflect.Map:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		// 预分配容量
		v2.Set(reflect.MakeMapWithSize(v1.Type(), v1.Len()))
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := reflect.New(val1.Type()).Elem()
			deepCopyOptimizedValue(val1, val2)
			v2.SetMapIndex(k, val2)
		}

	case reflect.Slice:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		v2.Set(reflect.MakeSlice(v1.Type(), v1.Len(), v1.Cap()))

		// 基本类型元素使用快速路径
		elemKind := v1.Type().Elem().Kind()
		if isBasicType(elemKind) {
			for i := 0; i < v1.Len(); i++ {
				v2.Index(i).Set(v1.Index(i))
			}
			return
		}

		// 复杂类型递归
		for i := 0; i < v1.Len(); i++ {
			deepCopyOptimizedValue(v1.Index(i), v2.Index(i))
		}

	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			if v2.Field(i).CanSet() {
				deepCopyOptimizedValue(v1.Field(i), v2.Field(i))
			}
		}

	case reflect.Interface:
		if v1.IsNil() {
			return
		}
		srcElem := v1.Elem()
		dstElem := reflect.New(srcElem.Type()).Elem()
		deepCopyOptimizedValue(srcElem, dstElem)
		v2.Set(dstElem)

	default:
		// 基本类型直接设置
		if v2.CanSet() {
			v2.Set(v1)
		}
	}
}

func Benchmark_v2_Optimized_SmallSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyOptimized(benchSmallSlice, &dst)
	}
}

func Benchmark_v2_Optimized_LargeSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyOptimized(benchLargeSlice, &dst)
	}
}

func Benchmark_v2_Optimized_SmallMap(b *testing.B) {
	var dst map[string]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyOptimized(benchSmallMap, &dst)
	}
}

func Benchmark_v2_Optimized_Struct(b *testing.B) {
	var dst DCBenchStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyOptimized(benchDCStruct, &dst)
	}
}

// ============================================================================
// 方案3: 使用 reflect.Copy
// ============================================================================

func DeepCopyReflectCopy(src, dst any) {
	v1, v2 := reflect.ValueOf(src), reflect.ValueOf(dst)
	deepCopyReflectCopyValue(v1, v2)
}

func deepCopyReflectCopyValue(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() {
		return
	}

	for v1.Kind() == reflect.Ptr {
		if v1.IsNil() {
			return
		}
		v1 = v1.Elem()
	}
	for v2.Kind() == reflect.Ptr {
		if v2.IsNil() {
			v2.Set(reflect.New(v2.Type().Elem()))
		}
		v2 = v2.Elem()
	}

	if v1.Type() != v2.Type() {
		return
	}

	switch v1.Kind() {
	case reflect.Slice:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		v2.Set(reflect.MakeSlice(v1.Type(), v1.Len(), v1.Cap()))
		reflect.Copy(v2, v1)
	case reflect.Map:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		v2.Set(reflect.MakeMapWithSize(v1.Type(), v1.Len()))
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := reflect.New(val1.Type()).Elem()
			deepCopyReflectCopyValue(val1, val2)
			v2.SetMapIndex(k, val2)
		}
	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			if v2.Field(i).CanSet() {
				deepCopyReflectCopyValue(v1.Field(i), v2.Field(i))
			}
		}
	case reflect.Interface:
		if v1.IsNil() {
			return
		}
		srcElem := v1.Elem()
		dstElem := reflect.New(srcElem.Type()).Elem()
		deepCopyReflectCopyValue(srcElem, dstElem)
		v2.Set(dstElem)
	default:
		if v2.CanSet() {
			v2.Set(v1)
		}
	}
}

func Benchmark_v3_ReflectCopy_SmallSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyReflectCopy(benchSmallSlice, &dst)
	}
}

func Benchmark_v3_ReflectCopy_LargeSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyReflectCopy(benchLargeSlice, &dst)
	}
}

func Benchmark_v3_ReflectCopy_SmallMap(b *testing.B) {
	var dst map[string]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyReflectCopy(benchSmallMap, &dst)
	}
}

func Benchmark_v3_ReflectCopy_Struct(b *testing.B) {
	var dst DCBenchStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyReflectCopy(benchDCStruct, &dst)
	}
}

// ============================================================================
// 方案4: 类型特化（针对常见类型）
// ============================================================================

func DeepCopySpecialized(src, dst any) {
	switch s := src.(type) {
	case []int:
		if d, ok := dst.(*[]int); ok {
			*d = make([]int, len(s))
			copy(*d, s)
			return
		}
	case []string:
		if d, ok := dst.(*[]string); ok {
			*d = make([]string, len(s))
			copy(*d, s)
			return
		}
	case map[string]int:
		if d, ok := dst.(*map[string]int); ok {
			*d = make(map[string]int, len(s))
			for k, v := range s {
				(*d)[k] = v
			}
			return
		}
	}
	// 回退到优化版本
	DeepCopyOptimized(src, dst)
}

func Benchmark_v4_Specialized_SmallSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopySpecialized(benchSmallSlice, &dst)
	}
}

func Benchmark_v4_Specialized_LargeSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopySpecialized(benchLargeSlice, &dst)
	}
}

func Benchmark_v4_Specialized_SmallMap(b *testing.B) {
	var dst map[string]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopySpecialized(benchSmallMap, &dst)
	}
}

func Benchmark_v4_Specialized_Struct(b *testing.B) {
	var dst DCBenchStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopySpecialized(benchDCStruct, &dst)
	}
}

// ============================================================================
// 方案5: 综合优化（类型特化 + 预分配 + 快速路径）
// ============================================================================

func DeepCopyUltimate(src, dst any) {
	// 类型特化快速路径
	switch s := src.(type) {
	case []int:
		if d, ok := dst.(*[]int); ok {
			*d = make([]int, len(s))
			copy(*d, s)
			return
		}
	case []string:
		if d, ok := dst.(*[]string); ok {
			*d = make([]string, len(s))
			copy(*d, s)
			return
		}
	case map[string]int:
		if d, ok := dst.(*map[string]int); ok {
			*d = make(map[string]int, len(s))
			for k, v := range s {
				(*d)[k] = v
			}
			return
		}
	}
	// 回退到反射优化版本
	DeepCopyOptimized(src, dst)
}

func Benchmark_v5_Ultimate_SmallSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyUltimate(benchSmallSlice, &dst)
	}
}

func Benchmark_v5_Ultimate_LargeSlice(b *testing.B) {
	var dst []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyUltimate(benchLargeSlice, &dst)
	}
}

func Benchmark_v5_Ultimate_SmallMap(b *testing.B) {
	var dst map[string]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyUltimate(benchSmallMap, &dst)
	}
}

func Benchmark_v5_Ultimate_Struct(b *testing.B) {
	var dst DCBenchStruct
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeepCopyUltimate(benchDCStruct, &dst)
	}
}
