package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMapFunctionsCoverage tests all 0% coverage map-related functions
func TestMapFunctionsCoverage(t *testing.T) {
	// Test KeyBy functions
	t.Run("KeyBy", func(t *testing.T) {
		type Person struct {
			ID   int
			Name string
		}

		people := []Person{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}

		// Test KeyBy with field name
		result := KeyBy(people, "ID")

		// KeyBy returns interface{}, so we need to cast it
		resultMap, ok := result.(map[int]Person)
		assert.True(t, ok)

		expected := map[int]Person{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
			3: {ID: 3, Name: "Charlie"},
		}

		assert.Equal(t, expected, resultMap)
	})

	t.Run("KeyByString", func(t *testing.T) {
		type User struct {
			Username string
			Email    string
		}

		users := []*User{
			{Username: "alice", Email: "alice@example.com"},
			{Username: "bob", Email: "bob@example.com"},
		}

		result := KeyByString(users, "Username")

		expected := map[string]*User{
			"alice": {Username: "alice", Email: "alice@example.com"},
			"bob":   {Username: "bob", Email: "bob@example.com"},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("KeyByInt32", func(t *testing.T) {
		type Item struct {
			ID   int32
			Name string
		}

		items := []*Item{
			{ID: 100, Name: "Item1"},
			{ID: 200, Name: "Item2"},
		}

		result := KeyByInt32(items, "ID")

		expected := map[int32]*Item{
			100: {ID: 100, Name: "Item1"},
			200: {ID: 200, Name: "Item2"},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("KeyByInt64", func(t *testing.T) {
		type Record struct {
			ID   int64
			Data string
		}

		records := []*Record{
			{ID: 1000, Data: "Data1"},
			{ID: 2000, Data: "Data2"},
		}

		result := KeyByInt64(records, "ID")

		expected := map[int64]*Record{
			1000: {ID: 1000, Data: "Data1"},
			2000: {ID: 2000, Data: "Data2"},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("KeyByUint64", func(t *testing.T) {
		type Entity struct {
			ID   uint64
			Name string
		}

		entities := []*Entity{
			{ID: 123456789, Name: "Entity1"},
			{ID: 987654321, Name: "Entity2"},
		}

		result := KeyByUint64(entities, "ID")

		expected := map[uint64]*Entity{
			123456789: {ID: 123456789, Name: "Entity1"},
			987654321: {ID: 987654321, Name: "Entity2"},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("KeyByGeneric", func(t *testing.T) {
		type Product struct {
			SKU  string
			Name string
		}

		products := []Product{
			{SKU: "ABC123", Name: "Product A"},
			{SKU: "XYZ789", Name: "Product X"},
		}

		result := KeyByGeneric(products, func(p Product) string {
			return p.SKU
		})

		expected := map[string]Product{
			"ABC123": {SKU: "ABC123", Name: "Product A"},
			"XYZ789": {SKU: "XYZ789", Name: "Product X"},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("KeyByPtr", func(t *testing.T) {
		type Node struct {
			ID   int
			Name string
		}

		nodes := []*Node{
			{ID: 1, Name: "Node1"},
			{ID: 2, Name: "Node2"},
		}

		result := KeyByPtr(nodes, func(n *Node) int {
			return n.ID
		})

		expected := map[int]*Node{
			1: {ID: 1, Name: "Node1"},
			2: {ID: 2, Name: "Node2"},
		}

		assert.Equal(t, expected, result)
	})
}

// TestMapKeysAndValuesFunctions tests all 0% coverage map keys/values functions
func TestMapKeysAndValuesFunctions(t *testing.T) {
	t.Run("MapKeysString", func(t *testing.T) {
		input := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		result := MapKeysString(input)
		expected := []string{"apple", "banana", "cherry"}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysInt", func(t *testing.T) {
		input := map[int]string{
			1: "one",
			2: "two",
			3: "three",
		}

		result := MapKeysInt(input)
		expected := []int{1, 2, 3}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysInt8", func(t *testing.T) {
		input := map[int8]string{
			1: "one",
			2: "two",
		}

		result := MapKeysInt8(input)
		expected := []int8{1, 2}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysInt16", func(t *testing.T) {
		input := map[int16]string{
			100: "hundred",
			200: "two hundred",
		}

		result := MapKeysInt16(input)
		expected := []int16{100, 200}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysInt32", func(t *testing.T) {
		input := map[int32]string{
			1000: "thousand",
			2000: "two thousand",
		}

		result := MapKeysInt32(input)
		expected := []int32{1000, 2000}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysInt64", func(t *testing.T) {
		input := map[int64]string{
			1000000: "million",
			2000000: "two million",
		}

		result := MapKeysInt64(input)
		expected := []int64{1000000, 2000000}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysUint", func(t *testing.T) {
		input := map[uint]string{
			10: "ten",
			20: "twenty",
		}

		result := MapKeysUint(input)
		expected := []uint{10, 20}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysUint8", func(t *testing.T) {
		input := map[uint8]string{
			5:  "five",
			10: "ten",
		}

		result := MapKeysUint8(input)
		expected := []uint8{5, 10}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysUint16", func(t *testing.T) {
		input := map[uint16]string{
			500:  "five hundred",
			1000: "thousand",
		}

		result := MapKeysUint16(input)
		expected := []uint16{500, 1000}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysUint32", func(t *testing.T) {
		input := map[uint32]string{
			50000:  "fifty thousand",
			100000: "hundred thousand",
		}

		result := MapKeysUint32(input)
		expected := []uint32{50000, 100000}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysUint64", func(t *testing.T) {
		input := map[uint64]string{
			5000000000:  "five billion",
			10000000000: "ten billion",
		}

		result := MapKeysUint64(input)
		expected := []uint64{5000000000, 10000000000}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysFloat32", func(t *testing.T) {
		input := map[float32]string{
			1.5: "one and half",
			2.5: "two and half",
		}

		result := MapKeysFloat32(input)
		expected := []float32{1.5, 2.5}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysFloat64", func(t *testing.T) {
		input := map[float64]string{
			3.14159: "pi",
			2.71828: "e",
		}

		result := MapKeysFloat64(input)
		expected := []float64{3.14159, 2.71828}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysInterface", func(t *testing.T) {
		input := map[interface{}]string{
			"key1": "value1",
			42:     "value2",
		}

		result := MapKeysInterface(input)

		assert.Len(t, result, 2)
		assert.Contains(t, result, "key1")
		assert.Contains(t, result, 42)
	})

	t.Run("MapKeysAny", func(t *testing.T) {
		input := map[any]string{
			"test": "value1",
			123:   "value2",
		}

		result := MapKeysAny(input)

		assert.Len(t, result, 2)
		assert.Contains(t, result, "test")
		assert.Contains(t, result, 123)
	})

	t.Run("MapKeysNumber", func(t *testing.T) {
		input := map[int]string{
			10: "ten",
			20: "twenty",
		}

		result := MapKeysNumber(input)
		expected := []int{10, 20}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapKeysGeneric", func(t *testing.T) {
		input := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		result := MapKeysGeneric(input)
		expected := []string{"a", "b", "c"}

		assert.ElementsMatch(t, expected, result)
	})
}

// TestMapValuesFunctions tests all 0% coverage map values functions
func TestMapValuesFunctions(t *testing.T) {
	t.Run("MapValues", func(t *testing.T) {
		input := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
		}

		result := MapValues(input)
		expected := []int{1, 2, 3}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapValuesGeneric", func(t *testing.T) {
		input := map[int]string{
			1: "one",
			2: "two",
			3: "three",
		}

		result := MapValuesGeneric(input)
		expected := []string{"one", "two", "three"}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapValuesAny", func(t *testing.T) {
		input := map[string]any{
			"number": 42,
			"text":   "hello",
			"bool":   true,
		}

		result := MapValuesAny(input)

		assert.Len(t, result, 3)
		assert.Contains(t, result, 42)
		assert.Contains(t, result, "hello")
		assert.Contains(t, result, true)
	})

	t.Run("MapValuesString", func(t *testing.T) {
		input := map[int]string{
			1: "first",
			2: "second",
			3: "third",
		}

		result := MapValuesString(input)
		expected := []string{"first", "second", "third"}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapValuesInt", func(t *testing.T) {
		input := map[string]int{
			"apple":  10,
			"banana": 20,
			"cherry": 30,
		}

		result := MapValuesInt(input)
		expected := []int{10, 20, 30}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("MapValuesFloat64", func(t *testing.T) {
		input := map[string]float64{
			"pi": 3.14159,
			"e":  2.71828,
		}

		result := MapValuesFloat64(input)
		expected := []float64{3.14159, 2.71828}

		assert.ElementsMatch(t, expected, result)
	})
}

// TestMapMergeFunctions tests all 0% coverage map merge functions
func TestMapMergeFunctions(t *testing.T) {
	t.Run("MergeMap", func(t *testing.T) {
		map1 := map[string]int{
			"a": 1,
			"b": 2,
		}

		map2 := map[string]int{
			"c": 3,
			"d": 4,
		}

		result := MergeMap(map1, map2)
		expected := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("MergeMapGeneric", func(t *testing.T) {
		map1 := map[int]string{
			1: "one",
			2: "two",
		}

		map2 := map[int]string{
			3: "three",
			4: "four",
		}

		result := MergeMapGeneric(map1, map2)
		expected := map[int]string{
			1: "one",
			2: "two",
			3: "three",
			4: "four",
		}

		assert.Equal(t, expected, result)
	})

	t.Run("MergeMapOverwrite", func(t *testing.T) {
		map1 := map[string]int{
			"a": 1,
			"b": 2,
		}

		map2 := map[string]int{
			"b": 20, // Should overwrite
			"c": 3,
		}

		result := MergeMapGeneric(map1, map2)
		expected := map[string]int{
			"a": 1,
			"b": 20,
			"c": 3,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("CloneMapShallow", func(t *testing.T) {
		original := map[string]int{
			"x": 10,
			"y": 20,
			"z": 30,
		}

		cloned := CloneMapShallow(original)

		// Should be equal
		assert.Equal(t, original, cloned)

		// Modifying clone shouldn't affect original
		cloned["w"] = 40
		assert.NotContains(t, original, "w")
		assert.Contains(t, cloned, "w")

		// Test with nil map
		var nilMap map[string]int
		clonedNil := CloneMapShallow(nilMap)
		assert.Nil(t, clonedNil)
	})
}

// TestMapUtilsFunctions tests 0% coverage map utility functions
func TestMapUtilsFunctions(t *testing.T) {
	t.Run("CheckValueType", func(t *testing.T) {
		// Test with string value
		result := CheckValueType("hello")
		assert.Equal(t, ValueString, result)

		// Test with int value
		result = CheckValueType(123)
		assert.Equal(t, ValueNumber, result)

		// Test with float value
		result = CheckValueType(3.14)
		assert.Equal(t, ValueNumber, result)

		// Test with bool value
		result = CheckValueType(true)
		assert.Equal(t, ValueBool, result)

		// Test with byte slice
		result = CheckValueType([]byte("hello"))
		assert.Equal(t, ValueString, result)

		// Test with unknown type
		result = CheckValueType(map[string]interface{}{})
		assert.Equal(t, ValueUnknown, result)

		// Test all integer types
		result = CheckValueType(int8(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(int16(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(int32(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(int64(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(uint(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(uint8(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(uint16(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(uint32(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(uint64(1))
		assert.Equal(t, ValueNumber, result)

		result = CheckValueType(float32(1.0))
		assert.Equal(t, ValueNumber, result)
	})
}