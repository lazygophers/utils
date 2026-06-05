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

// ============== KeyByInt8 测试 ==============

type benchStructInt8 struct {
	ID   int8
	Name string
}

func BenchmarkKeyByInt8_Baseline(b *testing.B) {
	data := make([]benchStructInt8, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt8{ID: int8(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KeyByInt8(data, "ID")
	}
}

func keyByInt8Unsafe[T any](ss []T, fieldName string) map[int8]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int8 {
		panic("field is not int8")
	}
	ret := make(map[int8]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
		id := *(*int8)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt8_Unsafe(b *testing.B) {
	data := make([]benchStructInt8, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt8{ID: int8(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt8Unsafe(data, "ID")
	}
}

var fieldCacheInt8 sync.Map

func keyByInt8Cached[T any](ss []T, fieldName string) map[int8]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	cacheKey := typ.String() + "." + fieldName
	if offset, ok := fieldCacheInt8.Load(cacheKey); ok {
		ret := make(map[int8]T, len(ss))
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset.(uintptr))
			id := *(*int8)(fieldPtr)
			ret[id] = item
		}
		return ret
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int8 {
		panic("field is not int8")
	}
	offset := field.Offset
	fieldCacheInt8.Store(cacheKey, offset)
	ret := make(map[int8]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset)
		id := *(*int8)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt8_Cached(b *testing.B) {
	data := make([]benchStructInt8, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt8{ID: int8(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt8Cached(data, "ID")
	}
}

func keyByInt8Optimized[T any](ss []T, fieldName string) map[int8]T {
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
	if !field.IsValid() || field.Kind() != reflect.Int8 {
		panic("field is not int8")
	}
	ret := make(map[int8]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item).Elem()
			id := int8(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			id := int8(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	}
	return ret
}

func BenchmarkKeyByInt8_Optimized(b *testing.B) {
	data := make([]benchStructInt8, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt8{ID: int8(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt8Optimized(data, "ID")
	}
}

func keyByInt8FieldIndex[T any](ss []T, fieldName string) map[int8]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int8 {
		panic("field is not int8")
	}
	index := field.Index
	ret := make(map[int8]T, len(ss))
	for _, item := range ss {
		v := reflect.ValueOf(item)
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fieldVal := v.FieldByIndex(index)
		ret[int8(fieldVal.Int())] = item
	}
	return ret
}

func BenchmarkKeyByInt8_FieldIndex(b *testing.B) {
	data := make([]benchStructInt8, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt8{ID: int8(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt8FieldIndex(data, "ID")
	}
}

// ============== KeyByInt16 测试 ==============

type benchStructInt16 struct {
	ID   int16
	Name string
}

func BenchmarkKeyByInt16_Baseline(b *testing.B) {
	data := make([]benchStructInt16, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt16{ID: int16(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KeyByInt16(data, "ID")
	}
}

func keyByInt16Unsafe[T any](ss []T, fieldName string) map[int16]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int16 {
		panic("field is not int16")
	}
	ret := make(map[int16]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
		id := *(*int16)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt16_Unsafe(b *testing.B) {
	data := make([]benchStructInt16, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt16{ID: int16(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt16Unsafe(data, "ID")
	}
}

var fieldCacheInt16 sync.Map

func keyByInt16Cached[T any](ss []T, fieldName string) map[int16]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	cacheKey := typ.String() + "." + fieldName
	if offset, ok := fieldCacheInt16.Load(cacheKey); ok {
		ret := make(map[int16]T, len(ss))
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset.(uintptr))
			id := *(*int16)(fieldPtr)
			ret[id] = item
		}
		return ret
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int16 {
		panic("field is not int16")
	}
	offset := field.Offset
	fieldCacheInt16.Store(cacheKey, offset)
	ret := make(map[int16]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset)
		id := *(*int16)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt16_Cached(b *testing.B) {
	data := make([]benchStructInt16, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt16{ID: int16(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt16Cached(data, "ID")
	}
}

func keyByInt16Optimized[T any](ss []T, fieldName string) map[int16]T {
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
	if !field.IsValid() || field.Kind() != reflect.Int16 {
		panic("field is not int16")
	}
	ret := make(map[int16]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item).Elem()
			id := int16(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			id := int16(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	}
	return ret
}

func BenchmarkKeyByInt16_Optimized(b *testing.B) {
	data := make([]benchStructInt16, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt16{ID: int16(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt16Optimized(data, "ID")
	}
}

func keyByInt16FieldIndex[T any](ss []T, fieldName string) map[int16]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int16 {
		panic("field is not int16")
	}
	index := field.Index
	ret := make(map[int16]T, len(ss))
	for _, item := range ss {
		v := reflect.ValueOf(item)
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fieldVal := v.FieldByIndex(index)
		ret[int16(fieldVal.Int())] = item
	}
	return ret
}

func BenchmarkKeyByInt16_FieldIndex(b *testing.B) {
	data := make([]benchStructInt16, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt16{ID: int16(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt16FieldIndex(data, "ID")
	}
}

// ============== KeyByInt32 测试 ==============

type benchStructInt32 struct {
	ID   int32
	Name string
}

func BenchmarkKeyByInt32_Baseline(b *testing.B) {
	data := make([]benchStructInt32, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt32{ID: int32(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KeyByInt32(data, "ID")
	}
}

func keyByInt32Unsafe[T any](ss []T, fieldName string) map[int32]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int32 {
		panic("field is not int32")
	}
	ret := make(map[int32]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
		id := *(*int32)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt32_Unsafe(b *testing.B) {
	data := make([]benchStructInt32, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt32{ID: int32(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt32Unsafe(data, "ID")
	}
}

var fieldCacheInt32 sync.Map

func keyByInt32Cached[T any](ss []T, fieldName string) map[int32]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	cacheKey := typ.String() + "." + fieldName
	if offset, ok := fieldCacheInt32.Load(cacheKey); ok {
		ret := make(map[int32]T, len(ss))
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset.(uintptr))
			id := *(*int32)(fieldPtr)
			ret[id] = item
		}
		return ret
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int32 {
		panic("field is not int32")
	}
	offset := field.Offset
	fieldCacheInt32.Store(cacheKey, offset)
	ret := make(map[int32]T, len(ss))
	for _, item := range ss {
		itemPtr := unsafe.Pointer(&item)
		fieldPtr := unsafe.Pointer(uintptr(itemPtr) + offset)
		id := *(*int32)(fieldPtr)
		ret[id] = item
	}
	return ret
}

func BenchmarkKeyByInt32_Cached(b *testing.B) {
	data := make([]benchStructInt32, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt32{ID: int32(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt32Cached(data, "ID")
	}
}

func keyByInt32Optimized[T any](ss []T, fieldName string) map[int32]T {
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
	if !field.IsValid() || field.Kind() != reflect.Int32 {
		panic("field is not int32")
	}
	ret := make(map[int32]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item).Elem()
			id := int32(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			id := int32(v.FieldByName(fieldName).Int())
			ret[id] = item
		}
	}
	return ret
}

func BenchmarkKeyByInt32_Optimized(b *testing.B) {
	data := make([]benchStructInt32, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt32{ID: int32(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt32Optimized(data, "ID")
	}
}

func keyByInt32FieldIndex[T any](ss []T, fieldName string) map[int32]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int32 {
		panic("field is not int32")
	}
	index := field.Index
	ret := make(map[int32]T, len(ss))
	for _, item := range ss {
		v := reflect.ValueOf(item)
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fieldVal := v.FieldByIndex(index)
		ret[int32(fieldVal.Int())] = item
	}
	return ret
}

func BenchmarkKeyByInt32_FieldIndex(b *testing.B) {
	data := make([]benchStructInt32, 100)
	for i := 0; i < 100; i++ {
		data[i] = benchStructInt32{ID: int32(i), Name: "test"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = keyByInt32FieldIndex(data, "ID")
	}
}
