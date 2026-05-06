package candy

import (
	"fmt"
	"testing"
)

type TestPerson struct {
	ID   int
	Name string
	Age  int
}

func TestSliceField2MapOptimized(t *testing.T) {
	people := []TestPerson{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	// Test String
	nameMap := SliceField2MapString(people, "Name")
	if len(nameMap) != 3 || !nameMap["Alice"] {
		t.Errorf("SliceField2MapString failed")
	}

	// Test Int
	idMap := SliceField2MapInt(people, "ID")
	if len(idMap) != 3 || !idMap[1] {
		t.Errorf("SliceField2MapInt failed")
	}

	// Test empty slice
	emptyPeople := []TestPerson{}
	emptyMap := SliceField2MapString(emptyPeople, "Name")
	if emptyMap != nil {
		t.Errorf("Expected nil for empty slice")
	}

	fmt.Println("✓ All optimized functions work correctly")
}

// 简单性能对比
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
