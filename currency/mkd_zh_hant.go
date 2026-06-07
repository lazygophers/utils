//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_mk || country_southern_europe || currency_all || currency_mkd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MKD.RegisterName(xlanguage.MustParse("zh-Hant"), "馬其頓代納爾")
}
