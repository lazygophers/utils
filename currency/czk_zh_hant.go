//go:build (lang_zh_hant || lang_all) && (country_all || country_cz || country_eastern_europe || country_europe || currency_all || currency_czk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Czk.RegisterName(xlanguage.MustParse("zh-Hant"), "捷克克朗")
}
