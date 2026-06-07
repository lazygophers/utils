//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_mm || country_south_eastern_asia || currency_all || currency_mmk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MMK.RegisterName(xlanguage.MustParse("zh-Hant"), "緬元")
}
