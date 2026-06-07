//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_vi)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "美屬維京群島")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "美屬維京群島")
	dataUsVirginIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "夏律第阿馬利亞")
}
