//go:build (lang_ar || lang_all) && (country_all || country_oceania || country_polynesia || country_ws || currency_all || currency_wst)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	WST.RegisterName(xlanguage.Arabic, "تالا ساموي")
}
