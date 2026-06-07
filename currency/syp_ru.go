//go:build (lang_ru || lang_all) && (country_all || country_asia || country_sy || country_western_asia || currency_all || currency_syp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SYP.RegisterName(xlanguage.Russian, "Сирийский фунт")
}
