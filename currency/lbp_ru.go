//go:build (lang_ru || lang_all) && (country_all || country_asia || country_lb || country_western_asia || currency_all || currency_lbp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lbp.RegisterName(xlanguage.Russian, "Ливанский фунт")
}
