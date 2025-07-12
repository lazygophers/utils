package anyx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluckInt(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	t.Run("正常情况", func(t *testing.T) {
		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		}
		ids := PluckInt(users, "ID")
		assert.Equal(t, []int{1, 2}, ids)
	})

	t.Run("空切片", func(t *testing.T) {
		empty := []User{}
		ids := PluckInt(empty, "ID")
		assert.Empty(t, ids)
	})

	t.Run("字段不存在", func(t *testing.T) {
		users := []User{{ID: 1}}
		assert.Panics(t, func() {
			PluckInt(users, "InvalidField")
		})
	})
}

func TestPluckString(t *testing.T) {
	type Product struct {
		Name string
		SKU  string
	}

	t.Run("正常情况", func(t *testing.T) {
		products := []Product{
			{Name: "Laptop", SKU: "LP123"},
			{Name: "Phone", SKU: "PH456"},
		}
		names := PluckString(products, "Name")
		assert.Equal(t, []string{"Laptop", "Phone"}, names)
	})
}

func TestPluckStringSlice(t *testing.T) {
	type User struct {
		Name   string
		Emails []string
	}

	t.Run("正常情况", func(t *testing.T) {
		users := []User{
			{Name: "Alice", Emails: []string{"a@test.com", "a@work.com"}},
			{Name: "Bob", Emails: []string{"b@test.com"}},
		}
		emails := PluckStringSlice(users, "Emails")
		assert.Equal(t, [][]string{
			{"a@test.com", "a@work.com"},
			{"b@test.com"},
		}, emails)
	})
}

func TestDiffSlice(t *testing.T) {
	t.Run("正常差异", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{2, 3, 4}
		added, removed := DiffSlice(a, b)
		assert.Equal(t, []int{1}, added)
		assert.Equal(t, []int{4}, removed)
	})

	t.Run("完全相同", func(t *testing.T) {
		a := []string{"a", "b"}
		b := []string{"a", "b"}
		added, removed := DiffSlice(a, b)
		assert.Empty(t, added)
		assert.Empty(t, removed)
	})

	t.Run("空切片", func(t *testing.T) {
		a := []int{}
		b := []int{1}
		added, removed := DiffSlice(a, b)
		assert.Empty(t, added)
		assert.Equal(t, []int{1}, removed)
	})
}

func TestRemoveSlice(t *testing.T) {
	t.Run("正常移除", func(t *testing.T) {
		src := []int{1, 2, 3, 4}
		rm := []int{2, 4}
		result := RemoveSlice(src, rm)
		assert.Equal(t, []int{1, 3}, result)
	})

	t.Run("无匹配项", func(t *testing.T) {
		src := []string{"a", "b"}
		rm := []string{"c"}
		result := RemoveSlice(src, rm)
		assert.Equal(t, src, result)
	})
}

func TestKeyBy(t *testing.T) {
	type Product struct {
		ID   string
		Name string
	}

	t.Run("正常情况", func(t *testing.T) {
		products := []Product{
			{ID: "p1", Name: "Product 1"},
			{ID: "p2", Name: "Product 2"},
		}
		result := KeyBy(products, "ID").(map[string]Product)
		assert.Equal(t, "Product 1", result["p1"].Name)
		assert.Equal(t, "Product 2", result["p2"].Name)
	})

	t.Run("空切片", func(t *testing.T) {
		var empty []Product
		result := KeyBy(empty, "ID").(map[string]Product)
		assert.Empty(t, result)
	})
}

func TestKeyByUint64(t *testing.T) {
	type Item struct {
		ID   uint64
		Name string
	}

	t.Run("正常情况", func(t *testing.T) {
		items := []*Item{
			{ID: 101, Name: "Item 101"},
			{ID: 102, Name: "Item 102"},
		}
		result := KeyByUint64(items, "ID")
		assert.Equal(t, "Item 101", result[101].Name)
		assert.Equal(t, "Item 102", result[102].Name)
	})
}

func TestKeyByString(t *testing.T) {
	type User struct {
		Username string
		Email    string
	}

	t.Run("正常情况", func(t *testing.T) {
		users := []*User{
			{Username: "alice", Email: "alice@example.com"},
			{Username: "bob", Email: "bob@example.com"},
		}
		result := KeyByString(users, "Username")
		assert.Equal(t, "alice@example.com", result["alice"].Email)
		assert.Equal(t, "bob@example.com", result["bob"].Email)
	})
}

func TestSlice2Map(t *testing.T) {
	t.Run("正常转换", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Slice2Map(slice)
		assert.True(t, result[1])
		assert.True(t, result[2])
		assert.True(t, result[3])
		assert.False(t, result[4])
	})

	t.Run("空切片", func(t *testing.T) {
		empty := []string{}
		result := Slice2Map(empty)
		assert.Empty(t, result)
	})

	t.Run("字符串切片", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		result := Slice2Map(slice)
		assert.True(t, result["a"])
		assert.True(t, result["b"])
		assert.False(t, result["d"])
	})
}
