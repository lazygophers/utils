//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.MustParse("zh-Hant"), "留尼旺")
	dataReunion.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "留尼旺島")
	dataReunion.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖但尼")
}
