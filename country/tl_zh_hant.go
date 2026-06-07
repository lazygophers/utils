//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_tl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.MustParse("zh-Hant"), "東帝汶")
	dataTimorLeste.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "東帝汶民主共和國")
	dataTimorLeste.RegisterCapital(xlanguage.MustParse("zh-Hant"), "狄力")
}
