package candy

import "reflect"

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