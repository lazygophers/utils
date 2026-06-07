//go:build (lang_es || lang_all) && (country_all || country_asia || country_western_asia || country_ye || currency_all || currency_yer)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Yer.RegisterName(xlanguage.Spanish, "Rial yemení")
}
