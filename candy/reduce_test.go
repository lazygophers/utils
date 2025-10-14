package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Run("sum integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reduce(input, func(acc, item int) int {
			return acc + item
		})
		assert.Equal(t, 15, result)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		input := []string{"hello", " ", "world"}
		result := Reduce(input, func(acc, item string) string {
			return acc + item
		})
		assert.Equal(t, "hello world", result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Reduce(input, func(acc, item int) int {
			return acc + item
		})
		assert.Equal(t, 42, result)
	})
}
