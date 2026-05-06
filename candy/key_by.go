package candy

import (
	"fmt"
	"reflect"
	"unsafe"
)

func validateStructFieldKind(elemType reflect.Type, fieldName string, expectedKind reflect.Kind) {
	// 解指针类型
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	// 检查是否为结构体
	if elemType.Kind() != reflect.Struct {
		panic("element is not a struct")
	}

	// 查找字段
	field, ok := elemType.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	// 检查字段类型
	if field.Type.Kind() != expectedKind {
		panic(fmt.Sprintf("field %s is not of type %s", fieldName, expectedKind))
	}
}

// getStructFieldValue 获取结构体字段的值
func getStructFieldValue(item any, fieldName string) reflect.Value {
	v := reflect.ValueOf(item)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fieldValue := v.FieldByName(fieldName)
	return fieldValue
}

func keyByField[T any, K comparable](ss []T, fieldName string, expectedKind reflect.Kind, converter func(reflect.Value) K) map[K]T {
	if len(ss) == 0 {
		return nil
	}

	validateStructFieldKind(reflect.TypeOf(ss[0]), fieldName, expectedKind)

	ret := make(map[K]T, len(ss))
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		ret[converter(fieldValue)] = item
	}

	return ret
}

// KeyByInt 将结构体切片按照指定int字段的值作为key重新组织为map
// 如果字段不存在或不是int类型，会panic
func KeyByInt[T any](ss []T, fieldName string) map[int]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int {
		panic(fmt.Sprintf("field %s is not of type int", fieldName))
	}
	ret := make(map[int]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*int)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*int)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByInt8 将结构体切片按照指定int8字段的值作为key重新组织为map
// 如果字段不存在或不是int8类型，会panic
func KeyByInt8[T any](ss []T, fieldName string) map[int8]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int8 {
		panic(fmt.Sprintf("field %s is not of type int8", fieldName))
	}
	ret := make(map[int8]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*int8)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*int8)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByInt16 将结构体切片按照指定int16字段的值作为key重新组织为map
// 如果字段不存在或不是int16类型，会panic
func KeyByInt16[T any](ss []T, fieldName string) map[int16]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int16 {
		panic(fmt.Sprintf("field %s is not of type int16", fieldName))
	}
	ret := make(map[int16]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*int16)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*int16)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByInt32 将结构体切片按照指定int32字段的值作为key重新组织为map
// 如果字段不存在或不是int32类型，会panic
func KeyByInt32[T any](ss []T, fieldName string) map[int32]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int32 {
		panic(fmt.Sprintf("field %s is not of type int32", fieldName))
	}
	ret := make(map[int32]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*int32)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*int32)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByInt64 将结构体切片按照指定int64字段的值作为key重新组织为map
// 如果字段不存在或不是int64类型，会panic
func KeyByInt64[T any](ss []T, fieldName string) map[int64]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Int64 {
		panic(fmt.Sprintf("field %s is not of type int64", fieldName))
	}
	ret := make(map[int64]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*int64)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*int64)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByUint8 将结构体切片按照指定uint8字段的值作为key重新组织为map
// 如果字段不存在或不是uint8类型，会panic
func KeyByUint8[T any](ss []T, fieldName string) map[uint8]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Uint8 {
		panic(fmt.Sprintf("field %s is not of type uint8", fieldName))
	}
	ret := make(map[uint8]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*uint8)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*uint8)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByUint16 将结构体切片按照指定uint16字段的值作为key重新组织为map
// 如果字段不存在或不是uint16类型，会panic
func KeyByUint16[T any](ss []T, fieldName string) map[uint16]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Uint16 {
		panic(fmt.Sprintf("field %s is not of type uint16", fieldName))
	}
	ret := make(map[uint16]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*uint16)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*uint16)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByUint32 将结构体切片按照指定uint32字段的值作为key重新组织为map
// 如果字段不存在或不是uint32类型，会panic
func KeyByUint32[T any](ss []T, fieldName string) map[uint32]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Uint32 {
		panic(fmt.Sprintf("field %s is not of type uint32", fieldName))
	}
	ret := make(map[uint32]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*uint32)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*uint32)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByUint64 将结构体切片按照指定uint64字段的值作为key重新组织为map
// 如果字段不存在或不是uint64类型，会panic
func KeyByUint64[T any](ss []T, fieldName string) map[uint64]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Uint64 {
		panic(fmt.Sprintf("field %s is not of type uint64", fieldName))
	}
	ret := make(map[uint64]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*uint64)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*uint64)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByUint 将结构体切片按照指定uint字段的值作为key重新组织为map
// 如果字段不存在或不是uint类型，会panic
func KeyByUint[T any](ss []T, fieldName string) map[uint]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Uint {
		panic(fmt.Sprintf("field %s is not of type uint", fieldName))
	}
	ret := make(map[uint]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*uint)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*uint)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByFloat32 将结构体切片按照指定float32字段的值作为key重新组织为map
// 如果字段不存在或不是float32类型，会panic
func KeyByFloat32[T any](ss []T, fieldName string) map[float32]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Float32 {
		panic(fmt.Sprintf("field %s is not of type float32", fieldName))
	}
	ret := make(map[float32]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*float32)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*float32)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByFloat64 将结构体切片按照指定float64字段的值作为key重新组织为map
// 如果字段不存在或不是float64类型，会panic
func KeyByFloat64[T any](ss []T, fieldName string) map[float64]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Float64 {
		panic(fmt.Sprintf("field %s is not of type float64", fieldName))
	}
	ret := make(map[float64]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*float64)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*float64)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByString 将结构体切片按照指定string字段的值作为key重新组织为map
// 如果字段不存在或不是string类型，会panic
func KeyByString[T any](ss []T, fieldName string) map[string]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.String {
		panic(fmt.Sprintf("field %s is not of type string", fieldName))
	}
	ret := make(map[string]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*string)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*string)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}

// KeyByBool 将结构体切片按照指定bool字段的值作为key重新组织为map
// 如果字段不存在或不是bool类型，会panic
func KeyByBool[T any](ss []T, fieldName string) map[bool]T {
	if len(ss) == 0 {
		return nil
	}
	typ := reflect.TypeOf(ss[0])
	isPtr := typ.Kind() == reflect.Ptr
	if isPtr {
		typ = typ.Elem()
	}
	field, ok := typ.FieldByName(fieldName)
	if !ok || field.Type.Kind() != reflect.Bool {
		panic(fmt.Sprintf("field %s is not of type bool", fieldName))
	}
	ret := make(map[bool]T, len(ss))
	if isPtr {
		for _, item := range ss {
			v := reflect.ValueOf(item)
			if v.IsNil() {
				panic("nil pointer in slice")
			}
			elemPtr := unsafe.Pointer(v.Pointer())
			fieldPtr := unsafe.Pointer(uintptr(elemPtr) + field.Offset)
			id := *(*bool)(fieldPtr)
			ret[id] = item
		}
	} else {
		for _, item := range ss {
			itemPtr := unsafe.Pointer(&item)
			fieldPtr := unsafe.Pointer(uintptr(itemPtr) + field.Offset)
			id := *(*bool)(fieldPtr)
			ret[id] = item
		}
	}
	return ret
}
