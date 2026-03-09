package candy

import (
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

func sliceField2Map[T any, K comparable](ss []T, fieldName string, expectedKind reflect.Kind, converter func(reflect.Value) K) map[K]bool {
	if len(ss) == 0 {
		return nil
	}

	validateStructFieldKind(reflect.TypeOf(ss[0]), fieldName, expectedKind)

	ret := make(map[K]bool, len(ss))
	for _, item := range ss {
		fieldValue := getStructFieldValue(item, fieldName)
		ret[converter(fieldValue)] = true
	}

	return ret
}

// SliceField2MapString extracts string field values from a struct slice and creates a map[string]bool
// Panics if the field doesn't exist or is not a string type
func SliceField2MapString[T any](ss []T, fieldName string) map[string]bool {
	return sliceField2Map(ss, fieldName, reflect.String, func(fieldValue reflect.Value) string {
		return fieldValue.String()
	})
}

// SliceField2MapInt extracts int field values from a struct slice and creates a map[int]bool
// Panics if the field doesn't exist or is not an int type
func SliceField2MapInt[T any](ss []T, fieldName string) map[int]bool {
	return sliceField2Map(ss, fieldName, reflect.Int, func(fieldValue reflect.Value) int {
		return int(fieldValue.Int())
	})
}

// SliceField2MapInt8 extracts int8 field values from a struct slice and creates a map[int8]bool
// Panics if the field doesn't exist or is not an int8 type
func SliceField2MapInt8[T any](ss []T, fieldName string) map[int8]bool {
	return sliceField2Map(ss, fieldName, reflect.Int8, func(fieldValue reflect.Value) int8 {
		return int8(fieldValue.Int())
	})
}

// SliceField2MapInt16 extracts int16 field values from a struct slice and creates a map[int16]bool
// Panics if the field doesn't exist or is not an int16 type
func SliceField2MapInt16[T any](ss []T, fieldName string) map[int16]bool {
	return sliceField2Map(ss, fieldName, reflect.Int16, func(fieldValue reflect.Value) int16 {
		return int16(fieldValue.Int())
	})
}

// SliceField2MapInt32 extracts int32 field values from a struct slice and creates a map[int32]bool
// Panics if the field doesn't exist or is not an int32 type
func SliceField2MapInt32[T any](ss []T, fieldName string) map[int32]bool {
	return sliceField2Map(ss, fieldName, reflect.Int32, func(fieldValue reflect.Value) int32 {
		return int32(fieldValue.Int())
	})
}

// SliceField2MapInt64 extracts int64 field values from a struct slice and creates a map[int64]bool
// Panics if the field doesn't exist or is not an int64 type
func SliceField2MapInt64[T any](ss []T, fieldName string) map[int64]bool {
	return sliceField2Map(ss, fieldName, reflect.Int64, func(fieldValue reflect.Value) int64 {
		return fieldValue.Int()
	})
}

// SliceField2MapUint extracts uint field values from a struct slice and creates a map[uint]bool
// Panics if the field doesn't exist or is not a uint type
func SliceField2MapUint[T any](ss []T, fieldName string) map[uint]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint, func(fieldValue reflect.Value) uint {
		return uint(fieldValue.Uint())
	})
}

// SliceField2MapUint8 extracts uint8 field values from a struct slice and creates a map[uint8]bool
// Panics if the field doesn't exist or is not a uint8 type
func SliceField2MapUint8[T any](ss []T, fieldName string) map[uint8]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint8, func(fieldValue reflect.Value) uint8 {
		return uint8(fieldValue.Uint())
	})
}

// SliceField2MapUint16 extracts uint16 field values from a struct slice and creates a map[uint16]bool
// Panics if the field doesn't exist or is not a uint16 type
func SliceField2MapUint16[T any](ss []T, fieldName string) map[uint16]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint16, func(fieldValue reflect.Value) uint16 {
		return uint16(fieldValue.Uint())
	})
}

// SliceField2MapUint32 extracts uint32 field values from a struct slice and creates a map[uint32]bool
// Panics if the field doesn't exist or is not a uint32 type
func SliceField2MapUint32[T any](ss []T, fieldName string) map[uint32]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint32, func(fieldValue reflect.Value) uint32 {
		return uint32(fieldValue.Uint())
	})
}

// SliceField2MapUint64 extracts uint64 field values from a struct slice and creates a map[uint64]bool
// Panics if the field doesn't exist or is not a uint64 type
func SliceField2MapUint64[T any](ss []T, fieldName string) map[uint64]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint64, func(fieldValue reflect.Value) uint64 {
		return fieldValue.Uint()
	})
}

// SliceField2MapFloat32 extracts float32 field values from a struct slice and creates a map[float32]bool
// Panics if the field doesn't exist or is not a float32 type
func SliceField2MapFloat32[T any](ss []T, fieldName string) map[float32]bool {
	return sliceField2Map(ss, fieldName, reflect.Float32, func(fieldValue reflect.Value) float32 {
		return float32(fieldValue.Float())
	})
}

// SliceField2MapFloat64 extracts float64 field values from a struct slice and creates a map[float64]bool
// Panics if the field doesn't exist or is not a float64 type
func SliceField2MapFloat64[T any](ss []T, fieldName string) map[float64]bool {
	return sliceField2Map(ss, fieldName, reflect.Float64, func(fieldValue reflect.Value) float64 {
		return fieldValue.Float()
	})
}

// SliceField2MapBool extracts bool field values from a struct slice and creates a map[bool]bool
// Panics if the field doesn't exist or is not a bool type
func SliceField2MapBool[T any](ss []T, fieldName string) map[bool]bool {
	return sliceField2Map(ss, fieldName, reflect.Bool, func(fieldValue reflect.Value) bool {
		return fieldValue.Bool()
	})
}
