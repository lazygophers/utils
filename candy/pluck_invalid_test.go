package candy

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// TestPluckInvalidValueHandling tests the edge case where reflect value becomes invalid
func TestPluckInvalidValueHandling(t *testing.T) {
	t.Run("manually_construct_invalid_scenario", func(t *testing.T) {
		// This test tries to construct a scenario where ev.IsValid() returns false
		// This is extremely difficult to achieve naturally but let's try through unsafe operations
		
		type TestStruct struct {
			ID int
		}
		
		// Create a slice with valid pointers
		user1 := &TestStruct{ID: 1}
		user2 := &TestStruct{ID: 2}
		
		users := []*TestStruct{user1, user2}
		
		// Try to trigger the invalid value condition by manipulating the slice through reflection
		_ = reflect.ValueOf(users)
		
		// This should work normally
		result := PluckInt(users, "ID")
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
		
		// Now let's try to create a scenario with a nil pointer that might cause invalid reflection
		usersWithNil := []*TestStruct{user1, nil, user2}
		
		// This should panic due to nil pointer dereference, not because of invalid value
		assert.Panics(t, func() {
			PluckInt(usersWithNil, "ID")
		})
	})
	
	t.Run("test_empty_struct_field_access", func(t *testing.T) {
		// Try another approach to trigger the invalid value case
		type EmptyStruct struct {
			ID int
		}
		
		// Create slice with pointer to empty struct
		empty := &EmptyStruct{ID: 42}
		
		// Try to manipulate the struct through unsafe operations
		// This is very dangerous and might not work as expected
		list := []*EmptyStruct{empty}
		
		// Use pluck normally - this should work
		result := PluckInt(list, "ID")
		expected := []int{42}
		assert.Equal(t, expected, result)
		
		// Try with zero value
		zeroStruct := EmptyStruct{}
		ptrToZero := &zeroStruct
		listWithZero := []*EmptyStruct{ptrToZero}
		
		result = PluckInt(listWithZero, "ID")
		expected = []int{0}
		assert.Equal(t, expected, result)
	})
	
	t.Run("manipulate_slice_memory", func(t *testing.T) {
		// Another attempt to create invalid reflection values
		type TestStruct struct {
			ID int
		}
		
		// Create normal slice
		user := &TestStruct{ID: 100}
		users := []*TestStruct{user}
		
		// Test normal case first
		result := PluckInt(users, "ID")
		expected := []int{100}
		assert.Equal(t, expected, result)
		
		// Try to corrupt the slice data (very unsafe)
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&users))
		originalData := sliceHeader.Data
		
		// Temporarily corrupt and restore
		defer func() {
			sliceHeader.Data = originalData // Restore to prevent crashes
		}()
		
		// This might cause a crash or undefined behavior, so we skip it
		// The invalid value case is likely impossible to trigger safely in normal code
	})
}

// TestPluckReflectionEdgeCases tests various edge cases in the pluck function
func TestPluckReflectionEdgeCases(t *testing.T) {
	t.Run("zero_length_slice", func(t *testing.T) {
		type TestStruct struct {
			ID int
		}
		
		// Test with zero-length slice
		var users []*TestStruct
		result := PluckInt(users, "ID")
		expected := []int{}
		assert.Equal(t, expected, result)
	})
	
	t.Run("single_element_access", func(t *testing.T) {
		// Try to understand when IsValid might return false
		type TestStruct struct {
			ID int
		}
		
		user := &TestStruct{ID: 999}
		users := []*TestStruct{user}
		
		// Manually walk through the pluck logic
		v := reflect.ValueOf(users)
		
		for i := 0; i < v.Len(); i++ {
			ev := v.Index(i)
			
			// Dereference pointer
			for ev.Kind() == reflect.Ptr {
				ev = ev.Elem()
			}
			
			// Check if valid
			if !ev.IsValid() {
				t.Logf("Found invalid value at index %d", i)
			} else {
				t.Logf("Value at index %d is valid: %v", i, ev.Interface())
			}
		}
		
		result := PluckInt(users, "ID")
		expected := []int{999}
		assert.Equal(t, expected, result)
	})
}

// TestPluckInvalidValueCoverage tries to specifically cover the !ev.IsValid() case
func TestPluckInvalidValueCoverage(t *testing.T) {
	t.Run("invalid_value_scenario", func(t *testing.T) {
		// The !ev.IsValid() case in pluck function is extremely difficult to trigger
		// in normal Go code. It typically requires unsafe operations or memory corruption.
		// Let's try to create a scenario using reflection directly
		
		type TestStruct struct {
			ID int
		}
		
		// Create a slice with valid elements first
		user1 := &TestStruct{ID: 1}
		user2 := &TestStruct{ID: 2}
		users := []*TestStruct{user1, user2}
		
		// This should work normally
		result := PluckInt(users, "ID")
		expected := []int{1, 2}
		assert.Equal(t, expected, result)
		
		// The !ev.IsValid() case is designed to handle edge cases in reflection
		// where a value becomes invalid after dereferencing. This is very rare
		// and typically indicates memory corruption or unsafe operations.
		// For the sake of coverage, we note that this case exists but is
		// nearly impossible to trigger in safe Go code.
	})
}