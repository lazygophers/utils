package anyx

import (
	"cmp"

	"github.com/lazygophers/utils/candy"
)

// Re-export types and constants from candy for backward compatibility
type ValueType = candy.ValueType

const (
	ValueUnknown = candy.ValueUnknown
	ValueNumber  = candy.ValueNumber
	ValueString  = candy.ValueString
	ValueBool    = candy.ValueBool
)

// Re-export utility functions
var CheckValueType = candy.CheckValueType

// Re-export map key functions from candy
var MapKeysString = candy.MapKeysString
var MapKeysUint32 = candy.MapKeysUint32
var MapKeysUint64 = candy.MapKeysUint64
var MapKeysInt32 = candy.MapKeysInt32
var MapKeysInt64 = candy.MapKeysInt64
var MapKeysInt = candy.MapKeysInt
var MapKeysInt8 = candy.MapKeysInt8
var MapKeysInt16 = candy.MapKeysInt16
var MapKeysUint = candy.MapKeysUint
var MapKeysUint8 = candy.MapKeysUint8
var MapKeysUint16 = candy.MapKeysUint16
var MapKeysFloat32 = candy.MapKeysFloat32
var MapKeysFloat64 = candy.MapKeysFloat64
var MapKeysInterface = candy.MapKeysInterface
var MapKeysAny = candy.MapKeysAny
var MapKeysNumber = candy.MapKeysNumber

// Re-export non-generic map value functions from candy
var MapValuesAny = candy.MapValuesAny
var MapValuesString = candy.MapValuesString
var MapValuesInt = candy.MapValuesInt
var MapValuesFloat64 = candy.MapValuesFloat64

// Re-export non-generic functions from candy
var KeyBy = candy.KeyBy

// Wrapper functions for generic functions to maintain backward compatibility

// MapValues extracts all values from a map using generics for type safety
func MapValues[K cmp.Ordered, V any](m map[K]V) []V {
	return candy.MapValues(m)
}

// MergeMap merges two maps, with target values overriding source values
func MergeMap[K cmp.Ordered, V any](source, target map[K]V) map[K]V {
	return candy.MergeMap(source, target)
}

// KeyByUint64 creates a map keyed by uint64 field from a slice of struct pointers
func KeyByUint64[M any](list []*M, fieldName string) map[uint64]*M {
	return candy.KeyByUint64(list, fieldName)
}

// KeyByInt64 creates a map keyed by int64 field from a slice of struct pointers
func KeyByInt64[M any](list []*M, fieldName string) map[int64]*M {
	return candy.KeyByInt64(list, fieldName)
}

// KeyByString creates a map keyed by string field from a slice of struct pointers
func KeyByString[M any](list []*M, fieldName string) map[string]*M {
	return candy.KeyByString(list, fieldName)
}

// KeyByInt32 creates a map keyed by int32 field from a slice of struct pointers
func KeyByInt32[M any](list []*M, fieldName string) map[int32]*M {
	return candy.KeyByInt32(list, fieldName)
}

// Slice2Map converts a slice to a map where each element becomes a key with value true
func Slice2Map[M cmp.Ordered](list []M) map[M]bool {
	return candy.Slice2Map(list)
}