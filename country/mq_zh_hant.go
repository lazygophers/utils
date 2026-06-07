//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_mq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.MustParse("zh-Hant"), "馬丁尼克")
	dataMartinique.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "馬丁尼克")
	dataMartinique.RegisterCapital(xlanguage.MustParse("zh-Hant"), "法蘭西堡")
}
