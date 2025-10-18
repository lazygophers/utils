package candy

import (
	"golang.org/x/exp/constraints"
	"reflect"
)

func MapKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapKeysInt(m interface{}) []int {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]int, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Int {
			keys = append(keys, int(key.Int()))
		}
	}
	return keys
}

func MapKeysInt8(m interface{}) []int8 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]int8, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Int8 {
			keys = append(keys, int8(key.Int()))
		}
	}
	return keys
}

func MapKeysInt16(m interface{}) []int16 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]int16, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Int16 {
			keys = append(keys, int16(key.Int()))
		}
	}
	return keys
}

func MapKeysInt32(m interface{}) []int32 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]int32, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Int32 {
			keys = append(keys, int32(key.Int()))
		}
	}
	return keys
}

func MapKeysInt64(m interface{}) []int64 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]int64, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Int64 {
			keys = append(keys, key.Int())
		}
	}
	return keys
}

func MapKeysUint(m interface{}) []uint {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]uint, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Uint {
			keys = append(keys, uint(key.Uint()))
		}
	}
	return keys
}

func MapKeysUint8(m interface{}) []uint8 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]uint8, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Uint8 {
			keys = append(keys, uint8(key.Uint()))
		}
	}
	return keys
}

func MapKeysUint16(m interface{}) []uint16 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]uint16, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Uint16 {
			keys = append(keys, uint16(key.Uint()))
		}
	}
	return keys
}

func MapKeysUint32(m interface{}) []uint32 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]uint32, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Uint32 {
			keys = append(keys, uint32(key.Uint()))
		}
	}
	return keys
}

func MapKeysUint64(m interface{}) []uint64 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]uint64, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Uint64 {
			keys = append(keys, key.Uint())
		}
	}
	return keys
}

func MapKeysFloat32(m interface{}) []float32 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]float32, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Float32 {
			keys = append(keys, float32(key.Float()))
		}
	}
	return keys
}

func MapKeysFloat64(m interface{}) []float64 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]float64, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Float64 {
			keys = append(keys, key.Float())
		}
	}
	return keys
}

func MapKeysString(m interface{}) []string {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]string, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.String {
			keys = append(keys, key.String())
		}
	}
	return keys
}

func MapKeysBytes(m interface{}) [][]byte {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([][]byte, 0, val.Len())
	for _, key := range val.MapKeys() {
		if key.Kind() == reflect.Slice && key.Type().Elem().Kind() == reflect.Uint8 {
			keys = append(keys, key.Bytes())
		}
	}
	return keys
}

func MapKeysAny(m interface{}) []interface{} {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	keys := make([]interface{}, 0, val.Len())
	for _, key := range val.MapKeys() {
		keys = append(keys, key.Interface())
	}
	return keys
}
