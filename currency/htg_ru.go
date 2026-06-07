//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_ht || currency_all || currency_htg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Htg.RegisterName(xlanguage.Russian, "Гурд")
}
