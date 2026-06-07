//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.MustParse("zh-Hant"), "牙買加")
	dataJamaica.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "牙買加")
	dataJamaica.RegisterCapital(xlanguage.MustParse("zh-Hant"), "金斯敦")
}
