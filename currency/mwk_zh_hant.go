//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw || currency_all || currency_mwk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mwk.RegisterName(xlanguage.MustParse("zh-Hant"), "馬拉威克瓦查")
}
