//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Czk.RegisterName(xlanguage.MustParse("zh-Hant"), "捷克克朗")
}
