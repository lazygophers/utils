package i18n

import (
	"errors"
	"testing"
)

func TestLocalizerBuiltins(t *testing.T) {
	for _, ext := range []string{"json", "yaml", "yml", "toml", ".json", ".YAML"} {
		if _, ok := GetLocalizer(ext); !ok {
			t.Errorf("GetLocalizer(%q) missing", ext)
		}
	}
	if _, ok := GetLocalizer("xml"); ok {
		t.Errorf("xml should not be registered")
	}
}

func TestLocalizerJsonYamlToml(t *testing.T) {
	type localizerCase struct {
		ext  string
		body []byte
	}
	cases := []localizerCase{
		{"json", []byte(`{"k":"v","n":{"a":"b"}}`)},
		{"yaml", []byte("k: v\nn:\n  a: b\n")},
		{"toml", []byte("k = \"v\"\n[n]\na = \"b\"\n")},
	}
	for _, c := range cases {
		loc, ok := GetLocalizer(c.ext)
		if !ok {
			t.Fatalf("no localizer for %s", c.ext)
		}
		var m map[string]any
		err := loc.Unmarshal(c.body, &m)
		if err != nil {
			t.Fatalf("%s unmarshal err: %v", c.ext, err)
		}
		if m["k"] != "v" {
			t.Errorf("%s: k=%v", c.ext, m["k"])
		}
	}
}

func TestRegisterLocalizer(t *testing.T) {
	custom := NewLocalizerHandle(func(b []byte, v any) error {
		if len(b) == 0 {
			return errors.New("empty")
		}
		mm, ok := v.(*map[string]any)
		if !ok {
			return errors.New("bad type")
		}
		*mm = map[string]any{"custom": string(b)}
		return nil
	})
	RegisterLocalizer("xyz", custom)
	loc, ok := GetLocalizer("xyz")
	if !ok {
		t.Fatal("xyz not registered")
	}
	var m map[string]any
	err := loc.Unmarshal([]byte("hello"), &m)
	if err != nil {
		t.Fatal(err)
	}
	if m["custom"] != "hello" {
		t.Errorf("custom=%v", m["custom"])
	}
}

func TestNormExt(t *testing.T) {
	for _, c := range []struct{ in, want string }{
		{"json", "json"},
		{".json", "json"},
		{".JSON", "json"},
		{"", ""},
	} {
		if got := normExt(c.in); got != c.want {
			t.Errorf("normExt(%q)=%q want %q", c.in, got, c.want)
		}
	}
}

func TestNewLocalizerHandleError(t *testing.T) {
	h := NewLocalizerHandle(func(b []byte, v any) error {
		return errors.New("boom")
	})
	err := h.Unmarshal(nil, nil)
	if err == nil || err.Error() != "boom" {
		t.Fatalf("err=%v", err)
	}
}
