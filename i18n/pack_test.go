package i18n

import (
	"testing"

	"github.com/lazygophers/utils/language"
)

func TestPackRegister(t *testing.T) {
	p := NewPack(language.Make("en"))
	if got := p.Tag().String(); got != "en" {
		t.Fatalf("Tag()=%q want en", got)
	}

	p.Register("hello", "world")
	if v, ok := p.Get("hello"); !ok || v != "world" {
		t.Fatalf("Get(hello)=%q,%v want world,true", v, ok)
	}

	// 覆盖
	p.Register("hello", "earth")
	if v, _ := p.Get("hello"); v != "earth" {
		t.Fatalf("override fail v=%q", v)
	}

	if _, ok := p.Get("missing"); ok {
		t.Fatalf("missing should return false")
	}
}

func TestPackRegisterBatchFlat(t *testing.T) {
	p := NewPack(language.Make("en"))
	p.RegisterBatch(map[string]any{
		"a": "A",
		"b": 42,
		"c": int64(64),
		"d": 3.14,
		"e": true,
		"f": nil,
	})
	cases := map[string]string{
		"a": "A",
		"b": "42",
		"c": "64",
		"d": "3.14",
		"e": "true",
		"f": "",
	}
	for k, want := range cases {
		if v, _ := p.Get(k); v != want {
			t.Errorf("Get(%q)=%q want %q", k, v, want)
		}
	}
}

func TestPackRegisterBatchNested(t *testing.T) {
	p := NewPack(language.Make("en"))
	p.RegisterBatch(map[string]any{
		"user": map[string]any{
			"name":  "Alice",
			"age":   30,
			"addr":  map[string]any{"city": "Beijing"},
		},
		"any_keyed": map[any]any{
			"k1": "v1",
			2:    "v2",
			3.5:  "v3",
			true: "v4",
		},
	})

	checks := map[string]string{
		"user.name":      "Alice",
		"user.age":       "30",
		"user.addr.city": "Beijing",
		"any_keyed.k1":   "v1",
		"any_keyed.2":    "v2",
		"any_keyed.3.5":  "v3",
		"any_keyed.true": "v4",
	}
	for k, want := range checks {
		if v, ok := p.Get(k); !ok || v != want {
			t.Errorf("Get(%q)=%q,%v want %q", k, v, ok, want)
		}
	}
}

func TestPackRegisterBatchOverride(t *testing.T) {
	p := NewPack(language.Make("en"))
	p.RegisterBatch(map[string]any{"k": "v1"})
	p.RegisterBatch(map[string]any{"k": "v2"})
	if v, _ := p.Get("k"); v != "v2" {
		t.Fatalf("override fail v=%q", v)
	}
}

func TestPackAll(t *testing.T) {
	p := NewPack(language.Make("en"))
	p.Register("a", "1")
	p.Register("b", "2")

	got := map[string]string{}
	for k, v := range p.All() {
		got[k] = v
	}
	if len(got) != 2 || got["a"] != "1" || got["b"] != "2" {
		t.Fatalf("All snapshot wrong: %v", got)
	}

	// All 返回快照，遍历过程中可安全写
	count := 0
	for range p.All() {
		count++
		p.Register("c", "3")
	}
	if count != 2 {
		t.Fatalf("snapshot iter count=%d want 2", count)
	}
}

func TestPackAllEarlyStop(t *testing.T) {
	p := NewPack(language.Make("en"))
	p.Register("a", "1")
	p.Register("b", "2")
	p.Register("c", "3")

	visited := 0
	for range p.All() {
		visited++
		if visited == 1 {
			break
		}
	}
	if visited != 1 {
		t.Fatalf("early stop fail: %d", visited)
	}
}

func TestScalarToString(t *testing.T) {
	cases := []struct {
		in   any
		want string
	}{
		{"s", "s"},
		{42, "42"},
		{int64(64), "64"},
		{3.14, "3.14"},
		{true, "true"},
		{nil, ""},
		{[]int{1, 2}, "[1 2]"}, // fmt.Sprint fallback
	}
	for _, c := range cases {
		if got := scalarToString(c.in); got != c.want {
			t.Errorf("scalarToString(%v)=%q want %q", c.in, got, c.want)
		}
	}
}

