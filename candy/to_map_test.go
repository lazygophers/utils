package candy

import (
	"testing"
)

func TestToMap(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := ToMap(nil); got != nil {
			t.Fatalf("expected nil, got %v", got)
		}
	})

	t.Run("json string", func(t *testing.T) {
		got := ToMap(`{"a":1,"b":"x"}`)
		if got["a"] == nil || got["b"] != "x" {
			t.Fatalf("unexpected: %v", got)
		}
	})

	t.Run("json bytes", func(t *testing.T) {
		got := ToMap([]byte(`{"a":1}`))
		if got["a"] == nil {
			t.Fatalf("unexpected: %v", got)
		}
	})

	t.Run("fallback map", func(t *testing.T) {
		got := ToMap(map[string]int{"a": 1})
		if got["a"] != 1 {
			t.Fatalf("unexpected: %v", got)
		}
	})
}

func TestToMapInt32String(t *testing.T) {
	t.Run("non-map", func(t *testing.T) {
		got := ToMapInt32String("x")
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %v", got)
		}
	})

	got := ToMapInt32String(map[any]any{int32(1): "a", "2": 3})
	if got[1] != "a" || got[2] != "3" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToMapInt64String(t *testing.T) {
	t.Run("non-map", func(t *testing.T) {
		got := ToMapInt64String("x")
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %v", got)
		}
	})

	got := ToMapInt64String(map[any]any{int64(1): "a", "2": 3})
	if got[1] != "a" || got[2] != "3" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToMapStringAny(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := ToMapStringAny(nil); got != nil {
			t.Fatalf("expected nil, got %v", got)
		}
	})

	t.Run("non-map", func(t *testing.T) {
		got := ToMapStringAny("x")
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %v", got)
		}
	})

	got := ToMapStringAny(map[any]any{1: "a", "b": 2})
	if got["1"] != "a" || got["b"] != 2 {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToMapStringArrayString(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := ToMapStringArrayString(nil); got != nil {
			t.Fatalf("expected nil, got %v", got)
		}
	})

	t.Run("panic non-map", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Fatalf("expected panic")
			}
		}()
		ToMapStringArrayString("x")
	})

	got := ToMapStringArrayString(map[any]any{"a": "x,y", 1: []any{"p", 2}})
	if len(got["a"]) != 2 || got["a"][0] != "x" || got["1"][0] != "p" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToMapStringInt64(t *testing.T) {
	t.Run("non-map", func(t *testing.T) {
		got := ToMapStringInt64("x")
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %v", got)
		}
	})

	got := ToMapStringInt64(map[any]any{1: "2", "b": 3})
	if got["1"] != 2 || got["b"] != 3 {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestToMapStringString(t *testing.T) {
	t.Run("non-map", func(t *testing.T) {
		got := ToMapStringString("x")
		if len(got) != 0 {
			t.Fatalf("expected empty map, got %v", got)
		}
	})

	got := ToMapStringString(map[any]any{1: 2, "b": "x"})
	if got["1"] != "2" || got["b"] != "x" {
		t.Fatalf("unexpected: %v", got)
	}
}
