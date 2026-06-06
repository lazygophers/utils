package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAllStringValidators verifies all 24 string tags are functional
func TestAllStringValidators(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	t.Run("alpha", func(t *testing.T) {
		assert.NoError(t, v.Var("abcDEF", "alpha"))
		assert.Error(t, v.Var("abc123", "alpha"))
	})

	t.Run("alphaspace", func(t *testing.T) {
		assert.NoError(t, v.Var("abc DEF", "alphaspace"))
		assert.Error(t, v.Var("abc123", "alphaspace"))
	})

	t.Run("alphanum", func(t *testing.T) {
		assert.NoError(t, v.Var("abc123", "alphanum"))
		assert.Error(t, v.Var("abc-123", "alphanum"))
	})

	t.Run("alphanumspace", func(t *testing.T) {
		assert.NoError(t, v.Var("abc 123", "alphanumspace"))
		assert.Error(t, v.Var("abc-123", "alphanumspace"))
	})

	t.Run("alphanumunicode", func(t *testing.T) {
		assert.NoError(t, v.Var("abc123你好", "alphanumunicode"))
		assert.Error(t, v.Var("abc-123", "alphanumunicode"))
	})

	t.Run("alphaunicode", func(t *testing.T) {
		assert.NoError(t, v.Var("abc你好", "alphaunicode"))
		assert.Error(t, v.Var("abc123", "alphaunicode"))
	})

	t.Run("ascii", func(t *testing.T) {
		assert.NoError(t, v.Var("hello world!", "ascii"))
		assert.Error(t, v.Var("hello你好", "ascii"))
	})

	t.Run("boolean", func(t *testing.T) {
		assert.NoError(t, v.Var("true", "boolean"))
		assert.NoError(t, v.Var("false", "boolean"))
		assert.NoError(t, v.Var("1", "boolean"))
		assert.NoError(t, v.Var("0", "boolean"))
		assert.Error(t, v.Var("maybe", "boolean"))
	})

	t.Run("contains", func(t *testing.T) {
		assert.NoError(t, v.Var("hello world", "contains=world"))
		assert.Error(t, v.Var("hello", "contains=world"))
	})

	t.Run("containsany", func(t *testing.T) {
		assert.NoError(t, v.Var("hello", "containsany=lo"))
		assert.Error(t, v.Var("abc", "containsany=xyz"))
	})

	t.Run("containsrune", func(t *testing.T) {
		assert.NoError(t, v.Var("café", "containsrune=é"))
		assert.Error(t, v.Var("cafe", "containsrune=é"))
	})

	t.Run("endswith", func(t *testing.T) {
		assert.NoError(t, v.Var("hello world", "endswith=world"))
		assert.Error(t, v.Var("hello", "endswith=world"))
	})

	t.Run("endsnotwith", func(t *testing.T) {
		assert.NoError(t, v.Var("hello", "endsnotwith=world"))
		assert.Error(t, v.Var("hello world", "endsnotwith=world"))
	})

	t.Run("excludes", func(t *testing.T) {
		assert.NoError(t, v.Var("hello", "excludes=world"))
		assert.Error(t, v.Var("hello world", "excludes=world"))
	})

	t.Run("excludesall", func(t *testing.T) {
		assert.NoError(t, v.Var("abc", "excludesall=xyz"))
		assert.Error(t, v.Var("axyz", "excludesall=xyz"))
	})

	t.Run("excludesrune", func(t *testing.T) {
		assert.NoError(t, v.Var("cafe", "excludesrune=é"))
		assert.Error(t, v.Var("café", "excludesrune=é"))
	})

	t.Run("lowercase", func(t *testing.T) {
		assert.NoError(t, v.Var("abc", "lowercase"))
		assert.Error(t, v.Var("ABC", "lowercase"))
	})

	t.Run("multibyte", func(t *testing.T) {
		assert.NoError(t, v.Var("你好", "multibyte"))
		assert.Error(t, v.Var("abc", "multibyte"))
	})

	t.Run("number", func(t *testing.T) {
		assert.NoError(t, v.Var("123", "number"))
		assert.Error(t, v.Var("abc", "number"))
	})

	t.Run("numeric", func(t *testing.T) {
		assert.NoError(t, v.Var("123.45", "numeric"))
		assert.Error(t, v.Var("abc", "numeric"))
	})

	t.Run("printascii", func(t *testing.T) {
		assert.NoError(t, v.Var("hello world!", "printascii"))
		assert.Error(t, v.Var("hello\tworld", "printascii"))
	})

	t.Run("startswith", func(t *testing.T) {
		assert.NoError(t, v.Var("hello world", "startswith=hello"))
		assert.Error(t, v.Var("world hello", "startswith=hello"))
	})

	t.Run("startsnotwith", func(t *testing.T) {
		assert.NoError(t, v.Var("world hello", "startsnotwith=hello"))
		assert.Error(t, v.Var("hello world", "startsnotwith=hello"))
	})

	t.Run("uppercase", func(t *testing.T) {
		assert.NoError(t, v.Var("ABC", "uppercase"))
		assert.Error(t, v.Var("abc", "uppercase"))
	})
}
