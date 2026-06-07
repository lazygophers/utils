//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_cf || country_cg || country_cm || country_ga || country_gq || country_middle_africa || country_td || currency_all || currency_xaf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	XAF.RegisterName(xlanguage.MustParse("zh-Hant"), "中非法郎")
}
