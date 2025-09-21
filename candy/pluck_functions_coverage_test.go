package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPluckFunctionsCoverage tests all 0% coverage pluck-related functions
func TestPluckFunctionsCoverage(t *testing.T) {
	// Test data structures
	type Person struct {
		ID   int
		Name string
		Age  int
		City string
	}

	type Product struct {
		SKU   string
		Name  string
		Price float64
	}

	people := []Person{
		{ID: 1, Name: "Alice", Age: 30, City: "New York"},
		{ID: 2, Name: "Bob", Age: 25, City: "Los Angeles"},
		{ID: 3, Name: "Charlie", Age: 35, City: "Chicago"},
		{ID: 1, Name: "Alice2", Age: 28, City: "New York"}, // Duplicate ID for testing
	}

	products := []Product{
		{SKU: "ABC123", Name: "Widget A", Price: 10.99},
		{SKU: "XYZ789", Name: "Widget X", Price: 25.50},
		{SKU: "DEF456", Name: "Widget D", Price: 15.75},
	}

	t.Run("Pluck", func(t *testing.T) {
		result := Pluck(people, func(p Person) string {
			return p.Name
		})

		expected := []string{"Alice", "Bob", "Charlie", "Alice2"}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckPtr", func(t *testing.T) {
		peoplePtr := []*Person{
			{ID: 1, Name: "Alice", Age: 30},
			{ID: 2, Name: "Bob", Age: 25},
		}

		result := PluckPtr(peoplePtr, func(p *Person) int {
			return p.ID
		}, 0) // default value for int

		expected := []int{1, 2}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckFilter", func(t *testing.T) {
		result := PluckFilter(people, func(p Person) string {
			return p.Name
		}, func(p Person) bool {
			return p.Age >= 30
		})

		expected := []string{"Alice", "Charlie"}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckUnique", func(t *testing.T) {
		result := PluckUnique(people, func(p Person) int {
			return p.ID
		})

		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckMap", func(t *testing.T) {
		result := PluckMap(people, func(p Person) int {
			return p.ID
		}, func(p Person) string {
			return p.Name
		})

		// Should contain latest mapping for duplicate keys
		expected := map[int]string{
			1: "Alice2", // Latest Alice
			2: "Bob",
			3: "Charlie",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckGroupBy", func(t *testing.T) {
		result := PluckGroupBy(people, func(p Person) string {
			return p.City
		})

		expected := map[string][]Person{
			"New York": {
				{ID: 1, Name: "Alice", Age: 30, City: "New York"},
				{ID: 1, Name: "Alice2", Age: 28, City: "New York"},
			},
			"Los Angeles": {
				{ID: 2, Name: "Bob", Age: 25, City: "Los Angeles"},
			},
			"Chicago": {
				{ID: 3, Name: "Charlie", Age: 35, City: "Chicago"},
			},
		}
		assert.Equal(t, expected, result)
	})

	// Test typed pluck functions
	t.Run("PluckInt", func(t *testing.T) {
		result := PluckInt(people, "Age")

		expected := []int{30, 25, 35, 28}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckString", func(t *testing.T) {
		result := PluckString(products, "SKU")

		expected := []string{"ABC123", "XYZ789", "DEF456"}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckInt32", func(t *testing.T) {
		type Item struct {
			ID   int32
			Name string
		}

		items := []Item{
			{ID: 100, Name: "Item1"},
			{ID: 200, Name: "Item2"},
		}

		result := PluckInt32(items, "ID")

		expected := []int32{100, 200}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckInt64", func(t *testing.T) {
		type Record struct {
			ID   int64
			Data string
		}

		records := []Record{
			{ID: 1000, Data: "Data1"},
			{ID: 2000, Data: "Data2"},
		}

		result := PluckInt64(records, "ID")

		expected := []int64{1000, 2000}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckUint32", func(t *testing.T) {
		type Entity struct {
			ID   uint32
			Name string
		}

		entities := []Entity{
			{ID: 12345, Name: "Entity1"},
			{ID: 67890, Name: "Entity2"},
		}

		result := PluckUint32(entities, "ID")

		expected := []uint32{12345, 67890}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckUint64", func(t *testing.T) {
		type BigEntity struct {
			ID   uint64
			Name string
		}

		entities := []BigEntity{
			{ID: 123456789012345, Name: "BigEntity1"},
			{ID: 987654321098765, Name: "BigEntity2"},
		}

		result := PluckUint64(entities, "ID")

		expected := []uint64{123456789012345, 987654321098765}
		assert.Equal(t, expected, result)
	})

	// Test generic pluck functions
	t.Run("PluckIntGeneric", func(t *testing.T) {
		result := PluckIntGeneric(people, func(p Person) int {
			return p.Age
		})

		expected := []int{30, 25, 35, 28}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckStringGeneric", func(t *testing.T) {
		result := PluckStringGeneric(people, func(p Person) string {
			return p.City
		})

		expected := []string{"New York", "Los Angeles", "Chicago", "New York"}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckInt32Generic", func(t *testing.T) {
		type TestStruct struct {
			Value int32
		}

		items := []TestStruct{
			{Value: 100},
			{Value: 200},
		}

		result := PluckInt32Generic(items, func(item TestStruct) int32 {
			return item.Value
		})

		expected := []int32{100, 200}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckInt64Generic", func(t *testing.T) {
		type TestStruct struct {
			Value int64
		}

		items := []TestStruct{
			{Value: 1000},
			{Value: 2000},
		}

		result := PluckInt64Generic(items, func(item TestStruct) int64 {
			return item.Value
		})

		expected := []int64{1000, 2000}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckUint32Generic", func(t *testing.T) {
		type TestStruct struct {
			Value uint32
		}

		items := []TestStruct{
			{Value: 300},
			{Value: 400},
		}

		result := PluckUint32Generic(items, func(item TestStruct) uint32 {
			return item.Value
		})

		expected := []uint32{300, 400}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckUint64Generic", func(t *testing.T) {
		type TestStruct struct {
			Value uint64
		}

		items := []TestStruct{
			{Value: 500},
			{Value: 600},
		}

		result := PluckUint64Generic(items, func(item TestStruct) uint64 {
			return item.Value
		})

		expected := []uint64{500, 600}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckStringSlice", func(t *testing.T) {
		type Category struct {
			Name string
			Tags []string
		}

		categories := []Category{
			{Name: "Tech", Tags: []string{"programming", "software"}},
			{Name: "Science", Tags: []string{"research", "experiment"}},
		}

		result := PluckStringSlice(categories, "Tags")

		expected := [][]string{
			{"programming", "software"},
			{"research", "experiment"},
		}
		assert.Equal(t, expected, result)
	})
}

// TestPluckEdgeCases tests edge cases for pluck functions
func TestPluckEdgeCases(t *testing.T) {
	t.Run("PluckEmptySlice", func(t *testing.T) {
		var empty []string
		result := Pluck(empty, func(s string) string {
			return s
		})

		assert.Empty(t, result)
	})

	t.Run("PluckFilterNoneMatch", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := PluckFilter(numbers, func(n int) int {
			return n
		}, func(n int) bool {
			return n > 10 // No numbers > 10
		})

		assert.Empty(t, result)
	})

	t.Run("PluckUniqueAllSame", func(t *testing.T) {
		sameValues := []string{"same", "same", "same"}
		result := PluckUnique(sameValues, func(s string) string {
			return s
		})

		expected := []string{"same"}
		assert.Equal(t, expected, result)
	})

	t.Run("PluckGroupByEmptyKey", func(t *testing.T) {
		type Item struct {
			Key   string
			Value int
		}

		items := []Item{
			{Key: "", Value: 1},
			{Key: "a", Value: 2},
			{Key: "", Value: 3},
		}

		result := PluckGroupBy(items, func(item Item) string {
			return item.Key
		})

		expected := map[string][]Item{
			"": {
				{Key: "", Value: 1},
				{Key: "", Value: 3},
			},
			"a": {
				{Key: "a", Value: 2},
			},
		}
		assert.Equal(t, expected, result)
	})
}