package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLTFieldLTEField(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		Min int `validate:"ltfield=Max"`
		Max int
	}
	assert.NoError(t, v.Struct(S{Min: 3, Max: 10}))   // 3 < 10
	assert.Error(t, v.Struct(S{Min: 10, Max: 3}))     // 10 > 3
	assert.NoError(t, v.Struct(S{Min: 3, Max: 10}))   

	type S2 struct {
		A int `validate:"ltefield=B"`
		B int
	}
	assert.NoError(t, v2_Struct(v, S2{A: 5, B: 5}))   // 5 <= 5
	assert.NoError(t, v2_Struct(v, S2{A: 3, B: 5}))   // 3 <= 5
	assert.Error(t, v2_Struct(v, S2{A: 10, B: 5}))    // 10 > 5
}

func TestNECSField(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A string `validate:"necsfield=B"`
		B string
	}
	assert.NoError(t, v.Struct(S{A: "hello", B: "world"}))
	assert.Error(t, v.Struct(S{A: "same", B: "same"}))
}

func TestGTCSField(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A int `validate:"gtcsfield=B"`
		B int
	}
	assert.NoError(t, v.Struct(S{A: 10, B: 5}))
	assert.Error(t, v.Struct(S{A: 3, B: 5}))
	assert.Error(t, v.Struct(S{A: 5, B: 5}))
}

func TestGTECSField(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A int `validate:"gtecsfield=B"`
		B int
	}
	assert.NoError(t, v.Struct(S{A: 10, B: 5}))
	assert.NoError(t, v.Struct(S{A: 5, B: 5}))
	assert.Error(t, v.Struct(S{A: 3, B: 5}))
}

func TestLTCSField(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A int `validate:"ltcsfield=B"`
		B int
	}
	assert.NoError(t, v.Struct(S{A: 3, B: 5}))
	assert.Error(t, v.Struct(S{A: 10, B: 5}))
}

func TestLTECSField(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type S struct {
		A int `validate:"ltecsfield=B"`
		B int
	}
	assert.NoError(t, v.Struct(S{A: 3, B: 5}))
	assert.NoError(t, v.Struct(S{A: 5, B: 5}))
	assert.Error(t, v.Struct(S{A: 10, B: 5}))
}

func TestFieldContainsExcludesEdgeCases(t *testing.T) {
	v, err := New()
	assert.NoError(t, err)

	type SC struct {
		Name string `validate:"fieldcontains=ab"`
	}
	// fieldcontains checks if Name contains any of the chars 'a' or 'b'
	assert.NoError(t, v.Struct(SC{Name: "apple"}))
	assert.Error(t, v.Struct(SC{Name: "xyz"}))

	type SE struct {
		Name string `validate:"fieldexcludes=ab"`
	}
	assert.NoError(t, v.Struct(SE{Name: "xyz"}))
	assert.Error(t, v.Struct(SE{Name: "apple"}))

	// Empty param
	type SC2 struct {
		Name string `validate:"fieldcontains="`
	}
	assert.Error(t, v.Struct(SC2{Name: "test"}))

	type SE2 struct {
		Name string `validate:"fieldexcludes="`
	}
	assert.NoError(t, v.Struct(SE2{Name: "test"}))
}

func TestResolveFieldPathDotNotation(t *testing.T) {
	// Test the dot notation path resolution
	// Uses Top() to navigate nested structs
	type Inner struct {
		Value int
	}
	type Outer struct {
		A     int `validate:"gtcsfield=Inner.Value"`
		Inner Inner
	}
	v, err := New()
	assert.NoError(t, err)

	assert.NoError(t, v.Struct(Outer{A: 20, Inner: Inner{Value: 10}}))
	assert.Error(t, v.Struct(Outer{A: 5, Inner: Inner{Value: 10}}))
}

// helper to avoid name collision with existing tests
func v2_Struct(v *Validator, s interface{}) error {
	return v.Struct(s)
}
