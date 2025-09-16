package candy

import (
	"cmp"
	"reflect"
)

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