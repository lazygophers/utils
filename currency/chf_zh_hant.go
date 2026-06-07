//go:build (lang_zh_hant || lang_all) && (country_all || country_ch || country_europe || country_li || country_western_europe || currency_all || currency_chf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CHF.RegisterName(xlanguage.MustParse("zh-Hant"), "瑞士法郎")
}
