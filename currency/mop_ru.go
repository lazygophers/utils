//go:build (lang_ru || lang_all) && (country_all || country_asia || country_eastern_asia || country_mo || currency_all || currency_mop)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MOP.RegisterName(xlanguage.Russian, "Патака")
}
