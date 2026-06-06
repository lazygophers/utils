package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInEmpty(t *testing.T) {
	fn := In()
	assert.False(t, fn(paramFL{field: reflect.ValueOf("x")}))
}

func TestInIntUnified(t *testing.T) {
	fn := In(1, 3, 5)
	assert.True(t, fn(paramFL{field: reflect.ValueOf(3)}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf(2)}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("3")}))
}

func TestInStringUnified(t *testing.T) {
	fn := In("red", "green", "blue")
	assert.True(t, fn(paramFL{field: reflect.ValueOf("green")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("yellow")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf(42)}))
}

func TestInFloatUnified(t *testing.T) {
	fn := In(1.1, 2.2, 3.3)
	assert.True(t, fn(paramFL{field: reflect.ValueOf(2.2)}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf(9.9)}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("2.2")}))
}

func TestInMixedTypes(t *testing.T) {
	fn := In("a", "b")
	// mixed type → linear compareFields with strings
	assert.True(t, fn(paramFL{field: reflect.ValueOf("a")}))
	assert.False(t, fn(paramFL{field: reflect.ValueOf("z")}))
}

func TestNotInEmpty(t *testing.T) {
	fn := NotIn()
	assert.True(t, fn(paramFL{field: reflect.ValueOf("x")}))
}

func TestNotInIntUnified(t *testing.T) {
	fn := NotIn(1, 3, 5)
	assert.False(t, fn(paramFL{field: reflect.ValueOf(3)}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(2)}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("3")}))
}

func TestNotInStringUnified(t *testing.T) {
	fn := NotIn("red", "green")
	assert.False(t, fn(paramFL{field: reflect.ValueOf("red")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("blue")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(42)}))
}

func TestNotInFloatUnified(t *testing.T) {
	fn := NotIn(1.1, 2.2)
	assert.False(t, fn(paramFL{field: reflect.ValueOf(1.1)}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf(3.3)}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("1.1")}))
}

func TestNotInMixedTypes(t *testing.T) {
	fn := NotIn("a", "b")
	assert.False(t, fn(paramFL{field: reflect.ValueOf("a")}))
	assert.True(t, fn(paramFL{field: reflect.ValueOf("z")}))
}
