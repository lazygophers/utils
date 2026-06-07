//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl || currency_all || currency_pln)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pln.RegisterName(xlanguage.Arabic, "زلوتي بولندي")
}
