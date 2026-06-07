//go:build (lang_ru || lang_all) && (country_al || country_all || country_europe || country_southern_europe || currency_all)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	All.RegisterName(xlanguage.Russian, "Лек")
}
