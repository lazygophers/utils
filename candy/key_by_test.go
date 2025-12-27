package candy

import (
	"reflect"
	"testing"
)

type TestPerson struct {
	ID      int
	ID8     int8
	ID16    int16
	ID32    int32
	ID64    int64
	UID     uint
	UID8    uint8
	UID16   uint16
	UID32   uint32
	UID64   uint64
	Name    string
	Age     int
	Score   float32
	Score64 float64
	Active  bool
	Code    string
}

func TestKeyByInt(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[int]TestPerson
	}{
		{
			name: "KeyBy ID field",
			ss: []TestPerson{
				{ID: 1, Name: "Alice", Age: 25},
				{ID: 2, Name: "Bob", Age: 30},
				{ID: 3, Name: "Charlie", Age: 35},
			},
			fieldName: "ID",
			expected: map[int]TestPerson{
				1: {ID: 1, Name: "Alice", Age: 25},
				2: {ID: 2, Name: "Bob", Age: 30},
				3: {ID: 3, Name: "Charlie", Age: 35},
			},
		},
		{
			name: "KeyBy Age field",
			ss: []TestPerson{
				{ID: 1, Name: "Alice", Age: 25},
				{ID: 2, Name: "Bob", Age: 30},
				{ID: 3, Name: "Charlie", Age: 25}, // Duplicate age
			},
			fieldName: "Age",
			expected: map[int]TestPerson{
				25: {ID: 3, Name: "Charlie", Age: 25}, // Last item wins for duplicate keys
				30: {ID: 2, Name: "Bob", Age: 30},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "ID",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByInt(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByInt() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByIntWithPointers(t *testing.T) {
	tests := []struct {
		name      string
		ss        []*TestPerson
		fieldName string
		expected  map[int]*TestPerson
	}{
		{
			name: "KeyBy ID field with pointers",
			ss: []*TestPerson{
				{ID: 1, Name: "Alice", Age: 25},
				{ID: 2, Name: "Bob", Age: 30},
				{ID: 3, Name: "Charlie", Age: 35},
			},
			fieldName: "ID",
			expected: map[int]*TestPerson{
				1: {ID: 1, Name: "Alice", Age: 25},
				2: {ID: 2, Name: "Bob", Age: 30},
				3: {ID: 3, Name: "Charlie", Age: 35},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByInt(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByInt() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{ID: 1, Name: "Alice", Age: 25},
	}
	KeyByInt(people, "NonExistentField")
}

func TestKeyByIntPanicWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{ID: 1, Name: "Alice", Age: 25},
	}
	KeyByInt(people, "Name") // Name is string, not int
}

func TestKeyByIntPanicNotStruct(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	numbers := []int{1, 2, 3}
	KeyByInt(numbers, "SomeField") // int is not a struct
}

func TestKeyByIntPanicWithPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	type Person struct {
		Name string
	}

	people := []*Person{
		{Name: "Alice"},
	}
	KeyByInt(people, "Name") // Should panic because element is not a struct
}

func TestKeyByIntPanicWithDoublePointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	type Person struct {
		Name string
	}

	person := Person{Name: "Alice"}
	people := []*Person{&person}
	KeyByInt(people, "Name") // Should panic because element is not a struct
}

func TestKeyByIntPanicWithNonStructElement(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	// Test element is not a struct
	people := []int{1, 2, 3}
	KeyByInt(people, "Invalid")
}

func TestKeyByIntPanicWithPointerToNonStruct(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	// Test pointer to non-struct
	num1 := 1
	num2 := 2
	people := []*int{&num1, &num2}
	KeyByInt(people, "Invalid")
}

func TestKeyByIntPanicWithInvalidField2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	type Person struct {
		Name string
	}

	people := []Person{{Name: "Alice"}}
	KeyByInt(people, "InvalidField")
}

func TestKeyByIntPanicWithNonStructElement2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	// Test element is not a struct
	people := []int{1, 2, 3}
	KeyByInt(people, "Invalid")
}

func TestKeyByIntPanicWithInvalidField3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	type Person struct {
		Name string
	}

	people := []Person{{Name: "Alice"}}
	KeyByInt(people, "InvalidField")
}

func TestKeyByInt8(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[int8]TestPerson
	}{
		{
			name: "KeyBy ID8 field",
			ss: []TestPerson{
				{ID8: 1, Name: "Alice", Age: 25},
				{ID8: 2, Name: "Bob", Age: 30},
				{ID8: 3, Name: "Charlie", Age: 35},
			},
			fieldName: "ID8",
			expected: map[int8]TestPerson{
				1: {ID8: 1, Name: "Alice", Age: 25},
				2: {ID8: 2, Name: "Bob", Age: 30},
				3: {ID8: 3, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "ID8",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByInt8(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByInt8() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByInt16(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[int16]TestPerson
	}{
		{
			name: "KeyBy ID16 field",
			ss: []TestPerson{
				{ID16: 100, Name: "Alice", Age: 25},
				{ID16: 200, Name: "Bob", Age: 30},
				{ID16: 300, Name: "Charlie", Age: 35},
			},
			fieldName: "ID16",
			expected: map[int16]TestPerson{
				100: {ID16: 100, Name: "Alice", Age: 25},
				200: {ID16: 200, Name: "Bob", Age: 30},
				300: {ID16: 300, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "ID16",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByInt16(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByInt16() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByInt32(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[int32]TestPerson
	}{
		{
			name: "KeyBy ID32 field",
			ss: []TestPerson{
				{ID32: 1000, Name: "Alice", Age: 25},
				{ID32: 2000, Name: "Bob", Age: 30},
				{ID32: 3000, Name: "Charlie", Age: 35},
			},
			fieldName: "ID32",
			expected: map[int32]TestPerson{
				1000: {ID32: 1000, Name: "Alice", Age: 25},
				2000: {ID32: 2000, Name: "Bob", Age: 30},
				3000: {ID32: 3000, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "ID32",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByInt32(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByInt32() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByInt64(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[int64]TestPerson
	}{
		{
			name: "KeyBy ID64 field",
			ss: []TestPerson{
				{ID64: 1000000, Name: "Alice", Age: 25},
				{ID64: 2000000, Name: "Bob", Age: 30},
				{ID64: 3000000, Name: "Charlie", Age: 35},
			},
			fieldName: "ID64",
			expected: map[int64]TestPerson{
				1000000: {ID64: 1000000, Name: "Alice", Age: 25},
				2000000: {ID64: 2000000, Name: "Bob", Age: 30},
				3000000: {ID64: 3000000, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "ID64",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByInt64(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByInt64() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByInt8Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{ID8: 1, Name: "Alice", Age: 25},
	}
	KeyByInt8(people, "NonExistentField")
}

func TestKeyByInt16Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{ID16: 1, Name: "Alice", Age: 25},
	}
	KeyByInt16(people, "NonExistentField")
}

func TestKeyByInt32Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{ID32: 1, Name: "Alice", Age: 25},
	}
	KeyByInt32(people, "NonExistentField")
}

func TestKeyByInt64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{ID64: 1, Name: "Alice", Age: 25},
	}
	KeyByInt64(people, "NonExistentField")
}

func TestKeyByUint(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[uint]TestPerson
	}{
		{
			name: "KeyBy UID field",
			ss: []TestPerson{
				{UID: 100, Name: "Alice", Age: 25},
				{UID: 200, Name: "Bob", Age: 30},
				{UID: 300, Name: "Charlie", Age: 35},
			},
			fieldName: "UID",
			expected: map[uint]TestPerson{
				100: {UID: 100, Name: "Alice", Age: 25},
				200: {UID: 200, Name: "Bob", Age: 30},
				300: {UID: 300, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "UID",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByUint(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByUint() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByUint8(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[uint8]TestPerson
	}{
		{
			name: "KeyBy UID8 field",
			ss: []TestPerson{
				{UID8: 10, Name: "Alice", Age: 25},
				{UID8: 20, Name: "Bob", Age: 30},
				{UID8: 30, Name: "Charlie", Age: 35},
			},
			fieldName: "UID8",
			expected: map[uint8]TestPerson{
				10: {UID8: 10, Name: "Alice", Age: 25},
				20: {UID8: 20, Name: "Bob", Age: 30},
				30: {UID8: 30, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "UID8",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByUint8(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByUint8() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByUint16(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[uint16]TestPerson
	}{
		{
			name: "KeyBy UID16 field",
			ss: []TestPerson{
				{UID16: 1000, Name: "Alice", Age: 25},
				{UID16: 2000, Name: "Bob", Age: 30},
				{UID16: 3000, Name: "Charlie", Age: 35},
			},
			fieldName: "UID16",
			expected: map[uint16]TestPerson{
				1000: {UID16: 1000, Name: "Alice", Age: 25},
				2000: {UID16: 2000, Name: "Bob", Age: 30},
				3000: {UID16: 3000, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "UID16",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByUint16(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByUint16() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByUint32(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[uint32]TestPerson
	}{
		{
			name: "KeyBy UID32 field",
			ss: []TestPerson{
				{UID32: 100000, Name: "Alice", Age: 25},
				{UID32: 200000, Name: "Bob", Age: 30},
				{UID32: 300000, Name: "Charlie", Age: 35},
			},
			fieldName: "UID32",
			expected: map[uint32]TestPerson{
				100000: {UID32: 100000, Name: "Alice", Age: 25},
				200000: {UID32: 200000, Name: "Bob", Age: 30},
				300000: {UID32: 300000, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "UID32",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByUint32(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByUint32() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByUint64(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[uint64]TestPerson
	}{
		{
			name: "KeyBy UID64 field",
			ss: []TestPerson{
				{UID64: 1000000000, Name: "Alice", Age: 25},
				{UID64: 2000000000, Name: "Bob", Age: 30},
				{UID64: 3000000000, Name: "Charlie", Age: 35},
			},
			fieldName: "UID64",
			expected: map[uint64]TestPerson{
				1000000000: {UID64: 1000000000, Name: "Alice", Age: 25},
				2000000000: {UID64: 2000000000, Name: "Bob", Age: 30},
				3000000000: {UID64: 3000000000, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "UID64",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByUint64(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByUint64() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByUintPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{UID: 1, Name: "Alice", Age: 25},
	}
	KeyByUint(people, "NonExistentField")
}

func TestKeyByUint8Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{UID8: 1, Name: "Alice", Age: 25},
	}
	KeyByUint8(people, "NonExistentField")
}

func TestKeyByUint16Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{UID16: 1, Name: "Alice", Age: 25},
	}
	KeyByUint16(people, "NonExistentField")
}

func TestKeyByUint32Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{UID32: 1, Name: "Alice", Age: 25},
	}
	KeyByUint32(people, "NonExistentField")
}

func TestKeyByUint64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{UID64: 1, Name: "Alice", Age: 25},
	}
	KeyByUint64(people, "NonExistentField")
}

func TestKeyByFloat32(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[float32]TestPerson
	}{
		{
			name: "KeyBy Score field",
			ss: []TestPerson{
				{Score: 95.5, Name: "Alice", Age: 25},
				{Score: 88.0, Name: "Bob", Age: 30},
				{Score: 92.3, Name: "Charlie", Age: 35},
			},
			fieldName: "Score",
			expected: map[float32]TestPerson{
				95.5: {Score: 95.5, Name: "Alice", Age: 25},
				88.0: {Score: 88.0, Name: "Bob", Age: 30},
				92.3: {Score: 92.3, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "Score",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByFloat32(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByFloat32() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByFloat64(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[float64]TestPerson
	}{
		{
			name: "KeyBy Score64 field",
			ss: []TestPerson{
				{Score64: 95.123456, Name: "Alice", Age: 25},
				{Score64: 88.987654, Name: "Bob", Age: 30},
				{Score64: 92.555555, Name: "Charlie", Age: 35},
			},
			fieldName: "Score64",
			expected: map[float64]TestPerson{
				95.123456: {Score64: 95.123456, Name: "Alice", Age: 25},
				88.987654: {Score64: 88.987654, Name: "Bob", Age: 30},
				92.555555: {Score64: 92.555555, Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "Score64",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByFloat64(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByFloat64() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByString(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[string]TestPerson
	}{
		{
			name: "KeyBy Name field",
			ss: []TestPerson{
				{Name: "Alice", Age: 25},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 35},
			},
			fieldName: "Name",
			expected: map[string]TestPerson{
				"Alice":   {Name: "Alice", Age: 25},
				"Bob":     {Name: "Bob", Age: 30},
				"Charlie": {Name: "Charlie", Age: 35},
			},
		},
		{
			name: "KeyBy Code field",
			ss: []TestPerson{
				{Code: "A001", Name: "Alice", Age: 25},
				{Code: "B002", Name: "Bob", Age: 30},
				{Code: "C003", Name: "Charlie", Age: 35},
			},
			fieldName: "Code",
			expected: map[string]TestPerson{
				"A001": {Code: "A001", Name: "Alice", Age: 25},
				"B002": {Code: "B002", Name: "Bob", Age: 30},
				"C003": {Code: "C003", Name: "Charlie", Age: 35},
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "Name",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByString(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByBool(t *testing.T) {
	tests := []struct {
		name      string
		ss        []TestPerson
		fieldName string
		expected  map[bool]TestPerson
	}{
		{
			name: "KeyBy Active field",
			ss: []TestPerson{
				{Active: true, Name: "Alice", Age: 25},
				{Active: false, Name: "Bob", Age: 30},
				{Active: true, Name: "Charlie", Age: 35},
			},
			fieldName: "Active",
			expected: map[bool]TestPerson{
				true:  {Active: true, Name: "Charlie", Age: 35}, // Last true wins
				false: {Active: false, Name: "Bob", Age: 30},
			},
		},
		{
			name: "All true",
			ss: []TestPerson{
				{Active: true, Name: "Alice", Age: 25},
				{Active: true, Name: "Bob", Age: 30},
			},
			fieldName: "Active",
			expected: map[bool]TestPerson{
				true: {Active: true, Name: "Bob", Age: 30}, // Last one wins
			},
		},
		{
			name:      "Empty slice",
			ss:        []TestPerson{},
			fieldName: "Active",
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KeyByBool(tt.ss, tt.fieldName)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("KeyByBool() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKeyByFloat32Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{Score: 95.5, Name: "Alice", Age: 25},
	}
	KeyByFloat32(people, "NonExistentField")
}

func TestKeyByFloat64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{Score64: 95.123456, Name: "Alice", Age: 25},
	}
	KeyByFloat64(people, "NonExistentField")
}

func TestKeyByStringPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{Name: "Alice", Age: 25},
	}
	KeyByString(people, "NonExistentField")
}

func TestKeyByBoolPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but didn't get one")
		}
	}()

	people := []TestPerson{
		{Active: true, Name: "Alice", Age: 25},
	}
	KeyByBool(people, "NonExistentField")
}
