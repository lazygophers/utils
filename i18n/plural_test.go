package i18n

import (
	"testing"

	"github.com/lazygophers/utils/language"
)

func TestPluralCategoryEn(t *testing.T) {
	en := language.Make("en")
	if got := pluralCategoryFor(en, 1); got != CategoryOne {
		t.Fatalf("en n=1 want one, got %s", got)
	}
	for _, n := range []int{0, 2, 3, 100, -1} {
		if got := pluralCategoryFor(en, n); got != CategoryOther {
			t.Fatalf("en n=%d want other, got %s", n, got)
		}
	}
}

func TestPluralCategoryZh(t *testing.T) {
	zh := language.Make("zh")
	for _, n := range []int{0, 1, 2, 5, 100} {
		if got := pluralCategoryFor(zh, n); got != CategoryOther {
			t.Fatalf("zh n=%d want other, got %s", n, got)
		}
	}
}

func TestPluralCategoryFallbackChain(t *testing.T) {
	// zh-CN 应回退到 zh 规则
	zhCN := language.Make("zh-CN")
	if got := pluralCategoryFor(zhCN, 1); got != CategoryOther {
		t.Fatalf("zh-CN n=1 want other (via zh), got %s", got)
	}
	// en-US 应回退到 en 规则
	enUS := language.Make("en-US")
	if got := pluralCategoryFor(enUS, 1); got != CategoryOne {
		t.Fatalf("en-US n=1 want one (via en), got %s", got)
	}
}

func TestPluralCategoryUnknownLocale(t *testing.T) {
	// 未注册规则 → other (qqq 不在 pluralRules 且 base 也为 qqq)
	qqq := language.Make("qqq")
	if got := pluralCategoryFor(qqq, 1); got != CategoryOther {
		t.Fatalf("qqq n=1 want other, got %s", got)
	}
}

func TestBundleTPluralEn(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "items.one", "{n} item")
	b.Register(en, "items.other", "{n} items")

	language.Set(en)
	defer language.Del()

	if got := b.TPlural("items", 1); got != "1 item" {
		t.Fatalf("n=1 got %q", got)
	}
	if got := b.TPlural("items", 0); got != "0 items" {
		t.Fatalf("n=0 got %q", got)
	}
	if got := b.TPlural("items", 5); got != "5 items" {
		t.Fatalf("n=5 got %q", got)
	}
}

func TestBundleTPluralZh(t *testing.T) {
	b := New()
	zh := language.Make("zh")
	b.Register(zh, "items.other", "{n} 个项目")

	language.Set(zh)
	defer language.Del()

	if got := b.TPlural("items", 1); got != "1 个项目" {
		t.Fatalf("zh n=1 got %q", got)
	}
	if got := b.TPlural("items", 100); got != "100 个项目" {
		t.Fatalf("zh n=100 got %q", got)
	}
}

func TestBundleTPluralFallbackToOther(t *testing.T) {
	b := New()
	en := language.Make("en")
	// 仅注册 .other，缺 .one
	b.Register(en, "msg.other", "many: {n}")

	language.Set(en)
	defer language.Del()

	if got := b.TPlural("msg", 1); got != "many: 1" {
		t.Fatalf("missing .one should fall back to .other, got %q", got)
	}
}

func TestBundleTPluralFallbackToKey(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "raw", "raw text")

	language.Set(en)
	defer language.Del()

	if got := b.TPlural("raw", 2); got != "raw text" {
		t.Fatalf("category missing should fall back to plain key, got %q", got)
	}
}

func TestBundleTPluralMissingKey(t *testing.T) {
	b := New()
	language.Set(language.Make("en"))
	defer language.Del()

	if got := b.TPlural("nothing", 1); got != "nothing" {
		t.Fatalf("missing everything should return key, got %q", got)
	}
}

func TestBundleTPluralNamedArgs(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "greet.one", "Hi {name}, {n} message")
	b.Register(en, "greet.other", "Hi {name}, {n} messages")

	language.Set(en)
	defer language.Del()

	if got := b.TPlural("greet", 1, "name", "Alice"); got != "Hi Alice, 1 message" {
		t.Fatalf("named args n=1 got %q", got)
	}
	if got := b.TPlural("greet", 3, "name", "Bob"); got != "Hi Bob, 3 messages" {
		t.Fatalf("named args n=3 got %q", got)
	}
}

func TestBundleTPluralPositionalArgs(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "pos.one", "{0} one")
	b.Register(en, "pos.other", "{0} other")

	language.Set(en)
	defer language.Del()

	// 位置模式下 {n} 不可用，但 {0} 应工作
	if got := b.TPlural("pos", 1, "x"); got != "x one" {
		t.Fatalf("positional n=1 got %q", got)
	}
}

func TestTPluralGlobal(t *testing.T) {
	en := language.Make("en")
	Register(en, "g.one", "g{n}-one")
	Register(en, "g.other", "g{n}-other")

	language.Set(en)
	defer language.Del()

	if got := TPlural("g", 1); got != "g1-one" {
		t.Fatalf("global n=1 got %q", got)
	}
	if got := TPlural("g", 4); got != "g4-other" {
		t.Fatalf("global n=4 got %q", got)
	}
}

func TestBundleTPluralDefaultLocaleFallback(t *testing.T) {
	b := New()
	en := language.Make("en")
	b.Register(en, "x.one", "1 x")
	b.Register(en, "x.other", "n x")

	// 当前 goroutine 设为 en-GB，无该 locale 翻译 → 走 fallback 链找到 en
	language.Set(language.Make("en-GB"))
	defer language.Del()

	if got := b.TPlural("x", 1); got != "1 x" {
		t.Fatalf("default locale fallback n=1 got %q", got)
	}
	if got := b.TPlural("x", 2); got != "n x" {
		t.Fatalf("default locale fallback n=2 got %q", got)
	}
}

func TestMergePluralArgsEmpty(t *testing.T) {
	got := mergePluralArgs(7, nil)
	if len(got) != 2 || got[0] != "n" || got[1] != 7 {
		t.Fatalf("empty args merge got %v", got)
	}
}

func TestMergePluralArgsNamed(t *testing.T) {
	got := mergePluralArgs(3, []any{"name", "alice"})
	if len(got) != 4 || got[0] != "n" || got[1] != 3 || got[2] != "name" || got[3] != "alice" {
		t.Fatalf("named merge got %v", got)
	}
}

func TestMergePluralArgsPositional(t *testing.T) {
	in := []any{1, 2, 3}
	got := mergePluralArgs(9, in)
	if len(got) != 3 {
		t.Fatalf("positional should be unchanged, got %v", got)
	}
}
