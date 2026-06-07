//go:build (lang_es || lang_all) && (country_all || country_fj || country_melanesia || country_oceania || currency_all || currency_fjd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	FJD.RegisterName(xlanguage.Spanish, "Dólar fiyiano")
}
