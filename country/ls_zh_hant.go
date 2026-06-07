//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_ls || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.MustParse("zh-Hant"), "賴索托")
	dataLesotho.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "賴索托王國")
	dataLesotho.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬賽魯")
}
