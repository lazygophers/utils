//go:build country_all || country_eastern_europe || country_europe || country_pl || currency_all || currency_pln

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pln.RegisterName(xlanguage.English, "Zloty")
}
