//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tjs.RegisterName(xlanguage.MustParse("zh-Hant"), "塔吉克索莫尼")
}
