package candy

import (
	"fmt"
	"reflect"
)

// KeyBy creates a map keyed by the specified field from a slice of structs
func KeyBy(list interface{}, fieldName string) interface{} {
	if list == nil {
		return nil
	}

	lv := reflect.ValueOf(list)

	switch lv.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("list required slice or array type")
	}

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	if evs.Kind() != reflect.Struct {
		panic("list element is not struct")
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, ev), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m.Interface()
}

// KeyByString creates a map keyed by string field from a slice of struct pointers
func KeyByString[M any](list []*M, fieldName string) map[string]*M {
	if len(list) == 0 {
		return map[string]*M{}
	}

	lv := reflect.ValueOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := make(map[string]*M, lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m[elemStruct.FieldByIndex(field.Index).String()] = elem.Interface().(*M)
	}

	return m
}

// KeyByInt32 creates a map keyed by int32 field from a slice of struct pointers
func KeyByInt32[M any](list []*M, fieldName string) map[int32]*M {
	if len(list) == 0 {
		return map[int32]*M{}
	}

	lv := reflect.ValueOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := make(map[int32]*M, lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m[int32(elemStruct.FieldByIndex(field.Index).Int())] = elem.Interface().(*M)
	}

	return m
}

// KeyByInt64 creates a map keyed by int64 field from a slice of struct pointers
func KeyByInt64[M any](list []*M, fieldName string) map[int64]*M {
	if len(list) == 0 {
		return map[int64]*M{}
	}

	lv := reflect.ValueOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := make(map[int64]*M, lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m[elemStruct.FieldByIndex(field.Index).Int()] = elem.Interface().(*M)
	}

	return m
}

// KeyByUint64 creates a map keyed by uint64 field from a slice of struct pointers
func KeyByUint64[M any](list []*M, fieldName string) map[uint64]*M {
	if len(list) == 0 {
		return map[uint64]*M{}
	}

	lv := reflect.ValueOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := make(map[uint64]*M, lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m[elemStruct.FieldByIndex(field.Index).Uint()] = elem.Interface().(*M)
	}

	return m
}

// KeyByGeneric creates a map keyed by a field using a selector function (type-safe version)
func KeyByGeneric[T any, K comparable](list []T, keySelector func(T) K) map[K]T {
	if len(list) == 0 {
		return make(map[K]T)
	}

	result := make(map[K]T, len(list))
	for _, item := range list {
		key := keySelector(item)
		result[key] = item
	}

	return result
}

// KeyByPtr creates a map keyed by a field using a selector function for pointer slices
func KeyByPtr[T any, K comparable](list []*T, keySelector func(*T) K) map[K]*T {
	if len(list) == 0 {
		return make(map[K]*T)
	}

	result := make(map[K]*T, len(list))
	for _, item := range list {
		if item != nil {
			key := keySelector(item)
			result[key] = item
		}
	}

	return result
}