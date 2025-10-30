package candy

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Slice2Map converts a slice to a map where each element becomes a key with value true
func Slice2Map[M constraints.Ordered](list []M) map[M]bool {
	m := make(map[M]bool, len(list))

	for _, v := range list {
		m[v] = true
	}

	return m
}

// validateSliceFieldForMap validates the field for SliceField2Map functions
func validateSliceFieldForMap(elemType reflect.Type, fieldName string, expectedKind reflect.Kind) {
	// Dereference pointer types
	for elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	// Check if it's a struct
	if elemType.Kind() != reflect.Struct {
		panic("element is not a struct")
	}

	// Find the field
	field, ok := elemType.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	// Check field type
	if field.Type.Kind() != expectedKind {
		panic(fmt.Sprintf("field %s is not of type %s", fieldName, expectedKind))
	}
}

// getSliceFieldValue gets the value of a struct field
func getSliceFieldValue(item interface{}, fieldName string) reflect.Value {
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

// SliceField2MapString extracts string field values from a struct slice and creates a map[string]bool
// Panics if the field doesn't exist or is not a string type
func SliceField2MapString[T any](ss []T, fieldName string) map[string]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.String)

	// Create result map
	ret := make(map[string]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := fieldValue.String()
		ret[key] = true
	}

	return ret
}

// SliceField2MapInt extracts int field values from a struct slice and creates a map[int]bool
// Panics if the field doesn't exist or is not an int type
func SliceField2MapInt[T any](ss []T, fieldName string) map[int]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Int)

	// Create result map
	ret := make(map[int]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := int(fieldValue.Int())
		ret[key] = true
	}

	return ret
}

// SliceField2MapInt8 extracts int8 field values from a struct slice and creates a map[int8]bool
// Panics if the field doesn't exist or is not an int8 type
func SliceField2MapInt8[T any](ss []T, fieldName string) map[int8]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Int8)

	// Create result map
	ret := make(map[int8]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := int8(fieldValue.Int())
		ret[key] = true
	}

	return ret
}

// SliceField2MapInt16 extracts int16 field values from a struct slice and creates a map[int16]bool
// Panics if the field doesn't exist or is not an int16 type
func SliceField2MapInt16[T any](ss []T, fieldName string) map[int16]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Int16)

	// Create result map
	ret := make(map[int16]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := int16(fieldValue.Int())
		ret[key] = true
	}

	return ret
}

// SliceField2MapInt32 extracts int32 field values from a struct slice and creates a map[int32]bool
// Panics if the field doesn't exist or is not an int32 type
func SliceField2MapInt32[T any](ss []T, fieldName string) map[int32]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Int32)

	// Create result map
	ret := make(map[int32]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := int32(fieldValue.Int())
		ret[key] = true
	}

	return ret
}

// SliceField2MapInt64 extracts int64 field values from a struct slice and creates a map[int64]bool
// Panics if the field doesn't exist or is not an int64 type
func SliceField2MapInt64[T any](ss []T, fieldName string) map[int64]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Int64)

	// Create result map
	ret := make(map[int64]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := fieldValue.Int()
		ret[key] = true
	}

	return ret
}

// SliceField2MapUint extracts uint field values from a struct slice and creates a map[uint]bool
// Panics if the field doesn't exist or is not a uint type
func SliceField2MapUint[T any](ss []T, fieldName string) map[uint]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Uint)

	// Create result map
	ret := make(map[uint]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := uint(fieldValue.Uint())
		ret[key] = true
	}

	return ret
}

// SliceField2MapUint8 extracts uint8 field values from a struct slice and creates a map[uint8]bool
// Panics if the field doesn't exist or is not a uint8 type
func SliceField2MapUint8[T any](ss []T, fieldName string) map[uint8]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Uint8)

	// Create result map
	ret := make(map[uint8]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := uint8(fieldValue.Uint())
		ret[key] = true
	}

	return ret
}

// SliceField2MapUint16 extracts uint16 field values from a struct slice and creates a map[uint16]bool
// Panics if the field doesn't exist or is not a uint16 type
func SliceField2MapUint16[T any](ss []T, fieldName string) map[uint16]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Uint16)

	// Create result map
	ret := make(map[uint16]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := uint16(fieldValue.Uint())
		ret[key] = true
	}

	return ret
}

// SliceField2MapUint32 extracts uint32 field values from a struct slice and creates a map[uint32]bool
// Panics if the field doesn't exist or is not a uint32 type
func SliceField2MapUint32[T any](ss []T, fieldName string) map[uint32]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Uint32)

	// Create result map
	ret := make(map[uint32]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := uint32(fieldValue.Uint())
		ret[key] = true
	}

	return ret
}

// SliceField2MapUint64 extracts uint64 field values from a struct slice and creates a map[uint64]bool
// Panics if the field doesn't exist or is not a uint64 type
func SliceField2MapUint64[T any](ss []T, fieldName string) map[uint64]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Uint64)

	// Create result map
	ret := make(map[uint64]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := fieldValue.Uint()
		ret[key] = true
	}

	return ret
}

// SliceField2MapFloat32 extracts float32 field values from a struct slice and creates a map[float32]bool
// Panics if the field doesn't exist or is not a float32 type
func SliceField2MapFloat32[T any](ss []T, fieldName string) map[float32]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Float32)

	// Create result map
	ret := make(map[float32]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := float32(fieldValue.Float())
		ret[key] = true
	}

	return ret
}

// SliceField2MapFloat64 extracts float64 field values from a struct slice and creates a map[float64]bool
// Panics if the field doesn't exist or is not a float64 type
func SliceField2MapFloat64[T any](ss []T, fieldName string) map[float64]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Float64)

	// Create result map
	ret := make(map[float64]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := fieldValue.Float()
		ret[key] = true
	}

	return ret
}

// SliceField2MapBool extracts bool field values from a struct slice and creates a map[bool]bool
// Panics if the field doesn't exist or is not a bool type
func SliceField2MapBool[T any](ss []T, fieldName string) map[bool]bool {
	if len(ss) == 0 {
		return nil
	}

	// Validate field
	validateSliceFieldForMap(reflect.TypeOf(ss[0]), fieldName, reflect.Bool)

	// Create result map
	ret := make(map[bool]bool, len(ss))

	// Iterate through slice and build map
	for _, item := range ss {
		fieldValue := getSliceFieldValue(item, fieldName)
		key := fieldValue.Bool()
		ret[key] = true
	}

	return ret
}
