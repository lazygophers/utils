package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("key by name", func(t *testing.T) {
		people := []Person{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}
		result := KeyBy(people, "Name")
		resultMap := result.(map[string]Person)
		assert.Len(t, resultMap, 3)
		assert.Equal(t, Person{Name: "Alice", Age: 30}, resultMap["Alice"])
		assert.Equal(t, Person{Name: "Bob", Age: 25}, resultMap["Bob"])
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []Person{}
		result := KeyBy(people, "Name")
		resultMap := result.(map[string]Person)
		assert.Empty(t, resultMap)
	})
}

func TestKeyByString(t *testing.T) {
	type Product struct {
		ID   string
		Name string
	}

	t.Run("basic key by string", func(t *testing.T) {
		products := []*Product{
			{ID: "p1", Name: "Product 1"},
			{ID: "p2", Name: "Product 2"},
		}
		result := KeyByString(products, "ID")
		assert.Len(t, result, 2)
		assert.Equal(t, "Product 1", result["p1"].Name)
	})
}

func TestKeyByInt32(t *testing.T) {
	type Item struct {
		ID   int32
		Name string
	}

	t.Run("basic key by int32", func(t *testing.T) {
		items := []*Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
		}
		result := KeyByInt32(items, "ID")
		assert.Len(t, result, 2)
		assert.Equal(t, "Item 1", result[1].Name)
	})
}

func TestKeyByInt64(t *testing.T) {
	type Item struct {
		ID   int64
		Name string
	}

	t.Run("basic key by int64", func(t *testing.T) {
		items := []*Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
		}
		result := KeyByInt64(items, "ID")
		assert.Len(t, result, 2)
		assert.Equal(t, "Item 1", result[1].Name)
	})
}

func TestKeyByUint64(t *testing.T) {
	type Item struct {
		ID   uint64
		Name string
	}

	t.Run("basic key by uint64", func(t *testing.T) {
		items := []*Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
		}
		result := KeyByUint64(items, "ID")
		assert.Len(t, result, 2)
		assert.Equal(t, "Item 1", result[1].Name)
	})
}

func TestKeyByGeneric(t *testing.T) {
	type User struct {
		Email string
		Name  string
	}

	t.Run("basic key by generic", func(t *testing.T) {
		users := []User{
			{Email: "alice@example.com", Name: "Alice"},
			{Email: "bob@example.com", Name: "Bob"},
		}
		result := KeyByGeneric(users, func(u User) string {
			return u.Email
		})
		assert.Len(t, result, 2)
		assert.Equal(t, "Alice", result["alice@example.com"].Name)
	})
}

func TestKeyByPtr(t *testing.T) {
	type Record struct {
		ID   int
		Data string
	}

	t.Run("basic key by ptr", func(t *testing.T) {
		records := []*Record{
			{ID: 1, Data: "Data 1"},
			{ID: 2, Data: "Data 2"},
		}
		result := KeyByPtr(records, func(r *Record) int {
			return r.ID
		})
		assert.Len(t, result, 2)
		assert.Equal(t, "Data 1", result[1].Data)
	})
}
