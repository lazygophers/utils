//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uzs.RegisterName(xlanguage.MustParse("zh-Hant"), "烏茲別克索姆")
}
