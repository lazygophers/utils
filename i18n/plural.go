package i18n

import "github.com/lazygophers/utils/language"

// PluralCategory 是 CLDR 复数分类。
type PluralCategory string

// 六种 CLDR 复数分类常量。
const (
	CategoryZero  PluralCategory = "zero"
	CategoryOne   PluralCategory = "one"
	CategoryTwo   PluralCategory = "two"
	CategoryFew   PluralCategory = "few"
	CategoryMany  PluralCategory = "many"
	CategoryOther PluralCategory = "other"
)

// pluralRule 把整数 n 映射到 CLDR 复数分类。
type pluralRule func(n int) PluralCategory

// pluralRules 是按 base language 注册的复数规则表，扩展位预留：
// 其他语言规则通过 messages_xx.go + //go:build lang_xx || lang_all 在 init 中追加。
var pluralRules = map[string]pluralRule{
	"en": pluralEn,
	"zh": pluralZh,
}

// pluralEn 实现英语 CLDR 规则：n==1 → one；其他 → other。
func pluralEn(n int) PluralCategory {
	if n == 1 {
		return CategoryOne
	}
	return CategoryOther
}

// pluralZh 实现中文 CLDR 规则：恒 other。
func pluralZh(_ int) PluralCategory {
	return CategoryOther
}

// pluralCategoryFor 按 locale fallback 链选规则，未命中默认 other。
// 跳过 "und"：xlanguage 把 und 的 Base 映射为 en，会污染未知 locale 的规则选择。
func pluralCategoryFor(locale *language.Tag, n int) PluralCategory {
	for _, t := range locale.FallbackChain() {
		s := t.String()
		if s == "und" {
			continue
		}
		if rule, ok := pluralRules[s]; ok {
			return rule(n)
		}
		if rule, ok := pluralRules[t.Base()]; ok {
			return rule(n)
		}
	}
	return CategoryOther
}

// TPlural 按当前 goroutine 语言选择复数 key 并插值，{n} 占位符可用。
func TPlural(key string, n int, args ...any) string {
	return defaultBundle.TPlural(key, n, args...)
}

// TPlural 按当前 goroutine 语言在该 Bundle 中执行复数查询。
func (b *Bundle) TPlural(key string, n int, args ...any) string {
	return b.pluralLookup(language.Get(), key, n, args)
}

// pluralLookup 选 category → 拼复数 key → 走 Bundle.lookup；未命中回退 .other 再回退 key 原文。
func (b *Bundle) pluralLookup(locale *language.Tag, key string, n int, args []any) string {
	cat := pluralCategoryFor(locale, n)
	merged := mergePluralArgs(n, args)

	candidate := key + "." + string(cat)
	if text, ok := b.findText(locale, candidate); ok {
		return interpolate(text, merged)
	}
	if cat != CategoryOther {
		fallback := key + "." + string(CategoryOther)
		if text, ok := b.findText(locale, fallback); ok {
			return interpolate(text, merged)
		}
	}
	if text, ok := b.findText(locale, key); ok {
		return interpolate(text, merged)
	}
	return key
}

// findText 沿 locale fallback 链查 Store，再回退默认 locale。
func (b *Bundle) findText(locale *language.Tag, key string) (string, bool) {
	for _, t := range locale.FallbackChain() {
		if text, ok := b.store.Get(t.String(), key); ok {
			return text, true
		}
	}
	if text, ok := b.store.Get(b.defaultLocale.String(), key); ok {
		return text, true
	}
	return "", false
}

// mergePluralArgs 把 n 作为命名参数 "n" 前置注入，让 {n} 占位符生效；
// 若原 args 已是位置模式（非命名），则改造为命名模式以容纳 "n"。
func mergePluralArgs(n int, args []any) []any {
	if len(args) == 0 || isNamed(args) {
		out := make([]any, 0, len(args)+2)
		out = append(out, "n", n)
		out = append(out, args...)
		return out
	}
	// 位置模式下保留原 args 行为，{n} 不可用，但 {0}/{1}... 仍按原顺序工作。
	return args
}
