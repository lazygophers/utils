package candy

import (
	"reflect"
	"sync"

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

// ==================== SliceField2Map 优化：反射字段缓存 ====================
// 使用缓存避免重复的字段查找，显著提升性能

// sliceField2MapCache 字段索引缓存
var sliceField2MapCache struct {
	sync.RWMutex
	cache map[reflect.Type]map[string][]int
}

func init() {
	sliceField2MapCache.cache = make(map[reflect.Type]map[string][]int)
}

// getFieldIndexCached 获取字段索引，使用缓存避免重复反射
func getFieldIndexCachedForSlice(elemType reflect.Type, fieldName string) ([]int, reflect.Type, bool) {
	// 读缓存
	sliceField2MapCache.RLock()
	if fields, ok := sliceField2MapCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			sliceField2MapCache.RUnlock()
			// 解析实际类型
			actualType := elemType
			for actualType.Kind() == reflect.Ptr {
				actualType = actualType.Elem()
			}
			if field, found := actualType.FieldByName(fieldName); found {
				return index, field.Type, true
			}
		}
	}
	sliceField2MapCache.RUnlock()

	// 缓存未命中，获取并缓存
	sliceField2MapCache.Lock()
	defer sliceField2MapCache.Unlock()

	// 双重检查
	if fields, ok := sliceField2MapCache.cache[elemType]; ok {
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

	// 解析指针类型
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

	// 初始化类型缓存
	if sliceField2MapCache.cache[elemType] == nil {
		sliceField2MapCache.cache[elemType] = make(map[string][]int)
	}

	fieldIndex := field.Index
	sliceField2MapCache.cache[elemType][fieldName] = fieldIndex

	return fieldIndex, field.Type, true
}

func sliceField2Map[T any, K comparable](ss []T, fieldName string, expectedKind reflect.Kind, converter func(reflect.Value) K) map[K]bool {
	if len(ss) == 0 {
		return nil
	}

	elemType := reflect.TypeOf(ss[0])
	fieldIndex, fieldType, ok := getFieldIndexCachedForSlice(elemType, fieldName)
	if !ok {
		panic("field not found or element is not a struct")
	}

	if fieldType.Kind() != expectedKind {
		panic("field type mismatch")
	}

	ret := make(map[K]bool, len(ss))
	for _, item := range ss {
		v := reflect.ValueOf(item)
		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		fieldValue := v.FieldByIndex(fieldIndex)
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
		return int8(fieldValue.Int()) // #nosec G115 -- intentional truncation for best-effort conversion
	})
}

// SliceField2MapInt16 extracts int16 field values from a struct slice and creates a map[int16]bool
// Panics if the field doesn't exist or is not an int16 type
func SliceField2MapInt16[T any](ss []T, fieldName string) map[int16]bool {
	return sliceField2Map(ss, fieldName, reflect.Int16, func(fieldValue reflect.Value) int16 {
		return int16(fieldValue.Int()) // #nosec G115 -- intentional truncation for best-effort conversion
	})
}

// SliceField2MapInt32 extracts int32 field values from a struct slice and creates a map[int32]bool
// Panics if the field doesn't exist or is not an int32 type
func SliceField2MapInt32[T any](ss []T, fieldName string) map[int32]bool {
	return sliceField2Map(ss, fieldName, reflect.Int32, func(fieldValue reflect.Value) int32 {
		return int32(fieldValue.Int()) // #nosec G115 -- intentional truncation for best-effort conversion
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
		return uint8(fieldValue.Uint()) // #nosec G115 -- intentional truncation for best-effort conversion
	})
}

// SliceField2MapUint16 extracts uint16 field values from a struct slice and creates a map[uint16]bool
// Panics if the field doesn't exist or is not a uint16 type
func SliceField2MapUint16[T any](ss []T, fieldName string) map[uint16]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint16, func(fieldValue reflect.Value) uint16 {
		return uint16(fieldValue.Uint()) // #nosec G115 -- intentional truncation for best-effort conversion
	})
}

// SliceField2MapUint32 extracts uint32 field values from a struct slice and creates a map[uint32]bool
// Panics if the field doesn't exist or is not a uint32 type
func SliceField2MapUint32[T any](ss []T, fieldName string) map[uint32]bool {
	return sliceField2Map(ss, fieldName, reflect.Uint32, func(fieldValue reflect.Value) uint32 {
		return uint32(fieldValue.Uint()) // #nosec G115 -- intentional truncation for best-effort conversion
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
