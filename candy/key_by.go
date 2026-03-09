package candy

import (
	"fmt"
	"reflect"
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
	return keyByField(ss, fieldName, reflect.Int, func(fieldValue reflect.Value) int {
		return int(fieldValue.Int())
	})
}

// KeyByInt8 将结构体切片按照指定int8字段的值作为key重新组织为map
// 如果字段不存在或不是int8类型，会panic
func KeyByInt8[T any](ss []T, fieldName string) map[int8]T {
	return keyByField(ss, fieldName, reflect.Int8, func(fieldValue reflect.Value) int8 {
		return int8(fieldValue.Int())
	})
}

// KeyByInt16 将结构体切片按照指定int16字段的值作为key重新组织为map
// 如果字段不存在或不是int16类型，会panic
func KeyByInt16[T any](ss []T, fieldName string) map[int16]T {
	return keyByField(ss, fieldName, reflect.Int16, func(fieldValue reflect.Value) int16 {
		return int16(fieldValue.Int())
	})
}

// KeyByInt32 将结构体切片按照指定int32字段的值作为key重新组织为map
// 如果字段不存在或不是int32类型，会panic
func KeyByInt32[T any](ss []T, fieldName string) map[int32]T {
	return keyByField(ss, fieldName, reflect.Int32, func(fieldValue reflect.Value) int32 {
		return int32(fieldValue.Int())
	})
}

// KeyByInt64 将结构体切片按照指定int64字段的值作为key重新组织为map
// 如果字段不存在或不是int64类型，会panic
func KeyByInt64[T any](ss []T, fieldName string) map[int64]T {
	return keyByField(ss, fieldName, reflect.Int64, func(fieldValue reflect.Value) int64 {
		return fieldValue.Int()
	})
}

// KeyByUint8 将结构体切片按照指定uint8字段的值作为key重新组织为map
// 如果字段不存在或不是uint8类型，会panic
func KeyByUint8[T any](ss []T, fieldName string) map[uint8]T {
	return keyByField(ss, fieldName, reflect.Uint8, func(fieldValue reflect.Value) uint8 {
		return uint8(fieldValue.Uint())
	})
}

// KeyByUint16 将结构体切片按照指定uint16字段的值作为key重新组织为map
// 如果字段不存在或不是uint16类型，会panic
func KeyByUint16[T any](ss []T, fieldName string) map[uint16]T {
	return keyByField(ss, fieldName, reflect.Uint16, func(fieldValue reflect.Value) uint16 {
		return uint16(fieldValue.Uint())
	})
}

// KeyByUint32 将结构体切片按照指定uint32字段的值作为key重新组织为map
// 如果字段不存在或不是uint32类型，会panic
func KeyByUint32[T any](ss []T, fieldName string) map[uint32]T {
	return keyByField(ss, fieldName, reflect.Uint32, func(fieldValue reflect.Value) uint32 {
		return uint32(fieldValue.Uint())
	})
}

// KeyByUint64 将结构体切片按照指定uint64字段的值作为key重新组织为map
// 如果字段不存在或不是uint64类型，会panic
func KeyByUint64[T any](ss []T, fieldName string) map[uint64]T {
	return keyByField(ss, fieldName, reflect.Uint64, func(fieldValue reflect.Value) uint64 {
		return fieldValue.Uint()
	})
}

// KeyByUint 将结构体切片按照指定uint字段的值作为key重新组织为map
// 如果字段不存在或不是uint类型，会panic
func KeyByUint[T any](ss []T, fieldName string) map[uint]T {
	return keyByField(ss, fieldName, reflect.Uint, func(fieldValue reflect.Value) uint {
		return uint(fieldValue.Uint())
	})
}

// KeyByFloat32 将结构体切片按照指定float32字段的值作为key重新组织为map
// 如果字段不存在或不是float32类型，会panic
func KeyByFloat32[T any](ss []T, fieldName string) map[float32]T {
	return keyByField(ss, fieldName, reflect.Float32, func(fieldValue reflect.Value) float32 {
		return float32(fieldValue.Float())
	})
}

// KeyByFloat64 将结构体切片按照指定float64字段的值作为key重新组织为map
// 如果字段不存在或不是float64类型，会panic
func KeyByFloat64[T any](ss []T, fieldName string) map[float64]T {
	return keyByField(ss, fieldName, reflect.Float64, func(fieldValue reflect.Value) float64 {
		return fieldValue.Float()
	})
}

// KeyByString 将结构体切片按照指定string字段的值作为key重新组织为map
// 如果字段不存在或不是string类型，会panic
func KeyByString[T any](ss []T, fieldName string) map[string]T {
	return keyByField(ss, fieldName, reflect.String, func(fieldValue reflect.Value) string {
		return fieldValue.String()
	})
}

// KeyByBool 将结构体切片按照指定bool字段的值作为key重新组织为map
// 如果字段不存在或不是bool类型，会panic
func KeyByBool[T any](ss []T, fieldName string) map[bool]T {
	return keyByField(ss, fieldName, reflect.Bool, func(fieldValue reflect.Value) bool {
		return fieldValue.Bool()
	})
}
