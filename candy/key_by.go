package candy

import (
	"fmt"
	"reflect"
)

// validateKeyByField 验证KeyBy函数的字段
func validateKeyByField(elemType reflect.Type, fieldName string, expectedKind reflect.Kind) reflect.StructField {
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

	return field
}

// getStructFieldValue 获取结构体字段的值
func getStructFieldValue(item interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(item)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("element is not a struct")
	}

	fieldValue := v.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		panic(fmt.Sprintf("field %s not found in struct", fieldName))
	}

	return fieldValue
}

// KeyByInt 将结构体切片按照指定int字段的值作为key重新组织为map
// 如果字段不存在或不是int类型，会panic
func KeyByInt[T any](ss []T, fieldName string) map[int]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Int)

	// 创建结果map
	ret := make(map[int]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := int(fieldValue.Int())
		ret[key] = item
	}

	return ret
}

// KeyByInt8 将结构体切片按照指定int8字段的值作为key重新组织为map
// 如果字段不存在或不是int8类型，会panic
func KeyByInt8[T any](ss []T, fieldName string) map[int8]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Int8)

	// 创建结果map
	ret := make(map[int8]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := int8(fieldValue.Int())
		ret[key] = item
	}

	return ret
}

// KeyByInt16 将结构体切片按照指定int16字段的值作为key重新组织为map
// 如果字段不存在或不是int16类型，会panic
func KeyByInt16[T any](ss []T, fieldName string) map[int16]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Int16)

	// 创建结果map
	ret := make(map[int16]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := int16(fieldValue.Int())
		ret[key] = item
	}

	return ret
}

// KeyByInt32 将结构体切片按照指定int32字段的值作为key重新组织为map
// 如果字段不存在或不是int32类型，会panic
func KeyByInt32[T any](ss []T, fieldName string) map[int32]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Int32)

	// 创建结果map
	ret := make(map[int32]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := int32(fieldValue.Int())
		ret[key] = item
	}

	return ret
}

// KeyByInt64 将结构体切片按照指定int64字段的值作为key重新组织为map
// 如果字段不存在或不是int64类型，会panic
func KeyByInt64[T any](ss []T, fieldName string) map[int64]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Int64)

	// 创建结果map
	ret := make(map[int64]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := fieldValue.Int()
		ret[key] = item
	}

	return ret
}

// KeyByUint8 将结构体切片按照指定uint8字段的值作为key重新组织为map
// 如果字段不存在或不是uint8类型，会panic
func KeyByUint8[T any](ss []T, fieldName string) map[uint8]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Uint8)

	// 创建结果map
	ret := make(map[uint8]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := uint8(fieldValue.Uint())
		ret[key] = item
	}

	return ret
}

// KeyByUint16 将结构体切片按照指定uint16字段的值作为key重新组织为map
// 如果字段不存在或不是uint16类型，会panic
func KeyByUint16[T any](ss []T, fieldName string) map[uint16]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Uint16)

	// 创建结果map
	ret := make(map[uint16]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := uint16(fieldValue.Uint())
		ret[key] = item
	}

	return ret
}

// KeyByUint32 将结构体切片按照指定uint32字段的值作为key重新组织为map
// 如果字段不存在或不是uint32类型，会panic
func KeyByUint32[T any](ss []T, fieldName string) map[uint32]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Uint32)

	// 创建结果map
	ret := make(map[uint32]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := uint32(fieldValue.Uint())
		ret[key] = item
	}

	return ret
}

// KeyByUint64 将结构体切片按照指定uint64字段的值作为key重新组织为map
// 如果字段不存在或不是uint64类型，会panic
func KeyByUint64[T any](ss []T, fieldName string) map[uint64]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Uint64)

	// 创建结果map
	ret := make(map[uint64]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := fieldValue.Uint()
		ret[key] = item
	}

	return ret
}

// KeyByUint 将结构体切片按照指定uint字段的值作为key重新组织为map
// 如果字段不存在或不是uint类型，会panic
func KeyByUint[T any](ss []T, fieldName string) map[uint]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Uint)

	// 创建结果map
	ret := make(map[uint]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := uint(fieldValue.Uint())
		ret[key] = item
	}

	return ret
}

// KeyByFloat32 将结构体切片按照指定float32字段的值作为key重新组织为map
// 如果字段不存在或不是float32类型，会panic
func KeyByFloat32[T any](ss []T, fieldName string) map[float32]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Float32)

	// 创建结果map
	ret := make(map[float32]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := float32(fieldValue.Float())
		ret[key] = item
	}

	return ret
}

// KeyByFloat64 将结构体切片按照指定float64字段的值作为key重新组织为map
// 如果字段不存在或不是float64类型，会panic
func KeyByFloat64[T any](ss []T, fieldName string) map[float64]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Float64)

	// 创建结果map
	ret := make(map[float64]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := fieldValue.Float()
		ret[key] = item
	}

	return ret
}

// KeyByString 将结构体切片按照指定string字段的值作为key重新组织为map
// 如果字段不存在或不是string类型，会panic
func KeyByString[T any](ss []T, fieldName string) map[string]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.String)

	// 创建结果map
	ret := make(map[string]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := fieldValue.String()
		ret[key] = item
	}

	return ret
}

// KeyByBool 将结构体切片按照指定bool字段的值作为key重新组织为map
// 如果字段不存在或不是bool类型，会panic
func KeyByBool[T any](ss []T, fieldName string) map[bool]T {
	if len(ss) == 0 {
		return nil
	}

	// 验证字段
	validateKeyByField(reflect.TypeOf(ss[0]), fieldName, reflect.Bool)

	// 创建结果map
	ret := make(map[bool]T, len(ss))

	// 遍历切片，构建map
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		key := fieldValue.Bool()
		ret[key] = item
	}

	return ret
}
