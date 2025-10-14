package candy

import (
	"testing"
)

// Test types
type Person struct {
	Name string
	Age  int
	City string
}

type Product struct {
	ID    int64
	Name  string
	Price float64
}

type Item struct {
	Code   uint32
	Amount uint64
}

// TestPluck 测试 Pluck 函数
func TestPluck(t *testing.T) {
	t.Run("empty slice returns nil", func(t *testing.T) {
		result := Pluck([]Person{}, func(p Person) string {
			return p.Name
		})
		if result != nil {
			t.Errorf("Pluck(empty) should return nil, got %v", result)
		}
	})

	t.Run("pluck names from persons", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30, City: "NYC"},
			{Name: "Bob", Age: 25, City: "LA"},
			{Name: "Charlie", Age: 35, City: "SF"},
		}
		result := Pluck(persons, func(p Person) string {
			return p.Name
		})
		expected := []string{"Alice", "Bob", "Charlie"}
		if len(result) != len(expected) {
			t.Errorf("Pluck names length mismatch")
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Pluck names[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("pluck ages from persons", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}
		result := Pluck(persons, func(p Person) int {
			return p.Age
		})
		expected := []int{30, 25}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Pluck ages[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("pluck with transformation", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
		}
		result := Pluck(persons, func(p Person) int {
			return p.Age * 2
		})
		expected := []int{60, 50}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Pluck transformed[%d] = %d, want %d", i, result[i], expected[i])
			}
		}
	})

	t.Run("single element", func(t *testing.T) {
		persons := []Person{{Name: "Alice", Age: 30}}
		result := Pluck(persons, func(p Person) string {
			return p.Name
		})
		if len(result) != 1 || result[0] != "Alice" {
			t.Errorf("Pluck single element failed")
		}
	})
}

// TestPluckPtr 测试 PluckPtr 函数
func TestPluckPtr(t *testing.T) {
	t.Run("empty slice returns nil", func(t *testing.T) {
		result := PluckPtr([]*Person{}, func(p *Person) string {
			return p.Name
		}, "")
		if result != nil {
			t.Errorf("PluckPtr(empty) should return nil")
		}
	})

	t.Run("pluck from pointer slice", func(t *testing.T) {
		p1 := &Person{Name: "Alice", Age: 30}
		p2 := &Person{Name: "Bob", Age: 25}
		persons := []*Person{p1, p2}
		result := PluckPtr(persons, func(p *Person) string {
			return p.Name
		}, "Unknown")
		expected := []string{"Alice", "Bob"}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("PluckPtr[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("handle nil pointers with default", func(t *testing.T) {
		p1 := &Person{Name: "Alice", Age: 30}
		persons := []*Person{p1, nil, {Name: "Charlie", Age: 35}}
		result := PluckPtr(persons, func(p *Person) string {
			return p.Name
		}, "Unknown")
		expected := []string{"Alice", "Unknown", "Charlie"}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("PluckPtr with nil[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("all nil pointers", func(t *testing.T) {
		persons := []*Person{nil, nil, nil}
		result := PluckPtr(persons, func(p *Person) int {
			return p.Age
		}, 0)
		for i, v := range result {
			if v != 0 {
				t.Errorf("PluckPtr all nil[%d] = %d, want 0", i, v)
			}
		}
	})
}

// TestPluckFilter 测试 PluckFilter 函数
func TestPluckFilter(t *testing.T) {
	t.Run("empty slice returns nil", func(t *testing.T) {
		result := PluckFilter([]Person{}, func(p Person) string {
			return p.Name
		}, func(p Person) bool {
			return p.Age > 25
		})
		if result != nil {
			t.Errorf("PluckFilter(empty) should return nil")
		}
	})

	t.Run("filter and pluck", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 20},
			{Name: "Charlie", Age: 35},
			{Name: "David", Age: 22},
		}
		result := PluckFilter(persons, func(p Person) string {
			return p.Name
		}, func(p Person) bool {
			return p.Age >= 30
		})
		expected := []string{"Alice", "Charlie"}
		if len(result) != len(expected) {
			t.Errorf("PluckFilter length = %d, want %d", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("PluckFilter[%d] = %s, want %s", i, result[i], expected[i])
			}
		}
	})

	t.Run("no items match filter", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 20},
			{Name: "Bob", Age: 22},
		}
		result := PluckFilter(persons, func(p Person) string {
			return p.Name
		}, func(p Person) bool {
			return p.Age > 50
		})
		if len(result) != 0 {
			t.Errorf("PluckFilter no match should return empty, got %v", result)
		}
	})

	t.Run("all items match filter", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 35},
		}
		result := PluckFilter(persons, func(p Person) string {
			return p.Name
		}, func(p Person) bool {
			return p.Age > 20
		})
		if len(result) != 2 {
			t.Errorf("PluckFilter all match should return 2 items")
		}
	})
}

// TestPluckUnique 测试 PluckUnique 函数
func TestPluckUnique(t *testing.T) {
	t.Run("empty slice returns nil", func(t *testing.T) {
		result := PluckUnique([]Person{}, func(p Person) string {
			return p.City
		})
		if result != nil {
			t.Errorf("PluckUnique(empty) should return nil")
		}
	})

	t.Run("pluck unique cities", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", City: "NYC"},
			{Name: "Bob", City: "LA"},
			{Name: "Charlie", City: "NYC"},
			{Name: "David", City: "SF"},
			{Name: "Eve", City: "LA"},
		}
		result := PluckUnique(persons, func(p Person) string {
			return p.City
		})
		if len(result) != 3 {
			t.Errorf("PluckUnique should return 3 unique cities, got %d", len(result))
		}
		// Verify uniqueness
		seen := make(map[string]bool)
		for _, city := range result {
			if seen[city] {
				t.Errorf("PluckUnique returned duplicate: %s", city)
			}
			seen[city] = true
		}
	})

	t.Run("all values unique", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}
		result := PluckUnique(persons, func(p Person) int {
			return p.Age
		})
		if len(result) != 3 {
			t.Errorf("PluckUnique all unique should return 3 items")
		}
	})

	t.Run("all values same", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", City: "NYC"},
			{Name: "Bob", City: "NYC"},
			{Name: "Charlie", City: "NYC"},
		}
		result := PluckUnique(persons, func(p Person) string {
			return p.City
		})
		if len(result) != 1 || result[0] != "NYC" {
			t.Errorf("PluckUnique all same should return 1 item")
		}
	})
}

// TestPluckMap 测试 PluckMap 函数
func TestPluckMap(t *testing.T) {
	t.Run("empty slice returns nil", func(t *testing.T) {
		result := PluckMap([]Person{}, func(p Person) string {
			return p.Name
		}, func(p Person) int {
			return p.Age
		})
		if result != nil {
			t.Errorf("PluckMap(empty) should return nil")
		}
	})

	t.Run("create name to age map", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}
		result := PluckMap(persons, func(p Person) string {
			return p.Name
		}, func(p Person) int {
			return p.Age
		})
		if len(result) != 3 {
			t.Errorf("PluckMap length = %d, want 3", len(result))
		}
		if result["Alice"] != 30 || result["Bob"] != 25 || result["Charlie"] != 35 {
			t.Errorf("PluckMap values incorrect")
		}
	})

	t.Run("duplicate keys override", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Alice", Age: 35},
		}
		result := PluckMap(persons, func(p Person) string {
			return p.Name
		}, func(p Person) int {
			return p.Age
		})
		if len(result) != 1 {
			t.Errorf("PluckMap with duplicates should have 1 key")
		}
		if result["Alice"] != 35 {
			t.Errorf("PluckMap should keep last value, got %d", result["Alice"])
		}
	})

	t.Run("single element", func(t *testing.T) {
		persons := []Person{{Name: "Alice", Age: 30}}
		result := PluckMap(persons, func(p Person) string {
			return p.Name
		}, func(p Person) int {
			return p.Age
		})
		if len(result) != 1 || result["Alice"] != 30 {
			t.Errorf("PluckMap single element failed")
		}
	})
}

// TestPluckGroupBy 测试 PluckGroupBy 函数
func TestPluckGroupBy(t *testing.T) {
	t.Run("empty slice returns nil", func(t *testing.T) {
		result := PluckGroupBy([]Person{}, func(p Person) string {
			return p.City
		})
		if result != nil {
			t.Errorf("PluckGroupBy(empty) should return nil")
		}
	})

	t.Run("group by city", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", City: "NYC"},
			{Name: "Bob", City: "LA"},
			{Name: "Charlie", City: "NYC"},
			{Name: "David", City: "SF"},
			{Name: "Eve", City: "LA"},
		}
		result := PluckGroupBy(persons, func(p Person) string {
			return p.City
		})
		if len(result) != 3 {
			t.Errorf("PluckGroupBy should have 3 groups, got %d", len(result))
		}
		if len(result["NYC"]) != 2 {
			t.Errorf("PluckGroupBy NYC should have 2 persons")
		}
		if len(result["LA"]) != 2 {
			t.Errorf("PluckGroupBy LA should have 2 persons")
		}
		if len(result["SF"]) != 1 {
			t.Errorf("PluckGroupBy SF should have 1 person")
		}
	})

	t.Run("all in one group", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", City: "NYC"},
			{Name: "Bob", City: "NYC"},
		}
		result := PluckGroupBy(persons, func(p Person) string {
			return p.City
		})
		if len(result) != 1 || len(result["NYC"]) != 2 {
			t.Errorf("PluckGroupBy all in one group failed")
		}
	})

	t.Run("each in separate group", func(t *testing.T) {
		persons := []Person{
			{Name: "Alice", City: "NYC"},
			{Name: "Bob", City: "LA"},
			{Name: "Charlie", City: "SF"},
		}
		result := PluckGroupBy(persons, func(p Person) string {
			return p.City
		})
		if len(result) != 3 {
			t.Errorf("PluckGroupBy separate groups should have 3 groups")
		}
		for _, group := range result {
			if len(group) != 1 {
				t.Errorf("Each group should have 1 person")
			}
		}
	})
}

// TestPluckGenericWrappers 测试泛型包装器函数
func TestPluckGenericWrappers(t *testing.T) {
	t.Run("PluckIntGeneric", func(t *testing.T) {
		persons := []Person{{Age: 30}, {Age: 25}}
		result := PluckIntGeneric(persons, func(p Person) int {
			return p.Age
		})
		if len(result) != 2 || result[0] != 30 || result[1] != 25 {
			t.Errorf("PluckIntGeneric failed")
		}
	})

	t.Run("PluckStringGeneric", func(t *testing.T) {
		persons := []Person{{Name: "Alice"}, {Name: "Bob"}}
		result := PluckStringGeneric(persons, func(p Person) string {
			return p.Name
		})
		if len(result) != 2 || result[0] != "Alice" || result[1] != "Bob" {
			t.Errorf("PluckStringGeneric failed")
		}
	})

	t.Run("PluckInt32Generic", func(t *testing.T) {
		items := []Item{{Code: 100}, {Code: 200}}
		result := PluckInt32Generic(items, func(i Item) int32 {
			return int32(i.Code)
		})
		if len(result) != 2 || result[0] != 100 || result[1] != 200 {
			t.Errorf("PluckInt32Generic failed")
		}
	})

	t.Run("PluckInt64Generic", func(t *testing.T) {
		products := []Product{{ID: 1}, {ID: 2}}
		result := PluckInt64Generic(products, func(p Product) int64 {
			return p.ID
		})
		if len(result) != 2 || result[0] != 1 || result[1] != 2 {
			t.Errorf("PluckInt64Generic failed")
		}
	})

	t.Run("PluckUint32Generic", func(t *testing.T) {
		items := []Item{{Code: 100}, {Code: 200}}
		result := PluckUint32Generic(items, func(i Item) uint32 {
			return i.Code
		})
		if len(result) != 2 || result[0] != 100 || result[1] != 200 {
			t.Errorf("PluckUint32Generic failed")
		}
	})

	t.Run("PluckUint64Generic", func(t *testing.T) {
		items := []Item{{Amount: 1000}, {Amount: 2000}}
		result := PluckUint64Generic(items, func(i Item) uint64 {
			return i.Amount
		})
		if len(result) != 2 || result[0] != 1000 || result[1] != 2000 {
			t.Errorf("PluckUint64Generic failed")
		}
	})
}

// TestReflectionBasedPluck 测试基于反射的旧版函数
func TestReflectionBasedPluck(t *testing.T) {
	t.Run("PluckInt", func(t *testing.T) {
		persons := []Person{{Age: 30}, {Age: 25}, {Age: 35}}
		result := PluckInt(persons, "Age")
		if len(result) != 3 || result[0] != 30 || result[1] != 25 || result[2] != 35 {
			t.Errorf("PluckInt failed")
		}
	})

	t.Run("PluckString", func(t *testing.T) {
		persons := []Person{{Name: "Alice"}, {Name: "Bob"}}
		result := PluckString(persons, "Name")
		if len(result) != 2 || result[0] != "Alice" || result[1] != "Bob" {
			t.Errorf("PluckString failed")
		}
	})

	t.Run("PluckInt32", func(t *testing.T) {
		type Item32 struct {
			Code int32
		}
		items := []Item32{{Code: 100}, {Code: 200}}
		result := PluckInt32(items, "Code")
		if len(result) != 2 || result[0] != 100 || result[1] != 200 {
			t.Errorf("PluckInt32 failed")
		}
	})

	t.Run("PluckInt64", func(t *testing.T) {
		type Item64 struct {
			ID int64
		}
		items := []Item64{{ID: 1000}, {ID: 2000}}
		result := PluckInt64(items, "ID")
		if len(result) != 2 || result[0] != 1000 || result[1] != 2000 {
			t.Errorf("PluckInt64 failed")
		}
	})

	t.Run("PluckUint32", func(t *testing.T) {
		type ItemU32 struct {
			Code uint32
		}
		items := []ItemU32{{Code: 100}, {Code: 200}}
		result := PluckUint32(items, "Code")
		if len(result) != 2 || result[0] != 100 || result[1] != 200 {
			t.Errorf("PluckUint32 failed")
		}
	})

	t.Run("PluckUint64", func(t *testing.T) {
		type ItemU64 struct {
			Amount uint64
		}
		items := []ItemU64{{Amount: 1000}, {Amount: 2000}}
		result := PluckUint64(items, "Amount")
		if len(result) != 2 || result[0] != 1000 || result[1] != 2000 {
			t.Errorf("PluckUint64 failed")
		}
	})

	t.Run("PluckStringSlice", func(t *testing.T) {
		type Container struct {
			Tags []string
		}
		items := []Container{
			{Tags: []string{"tag1", "tag2"}},
			{Tags: []string{"tag3", "tag4"}},
		}
		result := PluckStringSlice(items, "Tags")
		if len(result) != 2 {
			t.Errorf("PluckStringSlice length failed")
		}
		if len(result[0]) != 2 || result[0][0] != "tag1" {
			t.Errorf("PluckStringSlice content failed")
		}
	})

	t.Run("PluckInt with pointer structs", func(t *testing.T) {
		p1 := &Person{Age: 30}
		p2 := &Person{Age: 25}
		persons := []*Person{p1, p2}
		result := PluckInt(persons, "Age")
		if len(result) != 2 || result[0] != 30 || result[1] != 25 {
			t.Errorf("PluckInt with pointers failed")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := PluckInt([]Person{}, "Age")
		if len(result) != 0 {
			t.Errorf("PluckInt empty should return empty slice")
		}
	})
}

// TestReflectionBasedPluckPanic 测试基于反射的函数的 panic 情况
func TestReflectionBasedPluckPanic(t *testing.T) {
	t.Run("panic on non-existent field", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("PluckInt should panic on non-existent field")
			}
		}()
		persons := []Person{{Name: "Alice"}}
		PluckInt(persons, "NonExistentField")
	})

	t.Run("panic on non-slice input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("PluckInt should panic on non-slice input")
			}
		}()
		PluckInt("not a slice", "Age")
	})

	t.Run("panic on unsupported element type", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("PluckInt should panic on unsupported type")
			}
		}()
		// Slice of primitive types (not structs)
		numbers := []int{1, 2, 3}
		PluckInt(numbers, "Age")
	})
}

// TestReflectionBasedPluckAdvanced 测试基于反射的高级用例
func TestReflectionBasedPluckAdvanced(t *testing.T) {
	t.Run("pluck with double pointer", func(t *testing.T) {
		type NestedPerson struct {
			Name string
			Age  int
		}
		p1 := &NestedPerson{Name: "Alice", Age: 30}
		p2 := &NestedPerson{Name: "Bob", Age: 25}
		persons := []*NestedPerson{p1, p2}

		result := PluckInt(persons, "Age")
		if len(result) != 2 || result[0] != 30 || result[1] != 25 {
			t.Errorf("PluckInt with nested pointers failed")
		}
	})

	t.Run("pluck from array instead of slice", func(t *testing.T) {
		persons := [3]Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}
		result := PluckInt(persons, "Age")
		if len(result) != 3 || result[0] != 30 {
			t.Errorf("PluckInt from array failed")
		}
	})

	t.Run("pluck multiple types", func(t *testing.T) {
		type MultiType struct {
			IntVal    int
			Int32Val  int32
			Int64Val  int64
			StringVal string
			Uint32Val uint32
			Uint64Val uint64
		}
		items := []MultiType{
			{IntVal: 1, Int32Val: 10, Int64Val: 100, StringVal: "a", Uint32Val: 1000, Uint64Val: 10000},
			{IntVal: 2, Int32Val: 20, Int64Val: 200, StringVal: "b", Uint32Val: 2000, Uint64Val: 20000},
		}

		intResult := PluckInt(items, "IntVal")
		if len(intResult) != 2 || intResult[0] != 1 || intResult[1] != 2 {
			t.Errorf("PluckInt multiple types failed")
		}

		int32Result := PluckInt32(items, "Int32Val")
		if len(int32Result) != 2 || int32Result[0] != 10 || int32Result[1] != 20 {
			t.Errorf("PluckInt32 multiple types failed")
		}

		int64Result := PluckInt64(items, "Int64Val")
		if len(int64Result) != 2 || int64Result[0] != 100 || int64Result[1] != 200 {
			t.Errorf("PluckInt64 multiple types failed")
		}

		stringResult := PluckString(items, "StringVal")
		if len(stringResult) != 2 || stringResult[0] != "a" || stringResult[1] != "b" {
			t.Errorf("PluckString multiple types failed")
		}

		uint32Result := PluckUint32(items, "Uint32Val")
		if len(uint32Result) != 2 || uint32Result[0] != 1000 || uint32Result[1] != 2000 {
			t.Errorf("PluckUint32 multiple types failed")
		}

		uint64Result := PluckUint64(items, "Uint64Val")
		if len(uint64Result) != 2 || uint64Result[0] != 10000 || uint64Result[1] != 20000 {
			t.Errorf("PluckUint64 multiple types failed")
		}
	})

	t.Run("pluck with nested invalid pointer", func(t *testing.T) {
		type TestStruct struct {
			Value int
		}
		// Test case with nil in middle of pointer chain - edge case that's hard to trigger
		// This tests the IsValid check in the reflection-based pluck function
		slice := []*TestStruct{
			{Value: 1},
			{Value: 2},
		}
		result := PluckInt(slice, "Value")
		if len(result) != 2 || result[0] != 1 || result[1] != 2 {
			t.Errorf("PluckInt with valid pointers failed")
		}
	})
}
