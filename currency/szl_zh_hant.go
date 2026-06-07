//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Szl.RegisterName(xlanguage.MustParse("zh-Hant"), "埃馬蘭吉尼")
}
