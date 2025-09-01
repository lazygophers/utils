package anyx

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckValueType(t *testing.T) {
	type args struct {
		input  interface{}
		expect ValueType
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "string",
			args: args{
				input:  "hello",
				expect: ValueString,
			},
		},
		{
			name: "int",
			args: args{
				input:  123,
				expect: ValueNumber,
			},
		},
		{
			name: "float64",
			args: args{
				input:  123.456,
				expect: ValueNumber,
			},
		},
		{
			name: "bool",
			args: args{
				input:  true,
				expect: ValueBool,
			},
		},
		{
			name: "nil",
			args: args{
				input:  nil,
				expect: ValueUnknown,
			},
		},
		{
			name: "map",
			args: args{
				input:  map[string]interface{}{},
				expect: ValueUnknown,
			},
		},
		{
			name: "slice",
			args: args{
				input:  []interface{}{},
				expect: ValueUnknown,
			},
		},
		{
			name: "array",
			args: args{
				input:  [1]interface{}{},
				expect: ValueUnknown,
			},
		},
		{
			name: "struct",
			args: args{
				input:  struct{}{},
				expect: ValueUnknown,
			},
		},
		{
			name: "func",
			args: args{
				input:  func() {},
				expect: ValueUnknown,
			},
		},
		{
			name: "chan",
			args: args{
				input:  make(chan int),
				expect: ValueUnknown,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckValueType(tt.args.input); got != tt.args.expect {
				t.Errorf("CheckValueType() = %v, want %v", got, tt.args.expect)
			}
		})
	}
}

// Test MapKeys functions
func TestMapKeysString(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	keys := MapKeysString(m)

	sort.Strings(keys)
	expected := []string{"a", "b"}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysString() = %v, want %v", keys, expected)
	}
}

func TestMapKeysUint32(t *testing.T) {
	m := map[uint32]string{1: "a", 2: "b"}
	keys := MapKeysUint32(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []uint32{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysUint32() = %v, want %v", keys, expected)
	}
}

func TestMapKeysUint64(t *testing.T) {
	m := map[uint64]string{1: "a", 2: "b"}
	keys := MapKeysUint64(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []uint64{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysUint64() = %v, want %v", keys, expected)
	}
}

func TestMapKeysInt32(t *testing.T) {
	m := map[int32]string{1: "a", 2: "b"}
	keys := MapKeysInt32(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []int32{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysInt32() = %v, want %v", keys, expected)
	}
}

func TestMapKeysInt64(t *testing.T) {
	m := map[int64]string{1: "a", 2: "b"}
	keys := MapKeysInt64(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []int64{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysInt64() = %v, want %v", keys, expected)
	}
}

func TestMapKeysInt(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	keys := MapKeysInt(m)

	sort.Ints(keys)
	expected := []int{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysInt() = %v, want %v", keys, expected)
	}
}

func TestMapKeysInt8(t *testing.T) {
	m := map[int8]string{1: "a", 2: "b"}
	keys := MapKeysInt8(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []int8{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysInt8() = %v, want %v", keys, expected)
	}
}

func TestMapKeysInt16(t *testing.T) {
	m := map[int16]string{1: "a", 2: "b"}
	keys := MapKeysInt16(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []int16{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysInt16() = %v, want %v", keys, expected)
	}
}

func TestMapKeysUint(t *testing.T) {
	m := map[uint]string{1: "a", 2: "b"}
	keys := MapKeysUint(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []uint{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysUint() = %v, want %v", keys, expected)
	}
}

func TestMapKeysUint8(t *testing.T) {
	m := map[uint8]string{1: "a", 2: "b"}
	keys := MapKeysUint8(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []uint8{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysUint8() = %v, want %v", keys, expected)
	}
}

func TestMapKeysUint16(t *testing.T) {
	m := map[uint16]string{1: "a", 2: "b"}
	keys := MapKeysUint16(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []uint16{1, 2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysUint16() = %v, want %v", keys, expected)
	}
}

func TestMapKeysFloat32(t *testing.T) {
	m := map[float32]string{1.1: "a", 2.2: "b"}
	keys := MapKeysFloat32(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []float32{1.1, 2.2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysFloat32() = %v, want %v", keys, expected)
	}
}

func TestMapKeysFloat64(t *testing.T) {
	m := map[float64]string{1.1: "a", 2.2: "b"}
	keys := MapKeysFloat64(m)

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	expected := []float64{1.1, 2.2}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeysFloat64() = %v, want %v", keys, expected)
	}
}

func TestMapKeysInterface(t *testing.T) {
	m := map[interface{}]string{"a": "1", 1: "2"}
	keys := MapKeysInterface(m)

	sort.Slice(keys, func(i, j int) bool {
		return reflect.ValueOf(keys[i]).String() < reflect.ValueOf(keys[j]).String()
	})
	if len(keys) != 2 {
		t.Errorf("MapKeysInterface() length = %d, want 2", len(keys))
	}
}

func TestMapKeysAny(t *testing.T) {
	m := map[interface{}]string{"a": "1", 1: "2"}
	keys := MapKeysAny(m)

	sort.Slice(keys, func(i, j int) bool {
		return reflect.ValueOf(keys[i]).String() < reflect.ValueOf(keys[j]).String()
	})
	if len(keys) != 2 {
		t.Errorf("MapKeysAny() length = %d, want 2", len(keys))
	}
}

func TestMapKeysNumber(t *testing.T) {
	m := map[float64]string{1.1: "a", 2.2: "b"}
	keys := MapKeysNumber(m)

	sort.Slice(keys, func(i, j int) bool { return reflect.ValueOf(keys[i]).String() < reflect.ValueOf(keys[j]).String() })

	// Convert []interface{} to []float64
	result := make([]float64, len(keys))
	for i, v := range keys {
		result[i] = v.(float64)
	}

	expected := []float64{1.1, 2.2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MapKeysNumber() = %v, want %v", result, expected)
	}
}

// Test MapValues functions
func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	values := MapValues(m)
	sort.Ints(values)
	expected := []int{1, 2}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("MapValues() = %v, want %v", values, expected)
	}
}

func TestMapValuesAny(t *testing.T) {
	m := map[string]interface{}{"a": 1, "b": "2"}
	values := MapValuesAny(m)
	if len(values) != 2 {
		t.Errorf("MapValuesAny() length = %d, want 2", len(values))
	}
}

func TestMapValuesString(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	values := MapValuesString(m)
	sort.Strings(values)
	expected := []string{"1", "2"}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("MapValuesString() = %v, want %v", values, expected)
	}
}

func TestMapValuesInt(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	values := MapValuesInt(m)
	sort.Ints(values)
	expected := []int{1, 2}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("MapValuesInt() = %v, want %v", values, expected)
	}
}

func TestMapValuesFloat64(t *testing.T) {
	m := map[string]float64{"a": 1.1, "b": 2.2}
	values := MapValuesFloat64(m)
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	expected := []float64{1.1, 2.2}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("MapValuesFloat64() = %v, want %v", values, expected)
	}
}

// Test error cases
func TestMapKeysStringError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysString() expected panic for non-map input")
		}
	}()
	MapKeysString("not a map")
}

func TestMapKeysUint32Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysUint32() expected panic for non-map input")
		}
	}()
	MapKeysUint32("not a map")
}

func TestMapKeysUint64Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysUint64() expected panic for non-map input")
		}
	}()
	MapKeysUint64("not a map")
}

func TestMapKeysInt32Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysInt32() expected panic for non-map input")
		}
	}()
	MapKeysInt32("not a map")
}

func TestMapKeysInt64Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysInt64() expected panic for non-map input")
		}
	}()
	MapKeysInt64("not a map")
}

func TestMapKeysIntError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysInt() expected panic for non-map input")
		}
	}()
	MapKeysInt("not a map")
}

func TestMapKeysInt8Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysInt8() expected panic for non-map input")
		}
	}()
	MapKeysInt8("not a map")
}

func TestMapKeysInt16Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysInt16() expected panic for non-map input")
		}
	}()
	MapKeysInt16("not a map")
}

func TestMapKeysUintError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysUint() expected panic for non-map input")
		}
	}()
	MapKeysUint("not a map")
}

func TestMapKeysUint8Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysUint8() expected panic for non-map input")
		}
	}()
	MapKeysUint8("not a map")
}

func TestMapKeysUint16Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysUint16() expected panic for non-map input")
		}
	}()
	MapKeysUint16("not a map")
}

func TestMapKeysFloat32Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysFloat32() expected panic for non-map input")
		}
	}()
	MapKeysFloat32("not a map")
}

func TestMapKeysFloat64Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysFloat64() expected panic for non-map input")
		}
	}()
	MapKeysFloat64("not a map")
}

func TestMapKeysInterfaceError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysInterface() expected panic for non-map input")
		}
	}()
	MapKeysInterface("not a map")
}

func TestMapKeysAnyError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysAny() expected panic for non-map input")
		}
	}()
	MapKeysAny("not a map")
}

func TestMapKeysNumberError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapKeysNumber() expected panic for non-map input")
		}
	}()
	MapKeysNumber("not a map")
}

func TestMapValuesAnyError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapValuesAny() expected panic for non-map input")
		}
	}()
	MapValuesAny("not a map")
}

func TestMapValuesStringError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapValuesString() expected panic for non-map input")
		}
	}()
	MapValuesString("not a map")
}

func TestMapValuesIntError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapValuesInt() expected panic for non-map input")
		}
	}()
	MapValuesInt("not a map")
}

func TestMapValuesFloat64Error(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MapValuesFloat64() expected panic for non-map input")
		}
	}()
	MapValuesFloat64("not a map")
}

// Test nil cases
func TestMapKeysInt32Nil(t *testing.T) {
	var m map[int32]string
	keys := MapKeysInt32(m)
	if len(keys) != 0 {
		t.Error("MapKeysInt32() with nil map should return empty slice")
	}
}

func TestMapKeysInt64Nil(t *testing.T) {
	var m map[int64]string
	keys := MapKeysInt64(m)
	if len(keys) != 0 {
		t.Error("MapKeysInt64() with nil map should return empty slice")
	}
}

// Test nil cases
func TestMapKeysIntNil(t *testing.T) {
	var m map[int]string
	keys := MapKeysInt(m)
	if len(keys) != 0 {
		t.Error("MapKeysInt() with nil map should return empty slice")
	}
}

func TestMapKeysInt8Nil(t *testing.T) {
	var m map[int8]string
	keys := MapKeysInt8(m)
	if len(keys) != 0 {
		t.Error("MapKeysInt8() with nil map should return empty slice")
	}
}

func TestMapKeysInt16Nil(t *testing.T) {
	var m map[int16]string
	keys := MapKeysInt16(m)
	if len(keys) != 0 {
		t.Error("MapKeysInt16() with nil map should return empty slice")
	}
}

func TestMapKeysUintNil(t *testing.T) {
	var m map[uint]string
	keys := MapKeysUint(m)
	if len(keys) != 0 {
		t.Error("MapKeysUint() with nil map should return empty slice")
	}
}

func TestMapKeysUint8Nil(t *testing.T) {
	var m map[uint8]string
	keys := MapKeysUint8(m)
	if len(keys) != 0 {
		t.Error("MapKeysUint8() with nil map should return empty slice")
	}
}

func TestMapKeysUint16Nil(t *testing.T) {
	var m map[uint16]string
	keys := MapKeysUint16(m)
	if len(keys) != 0 {
		t.Error("MapKeysUint16() with nil map should return empty slice")
	}
}

func TestMapKeysFloat32Nil(t *testing.T) {
	var m map[float32]string
	keys := MapKeysFloat32(m)
	if len(keys) != 0 {
		t.Error("MapKeysFloat32() with nil map should return empty slice")
	}
}

func TestMapKeysFloat64Nil(t *testing.T) {
	var m map[float64]string
	keys := MapKeysFloat64(m)
	if len(keys) != 0 {
		t.Error("MapKeysFloat64() with nil map should return empty slice")
	}
}

func TestMapKeysInterfaceNil(t *testing.T) {
	var m map[interface{}]string
	keys := MapKeysInterface(m)
	if len(keys) != 0 {
		t.Error("MapKeysInterface() with nil map should return empty slice")
	}
}

func TestMapKeysAnyNil(t *testing.T) {
	var m map[interface{}]string
	keys := MapKeysAny(m)
	if len(keys) != 0 {
		t.Error("MapKeysAny() with nil map should return empty slice")
	}
}

func TestMapKeysNumberNil(t *testing.T) {
	var m map[float64]string
	keys := MapKeysNumber(m)
	if len(keys) != 0 {
		t.Error("MapKeysNumber() with nil map should return empty slice")
	}
}

func TestMapValuesAnyNil(t *testing.T) {
	var m map[string]interface{}
	values := MapValuesAny(m)
	if len(values) != 0 {
		t.Error("MapValuesAny() with nil map should return empty slice")
	}
}

func TestMapValuesStringNil(t *testing.T) {
	var m map[string]string
	values := MapValuesString(m)
	if len(values) != 0 {
		t.Error("MapValuesString() with nil map should return empty slice")
	}
}

func TestMapValuesIntNil(t *testing.T) {
	var m map[string]int
	values := MapValuesInt(m)
	if len(values) != 0 {
		t.Error("MapValuesInt() with nil map should return empty slice")
	}
}

func TestMapValuesFloat64Nil(t *testing.T) {
	var m map[string]float64
	values := MapValuesFloat64(m)
	if len(values) != 0 {
		t.Error("MapValuesFloat64() with nil map should return empty slice")
	}
}

// Test ToMap functions
func TestToMapStringArrayString(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[string][]string
		panic bool
	}{
		{
			name: "正常情况",
			give: map[string]interface{}{
				"a": []string{"1", "2"},
				"b": []string{"3"},
			},
			want: map[string][]string{
				"a": {"1", "2"},
				"b": {"3"},
			},
		},
		{
			name: "空map",
			give: map[string]interface{}{},
			want: map[string][]string{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name: "类型转换 - string slice",
			give: map[string]interface{}{
				"a": []interface{}{"1", "2"},
			},
			want: map[string][]string{
				"a": {"1", "2"},
			},
		},
		{
			name: "类型转换 - int slice",
			give: map[string]interface{}{
				"a": []interface{}{1, 2},
			},
			want: map[string][]string{
				"a": {"1", "2"},
			},
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
		{
			name: "panic - 值类型不匹配",
			give: map[string]interface{}{
				"a": "not a slice",
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMapStringArrayString() expected panic")
					}
				}()
			}
			got := ToMapStringArrayString(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStringArrayString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test MergeMap
func TestMergeMap(t *testing.T) {
	tests := []struct {
		name   string
		source map[string]int
		target map[string]int
		want   map[string]int
	}{
		{
			name:   "正常合并",
			source: map[string]int{"a": 1, "b": 2},
			target: map[string]int{"b": 3, "c": 4},
			want:   map[string]int{"a": 1, "b": 3, "c": 4},
		},
		{
			name:   "source为空",
			source: map[string]int{},
			target: map[string]int{"a": 1},
			want:   map[string]int{"a": 1},
		},
		{
			name:   "target为空",
			source: map[string]int{"a": 1},
			target: map[string]int{},
			want:   map[string]int{"a": 1},
		},
		{
			name:   "都为空",
			source: map[string]int{},
			target: map[string]int{},
			want:   map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试 string:int 类型
			got := MergeMap(tt.source, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Test Slice2Map
func TestSlice2Map(t *testing.T) {
	tests := []struct {
		name string
		list []string
		want map[string]bool
	}{
		{
			name: "正常列表",
			list: []string{"a", "b", "c"},
			want: map[string]bool{"a": true, "b": true, "c": true},
		},
		{
			name: "空列表",
			list: []string{},
			want: map[string]bool{},
		},
		{
			name: "nil列表",
			list: nil,
			want: map[string]bool{},
		},
		{
			name: "有重复元素",
			list: []string{"a", "b", "a"},
			want: map[string]bool{"a": true, "b": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试 string 类型
			if tt.list == nil {
				var list []string
				got1 := Slice2Map(list)
				assert.Equal(t, map[string]bool{}, got1)
			} else {
				got1 := Slice2Map(tt.list)
				assert.Equal(t, tt.want, got1)
			}
		})
	}
}

// Test KeyBy
type TestUser struct {
	ID   int
	Name string
	Age  int
}

func TestKeyBy(t *testing.T) {
	users := []*TestUser{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
		{ID: 3, Name: "Alice", Age: 25},
	}

	tests := []struct {
		name      string
		list      interface{}
		fieldName string
		want      interface{}
		panic     bool
	}{
		{
			name:      "按ID字段分组",
			list:      users,
			fieldName: "ID",
			want: map[int]*TestUser{
				1: {ID: 1, Name: "Alice", Age: 25},
				2: {ID: 2, Name: "Bob", Age: 30},
				3: {ID: 3, Name: "Alice", Age: 25},
			},
		},
		{
			name:      "按Name字段分组",
			list:      users,
			fieldName: "Name",
			want: map[string]*TestUser{
				"Alice": {ID: 3, Name: "Alice", Age: 25}, // 后面的覆盖前面的
				"Bob":   {ID: 2, Name: "Bob", Age: 30},
			},
		},
		{
			name:      "空列表",
			list:      []*TestUser{},
			fieldName: "ID",
			want:      map[int]*TestUser{},
		},
		{
			name:      "nil列表",
			list:      ([]*TestUser)(nil),
			fieldName: "ID",
			panic:     true,
		},
		{
			name:      "非slice类型",
			list:      "not a slice",
			fieldName: "ID",
			panic:     true,
		},
		{
			name:      "字段不存在",
			list:      users,
			fieldName: "NonExistent",
			panic:     true,
		},
		{
			name:      "非struct元素",
			list:      []int{1, 2, 3},
			fieldName: "ID",
			panic:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() {
					KeyBy(tt.list, tt.fieldName)
				})
			} else {
				got := KeyBy(tt.list, tt.fieldName)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Test KeyByUint64
func TestKeyByUint64(t *testing.T) {
	type User struct {
		ID   uint64
		Name string
	}

	users := []*User{
		{ID: 100, Name: "Alice"},
		{ID: 200, Name: "Bob"},
		{ID: 300, Name: "Charlie"},
	}

	tests := []struct {
		name      string
		list      []*User
		fieldName string
		want      map[uint64]*User
		panic     bool
	}{
		{
			name:      "正常分组",
			list:      users,
			fieldName: "ID",
			want: map[uint64]*User{
				100: {ID: 100, Name: "Alice"},
				200: {ID: 200, Name: "Bob"},
				300: {ID: 300, Name: "Charlie"},
			},
		},
		{
			name:      "空列表",
			list:      []*User{},
			fieldName: "ID",
			want:      map[uint64]*User{},
		},
		{
			name:      "nil列表",
			list:      nil,
			fieldName: "ID",
			want:      map[uint64]*User{},
		},
		{
			name:      "字段不存在",
			list:      users,
			fieldName: "NonExistent",
			panic:     true,
		},
		{
			name:      "字段类型不是uint64",
			list:      []*User{{ID: 100, Name: "Alice"}},
			fieldName: "Name",
			panic:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				assert.Panics(t, func() {
					KeyByUint64(tt.list, tt.fieldName)
				})
			} else {
				got := KeyByUint64(tt.list, tt.fieldName)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Test KeyByInt64
func TestKeyByInt64(t *testing.T) {
	type User struct {
		ID   int64
		Name string
	}

	users := []*User{
		{ID: -100, Name: "Alice"},
		{ID: 0, Name: "Bob"},
		{ID: 100, Name: "Charlie"},
	}

	tests := []struct {
		name      string
		list      []*User
		fieldName string
		want      map[int64]*User
		panic     bool
	}{
		{
			name:      "正常分组",
			list:      users,
			fieldName: "ID",
			want: map[int64]*User{
				-100: {ID: -100, Name: "Alice"},
				0:    {ID: 0, Name: "Bob"},
				100:  {ID: 100, Name: "Charlie"},
			},
		},
		{
			name:      "空列表",
			list:      []*User{},
			fieldName: "ID",
			want:      map[int64]*User{},
		},
		{
			name:      "nil列表",
			list:      nil,
			fieldName: "ID",
			want:      map[int64]*User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeyByInt64(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Test KeyByString
func TestKeyByString(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	users := []*User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	tests := []struct {
		name      string
		list      []*User
		fieldName string
		want      map[string]*User
		panic     bool
	}{
		{
			name:      "正常分组",
			list:      users,
			fieldName: "Name",
			want: map[string]*User{
				"Alice":   {Name: "Alice", Age: 25},
				"Bob":     {Name: "Bob", Age: 30},
				"Charlie": {Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "空字符串",
			list:      []*User{{Name: "", Age: 25}},
			fieldName: "Name",
			want:      map[string]*User{"": {Name: "", Age: 25}},
		},
		{
			name:      "nil列表",
			list:      nil,
			fieldName: "Name",
			want:      map[string]*User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeyByString(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Test KeyByInt32
func TestKeyByInt32(t *testing.T) {
	type User struct {
		ID   int32
		Name string
	}

	users := []*User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	tests := []struct {
		name      string
		list      []*User
		fieldName string
		want      map[int32]*User
		panic     bool
	}{
		{
			name:      "正常分组",
			list:      users,
			fieldName: "ID",
			want: map[int32]*User{
				1: {ID: 1, Name: "Alice"},
				2: {ID: 2, Name: "Bob"},
				3: {ID: 3, Name: "Charlie"},
			},
		},
		{
			name:      "负数",
			list:      []*User{{ID: -1, Name: "Alice"}},
			fieldName: "ID",
			want:      map[int32]*User{-1: {ID: -1, Name: "Alice"}},
		},
		{
			name:      "nil列表",
			list:      nil,
			fieldName: "ID",
			want:      map[int32]*User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeyByInt32(tt.list, tt.fieldName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToMapInt32String(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[int32]string
		panic bool
	}{
		{
			name: "正常情况",
			give: map[interface{}]interface{}{
				1: "a",
				2: "b",
			},
			want: map[int32]string{
				1: "a",
				2: "b",
			},
		},
		{
			name: "类型转换 - int32",
			give: map[interface{}]interface{}{
				int32(1): "a",
				int32(2): "b",
			},
			want: map[int32]string{
				1: "a",
				2: "b",
			},
		},
		{
			name: "类型转换 - int64",
			give: map[interface{}]interface{}{
				int64(1): "a",
				int64(2): "b",
			},
			want: map[int32]string{
				1: "a",
				2: "b",
			},
		},
		{
			name: "空map",
			give: map[interface{}]interface{}{},
			want: map[int32]string{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
		{
			name: "panic - 键类型不匹配",
			give: map[interface{}]interface{}{
				"not a number": "a",
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMapInt32String() expected panic")
					}
				}()
			}
			got := ToMapInt32String(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapInt32String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapInt64String(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[int64]string
		panic bool
	}{
		{
			name: "正常情况",
			give: map[interface{}]interface{}{
				1: "a",
				2: "b",
			},
			want: map[int64]string{
				1: "a",
				2: "b",
			},
		},
		{
			name: "类型转换 - int64",
			give: map[interface{}]interface{}{
				int64(1): "a",
				int64(2): "b",
			},
			want: map[int64]string{
				1: "a",
				2: "b",
			},
		},
		{
			name: "类型转换 - int32",
			give: map[interface{}]interface{}{
				int32(1): "a",
				int32(2): "b",
			},
			want: map[int64]string{
				1: "a",
				2: "b",
			},
		},
		{
			name: "空map",
			give: map[interface{}]interface{}{},
			want: map[int64]string{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
		{
			name: "panic - 键类型不匹配",
			give: map[interface{}]interface{}{
				"not a number": "a",
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMapInt64String() expected panic")
					}
				}()
			}
			got := ToMapInt64String(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapInt64String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapStringInt64(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[string]int64
		panic bool
	}{
		{
			name: "正常情况",
			give: map[interface{}]interface{}{
				"a": 1,
				"b": 2,
			},
			want: map[string]int64{
				"a": 1,
				"b": 2,
			},
		},
		{
			name: "类型转换 - int64",
			give: map[interface{}]interface{}{
				"a": int64(1),
				"b": int64(2),
			},
			want: map[string]int64{
				"a": 1,
				"b": 2,
			},
		},
		{
			name: "类型转换 - int32",
			give: map[interface{}]interface{}{
				"a": int32(1),
				"b": int32(2),
			},
			want: map[string]int64{
				"a": 1,
				"b": 2,
			},
		},
		{
			name: "类型转换 - float64",
			give: map[interface{}]interface{}{
				"a": 1.0,
				"b": 2.0,
			},
			want: map[string]int64{
				"a": 1,
				"b": 2,
			},
		},
		{
			name: "空map",
			give: map[interface{}]interface{}{},
			want: map[string]int64{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
		{
			name: "panic - 键类型不匹配",
			give: map[interface{}]interface{}{
				1: "not a string key",
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMapStringInt64() expected panic")
					}
				}()
			}
			got := ToMapStringInt64(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStringInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapStringString(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[string]string
		panic bool
	}{
		{
			name: "正常情况",
			give: map[interface{}]interface{}{
				"a": "1",
				"b": "2",
			},
			want: map[string]string{
				"a": "1",
				"b": "2",
			},
		},
		{
			name: "类型转换 - 数字转字符串",
			give: map[interface{}]interface{}{
				"a": 1,
				"b": 2,
			},
			want: map[string]string{
				"a": "1",
				"b": "2",
			},
		},
		{
			name: "类型转换 - bool转字符串",
			give: map[interface{}]interface{}{
				"a": true,
				"b": false,
			},
			want: map[string]string{
				"a": "true",
				"b": "false",
			},
		},
		{
			name: "空map",
			give: map[interface{}]interface{}{},
			want: map[string]string{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
		{
			name: "panic - 键类型不匹配",
			give: map[interface{}]interface{}{
				1: "not a string key",
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMapStringString() expected panic")
					}
				}()
			}
			got := ToMapStringString(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStringString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[interface{}]interface{}
		panic bool
	}{
		{
			name: "正常情况",
			give: map[string]interface{}{
				"a": 1,
				"b": "2",
			},
			want: map[interface{}]interface{}{
				"a": 1,
				"b": "2",
			},
		},
		{
			name: "嵌套map",
			give: map[string]interface{}{
				"a": map[string]interface{}{
					"b": 1,
				},
			},
			want: map[interface{}]interface{}{
				"a": map[string]interface{}{
					"b": 1,
				},
			},
		},
		{
			name: "空map",
			give: map[string]interface{}{},
			want: map[interface{}]interface{}{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMap() expected panic")
					}
				}()
			}
			got := ToMap(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapStringAny(t *testing.T) {
	tests := []struct {
		name  string
		give  interface{}
		want  map[string]interface{}
		panic bool
	}{
		{
			name: "正常情况",
			give: map[interface{}]interface{}{
				"a": 1,
				"b": "2",
				"c": true,
			},
			want: map[string]interface{}{
				"a": 1,
				"b": "2",
				"c": true,
			},
		},
		{
			name: "嵌套map",
			give: map[interface{}]interface{}{
				"a": map[interface{}]interface{}{
					"b": 1,
				},
			},
			want: map[string]interface{}{
				"a": map[interface{}]interface{}{
					"b": 1,
				},
			},
		},
		{
			name: "空map",
			give: map[interface{}]interface{}{},
			want: map[string]interface{}{},
		},
		{
			name: "nil map",
			give: nil,
			want: nil,
		},
		{
			name:  "panic - 非map",
			give:  "not a map",
			panic: true,
		},
		{
			name: "panic - 键类型不匹配",
			give: map[interface{}]interface{}{
				1: "not a string key",
			},
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("ToMapStringAny() expected panic")
					}
				}()
			}
			got := ToMapStringAny(tt.give)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapStringAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
