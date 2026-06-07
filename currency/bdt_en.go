//go:build country_all || country_asia || country_bd || country_southern_asia || currency_all || currency_bdt

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bdt.RegisterName(xlanguage.English, "Taka")
}
