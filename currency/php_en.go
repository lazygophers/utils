//go:build country_all || country_asia || country_ph || country_south_eastern_asia || currency_all || currency_php

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	PHP.RegisterName(xlanguage.English, "Philippine Peso")
}
