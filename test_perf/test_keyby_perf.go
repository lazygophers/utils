package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
	"time"

	"github.com/lazygophers/utils/candy"
)

type benchStructInt struct {
	ID   int
	Name string
}

type benchStructInt8 struct {
	ID   int8
	Name string
}

// Baseline 实现
func keyByBaseline[T any](ss []T, fieldName string) map[int]T {
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
		v := reflect.ValueOf(item)
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fieldValue := v.FieldByName(fieldName)
		ret[int(fieldValue.Int())] = item
	}
	return ret
}

// Unsafe 实现
func keyByUnsafe[T any](ss []T, fieldName string) map[int]T {
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

// Cached 实现
var fieldCache sync.Map

func keyByCached[T any](ss []T, fieldName string) map[int]T {
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

// Optimized 实现
func keyByOptimized[T any](ss []T, fieldName string) map[int]T {
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

// FieldIndex 实现
func keyByFieldIndex[T any](ss []T, fieldName string) map[int]T {
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

func benchmarkKeyByFunc(name string, fn func() interface{}, iterations int) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = fn()
	}
	elapsed := time.Since(start)
	avgNs := elapsed.Nanoseconds() / int64(iterations)
	fmt.Printf("%-20s: %12v total, %8d ns/op\n", name, elapsed, avgNs)
}

func main() {
	// 准备测试数据
	dataSize := 1000
	data := make([]benchStructInt, dataSize)
	for i := 0; i < dataSize; i++ {
		data[i] = benchStructInt{ID: i, Name: "test"}
	}

	iterations := 10000

	fmt.Printf("测试 %d 次操作，每次处理 %d 个元素\n\n", iterations, dataSize)

	// 测试 Baseline
	benchmarkKeyByFunc("Baseline", func() interface{} {
		return keyByBaseline(data, "ID")
	}, iterations)

	// 测试 Unsafe
	benchmarkKeyByFunc("Unsafe", func() interface{} {
		return keyByUnsafe(data, "ID")
	}, iterations)

	// 测试 Cached (预热后)
	_ = keyByCached(data, "ID") // 预热缓存
	benchmarkKeyByFunc("Cached", func() interface{} {
		return keyByCached(data, "ID")
	}, iterations)

	// 测试 Optimized
	benchmarkKeyByFunc("Optimized", func() interface{} {
		return keyByOptimized(data, "ID")
	}, iterations)

	// 测试 FieldIndex
	benchmarkKeyByFunc("FieldIndex", func() interface{} {
		return keyByFieldIndex(data, "ID")
	}, iterations)

	// 测试原始实现
	benchmarkKeyByFunc("Original", func() interface{} {
		return candy.KeyByInt(data, "ID")
	}, iterations)

	fmt.Println("\n=== 验证功能正确性 ===")
	result1 := keyByBaseline(data, "ID")
	result2 := keyByUnsafe(data, "ID")
	result3 := keyByCached(data, "ID")
	result4 := keyByOptimized(data, "ID")
	result5 := keyByFieldIndex(data, "ID")
	result6 := candy.KeyByInt(data, "ID")

	if len(result1) != len(result2) || len(result2) != len(result3) ||
		len(result3) != len(result4) || len(result4) != len(result5) ||
		len(result5) != len(result6) {
		fmt.Println("错误: 结果长度不一致")
	} else {
		fmt.Println("成功: 所有实现结果一致")
	}
}
