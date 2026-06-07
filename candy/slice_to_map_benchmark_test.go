package candy

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"golang.org/x/exp/constraints"
)

func BenchmarkSliceField2MapString_Simple(b *testing.B) {
	people := make([]TestPerson, 1000)
	for i := 0; i < 1000; i++ {
		people[i] = TestPerson{ID: i, Name: fmt.Sprintf("user%d", i), Age: i}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapString(people, "Name")
	}
}

func BenchmarkSliceField2MapInt_Simple(b *testing.B) {
	people := make([]TestPerson, 1000)
	for i := 0; i < 1000; i++ {
		people[i] = TestPerson{ID: i, Name: fmt.Sprintf("user%d", i), Age: i}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapInt(people, "ID")
	}
}

// ==================== 测试数据结构 ====================

type Person struct {
	ID        int
	Name      string
	Age       int8
	Score     int16
	Balance   int32
	Amount    int64
	UID       uint
	ByteValue uint8
	ShortVal  uint16
	Value32   uint32
	Value64   uint64
	Rate      float32
	Ratio     float64
	Active    bool
}

var testPeople = make([]Person, 1000)

func init() {
	for i := 0; i < 1000; i++ {
		testPeople[i] = Person{
			ID:        i,
			Name:      fmt.Sprintf("user%d", i),
			Age:       int8(i % 100),
			Score:     int16(i * 10),
			Balance:   int32(i * 100),
			Amount:    int64(i * 1000),
			UID:       uint(i),
			ByteValue: uint8(i % 256),
			ShortVal:  uint16(i * 5),
			Value32:   uint32(i * 100),
			Value64:   uint64(i * 1000),
			Rate:      float32(i) * 1.5,
			Ratio:     float64(i) * 2.5,
			Active:    i%2 == 0,
		}
	}
}

// ==================== 基准测试：原始实现 ====================

func BenchmarkSliceField2MapString_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapString(testPeople, "Name")
	}
}

func BenchmarkSliceField2MapInt_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapInt(testPeople, "ID")
	}
}

func BenchmarkSliceField2MapInt8_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapInt8(testPeople, "Age")
	}
}

func BenchmarkSliceField2MapInt16_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapInt16(testPeople, "Score")
	}
}

func BenchmarkSliceField2MapInt32_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapInt32(testPeople, "Balance")
	}
}

func BenchmarkSliceField2MapInt64_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapInt64(testPeople, "Amount")
	}
}

func BenchmarkSliceField2MapUint_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapUint(testPeople, "UID")
	}
}

func BenchmarkSliceField2MapUint8_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapUint8(testPeople, "ByteValue")
	}
}

func BenchmarkSliceField2MapUint16_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapUint16(testPeople, "ShortVal")
	}
}

func BenchmarkSliceField2MapUint32_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapUint32(testPeople, "Value32")
	}
}

func BenchmarkSliceField2MapUint64_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapUint64(testPeople, "Value64")
	}
}

func BenchmarkSliceField2MapFloat32_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapFloat32(testPeople, "Rate")
	}
}

func BenchmarkSliceField2MapFloat64_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapFloat64(testPeople, "Ratio")
	}
}

func BenchmarkSliceField2MapBool_Original(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = SliceField2MapBool(testPeople, "Active")
	}
}

// ==================== 优化方案 1：反射字段索引缓存 ====================

// fieldIndexCacheT 字段索引缓存类型
type fieldIndexCacheT struct {
	sync.RWMutex
	cache map[reflect.Type]map[string][]int
}

// fieldIndexCache 字段索引缓存
var fieldIndexCache fieldIndexCacheT

func init() {
	fieldIndexCache.cache = make(map[reflect.Type]map[string][]int)
}

// getFieldIndexCached 获取字段索引，使用缓存
func getFieldIndexCached(elemType reflect.Type, fieldName string) ([]int, reflect.Type, bool) {
	// 读缓存
	fieldIndexCache.RLock()
	if fields, ok := fieldIndexCache.cache[elemType]; ok {
		if index, exists := fields[fieldName]; exists {
			fieldIndexCache.RUnlock()
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
	fieldIndexCache.RUnlock()

	// 缓存未命中，获取并缓存
	fieldIndexCache.Lock()
	defer fieldIndexCache.Unlock()

	// 双重检查
	if fields, ok := fieldIndexCache.cache[elemType]; ok {
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
	if fieldIndexCache.cache[elemType] == nil {
		fieldIndexCache.cache[elemType] = make(map[string][]int)
	}

	fieldIndex := field.Index
	fieldIndexCache.cache[elemType][fieldName] = fieldIndex

	return fieldIndex, field.Type, true
}

func BenchmarkSliceField2MapString_Cached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_Cached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapInt32_Cached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Balance")
		if !ok || fieldType.Kind() != reflect.Int32 {
			continue
		}

		ret := make(map[int32]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int32(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_Cached(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 2：直接访问（已知类型） ====================

func BenchmarkSliceField2MapString_DirectAccess(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ret := make(map[string]bool, len(testPeople))
		for _, p := range testPeople {
			ret[p.Name] = true
		}
	}
}

func BenchmarkSliceField2MapInt_DirectAccess(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ret := make(map[int]bool, len(testPeople))
		for _, p := range testPeople {
			ret[p.ID] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_DirectAccess(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ret := make(map[float64]bool, len(testPeople))
		for _, p := range testPeople {
			ret[p.Ratio] = true
		}
	}
}

// ==================== 优化方案 3：反射值复用 ====================

func BenchmarkSliceField2MapString_ValueReuse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople))
		var v reflect.Value
		for _, item := range testPeople {
			v = reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_ValueReuse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople))
		var v reflect.Value
		for _, item := range testPeople {
			v = reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_ValueReuse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople))
		var v reflect.Value
		for _, item := range testPeople {
			v = reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 4：批量处理（减少反射调用） ====================

func BenchmarkSliceField2MapString_Batch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople))
		sliceValue := reflect.ValueOf(testPeople)

		for j := 0; j < sliceValue.Len(); j++ {
			elem := sliceValue.Index(j)
			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			fieldValue := elem.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

// ==================== 优化方案 5：切片级别反射（减少反射调用） ====================

func BenchmarkSliceField2MapString_SliceReflect(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople))
		sliceValue := reflect.ValueOf(testPeople)

		for j := 0; j < sliceValue.Len(); j++ {
			elem := sliceValue.Index(j)
			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			fieldValue := elem.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_SliceReflect(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople))
		sliceValue := reflect.ValueOf(testPeople)

		for j := 0; j < sliceValue.Len(); j++ {
			elem := sliceValue.Index(j)
			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			fieldValue := elem.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_SliceReflect(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople))
		sliceValue := reflect.ValueOf(testPeople)

		for j := 0; j < sliceValue.Len(); j++ {
			elem := sliceValue.Index(j)
			for elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			fieldValue := elem.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 6：避免指针解包（优化分支预测） ====================

func BenchmarkSliceField2MapString_NoUnwrap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			// 假设非指针类型，直接访问
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_NoUnwrap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_NoUnwrap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 7：批量分配（减少 map 扩容） ====================

func BenchmarkSliceField2MapString_BatchAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		// 稍微多分配一些空间以减少扩容
		ret := make(map[string]bool, len(testPeople)*3/2)
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_BatchAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople)*3/2)
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_BatchAlloc(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople)*3/2)
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 8：完整优化实现 ====================

func BenchmarkSliceField2MapString_FullOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_FullOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_FullOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople))
		for _, item := range testPeople {
			v := reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 9：组合优化（缓存 + 值复用 + 批量分配） ====================

func BenchmarkSliceField2MapString_Combo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Name")
		if !ok || fieldType.Kind() != reflect.String {
			continue
		}

		ret := make(map[string]bool, len(testPeople)*3/2)
		var v reflect.Value
		for _, item := range testPeople {
			v = reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.String()] = true
		}
	}
}

func BenchmarkSliceField2MapInt_Combo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "ID")
		if !ok || fieldType.Kind() != reflect.Int {
			continue
		}

		ret := make(map[int]bool, len(testPeople)*3/2)
		var v reflect.Value
		for _, item := range testPeople {
			v = reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[int(fieldValue.Int())] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_Combo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(testPeople) == 0 {
			continue
		}

		elemType := reflect.TypeOf(testPeople[0])
		fieldIndex, fieldType, ok := getFieldIndexCached(elemType, "Ratio")
		if !ok || fieldType.Kind() != reflect.Float64 {
			continue
		}

		ret := make(map[float64]bool, len(testPeople)*3/2)
		var v reflect.Value
		for _, item := range testPeople {
			v = reflect.ValueOf(item)
			for v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			fieldValue := v.FieldByIndex(fieldIndex)
			ret[fieldValue.Float()] = true
		}
	}
}

// ==================== 优化方案 10：类型特化（已知类型时的最佳性能） ====================

func BenchmarkSliceField2MapString_Specialized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ret := make(map[string]bool, len(testPeople)*3/2)
		for _, p := range testPeople {
			ret[p.Name] = true
		}
	}
}

func BenchmarkSliceField2MapInt_Specialized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ret := make(map[int]bool, len(testPeople)*3/2)
		for _, p := range testPeople {
			ret[p.ID] = true
		}
	}
}

func BenchmarkSliceField2MapFloat64_Specialized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ret := make(map[float64]bool, len(testPeople)*3/2)
		for _, p := range testPeople {
			ret[p.Ratio] = true
		}
	}
}

// Benchmark Slice2Map 函数的各种实现方案

// 方案1: 当前实现（使用 range）
func BenchmarkSlice2Map_Current_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

// 方案2: 使用索引循环
func Slice2MapIndex[M constraints.Ordered](list []M) map[M]bool {
	if len(list) == 0 {
		return make(map[M]bool)
	}

	result := make(map[M]bool, len(list))
	for i := 0; i < len(list); i++ {
		result[list[i]] = true
	}
	return result
}

func BenchmarkSlice2Map_Index_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

// 方案3: 不预分配容量
func Slice2MapNoPrealloc[M constraints.Ordered](list []M) map[M]bool {
	if len(list) == 0 {
		return make(map[M]bool)
	}

	result := make(map[M]bool)
	for _, item := range list {
		result[item] = true
	}
	return result
}

func BenchmarkSlice2Map_NoPrealloc_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapNoPrealloc(data)
	}
}

// 方案4: 使用 struct{} 作为值（节省内存）
func Slice2MapStruct[M constraints.Ordered](list []M) map[M]struct{} {
	if len(list) == 0 {
		return make(map[M]struct{})
	}

	result := make(map[M]struct{}, len(list))
	for i := 0; i < len(list); i++ {
		result[list[i]] = struct{}{}
	}
	return result
}

func BenchmarkSlice2Map_Struct_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

// 方案5: 检查是否已存在再设置
func Slice2MapCheck[M constraints.Ordered](list []M) map[M]bool {
	if len(list) == 0 {
		return make(map[M]bool)
	}

	result := make(map[M]bool, len(list))
	for i := 0; i < len(list); i++ {
		if !result[list[i]] {
			result[list[i]] = true
		}
	}
	return result
}

func BenchmarkSlice2Map_Check_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}

// 中等数据集测试
func BenchmarkSlice2Map_Current_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

func BenchmarkSlice2Map_Index_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

func BenchmarkSlice2Map_NoPrealloc_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapNoPrealloc(data)
	}
}

func BenchmarkSlice2Map_Struct_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

func BenchmarkSlice2Map_Check_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}

// 大数据集测试
func BenchmarkSlice2Map_Current_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

func BenchmarkSlice2Map_Index_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

func BenchmarkSlice2Map_NoPrealloc_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapNoPrealloc(data)
	}
}

func BenchmarkSlice2Map_Struct_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

func BenchmarkSlice2Map_Check_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}

// 重复数据测试
func BenchmarkSlice2Map_Current_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10 // 创建重复数据
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2Map(data)
	}
}

func BenchmarkSlice2Map_Index_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapIndex(data)
	}
}

func BenchmarkSlice2Map_Struct_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapStruct(data)
	}
}

func BenchmarkSlice2Map_Check_Duplicates(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice2MapCheck(data)
	}
}
