//go:build (lang_zh_hant || lang_all) && (country_all || country_melanesia || country_nc || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.MustParse("zh-Hant"), "新喀里多尼亞")
	dataNewCaledonia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "新喀里多尼亞")
	dataNewCaledonia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "諾美亞")
}
