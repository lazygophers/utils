//go:build (lang_zh_hant || lang_all) && (country_all || country_at || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.MustParse("zh-Hant"), "奧地利")
	dataAustria.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "奧地利共和國")
	dataAustria.RegisterCapital(xlanguage.MustParse("zh-Hant"), "維也納")
}
