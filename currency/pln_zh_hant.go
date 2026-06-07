//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pln.RegisterName(xlanguage.MustParse("zh-Hant"), "波蘭茲羅提")
}
