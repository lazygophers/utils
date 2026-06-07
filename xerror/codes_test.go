package xerror

import (
	"testing"

	"github.com/lazygophers/utils/i18n"
	"github.com/lazygophers/utils/language"
	xlanguage "golang.org/x/text/language"
)

type codeShortcutCase struct {
	name    string
	ctor    func(args ...any) *Error
	wantCode int
}

func TestCodeShortcutConstructors(t *testing.T) {
	cases := []codeShortcutCase{
		{"SystemError", NewSystemError, CodeSystem},
		{"InvalidParam", NewInvalidParam, CodeInvalidParam},
		{"NoAuth", NewNoAuth, CodeNoAuth},
		{"NoData", NewNoData, CodeNoData},
		{"Conflict", NewConflict, CodeConflict},
		{"NotLogin", NewNotLogin, CodeNotLogin},
		{"Timeout", NewTimeout, CodeTimeout},
		{"RateLimited", NewRateLimited, CodeRateLimited},
		{"Forbidden", NewForbidden, CodeForbidden},
		{"Unavailable", NewUnavailable, CodeUnavailable},
		{"DataCorrupted", NewDataCorrupted, CodeDataCorrupted},
	}
	for _, c := range cases {
		err := c.ctor("detail")
		if err.Code() != c.wantCode {
			t.Errorf("%s: code=%d, want %d", c.name, err.Code(), c.wantCode)
		}
		if err.Error() != "detail" {
			t.Errorf("%s: Error()=%q, want detail", c.name, err.Error())
		}
	}
}

func TestCodeConstantsUnique(t *testing.T) {
	codes := []int{
		CodeSystem, CodeSuccess,
		CodeInvalidParam, CodeNoAuth, CodeNoData, CodeConflict, CodeNotLogin,
		CodeTimeout, CodeRateLimited, CodeForbidden, CodeUnavailable, CodeDataCorrupted,
	}
	seen := map[int]bool{}
	for _, c := range codes {
		if seen[c] {
			t.Errorf("duplicate code: %d", c)
		}
		seen[c] = true
	}
}

type builtinLangCase struct {
	lang string
	code int
	want string
}

func TestBuiltinMessagesAutoRegistered(t *testing.T) {
	// 内置翻译由 codes_en.go / codes_zh.go init() 注册到 i18n.Default；直接查 i18n.Default。
	cases := []builtinLangCase{
		{"en", CodeSystem, "system error"},
		{"en", CodeInvalidParam, "invalid parameter"},
		{"zh", CodeSystem, "系统错误"},
		{"zh", CodeNoAuth, "未授权"},
	}
	for _, c := range cases {
		got := i18n.Default.LocalizeWithLang(makeTag(c.lang), errorKey(c.code))
		if got != c.want {
			t.Errorf("%s/%d: got %q, want %q", c.lang, c.code, got, c.want)
		}
	}
}

func TestBuiltinMessagesNewIntegration(t *testing.T) {
	// 通过 New 构造时按当前 goroutine 语言渲染，需保证默认 Localizer = i18n.Default。
	SetLocalizer(i18n.Default)
	language.Set(language.Make("zh"))
	defer language.Del()
	if got := NewInvalidParam().Error(); got != "参数无效" {
		t.Errorf("zh InvalidParam: %q", got)
	}
	language.Set(language.Make("en"))
	if got := NewInvalidParam().Error(); got != "invalid parameter" {
		t.Errorf("en InvalidParam: %q", got)
	}
}

func makeTag(s string) xlanguage.Tag {
	return xlanguage.Make(s)
}
