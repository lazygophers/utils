//go:build country_all || country_asia || country_az || country_western_asia || currency_all || currency_azn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AZN.RegisterName(xlanguage.English, "Azerbaijani Manat")
}
