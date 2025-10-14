package candy

import (
	"cmp"
	"fmt"
	"reflect"
)

// ValueType represents the type of a value for type checking
type ValueType int

const (
	// ValueUnknown represents an unknown or unsupported type
	ValueUnknown ValueType = iota
	// ValueNumber represents numeric types (int, float, etc.)
	ValueNumber
	// ValueString represents string and byte slice types
	ValueString
	// ValueBool represents boolean type
	ValueBool
)

// CheckValueType determines the general category of a value's type
func CheckValueType(val interface{}) ValueType {
	switch val.(type) {
	case bool:
		return ValueBool
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return ValueNumber
	case string, []byte:
		return ValueString
	default:
		return ValueUnknown
	}
}

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

// MapKeysString extracts all string keys from a map
func MapKeysString(m interface{}) []string {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		panic("nil map")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.String {
		panic("map key type required string")
	}

	result := make([]string, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.String())
	}

	return result
}

// MapKeysInt extracts all int keys from a map
func MapKeysInt(m interface{}) []int {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []int{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int {
		panic("map key type required int")
	}

	result := make([]int, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, int(v.Int()))
	}

	return result
}

// MapKeysInt8 extracts all int8 keys from a map
func MapKeysInt8(m interface{}) []int8 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []int8{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int8 {
		panic("map key type required int8")
	}

	result := make([]int8, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, int8(v.Int()))
	}

	return result
}

// MapKeysInt16 extracts all int16 keys from a map
func MapKeysInt16(m interface{}) []int16 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []int16{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int16 {
		panic("map key type required int16")
	}

	result := make([]int16, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, int16(v.Int()))
	}

	return result
}

// MapKeysInt32 extracts all int32 keys from a map
func MapKeysInt32(m interface{}) []int32 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []int32{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int32 {
		panic("map key type required int32")
	}

	result := make([]int32, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, int32(v.Int()))
	}

	return result
}

// MapKeysInt64 extracts all int64 keys from a map
func MapKeysInt64(m interface{}) []int64 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []int64{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int64 {
		panic("map key type required int64")
	}

	result := make([]int64, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Int())
	}

	return result
}

// MapKeysUint extracts all uint keys from a map
func MapKeysUint(m interface{}) []uint {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []uint{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint {
		panic("map key type required uint")
	}

	result := make([]uint, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, uint(v.Uint()))
	}

	return result
}

// MapKeysUint8 extracts all uint8 keys from a map
func MapKeysUint8(m interface{}) []uint8 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []uint8{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint8 {
		panic("map key type required uint8")
	}

	result := make([]uint8, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, uint8(v.Uint()))
	}

	return result
}

// MapKeysUint16 extracts all uint16 keys from a map
func MapKeysUint16(m interface{}) []uint16 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []uint16{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint16 {
		panic("map key type required uint16")
	}

	result := make([]uint16, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, uint16(v.Uint()))
	}

	return result
}

// MapKeysUint32 extracts all uint32 keys from a map
func MapKeysUint32(m interface{}) []uint32 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		panic("nil map")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint32 {
		panic("map key type required uint32")
	}

	result := make([]uint32, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, uint32(v.Uint()))
	}

	return result
}

// MapKeysUint64 extracts all uint64 keys from a map
func MapKeysUint64(m interface{}) []uint64 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Uint64 {
		panic("map key type required uint64")
	}

	result := make([]uint64, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Uint())
	}

	return result
}

// MapKeysFloat32 extracts all float32 keys from a map
func MapKeysFloat32(m interface{}) []float32 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []float32{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Float32 {
		panic("map key type required float32")
	}

	result := make([]float32, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, float32(v.Float()))
	}

	return result
}

// MapKeysFloat64 extracts all float64 keys from a map
func MapKeysFloat64(m interface{}) []float64 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []float64{}
	}

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Float64 {
		panic("map key type required float64")
	}

	result := make([]float64, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Float())
	}

	return result
}

// MapKeysInterface extracts all keys from a map as interface{} slice
func MapKeysInterface(m interface{}) []interface{} {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []interface{}{}
	}

	result := make([]interface{}, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Interface())
	}

	return result
}

// MapKeysAny extracts all keys from a map as interface{} slice (alias for MapKeysInterface)
func MapKeysAny(m interface{}) []interface{} {
	return MapKeysInterface(m)
}

// MapKeysNumber extracts all numeric keys from a map as interface{} slice
func MapKeysNumber(m interface{}) []interface{} {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []interface{}{}
	}

	keyType := t.Type().Key()
	switch keyType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		// valid number types
	default:
		panic("map key type required number")
	}

	result := make([]interface{}, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Interface())
	}

	return result
}

// MapKeysGeneric extracts all keys from a map using generics for type safety
func MapKeysGeneric[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}

	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}

	return result
}

// MapValues extracts all values from a map using generics for type safety
func MapValues[K cmp.Ordered, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// MapValuesGeneric extracts all values from a map using generics for any comparable key type
func MapValuesGeneric[K comparable, V any](m map[K]V) []V {
	if m == nil {
		return nil
	}

	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// MapValuesAny extracts all values from a map as interface{} slice
func MapValuesAny(m interface{}) []interface{} {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []interface{}{}
	}

	result := make([]interface{}, 0, t.Len())
	iter := t.MapRange()
	for iter.Next() {
		result = append(result, iter.Value().Interface())
	}

	return result
}

// MapValuesString extracts all string values from a map
func MapValuesString(m interface{}) []string {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []string{}
	}

	result := make([]string, 0, t.Len())
	iter := t.MapRange()
	for iter.Next() {
		result = append(result, iter.Value().String())
	}

	return result
}

// MapValuesInt extracts all int values from a map
func MapValuesInt(m interface{}) []int {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []int{}
	}

	result := make([]int, 0, t.Len())
	iter := t.MapRange()
	for iter.Next() {
		result = append(result, int(iter.Value().Int()))
	}

	return result
}

// MapValuesFloat64 extracts all float64 values from a map
func MapValuesFloat64(m interface{}) []float64 {
	t := reflect.ValueOf(m)
	if t.Kind() != reflect.Map {
		panic("required map type")
	}
	if t.IsNil() {
		return []float64{}
	}

	result := make([]float64, 0, t.Len())
	iter := t.MapRange()
	for iter.Next() {
		result = append(result, iter.Value().Float())
	}

	return result
}

// MergeMap merges two maps, with target values overriding source values
func MergeMap[K cmp.Ordered, V any](source, target map[K]V) map[K]V {
	res := make(map[K]V, len(source))

	// Manual clone implementation
	for k, v := range source {
		res[k] = v
	}

	if len(target) > 0 {
		for k, v := range target {
			res[k] = v
		}
	}

	return res
}

// MergeMapGeneric merges two maps with any comparable key type
func MergeMapGeneric[K comparable, V any](source, target map[K]V) map[K]V {
	if source == nil && target == nil {
		return nil
	}
	if source == nil {
		return CloneMapShallow(target)
	}
	if target == nil {
		return CloneMapShallow(source)
	}

	res := make(map[K]V, len(source)+len(target))

	// Copy source map
	for k, v := range source {
		res[k] = v
	}

	// Override with target map values
	for k, v := range target {
		res[k] = v
	}

	return res
}

// CloneMapShallow creates a shallow copy of a map
func CloneMapShallow[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}
