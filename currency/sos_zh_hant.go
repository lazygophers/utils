//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sos.RegisterName(xlanguage.MustParse("zh-Hant"), "索馬利亞先令")
}
