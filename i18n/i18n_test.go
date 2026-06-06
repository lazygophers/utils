package i18n

import (
	"strings"
	"sync"
	"testing"
	"text/template"

	"github.com/lazygophers/utils/language"
)

func TestI18nRegisterAndLocalize(t *testing.T) {
	p := New()
	en := language.Make("en")
	zh := language.Make("zh")

	p.Register(en, "hello", "Hello")
	p.Register(zh, "hello", "你好")

	if got := p.LocalizeWithLang(en, "hello"); got != "Hello" {
		t.Errorf("en hello=%q", got)
	}
	if got := p.LocalizeWithLang(zh, "hello"); got != "你好" {
		t.Errorf("zh hello=%q", got)
	}
}

func TestI18nFallbackChain(t *testing.T) {
	p := New(WithDefaultLang(language.Make("en")))
	p.Register(language.Make("zh"), "hello", "你好")
	p.Register(language.Make("en"), "bye", "Bye")

	// zh-CN → 命中 zh
	if got := p.LocalizeWithLang(language.Make("zh-CN"), "hello"); got != "你好" {
		t.Errorf("zh-CN→zh fallback failed: %q", got)
	}
	// fr → fallback 到 default(en)
	if got := p.LocalizeWithLang(language.Make("fr"), "bye"); got != "Bye" {
		t.Errorf("fr→en fallback failed: %q", got)
	}
	// 全部 miss 返回 key
	if got := p.LocalizeWithLang(language.Make("fr"), "missing"); got != "missing" {
		t.Errorf("missing key should return key, got %q", got)
	}
}

func TestI18nFallbackDefaultRegion(t *testing.T) {
	p := New(WithDefaultLang(language.Make("en-US")))
	p.Register(language.Make("en"), "k", "v")
	// fr 未注册 → defaultLang en-US 未注册 → en-US.base en 命中
	if got := p.LocalizeWithLang(language.Make("fr"), "k"); got != "v" {
		t.Errorf("en-US→en fallback failed: %q", got)
	}
}

func TestI18nNilLangUsesDefault(t *testing.T) {
	p := New(WithDefaultLang(language.Make("en")))
	p.Register(language.Make("en"), "k", "v")
	if got := p.LocalizeWithLang(nil, "k"); got != "v" {
		t.Errorf("nil lang should use default: %q", got)
	}
}

func TestI18nTemplate(t *testing.T) {
	p := New()
	en := language.Make("en")
	p.Register(en, "greet", "Hello {{.Name}}, you have {{.Count}} msg")
	got := p.LocalizeWithLang(en, "greet", map[string]any{"Name": "Alice", "Count": 3})
	if got != "Hello Alice, you have 3 msg" {
		t.Errorf("template=%q", got)
	}
}

func TestI18nTemplateBadParse(t *testing.T) {
	p := New()
	en := language.Make("en")
	// 解析失败 → 返回原 value
	p.Register(en, "bad", "Hello {{.Name")
	got := p.LocalizeWithLang(en, "bad", map[string]any{"Name": "A"})
	if got != "Hello {{.Name" {
		t.Errorf("bad parse should return raw value: %q", got)
	}
}

func TestI18nTemplateBadExec(t *testing.T) {
	p := New()
	en := language.Make("en")
	p.Register(en, "x", "{{call .Fn}}")
	// .Fn 不存在 → 执行失败 → 返回原 value
	got := p.LocalizeWithLang(en, "x", map[string]any{})
	if got != "{{call .Fn}}" {
		t.Errorf("bad exec should return raw value: %q", got)
	}
}

func TestI18nAddTemplateFunc(t *testing.T) {
	p := New()
	p.AddTemplateFunc("upper", strings.ToUpper)
	p.Register(language.Make("en"), "k", `{{upper .Name}}`)
	got := p.LocalizeWithLang(language.Make("en"), "k", map[string]any{"Name": "abc"})
	if got != "ABC" {
		t.Errorf("AddTemplateFunc fail: %q", got)
	}
}

func TestWithTemplateFuncs(t *testing.T) {
	p := New(WithTemplateFuncs(template.FuncMap{
		"shout": func(s string) string { return s + "!" },
	}))
	p.Register(language.Make("en"), "k", `{{shout .Name}}`)
	got := p.LocalizeWithLang(language.Make("en"), "k", map[string]any{"Name": "hi"})
	if got != "hi!" {
		t.Errorf("WithTemplateFuncs fail: %q", got)
	}
}

func TestI18nSetDefaultLang(t *testing.T) {
	p := New()
	tag := language.Make("zh")
	if p.SetDefaultLang(tag) != p {
		t.Error("SetDefaultLang should return self")
	}
	if p.DefaultLang().String() != "zh" {
		t.Errorf("DefaultLang=%q", p.DefaultLang().String())
	}
}

func TestI18nRegisterBatch(t *testing.T) {
	p := New()
	en := language.Make("en")
	p.RegisterBatch(en, map[string]any{
		"a": "A",
		"nested": map[string]any{
			"b": "B",
		},
	})
	if got := p.LocalizeWithLang(en, "nested.b"); got != "B" {
		t.Errorf("nested=%q", got)
	}
}

func TestI18nLocalizeUsesGoroutineLang(t *testing.T) {
	p := New()
	en := language.Make("en")
	zh := language.Make("zh")
	p.Register(en, "hi", "Hello")
	p.Register(zh, "hi", "你好")

	// 测试不依赖全局 default，因为可能被其他测试改过
	language.Set(zh)
	defer language.Del()

	if got := p.Localize("hi"); got != "你好" {
		t.Errorf("Localize=%q", got)
	}
}

func TestI18nConcurrent(t *testing.T) {
	p := New()
	en := language.Make("en")
	p.Register(en, "init", "Hello")

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			p.Register(en, "k", "v")
		}(i)
		go func() {
			defer wg.Done()
			_ = p.LocalizeWithLang(en, "init")
		}()
	}
	wg.Wait()
}

func TestNormalizeLang(t *testing.T) {
	if got := normalizeLang(nil); got != "" {
		t.Errorf("nil→%q", got)
	}
	if got := normalizeLang(language.Make("ZH-cn")); got != "zh-cn" {
		t.Errorf("ZH-cn→%q", got)
	}
}

func TestBaseLang(t *testing.T) {
	if base := baseLang(language.Make("zh-CN")); base == nil || base.String() != "zh" {
		t.Errorf("baseLang(zh-CN)=%v", base)
	}
	if base := baseLang(language.Make("en")); base != nil {
		t.Errorf("baseLang(en) should be nil, got %v", base)
	}
}

func TestSetGetDelLanguage(t *testing.T) {
	zh := language.Make("zh")
	SetLanguage(zh)
	if got := GetLanguage(); got.String() != "zh" {
		t.Errorf("GetLanguage=%q", got.String())
	}
	DelLanguage()
	// 删除后回到全局 default
	if got := GetLanguage(); got.String() == "zh" {
		// 仅当全局 default 也是 zh 才可能；按 store.go 默认 en
		t.Errorf("DelLanguage failed, still %q", got.String())
	}
}

func TestDefaultPackageLevel(t *testing.T) {
	// 用一个全新 I18n 隔离 Default 副作用：临时替换 Default
	original := Default
	Default = New(WithDefaultLang(language.Make("en")))
	defer func() { Default = original }()

	en := language.Make("en")
	Register(en, "pk", "PV")
	RegisterBatch(en, map[string]any{"batch": "B"})

	if got := LocalizeWithLang(en, "pk"); got != "PV" {
		t.Errorf("LocalizeWithLang=%q", got)
	}
	if got := LocalizeWithLang(en, "batch"); got != "B" {
		t.Errorf("batch=%q", got)
	}

	language.Set(en)
	defer language.Del()
	if got := Localize("pk"); got != "PV" {
		t.Errorf("Localize=%q", got)
	}
}
