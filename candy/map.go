package candy

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

func MapKeys[K constraints.Ordered, V any](m map[K]V) []K {
	if len(m) == 0 {
		return nil
	}
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func MapKeysInt(m interface{}) []int {
	// 快速路径：类型断言避免反射开销
	switch m := m.(type) {
	case map[int]int:
		if len(m) == 0 {
			return nil
		}
		keys := make([]int, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[int]interface{}:
		if len(m) == 0 {
			return nil
		}
		keys := make([]int, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[int]string:
		if len(m) == 0 {
			return nil
		}
		keys := make([]int, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[int]bool:
		if len(m) == 0 {
			return nil
		}
		keys := make([]int, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	}

	// 反射路径
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：检查 key 类型是否匹配，如果匹配则使用索引访问
	mapKeys := val.MapKeys()
	if len(mapKeys) > 0 && mapKeys[0].Kind() == reflect.Int {
		// 所有 key 都是 int 类型，直接使用索引
		keys := make([]int, len(mapKeys))
		for i, key := range mapKeys {
			keys[i] = int(key.Int())
		}
		return keys
	}

	// key 类型不匹配，需要过滤
	keys := make([]int, 0, val.Len())
	for _, key := range mapKeys {
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

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]int8, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Int8 {
			keys[i] = int8(key.Int()) // #nosec G115 -- intentional truncation for best-effort conversion
		}
	}
	return keys
}

func MapKeysInt16(m interface{}) []int16 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]int16, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Int16 {
			keys[i] = int16(key.Int()) // #nosec G115 -- intentional truncation for best-effort conversion
		}
	}
	return keys
}

func MapKeysInt32(m interface{}) []int32 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]int32, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Int32 {
			keys[i] = int32(key.Int()) // #nosec G115 -- intentional truncation for best-effort conversion
		}
	}
	return keys
}

func MapKeysInt64(m interface{}) []int64 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]int64, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Int64 {
			keys[i] = key.Int()
		}
	}
	return keys
}

func MapKeysUint(m interface{}) []uint {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]uint, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Uint {
			keys[i] = uint(key.Uint())
		}
	}
	return keys
}

func MapKeysUint8(m interface{}) []uint8 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]uint8, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Uint8 {
			keys[i] = uint8(key.Uint()) // #nosec G115 -- intentional truncation for best-effort conversion
		}
	}
	return keys
}

func MapKeysUint16(m interface{}) []uint16 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]uint16, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Uint16 {
			keys[i] = uint16(key.Uint()) // #nosec G115 -- intentional truncation for best-effort conversion
		}
	}
	return keys
}

func MapKeysUint32(m interface{}) []uint32 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]uint32, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Uint32 {
			keys[i] = uint32(key.Uint()) // #nosec G115 -- intentional truncation for best-effort conversion
		}
	}
	return keys
}

func MapKeysUint64(m interface{}) []uint64 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]uint64, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Uint64 {
			keys[i] = key.Uint()
		}
	}
	return keys
}

func MapKeysFloat32(m interface{}) []float32 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]float32, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Float32 {
			keys[i] = float32(key.Float())
		}
	}
	return keys
}

func MapKeysFloat64(m interface{}) []float64 {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]float64, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.Float64 {
			keys[i] = key.Float()
		}
	}
	return keys
}

func MapKeysString(m interface{}) []string {
	// 快速路径：类型断言避免反射开销
	switch m := m.(type) {
	case map[string]int:
		if len(m) == 0 {
			return nil
		}
		keys := make([]string, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[string]string:
		if len(m) == 0 {
			return nil
		}
		keys := make([]string, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[string]interface{}:
		if len(m) == 0 {
			return nil
		}
		keys := make([]string, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[string]bool:
		if len(m) == 0 {
			return nil
		}
		keys := make([]string, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	case map[string]float64:
		if len(m) == 0 {
			return nil
		}
		keys := make([]string, len(m))
		i := 0
		for k := range m {
			keys[i] = k
			i++
		}
		return keys
	}

	// 反射路径
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]string, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		if key.Kind() == reflect.String {
			keys[i] = key.String()
		}
	}
	return keys
}

func MapKeysAny(m interface{}) []interface{} {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		return nil
	}

	if val.Len() == 0 {
		return nil
	}

	// 优化：预分配精确容量
	keys := make([]interface{}, val.Len())
	mapKeys := val.MapKeys()
	for i, key := range mapKeys {
		keys[i] = key.Interface()
	}
	return keys
}

func MapValues[K constraints.Ordered, V any](m map[K]V) []V {
	if len(m) == 0 {
		return nil
	}
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}
