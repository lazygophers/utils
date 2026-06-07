//go:build (lang_es || lang_all) && (country_al || country_all || country_europe || country_southern_europe || currency_all)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ALL.RegisterName(xlanguage.Spanish, "Lek")
}
