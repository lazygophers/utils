//go:build (lang_fr || lang_all) && (country_all || country_asia || country_az || country_western_asia || currency_all || currency_azn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AZN.RegisterName(xlanguage.French, "Manat azerbaïdjanais")
}
