//go:build (lang_es || lang_all) && (country_africa || country_all || country_na || country_southern_africa || currency_all || currency_nad)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nad.RegisterName(xlanguage.Spanish, "Dólar namibio")
}
