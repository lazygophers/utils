//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw || currency_all || currency_zwl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zwl.RegisterName(xlanguage.MustParse("zh-Hant"), "辛巴威元")
}
