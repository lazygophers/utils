//go:build (lang_zh_hant || lang_all) && (country_all || country_gu || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.MustParse("zh-Hant"), "關島")
	dataGuam.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "關島領地")
	dataGuam.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿加尼亞")
}
