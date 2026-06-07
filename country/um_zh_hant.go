//go:build (lang_zh_hant || lang_all) && (country_all || country_micronesia || country_oceania || country_um)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "美國本土外小島嶼")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "美國本土外小島嶼")
}
