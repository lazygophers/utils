//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_th)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.MustParse("zh-Hant"), "泰國")
	dataThailand.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "泰王國")
	dataThailand.RegisterCapital(xlanguage.MustParse("zh-Hant"), "曼谷")
}
