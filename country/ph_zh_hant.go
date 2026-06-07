//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_ph || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.MustParse("zh-Hant"), "菲律賓")
	dataPhilippines.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "菲律賓共和國")
	dataPhilippines.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬尼拉")
}
