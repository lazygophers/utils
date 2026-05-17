package candy

import (
	"fmt"
	"reflect"
	"sync"
)

var unifiedPluckCache struct {
	sync.RWMutex
	cache map[reflect.Type]map[string][]int
}

func init() {
	unifiedPluckCache.cache = make(map[reflect.Type]map[string][]int)
}

func getPluckFieldIndex(elemType reflect.Type, fieldName string) ([]int, reflect.Type, bool) {
	unifiedPluckCache.RLock()
	if fields, ok := unifiedPluckCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			unifiedPluckCache.RUnlock()
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}
	unifiedPluckCache.RUnlock()

	unifiedPluckCache.Lock()
	defer unifiedPluckCache.Unlock()

	if fields, ok := unifiedPluckCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}

	actualType := elemType
	for actualType.Kind() == reflect.Ptr {
		actualType = actualType.Elem()
	}
	if actualType.Kind() != reflect.Struct {
		return nil, nil, false
	}

	field, found := actualType.FieldByName(fieldName)
	if !found {
		return nil, nil, false
	}

	if unifiedPluckCache.cache[elemType] == nil {
		unifiedPluckCache.cache[elemType] = make(map[string][]int)
	}
	unifiedPluckCache.cache[elemType][fieldName] = field.Index
	return field.Index, field.Type, true
}

func pluckTyped[T any](list interface{}, fieldName string, expectedKind reflect.Kind, convert func(reflect.Value) T) []T {
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic("list must be an array or slice")
	}
	if v.Len() == 0 {
		return nil
	}

	elemType := v.Type().Elem()
	fieldIndex, fieldValueType, ok := getPluckFieldIndex(elemType, fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}
	if fieldValueType.Kind() != expectedKind {
		panic(fmt.Sprintf("field %s is not of type %s", fieldName, expectedKind))
	}

	result := make([]T, v.Len())
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		for elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		if elem.Kind() != reflect.Struct {
			panic("element is not a struct")
		}
		result[i] = convert(elem.FieldByIndex(fieldIndex))
	}
	return result
}

func PluckInt8(list interface{}, fieldName string) []int8 {
	return pluckTyped(list, fieldName, reflect.Int8, func(v reflect.Value) int8 { return int8(v.Int()) })
}

func PluckInt16(list interface{}, fieldName string) []int16 {
	return pluckTyped(list, fieldName, reflect.Int16, func(v reflect.Value) int16 { return int16(v.Int()) })
}

func PluckInt32(list interface{}, fieldName string) []int32 {
	return pluckTyped(list, fieldName, reflect.Int32, func(v reflect.Value) int32 { return int32(v.Int()) })
}

func PluckInt64(list interface{}, fieldName string) []int64 {
	return pluckTyped(list, fieldName, reflect.Int64, func(v reflect.Value) int64 { return v.Int() })
}

func PluckUint(list interface{}, fieldName string) []uint {
	return pluckTyped(list, fieldName, reflect.Uint, func(v reflect.Value) uint { return uint(v.Uint()) })
}

func PluckUint8(list interface{}, fieldName string) []uint8 {
	return pluckTyped(list, fieldName, reflect.Uint8, func(v reflect.Value) uint8 { return uint8(v.Uint()) })
}

func PluckUint16(list interface{}, fieldName string) []uint16 {
	return pluckTyped(list, fieldName, reflect.Uint16, func(v reflect.Value) uint16 { return uint16(v.Uint()) })
}

func PluckUint32(list interface{}, fieldName string) []uint32 {
	return pluckTyped(list, fieldName, reflect.Uint32, func(v reflect.Value) uint32 { return uint32(v.Uint()) })
}

func PluckUint64(list interface{}, fieldName string) []uint64 {
	return pluckTyped(list, fieldName, reflect.Uint64, func(v reflect.Value) uint64 { return v.Uint() })
}

func PluckFloat32(list interface{}, fieldName string) []float32 {
	return pluckTyped(list, fieldName, reflect.Float32, func(v reflect.Value) float32 { return float32(v.Float()) })
}

func PluckFloat64(list interface{}, fieldName string) []float64 {
	return pluckTyped(list, fieldName, reflect.Float64, func(v reflect.Value) float64 { return v.Float() })
}

func PluckBool(list interface{}, fieldName string) []bool {
	return pluckTyped(list, fieldName, reflect.Bool, func(v reflect.Value) bool { return v.Bool() })
}
