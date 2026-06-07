//go:build (lang_zh_hant || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.MustParse("zh-Hant"), "波蘭")
	dataPoland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "波蘭共和國")
	dataPoland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "華沙")
}
