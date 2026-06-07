//go:build (lang_zh_hant || lang_all) && (country_all || country_cz || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.MustParse("zh-Hant"), "捷克")
	dataCzechia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "捷克共和國")
	dataCzechia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布拉格")
}
