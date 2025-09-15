package anyx

import (
	"fmt"
	"reflect"

	"cmp"
)

type ValueType int

const (
	ValueUnknown ValueType = iota
	ValueNumber
	ValueString
	ValueBool
)

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

func MapKeysAny(m interface{}) []interface{} {
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

func MapValues[K cmp.Ordered, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

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

func Slice2Map[M cmp.Ordered](list []M) map[M]bool {
	m := make(map[M]bool, len(list))

	for _, v := range list {
		m[v] = true
	}

	return m
}
