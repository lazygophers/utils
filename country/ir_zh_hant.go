//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_ir || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.MustParse("zh-Hant"), "伊朗")
	dataIran.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "伊朗伊斯蘭共和國")
	dataIran.RegisterCapital(xlanguage.MustParse("zh-Hant"), "德黑蘭")
}
