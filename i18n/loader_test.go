package i18n

import (
	"testing"

	"github.com/lazygophers/utils/language"
)

func TestMessagesBuiltinEn(t *testing.T) {
	en := language.Make("en")
	language.Set(en)
	defer language.Del()

	if got := T("greeting", "name", "world"); got != "Hello, world!" {
		t.Fatalf("en greeting = %q, want Hello, world!", got)
	}
	if got := T("farewell"); got != "Goodbye" {
		t.Fatalf("en farewell = %q, want Goodbye", got)
	}
	if got := T("apple.one", "n", 1); got != "1 apple" {
		t.Fatalf("en apple.one = %q", got)
	}
	if got := T("apple.other", "n", 3); got != "3 apples" {
		t.Fatalf("en apple.other = %q", got)
	}
}

func TestMessagesBuiltinZh(t *testing.T) {
	zh := language.Make("zh")
	language.Set(zh)
	defer language.Del()

	if got := T("greeting", "name", "世界"); got != "你好，世界！" {
		t.Fatalf("zh greeting = %q", got)
	}
	if got := T("farewell"); got != "再见" {
		t.Fatalf("zh farewell = %q", got)
	}
	if got := T("apple.other", "n", 5); got != "5 个苹果" {
		t.Fatalf("zh apple.other = %q", got)
	}
}

func TestLoadSingleLocale(t *testing.T) {
	Load(map[string]map[string]string{
		"fr": {
			"greeting": "Bonjour, {name}!",
			"farewell": "Au revoir",
		},
	})

	fr := language.Make("fr")
	if got := TLocale(fr, "greeting", "name", "monde"); got != "Bonjour, monde!" {
		t.Fatalf("fr greeting = %q", got)
	}
	if got := TLocale(fr, "farewell"); got != "Au revoir" {
		t.Fatalf("fr farewell = %q", got)
	}
}

func TestLoadMultiLocale(t *testing.T) {
	Load(map[string]map[string]string{
		"de": {"hi": "Hallo"},
		"es": {"hi": "Hola"},
		"it": {"hi": "Ciao"},
	})

	cases := map[string]string{
		"de": "Hallo",
		"es": "Hola",
		"it": "Ciao",
	}
	for loc, want := range cases {
		tag := language.Make(loc)
		if got := TLocale(tag, "hi"); got != want {
			t.Fatalf("%s hi = %q, want %q", loc, got, want)
		}
	}
}

func TestLoadEmpty(t *testing.T) {
	Load(map[string]map[string]string{})
	Load(nil)
	// 不应 panic；无副作用断言。
}

func TestLoadOverridesExisting(t *testing.T) {
	Load(map[string]map[string]string{
		"pt": {"greeting": "Olá"},
	})
	pt := language.Make("pt")
	if got := TLocale(pt, "greeting"); got != "Olá" {
		t.Fatalf("pt initial = %q", got)
	}
	Load(map[string]map[string]string{
		"pt": {"greeting": "Oi"},
	})
	if got := TLocale(pt, "greeting"); got != "Oi" {
		t.Fatalf("pt overridden = %q", got)
	}
}
