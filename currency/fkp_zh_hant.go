//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Fkp.RegisterName(xlanguage.MustParse("zh-Hant"), "福克蘭群島鎊")
}
