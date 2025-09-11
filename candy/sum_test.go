package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("positive integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Sum(input)
		assert.Equal(t, 15, result)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Sum(input)
		assert.Equal(t, 0, result)
	})

	t.Run("single element", func(t *testing.T) {
		input := []int{42}
		result := Sum(input)
		assert.Equal(t, 42, result)
	})

	t.Run("negative integers", func(t *testing.T) {
		input := []int{-1, -2, -3}
		result := Sum(input)
		assert.Equal(t, -6, result)
	})

	t.Run("mixed positive and negative", func(t *testing.T) {
		input := []int{-5, 10, -3, 8}
		result := Sum(input)
		assert.Equal(t, 10, result)
	})

	t.Run("zeros", func(t *testing.T) {
		input := []int{0, 0, 0}
		result := Sum(input)
		assert.Equal(t, 0, result)
	})

	t.Run("float64 slice", func(t *testing.T) {
		input := []float64{1.5, 2.5, 3.0}
		result := Sum(input)
		assert.Equal(t, 7.0, result)
	})

	t.Run("float32 slice", func(t *testing.T) {
		input := []float32{1.1, 2.2, 3.3}
		result := Sum(input)
		assert.InDelta(t, 6.6, result, 1e-6)
	})

	t.Run("int32 slice", func(t *testing.T) {
		input := []int32{10, 20, 30}
		result := Sum(input)
		assert.Equal(t, int32(60), result)
	})

	t.Run("int64 slice", func(t *testing.T) {
		input := []int64{100, 200, 300}
		result := Sum(input)
		assert.Equal(t, int64(600), result)
	})

	t.Run("uint slice", func(t *testing.T) {
		input := []uint{1, 2, 3}
		result := Sum(input)
		assert.Equal(t, uint(6), result)
	})

	t.Run("large numbers", func(t *testing.T) {
		input := []int64{1000000, 2000000, 3000000}
		result := Sum(input)
		assert.Equal(t, int64(6000000), result)
	})
}