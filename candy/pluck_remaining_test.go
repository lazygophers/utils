package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type PersonWithID struct {
	ID   int
	Name string
	Tags []string
}

func TestPluckInt(t *testing.T) {
	t.Run("basic pluck int", func(t *testing.T) {
		people := []PersonWithID{
			{ID: 1, Name: "Alice", Tags: []string{"admin"}},
			{ID: 2, Name: "Bob", Tags: []string{"user"}},
			{ID: 3, Name: "Charlie", Tags: []string{"guest"}},
		}
		result := PluckInt(people, "ID")
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []PersonWithID{}
		result := PluckInt(people, "ID")
		expected := []int{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-int field should panic", func(t *testing.T) {
		people := []PersonWithID{
			{ID: 1, Name: "Alice", Tags: []string{"admin"}},
		}
		assert.Panics(t, func() {
			PluckInt(people, "Name") // Name is string, not int - should panic
		})
	})
}

func TestPluckString(t *testing.T) {
	t.Run("basic pluck string", func(t *testing.T) {
		people := []PersonWithID{
			{ID: 1, Name: "Alice", Tags: []string{"admin"}},
			{ID: 2, Name: "Bob", Tags: []string{"user"}},
			{ID: 3, Name: "Charlie", Tags: []string{"guest"}},
		}
		result := PluckString(people, "Name")
		expected := []string{"Alice", "Bob", "Charlie"}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []PersonWithID{}
		result := PluckString(people, "Name")
		expected := []string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-string field should panic", func(t *testing.T) {
		people := []PersonWithID{
			{ID: 1, Name: "Alice", Tags: []string{"admin"}},
		}
		assert.Panics(t, func() {
			PluckString(people, "ID") // ID is int, not string - should panic
		})
	})
}

func TestPluckStringSlice(t *testing.T) {
	t.Run("basic pluck string slice", func(t *testing.T) {
		people := []PersonWithID{
			{ID: 1, Name: "Alice", Tags: []string{"admin", "user"}},
			{ID: 2, Name: "Bob", Tags: []string{"user"}},
			{ID: 3, Name: "Charlie", Tags: []string{"guest", "temp"}},
		}
		result := PluckStringSlice(people, "Tags")
		expected := [][]string{{"admin", "user"}, {"user"}, {"guest", "temp"}}
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []PersonWithID{}
		result := PluckStringSlice(people, "Tags")
		expected := [][]string{}
		assert.Equal(t, expected, result)
	})

	t.Run("non-string-slice field should panic", func(t *testing.T) {
		people := []PersonWithID{
			{ID: 1, Name: "Alice", Tags: []string{"admin"}},
		}
		assert.Panics(t, func() {
			PluckStringSlice(people, "Name") // Name is string, not []string - should panic
		})
	})
}