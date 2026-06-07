//go:build (lang_fr || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp || currency_all || currency_kpw)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	KPW.RegisterName(xlanguage.French, "Won nord-coréen")
}
