//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_mu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.MustParse("zh-Hant"), "模里西斯")
	dataMauritius.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "模里西斯共和國")
	dataMauritius.RegisterCapital(xlanguage.MustParse("zh-Hant"), "路易港")
}
