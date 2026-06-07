//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.MustParse("zh-Hant"), "柬埔寨")
	dataCambodia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "柬埔寨王國")
	dataCambodia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "金邊")
}
