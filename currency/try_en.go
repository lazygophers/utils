//go:build country_all || country_asia || country_tr || country_western_asia || currency_all || currency_try

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TRY.RegisterName(xlanguage.English, "Turkish Lira")
}
