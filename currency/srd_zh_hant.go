//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Srd.RegisterName(xlanguage.MustParse("zh-Hant"), "蘇利南元")
}
