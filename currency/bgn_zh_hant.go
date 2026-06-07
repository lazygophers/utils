//go:build (lang_zh_hant || lang_all) && (country_all || country_bg || country_eastern_europe || country_europe || currency_all || currency_bgn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bgn.RegisterName(xlanguage.MustParse("zh-Hant"), "保加利亞列弗")
}
