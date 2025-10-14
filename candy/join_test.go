package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	t.Run("join integers with comma", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Join(input, ", ")
		expected := "1, 2, 3, 4, 5"
		assert.Equal(t, expected, result)
	})

	t.Run("join strings with space", func(t *testing.T) {
		input := []string{"hello", "world", "test"}
		result := Join(input, " ")
		expected := "hello world test"
		assert.Equal(t, expected, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"only"}
		result := Join(input, ", ")
		expected := "only"
		assert.Equal(t, expected, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []string{}
		result := Join(input, ", ")
		expected := ""
		assert.Equal(t, expected, result)
	})
}
