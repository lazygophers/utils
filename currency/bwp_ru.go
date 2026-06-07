//go:build (lang_ru || lang_all) && (country_africa || country_all || country_bw || country_southern_africa || currency_all || currency_bwp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bwp.RegisterName(xlanguage.Russian, "Пула")
}
