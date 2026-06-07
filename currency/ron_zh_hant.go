//go:build (lang_zh_hant || lang_all) && (country_all || country_eastern_europe || country_europe || country_ro || currency_all || currency_ron)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	RON.RegisterName(xlanguage.MustParse("zh-Hant"), "羅馬尼亞列伊")
}
