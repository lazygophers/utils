//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_ge || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.MustParse("zh-Hant"), "喬治亞")
	dataGeorgia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "喬治亞")
	dataGeorgia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "提比里西")
}
