package candy

import (
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

type benchStruct struct {
	ID   int
	Name string
}

// 1. Baseline - 当前实现
func BenchmarkKeyByInt_Baseline(b *testing.B) {
	data := make([]benchStruct, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KeyByInt(data, "ID")
	}
}

// 1.1 旧实现对比（使用原始的getStructFieldValue）
func BenchmarkKeyByInt_OldImpl(b *testing.B) {
	type oldBenchStruct struct {
		ID   int
		Name string
	}
	data := make([]oldBenchStruct, 100)
	for i := 0; i < 100; i++ {
		data[i] = oldBenchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 使用旧的反射方式
		ret := make(map[int]oldBenchStruct, len(data))
		for _, item := range data {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByName("ID")
			ret[int(fieldValue.Int())] = item
		}
		_ = ret
	}
}

// 2. Unsafe直接访问
func keyByIntUnsafe[T any](ss []T, fieldName string) map[int]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int {
		panic("field is not int")
	}
	ret := make(map[int]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
		id := *(*int)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt_Unsafe(b *testing.B) {
	data := make([]benchStruct, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntUnsafe(data, "ID")
	}
}

// 3. 缓存字段偏移量
var fieldCache sync.Map

func keyByIntCached[T any](ss []T, fieldName string) map[int]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	cacheKey := typ.String() + "." + fieldName
	if offset, ok := fieldCache.Load(cacheKey); ok {
		ret := make(map[int]T, len(ss))
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset.(uintptr))
			id := *(*int)(fieldPtr)
			ret[id] = item
		}
		return ret
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int {
		panic("field is not int")
	}
	offset := field.Offset
	fieldCache.Store(cacheKey, offset)
	ret := make(map[int]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset)
		id := *(*int)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt_Cached(b *testing.B) {
	data := make([]benchStruct, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntCached(data, "ID")
	}
}

// 4. 优化反射 - 减少重复解引用
func keyByIntOptimized[T any](ss []T, fieldName string) map[int]T {
	if len(ss) == 0 {
		return nil
	}
	v0 := reflect.ValueOf(ss[0])
	isPtr := v0.Kind() == reflect.Ptr
	if isPtr {
		v0 = v0.Elem()
	}
	if v0.Kind() != reflect.Struct {
		panic("not a struct")
	}
	field := v0.FieldByName(fieldName)
	if !field.IsValid() || field.Kind() != reflect.Int {
		panic("field is not int")
	}
	ret := make(map[int]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item).Elem()
			id := int(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			id := int(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	}
	return ret
}

func BenchmarkKeyByInt_Optimized(b *testing.B) {
	data := make([]benchStruct, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntOptimized(data, "ID")
	}
}

// 5. 使用字段索引
func keyByIntFieldIndex[T any](ss []T, fieldName string) map[int]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int {
		panic("field is not int")
	}
	index := field.Index
	ret := make(map[int]T, len(ss))
	for _, item := range ss {
		v := reflect.ValueOf(item)
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fieldVal := v.FieldByIndex(index)
		ret[int(fieldVal.Int())] = item
	}
	return ret
}

func BenchmarkKeyByInt_FieldIndex(b *testing.B) {
	data := make([]benchStruct, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntFieldIndex(data, "ID")
	}
}

// 不同数据规模测试
func BenchmarkKeyByInt_Large_Baseline(b *testing.B) {
	data := make([]benchStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KeyByInt(data, "ID")
	}
}

func BenchmarkKeyByInt_Large_Unsafe(b *testing.B) {
	data := make([]benchStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntUnsafe(data, "ID")
	}
}

func BenchmarkKeyByInt_Large_Cached(b *testing.B) {
	data := make([]benchStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntCached(data, "ID")
	}
}

func BenchmarkKeyByInt_Large_Optimized(b *testing.B) {
	data := make([]benchStruct, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = benchStruct{ID: i, Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByIntOptimized(data, "ID")
	}
}
