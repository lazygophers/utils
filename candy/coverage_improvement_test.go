package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPluckIntCoverage targets the 43.6% coverage pluck function in pluck_int.go
func TestPluckIntCoverage(t *testing.T) {
	t.Run("pluck int function coverage", func(t *testing.T) {
		// Test with various struct types to hit different code paths
		type Person struct {
			ID   int
			Name string
			Age  int
		}

		people := []Person{
			{ID: 1, Name: "Alice", Age: 25},
			{ID: 2, Name: "Bob", Age: 30},
			{ID: 3, Name: "Charlie", Age: 35},
		}

		// Test plucking int field using selector function
		ids := Pluck(people, func(p Person) int { return p.ID })
		expectedIDs := []int{1, 2, 3}
		assert.Equal(t, expectedIDs, ids)

		// Test plucking string field to ensure proper type handling
		names := Pluck(people, func(p Person) string { return p.Name })
		expectedNames := []string{"Alice", "Bob", "Charlie"}
		assert.Equal(t, expectedNames, names)

		// Test with empty slice
		var emptyPeople []Person
		emptyIDs := Pluck(emptyPeople, func(p Person) int { return p.ID })
		assert.Empty(t, emptyIDs)

		// Test with nil slice
		var nilPeople []Person = nil
		nilIDs := Pluck(nilPeople, func(p Person) int { return p.ID })
		assert.Empty(t, nilIDs)

		// Test with different struct types
		type Product struct {
			ProductID int
			Price     float64
			Quantity  int
		}

		products := []Product{
			{ProductID: 101, Price: 19.99, Quantity: 10},
			{ProductID: 102, Price: 29.99, Quantity: 5},
		}

		productIDs := Pluck(products, func(p Product) int { return p.ProductID })
		expectedProductIDs := []int{101, 102}
		assert.Equal(t, expectedProductIDs, productIDs)

		quantities := Pluck(products, func(p Product) int { return p.Quantity })
		expectedQuantities := []int{10, 5}
		assert.Equal(t, expectedQuantities, quantities)
	})
}

// TestConvertCoverageImproved targets the 57.1% coverage Convert function
func TestConvertCoverageImproved(t *testing.T) {
	t.Run("convert function coverage", func(t *testing.T) {
		// Test various conversion scenarios
		t.Run("basic conversions", func(t *testing.T) {
			// Test string to int (Convert function has different signature)
			result := Convert[string, int]("123")
			assert.Equal(t, 123, result)

			// Test string to float64
			resultFloat := Convert[string, float64]("123.45")
			assert.Equal(t, 123.45, resultFloat)

			// Test bool to int
			resultBool := Convert[bool, int](true)
			assert.Equal(t, 1, resultBool)
		})

		t.Run("numeric conversions", func(t *testing.T) {
			// Test int to float64
			resultFloat := Convert[int, float64](456)
			assert.Equal(t, 456.0, resultFloat)

			// Test float to int
			resultInt := Convert[float64, int](789.123)
			assert.Equal(t, 789, resultInt)

			// Test int to int (same type)
			resultSame := Convert[int, int](100)
			assert.Equal(t, 100, resultSame)
		})

		t.Run("edge cases", func(t *testing.T) {
			// Test bool false to int
			resultZero := Convert[bool, int](false)
			assert.Equal(t, 0, resultZero)

			// Test zero values
			resultZero2 := Convert[int, int](0)
			assert.Equal(t, 0, resultZero2)

			// Test negative numbers
			resultNeg := Convert[int, float64](-42)
			assert.Equal(t, -42.0, resultNeg)
		})

		t.Run("various numeric types", func(t *testing.T) {
			// Test various integer types
			assert.Equal(t, int8(42), Convert[int, int8](42))
			assert.Equal(t, int16(42), Convert[int, int16](42))
			assert.Equal(t, int32(42), Convert[int, int32](42))
			assert.Equal(t, int64(42), Convert[int, int64](42))
			assert.Equal(t, uint8(42), Convert[int, uint8](42))
			assert.Equal(t, uint16(42), Convert[int, uint16](42))
			assert.Equal(t, uint32(42), Convert[int, uint32](42))
			assert.Equal(t, uint64(42), Convert[int, uint64](42))
		})
	})
}

// TestToStringGenericCoverage targets the 50% coverage ToStringGeneric function
func TestToStringGenericCoverage(t *testing.T) {
	t.Run("to string generic coverage", func(t *testing.T) {
		// Test various types to hit different code paths
		t.Run("basic types", func(t *testing.T) {
			assert.Equal(t, "123", ToStringGeneric(123))
			assert.Equal(t, "123.45", ToStringGeneric(123.45))
			assert.Equal(t, "true", ToStringGeneric(true))
			assert.Equal(t, "false", ToStringGeneric(false))
			assert.Equal(t, "hello", ToStringGeneric("hello"))
		})

		t.Run("numeric types", func(t *testing.T) {
			assert.Equal(t, "42", ToStringGeneric(int8(42)))
			assert.Equal(t, "42", ToStringGeneric(int16(42)))
			assert.Equal(t, "42", ToStringGeneric(int32(42)))
			assert.Equal(t, "42", ToStringGeneric(int64(42)))
			assert.Equal(t, "42", ToStringGeneric(uint8(42)))
			assert.Equal(t, "42", ToStringGeneric(uint16(42)))
			assert.Equal(t, "42", ToStringGeneric(uint32(42)))
			assert.Equal(t, "42", ToStringGeneric(uint64(42)))
			assert.Equal(t, "42", ToStringGeneric(uint(42)))
		})

		t.Run("float types", func(t *testing.T) {
			assert.Equal(t, "3.14", ToStringGeneric(float32(3.14)))
			assert.Equal(t, "3.14159", ToStringGeneric(float64(3.14159)))
		})

		t.Run("complex types", func(t *testing.T) {
			// Test slice - ToStringGeneric returns empty string for complex types
			slice := []int{1, 2, 3}
			result := ToStringGeneric(slice)
			assert.Equal(t, "", result)

			// Test map - ToStringGeneric returns empty string for complex types
			m := map[string]int{"a": 1, "b": 2}
			result = ToStringGeneric(m)
			assert.Equal(t, "", result)

			// Test struct - ToStringGeneric returns empty string for complex types
			type TestStruct struct {
				Name string
				Age  int
			}
			s := TestStruct{Name: "Alice", Age: 30}
			result = ToStringGeneric(s)
			assert.Equal(t, "", result)
		})

		t.Run("edge cases", func(t *testing.T) {
			// Test nil interface - ToStringGeneric returns empty string for invalid values
			var nilInterface interface{} = nil
			assert.Equal(t, "", ToStringGeneric(nilInterface))

			// Test zero values
			assert.Equal(t, "0", ToStringGeneric(0))
			assert.Equal(t, "", ToStringGeneric(""))
			assert.Equal(t, "false", ToStringGeneric(false))
		})
	})
}

// TestToUint32Coverage targets the 72% coverage ToUint32 function
func TestToUint32Coverage(t *testing.T) {
	t.Run("to uint32 coverage", func(t *testing.T) {
		// Test various input types to improve coverage
		t.Run("valid conversions", func(t *testing.T) {
			// Test from int
			result := ToUint32(123)
			assert.Equal(t, uint32(123), result)

			// Test from string
			result = ToUint32("456")
			assert.Equal(t, uint32(456), result)

			// Test from float
			result = ToUint32(789.123)
			assert.Equal(t, uint32(789), result)

			// Test from bool
			result = ToUint32(true)
			assert.Equal(t, uint32(1), result)
			result = ToUint32(false)
			assert.Equal(t, uint32(0), result)
		})

		t.Run("edge cases", func(t *testing.T) {
			// Test zero values
			assert.Equal(t, uint32(0), ToUint32(0))
			assert.Equal(t, uint32(0), ToUint32("0"))

			// Test invalid string
			result := ToUint32("invalid")
			assert.Equal(t, uint32(0), result) // Should default to 0

			// Test negative numbers (should be handled gracefully)
			result = ToUint32(-123)
			// Depends on implementation, but should handle gracefully

			// Test large numbers
			result = ToUint32(4294967295) // Max uint32
			assert.Equal(t, uint32(4294967295), result)
		})

		t.Run("various numeric types", func(t *testing.T) {
			assert.Equal(t, uint32(42), ToUint32(int8(42)))
			assert.Equal(t, uint32(42), ToUint32(int16(42)))
			assert.Equal(t, uint32(42), ToUint32(int32(42)))
			assert.Equal(t, uint32(42), ToUint32(int64(42)))
			assert.Equal(t, uint32(42), ToUint32(uint8(42)))
			assert.Equal(t, uint32(42), ToUint32(uint16(42)))
			assert.Equal(t, uint32(42), ToUint32(uint32(42)))
			assert.Equal(t, uint32(42), ToUint32(uint64(42)))
		})
	})
}

// TestToUint64Coverage targets the 68% coverage ToUint64 function
func TestToUint64Coverage(t *testing.T) {
	t.Run("to uint64 coverage", func(t *testing.T) {
		// Test various input types to improve coverage
		t.Run("valid conversions", func(t *testing.T) {
			// Test from int
			result := ToUint64(123)
			assert.Equal(t, uint64(123), result)

			// Test from string
			result = ToUint64("456")
			assert.Equal(t, uint64(456), result)

			// Test from float
			result = ToUint64(789.123)
			assert.Equal(t, uint64(789), result)

			// Test from bool
			result = ToUint64(true)
			assert.Equal(t, uint64(1), result)
			result = ToUint64(false)
			assert.Equal(t, uint64(0), result)
		})

		t.Run("edge cases", func(t *testing.T) {
			// Test zero values
			assert.Equal(t, uint64(0), ToUint64(0))
			assert.Equal(t, uint64(0), ToUint64("0"))

			// Test invalid string
			result := ToUint64("invalid")
			assert.Equal(t, uint64(0), result) // Should default to 0

			// Test large numbers
			result = ToUint64(uint64(18446744073709551615)) // Max uint64
			assert.Equal(t, uint64(18446744073709551615), result)
		})

		t.Run("various numeric types", func(t *testing.T) {
			assert.Equal(t, uint64(42), ToUint64(int8(42)))
			assert.Equal(t, uint64(42), ToUint64(int16(42)))
			assert.Equal(t, uint64(42), ToUint64(int32(42)))
			assert.Equal(t, uint64(42), ToUint64(int64(42)))
			assert.Equal(t, uint64(42), ToUint64(uint8(42)))
			assert.Equal(t, uint64(42), ToUint64(uint16(42)))
			assert.Equal(t, uint64(42), ToUint64(uint32(42)))
			assert.Equal(t, uint64(42), ToUint64(uint64(42)))
		})
	})
}

// TestKeyByFunctionsCoverage tests various KeyBy functions that have 69-85% coverage
func TestKeyByFunctionsCoverage(t *testing.T) {
	t.Run("KeyBy generic function coverage", func(t *testing.T) {
		// Test basic KeyBy functionality
		type Person struct {
			ID   int
			Name string
		}

		people := []Person{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Alice"}, // Duplicate name to test overwrite
		}

		// Test KeyBy with string key using field name
		nameMapInterface := KeyBy(people, "Name")
		nameMap, ok := nameMapInterface.(map[string]Person)
		assert.True(t, ok)
		assert.Contains(t, nameMap, "Alice")
		assert.Contains(t, nameMap, "Bob")
		assert.Equal(t, Person{ID: 3, Name: "Alice"}, nameMap["Alice"]) // Should be overwritten

		// Test KeyBy with int key using field name
		idMapInterface := KeyBy(people, "ID")
		idMap, ok := idMapInterface.(map[int]Person)
		assert.True(t, ok)
		assert.Contains(t, idMap, 1)
		assert.Contains(t, idMap, 2)
		assert.Contains(t, idMap, 3)

		// Test with empty slice - KeyBy returns empty map, not nil
		var emptyPeople []Person
		emptyMap := KeyBy(emptyPeople, "Name")
		assert.NotNil(t, emptyMap)
		// The function returns an empty map, not nil
		emptyMapTyped := emptyMap.(map[string]Person)
		assert.Equal(t, 0, len(emptyMapTyped))
	})

	t.Run("KeyBy with different types", func(t *testing.T) {
		// Test with different struct types to improve coverage
		type Item struct {
			Code string
			Name string
		}

		items := []Item{
			{Code: "A001", Name: "Item 1"},
			{Code: "A002", Name: "Item 2"},
		}

		// Test KeyBy with field names
		codeMapInterface := KeyBy(items, "Code")
		codeMap, ok := codeMapInterface.(map[string]Item)
		assert.True(t, ok)
		assert.Contains(t, codeMap, "A001")
		assert.Contains(t, codeMap, "A002")
		assert.Equal(t, Item{Code: "A001", Name: "Item 1"}, codeMap["A001"])

		// Test with pointer slice
		itemPtrs := []*Item{
			{Code: "B001", Name: "Item 3"},
			{Code: "B002", Name: "Item 4"},
		}

		ptrMapInterface := KeyBy(itemPtrs, "Code")
		ptrMap, ok := ptrMapInterface.(map[string]*Item)
		assert.True(t, ok)
		assert.Contains(t, ptrMap, "B001")
		assert.Contains(t, ptrMap, "B002")
	})
}

// TestMapKeysFunctionsCoverage tests various MapKeys functions that have 75% coverage
func TestMapKeysFunctionsCoverage(t *testing.T) {
	t.Run("MapKeysString coverage", func(t *testing.T) {
		m := map[string]int{
			"apple":  1,
			"banana": 2,
			"cherry": 3,
		}

		keys := MapKeysString(m)
		assert.Len(t, keys, 3)
		assert.Contains(t, keys, "apple")
		assert.Contains(t, keys, "banana")
		assert.Contains(t, keys, "cherry")

		// Test with empty map
		emptyMap := make(map[string]int)
		emptyKeys := MapKeysString(emptyMap)
		assert.Empty(t, emptyKeys)
	})

	t.Run("MapKeysInt coverage", func(t *testing.T) {
		m := map[int]string{
			1: "one",
			2: "two",
			3: "three",
		}

		keys := MapKeysInt(m)
		assert.Len(t, keys, 3)
		assert.Contains(t, keys, 1)
		assert.Contains(t, keys, 2)
		assert.Contains(t, keys, 3)
	})

	t.Run("MapKeysInt64 coverage", func(t *testing.T) {
		m := map[int64]string{
			int64(1000): "thousand",
			int64(2000): "two thousand",
		}

		keys := MapKeysInt64(m)
		assert.Len(t, keys, 2)
		assert.Contains(t, keys, int64(1000))
		assert.Contains(t, keys, int64(2000))
	})

	t.Run("MapKeysFloat64 coverage", func(t *testing.T) {
		m := map[float64]string{
			3.14:  "pi",
			2.718: "e",
		}

		keys := MapKeysFloat64(m)
		assert.Len(t, keys, 2)
		assert.Contains(t, keys, 3.14)
		assert.Contains(t, keys, 2.718)
	})

	t.Run("MapKeysGeneric coverage", func(t *testing.T) {
		// Test with custom type
		type CustomKey struct {
			ID   int
			Name string
		}

		m := map[CustomKey]string{
			{ID: 1, Name: "key1"}: "value1",
			{ID: 2, Name: "key2"}: "value2",
		}

		keys := MapKeysGeneric(m)
		assert.Len(t, keys, 2)
	})
}