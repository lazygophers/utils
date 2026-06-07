//go:build country_all || country_eastern_europe || country_europe || country_ro || currency_all || currency_ron

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ron.RegisterName(xlanguage.English, "Romanian Leu")
}
