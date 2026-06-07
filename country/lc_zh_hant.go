//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_lc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.MustParse("zh-Hant"), "聖露西亞")
	dataSaintLucia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖露西亞")
	dataSaintLucia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "卡斯翠")
}
