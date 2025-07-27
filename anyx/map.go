package anyx

import (
	"fmt"
	"maps"
	"reflect"

	"github.com/lazygophers/utils/json"
	"golang.org/x/exp/constraints"
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
		panic("map key type required string")
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
		panic("map key type required string")
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

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int32 {
		panic("map key type required string")
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

	keyType := t.Type().Key()
	if keyType.Kind() != reflect.Int64 {
		panic("map key type required string")
	}

	result := make([]int64, 0, t.Len())
	for _, v := range t.MapKeys() {
		result = append(result, v.Int())
	}

	return result
}

func MapValues[K constraints.Ordered, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func MergeMap[K constraints.Ordered, V any](source, target map[K]V) map[K]V {
	res := maps.Clone(source)

	if len(target) > 0 {
		for k, v := range target {
			res[k] = v
		}
	}

	return res
}

func KeyBy(list interface{}, fieldName string) interface{} {
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

func Slice2Map[M constraints.Ordered](list []M) map[M]bool {
	m := make(map[M]bool, len(list))

	for _, v := range list {
		m[v] = true
	}

	return m
}

func ToMapStringAny(v interface{}) map[string]interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]interface{}{}
	}

	m := make(map[string]any)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = mg.Value().Interface()
	}

	return m
}

func ToMap(v interface{}) map[string]interface{} {
	switch x := v.(type) {
	case []byte:
		var m map[string]any
		err := json.Unmarshal(x, &m)
		if err == nil {
			return m
		}

	case string:
		var m map[string]any
		err := json.UnmarshalString(x, &m)
		if err == nil {
			return m
		}

	}

	return ToMapStringAny(v)
}

func ToMapStringString(v interface{}) map[string]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]string{}
	}

	m := make(map[string]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToString(mg.Value().Interface())
	}

	return m
}

func ToMapStringInt64(v interface{}) map[string]int64 {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string]int64{}
	}

	m := make(map[string]int64)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToInt64(mg.Value().Interface())
	}

	return m
}

func ToMapInt64String(v interface{}) map[int64]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[int64]string{}
	}

	m := make(map[int64]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToInt64(mg.Key().Interface())] = ToString(mg.Value().Interface())
	}

	return m
}

func ToMapInt32String(v interface{}) map[int32]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[int32]string{}
	}

	m := make(map[int32]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToInt32(mg.Key().Interface())] = ToString(mg.Value().Interface())
	}

	return m
}

func ToMapStringArrayString(v interface{}) map[string][]string {
	vv := reflect.ValueOf(v)
	if vv.Kind() != reflect.Map {
		return map[string][]string{}
	}

	m := make(map[string][]string)

	mg := vv.MapRange()

	for mg.Next() {
		m[ToString(mg.Key().Interface())] = ToArrayString(mg.Value().Interface())
	}

	return m
}
