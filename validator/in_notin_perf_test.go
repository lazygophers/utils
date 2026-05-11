package validator

import (
	"reflect"
	"sort"
	"sync"
	"testing"
)

// ============== 方案1: 预转换 reflect.Value ==============
func InV1(values ...interface{}) ValidatorFunc {
	// 预先转换所有值为 reflect.Value，避免重复转换
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV1(values ...interface{}) ValidatorFunc {
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 方案2: int 类型专用快速路径 ==============
func InV2(values ...interface{}) ValidatorFunc {
	// 检测是否全是 int
	allInt := true
	intValues := make([]int, 0, len(values))

	for _, v := range values {
		if i, ok := v.(int); ok {
			intValues = append(intValues, i)
		} else {
			allInt = false
			break
		}
	}

	if allInt {
		// 使用 map 优化
		intMap := make(map[int]bool, len(intValues))
		for _, v := range intValues {
			intMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.Int {
				return intMap[int(field.Int())]
			}
			// 降级到原始方法
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return true
				}
			}
			return false
		}
	}

	// 非纯 int，使用原始方法
	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV2(values ...interface{}) ValidatorFunc {
	allInt := true
	intValues := make([]int, 0, len(values))

	for _, v := range values {
		if i, ok := v.(int); ok {
			intValues = append(intValues, i)
		} else {
			allInt = false
			break
		}
	}

	if allInt {
		intMap := make(map[int]bool, len(intValues))
		for _, v := range intValues {
			intMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.Int {
				return !intMap[int(field.Int())]
			}
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return false
				}
			}
			return true
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 方案3: string 类型专用快速路径 ==============
func InV3(values ...interface{}) ValidatorFunc {
	// 检测是否全是 string
	allString := true
	stringValues := make([]string, 0, len(values))

	for _, v := range values {
		if s, ok := v.(string); ok {
			stringValues = append(stringValues, s)
		} else {
			allString = false
			break
		}
	}

	if allString {
		// 使用 map 优化
		stringMap := make(map[string]bool, len(stringValues))
		for _, v := range stringValues {
			stringMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.String {
				return stringMap[field.String()]
			}
			// 降级
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return true
				}
			}
			return false
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV3(values ...interface{}) ValidatorFunc {
	allString := true
	stringValues := make([]string, 0, len(values))

	for _, v := range values {
		if s, ok := v.(string); ok {
			stringValues = append(stringValues, s)
		} else {
			allString = false
			break
		}
	}

	if allString {
		stringMap := make(map[string]bool, len(stringValues))
		for _, v := range stringValues {
			stringMap[v] = true
		}

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if field.Kind() == reflect.String {
				return !stringMap[field.String()]
			}
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return false
				}
			}
			return true
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, v := range values {
			if compareFields(field, reflect.ValueOf(v)) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 方案4: 统一类型用 map ==============
func InV4(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 检查类型是否统一
	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return intMap[int(field.Int())]
				}
				return false
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return stringMap[field.String()]
				}
				return false
			}
		}
	}

	// 混合类型或未处理类型，使用线性查找
	return InV1(values...)
}

func NotInV4(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return !intMap[int(field.Int())]
				}
				return true
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return !stringMap[field.String()]
				}
				return true
			}
		}
	}

	return NotInV1(values...)
}

// ============== 方案5: 少量枚举用展开 switch (硬编码 3 个) ==============
func InV5(values ...interface{}) ValidatorFunc {
	if len(values) == 3 {
		v0, v1, v2 := values[0], values[1], values[2]
		rv0, rv1, rv2 := reflect.ValueOf(v0), reflect.ValueOf(v1), reflect.ValueOf(v2)

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if compareFields(field, rv0) == 0 {
				return true
			}
			if compareFields(field, rv1) == 0 {
				return true
			}
			if compareFields(field, rv2) == 0 {
				return true
			}
			return false
		}
	}

	// 降级
	return InV1(values...)
}

func NotInV5(values ...interface{}) ValidatorFunc {
	if len(values) == 3 {
		v0, v1, v2 := values[0], values[1], values[2]
		rv0, rv1, rv2 := reflect.ValueOf(v0), reflect.ValueOf(v1), reflect.ValueOf(v2)

		return func(fl FieldLevel) bool {
			field := fl.Field()
			if compareFields(field, rv0) == 0 {
				return false
			}
			if compareFields(field, rv1) == 0 {
				return false
			}
			if compareFields(field, rv2) == 0 {
				return false
			}
			return true
		}
	}

	return NotInV1(values...)
}

// ============== 方案6: 直接 interface{} 比较（避免反射） ==============
func InV6(values ...interface{}) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 尝试直接获取 interface{} 值
		var fieldInterface interface{}
		if field.CanInterface() {
			fieldInterface = field.Interface()
		} else {
			// 无法直接获取，使用反射
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return true
				}
			}
			return false
		}

		// 直接比较
		for _, v := range values {
			if fieldInterface == v {
				return true
			}
		}
		return false
	}
}

func NotInV6(values ...interface{}) ValidatorFunc {
	return func(fl FieldLevel) bool {
		field := fl.Field()

		var fieldInterface interface{}
		if field.CanInterface() {
			fieldInterface = field.Interface()
		} else {
			for _, v := range values {
				if compareFields(field, reflect.ValueOf(v)) == 0 {
					return false
				}
			}
			return true
		}

		for _, v := range values {
			if fieldInterface == v {
				return false
			}
		}
		return true
	}
}

// ============== 方案7: sort + binary search（有序枚举） ==============
func InV7(values ...interface{}) ValidatorFunc {
	if len(values) <= 10 {
		// 少量值不排序，直接线性查找
		return InV1(values...)
	}

	// 复制并排序
	sortedValues := make([]interface{}, len(values))
	copy(sortedValues, values)

	sort.Slice(sortedValues, func(i, j int) bool {
		vi := reflect.ValueOf(sortedValues[i])
		vj := reflect.ValueOf(sortedValues[j])
		return compareFields(vi, vj) < 0
	})

	reflectValues := make([]reflect.Value, len(sortedValues))
	for i, v := range sortedValues {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 二分查找
		low, high := 0, len(reflectValues)-1
		for low <= high {
			mid := (low + high) / 2
			cmp := compareFields(field, reflectValues[mid])
			if cmp == 0 {
				return true
			} else if cmp < 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return false
	}
}

func NotInV7(values ...interface{}) ValidatorFunc {
	if len(values) <= 10 {
		return NotInV1(values...)
	}

	sortedValues := make([]interface{}, len(values))
	copy(sortedValues, values)

	sort.Slice(sortedValues, func(i, j int) bool {
		vi := reflect.ValueOf(sortedValues[i])
		vj := reflect.ValueOf(sortedValues[j])
		return compareFields(vi, vj) < 0
	})

	reflectValues := make([]reflect.Value, len(sortedValues))
	for i, v := range sortedValues {
		reflectValues[i] = reflect.ValueOf(v)
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		low, high := 0, len(reflectValues)-1
		for low <= high {
			mid := (low + high) / 2
			cmp := compareFields(field, reflectValues[mid])
			if cmp == 0 {
				return false
			} else if cmp < 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return true
	}
}

// ============== 方案8: 混合优化（智能选择策略） ==============
func InV8(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 少量枚举：直接线性
	if len(values) <= 5 {
		return InV1(values...)
	}

	// 检查是否统一类型
	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return intMap[int(field.Int())]
				}
				return false
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return stringMap[field.String()]
				}
				return false
			}
		}
	}

	// 大量混合类型：二分查找
	return InV7(values...)
}

func NotInV8(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	if len(values) <= 5 {
		return NotInV1(values...)
	}

	firstType := reflect.TypeOf(values[0])
	allSameType := true
	for _, v := range values {
		if reflect.TypeOf(v) != firstType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch firstType.Kind() {
		case reflect.Int:
			intMap := make(map[int]bool, len(values))
			for _, v := range values {
				intMap[v.(int)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return !intMap[int(field.Int())]
				}
				return true
			}
		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, v := range values {
				stringMap[v.(string)] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return !stringMap[field.String()]
				}
				return true
			}
		}
	}

	return NotInV7(values...)
}

// ============== 方案9: 使用 sync.Map 缓存验证结果 ==============
func InV9(values ...interface{}) ValidatorFunc {
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	cache := &sync.Map{}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		// 尝试从缓存获取
		fieldKey := field.Interface()
		if cached, ok := cache.Load(fieldKey); ok {
			return cached.(bool)
		}

		// 计算并缓存
		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				cache.Store(fieldKey, true)
				return true
			}
		}
		cache.Store(fieldKey, false)
		return false
	}
}

func NotInV9(values ...interface{}) ValidatorFunc {
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	cache := &sync.Map{}

	return func(fl FieldLevel) bool {
		field := fl.Field()

		fieldKey := field.Interface()
		if cached, ok := cache.Load(fieldKey); ok {
			return !cached.(bool)
		}

		for _, v := range reflectValues {
			if compareFields(field, v) == 0 {
				cache.Store(fieldKey, true)
				return false
			}
		}
		cache.Store(fieldKey, false)
		return true
	}
}

// ============== 方案10: 分组优化（按类型分组） ==============
func InV10(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 按类型分组
	typeGroups := make(map[reflect.Type][]reflect.Value)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		typeGroups[rt] = append(typeGroups[rt], rv)
	}

	// 为每种类型创建优化的查找器
	typeCheckers := make(map[reflect.Kind]func(reflect.Value) bool)

	for rt, rvs := range typeGroups {
		kind := rt.Kind()

		switch kind {
		case reflect.Int:
			intMap := make(map[int64]bool, len(rvs))
			for _, rv := range rvs {
				intMap[rv.Int()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return intMap[field.Int()]
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(rvs))
			for _, rv := range rvs {
				stringMap[rv.String()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return stringMap[field.String()]
			}

		default:
			// 其他类型使用线性查找
			typeCheckers[kind] = func(field reflect.Value) bool {
				for _, rv := range rvs {
					if compareFields(field, rv) == 0 {
						return true
					}
				}
				return false
			}
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		if checker, ok := typeCheckers[field.Kind()]; ok {
			return checker(field)
		}
		return false
	}
}

func NotInV10(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	typeGroups := make(map[reflect.Type][]reflect.Value)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		typeGroups[rt] = append(typeGroups[rt], rv)
	}

	typeCheckers := make(map[reflect.Kind]func(reflect.Value) bool)

	for rt, rvs := range typeGroups {
		kind := rt.Kind()

		switch kind {
		case reflect.Int:
			intMap := make(map[int64]bool, len(rvs))
			for _, rv := range rvs {
				intMap[rv.Int()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return !intMap[field.Int()]
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(rvs))
			for _, rv := range rvs {
				stringMap[rv.String()] = true
			}
			typeCheckers[kind] = func(field reflect.Value) bool {
				return !stringMap[field.String()]
			}

		default:
			typeCheckers[kind] = func(field reflect.Value) bool {
				for _, rv := range rvs {
					if compareFields(field, rv) == 0 {
						return false
					}
				}
				return true
			}
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		if checker, ok := typeCheckers[field.Kind()]; ok {
			return checker(field)
		}
		return true
	}
}

// ============== 方案11: 组合优化（预转换 + 类型检测 + map） ==============
func InV11(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return false }
	}

	// 预转换并分析
	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	// 检查是否统一类型
	unifiedType := reflectValues[0].Type()
	allSameType := true
	for _, rv := range reflectValues {
		if rv.Type() != unifiedType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch unifiedType.Kind() {
		case reflect.Int:
			intMap := make(map[int64]bool, len(values))
			for _, rv := range reflectValues {
				intMap[rv.Int()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return intMap[field.Int()]
				}
				return false
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, rv := range reflectValues {
				stringMap[rv.String()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return stringMap[field.String()]
				}
				return false
			}

		case reflect.Float64, reflect.Float32:
			floatMap := make(map[float64]bool, len(values))
			for _, rv := range reflectValues {
				floatMap[rv.Float()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
					return floatMap[field.Float()]
				}
				return false
			}
		}
	}

	// 混合类型，使用预转换的线性查找
	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, rv := range reflectValues {
			if compareFields(field, rv) == 0 {
				return true
			}
		}
		return false
	}
}

func NotInV11(values ...interface{}) ValidatorFunc {
	if len(values) == 0 {
		return func(fl FieldLevel) bool { return true }
	}

	reflectValues := make([]reflect.Value, len(values))
	for i, v := range values {
		reflectValues[i] = reflect.ValueOf(v)
	}

	unifiedType := reflectValues[0].Type()
	allSameType := true
	for _, rv := range reflectValues {
		if rv.Type() != unifiedType {
			allSameType = false
			break
		}
	}

	if allSameType {
		switch unifiedType.Kind() {
		case reflect.Int:
			intMap := make(map[int64]bool, len(values))
			for _, rv := range reflectValues {
				intMap[rv.Int()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Int {
					return !intMap[field.Int()]
				}
				return true
			}

		case reflect.String:
			stringMap := make(map[string]bool, len(values))
			for _, rv := range reflectValues {
				stringMap[rv.String()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.String {
					return !stringMap[field.String()]
				}
				return true
			}

		case reflect.Float64, reflect.Float32:
			floatMap := make(map[float64]bool, len(values))
			for _, rv := range reflectValues {
				floatMap[rv.Float()] = true
			}
			return func(fl FieldLevel) bool {
				field := fl.Field()
				if field.Kind() == reflect.Float32 || field.Kind() == reflect.Float64 {
					return !floatMap[field.Float()]
				}
				return true
			}
		}
	}

	return func(fl FieldLevel) bool {
		field := fl.Field()
		for _, rv := range reflectValues {
			if compareFields(field, rv) == 0 {
				return false
			}
		}
		return true
	}
}

// ============== 基准测试 ==============

// 测试数据准备
type TestStruct struct {
	Value int
}

func (t TestStruct) Field() reflect.Value {
	return reflect.ValueOf(t.Value)
}

func (t TestStruct) FieldName() string {
	return "Field"
}

func (t TestStruct) StructFieldName() string {
	return "TestStruct.Field"
}

func (t TestStruct) Param() string {
	return ""
}

func (t TestStruct) GetTag(key string) string {
	return ""
}

func (t TestStruct) Top() reflect.Value {
	return reflect.ValueOf(t)
}

func (t TestStruct) Parent() reflect.Value {
	return reflect.Value{}
}

func (t TestStruct) GetFieldByName(name string) reflect.Value {
	return reflect.Value{}
}

// ============== In 函数基准测试 ==============

// 少量 int 枚举（3个）
func BenchmarkIn_Original_3Int(b *testing.B) {
	validator := In(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V1_3Int(b *testing.B) {
	validator := InV1(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V2_3Int(b *testing.B) {
	validator := InV2(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V4_3Int(b *testing.B) {
	validator := InV4(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V5_3Int(b *testing.B) {
	validator := InV5(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V6_3Int(b *testing.B) {
	validator := InV6(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V8_3Int(b *testing.B) {
	validator := InV8(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V10_3Int(b *testing.B) {
	validator := InV10(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V11_3Int(b *testing.B) {
	validator := InV11(1, 2, 3)
	fl := TestStruct{Value: 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 中等 int 枚举（15个）
func BenchmarkIn_Original_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := In(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V1_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV1(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V2_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV2(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V4_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV4(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V7_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV7(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V8_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV8(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V10_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV10(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V11_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := InV11(values...)
	fl := TestStruct{Value: 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 大量 int 枚举（100个）
func BenchmarkIn_Original_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := In(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V1_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV1(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V2_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV2(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V4_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV4(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V7_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV7(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V8_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV8(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V10_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV10(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkIn_V11_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := InV11(values...)
	fl := TestStruct{Value: 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// String 类型测试（中等枚举）
func BenchmarkIn_Original_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := In(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := struct {
			Field string
		}{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V1_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV1(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := struct {
			Field string
		}{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V3_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV3(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := struct {
			Field string
		}{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V4_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV4(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := struct {
			Field string
		}{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

func BenchmarkIn_V11_15String(b *testing.B) {
	values := []interface{}{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya", "quince"}
	validator := InV11(values...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFl := struct {
			Field string
		}{Field: "mango"}
		rv := reflect.ValueOf(stringFl.Field)
		flWrapper := &fieldLevel{field: rv}
		validator(flWrapper)
	}
}

// ============== NotIn 函数基准测试 ==============

// 少量 int 枚举
func BenchmarkNotIn_Original_3Int(b *testing.B) {
	validator := NotIn(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V1_3Int(b *testing.B) {
	validator := NotInV1(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V2_3Int(b *testing.B) {
	validator := NotInV2(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V11_3Int(b *testing.B) {
	validator := NotInV11(1, 2, 3)
	fl := TestStruct{Value: 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 中等 int 枚举
func BenchmarkNotIn_Original_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotIn(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V1_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotInV1(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V2_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotInV2(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V11_15Int(b *testing.B) {
	values := make([]interface{}, 15)
	for i := 0; i < 15; i++ {
		values[i] = i
	}
	validator := NotInV11(values...)
	fl := TestStruct{Value: 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

// 大量 int 枚举
func BenchmarkNotIn_Original_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotIn(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V1_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotInV1(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V2_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotInV2(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkNotIn_V11_100Int(b *testing.B) {
	values := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		values[i] = i
	}
	validator := NotInV11(values...)
	fl := TestStruct{Value: 150}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}
