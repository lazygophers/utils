//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_ge || country_western_asia || currency_all || currency_gel)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gel.RegisterName(xlanguage.MustParse("zh-Hant"), "喬治亞拉里")
}
