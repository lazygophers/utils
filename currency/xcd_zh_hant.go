//go:build (lang_zh_hant || lang_all) && (country_ag || country_ai || country_all || country_americas || country_caribbean || country_dm || country_gd || country_kn || country_lc || country_ms || country_vc || currency_all || currency_xcd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xcd.RegisterName(xlanguage.MustParse("zh-Hant"), "東加勒比元")
}
