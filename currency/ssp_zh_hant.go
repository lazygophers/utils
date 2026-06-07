//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss || currency_all || currency_ssp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ssp.RegisterName(xlanguage.MustParse("zh-Hant"), "南蘇丹鎊")
}
