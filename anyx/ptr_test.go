package anyx_test

import (
	"testing"

	"github.com/lazygophers/utils/anyx"
	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := 42
		p := anyx.ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("string", func(t *testing.T) {
		v := "hello"
		p := anyx.ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("bool", func(t *testing.T) {
		v := true
		p := anyx.ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("zero-value-int", func(t *testing.T) {
		v := 0
		p := anyx.ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("struct", func(t *testing.T) {
		type myStruct struct {
			Field int
		}
		v := myStruct{Field: 100}
		p := anyx.ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})

	t.Run("slice", func(t *testing.T) {
		v := []int{1, 2, 3}
		p := anyx.ToPtr(v)
		assert.NotNil(t, p)
		assert.Equal(t, v, *p)
	})
}
