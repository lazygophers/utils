//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_jo || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.MustParse("zh-Hant"), "約旦")
	dataJordan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "約旦哈希米王國")
	dataJordan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "安曼")
}
