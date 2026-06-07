//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_is || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.MustParse("zh-Hant"), "冰島")
	dataIceland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "冰島")
	dataIceland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "雷克雅維克")
}
