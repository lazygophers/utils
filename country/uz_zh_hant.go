//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_central_asia || country_uz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.MustParse("zh-Hant"), "烏茲別克")
	dataUzbekistan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "烏茲別克共和國")
	dataUzbekistan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "塔什干")
}
