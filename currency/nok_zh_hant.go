//go:build (lang_zh_hant || lang_all) && (country_all || country_antarctic || country_bv || country_europe || country_no || country_northern_europe || country_sj || currency_all || currency_nok)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	NOK.RegisterName(xlanguage.MustParse("zh-Hant"), "挪威克朗")
}
