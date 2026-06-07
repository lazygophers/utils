//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_do || currency_all || currency_dop)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dop.RegisterName(xlanguage.Arabic, "بيزو دومينيكي")
}
