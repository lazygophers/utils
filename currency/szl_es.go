//go:build (lang_es || lang_all) && (country_africa || country_all || country_southern_africa || country_sz || currency_all || currency_szl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Szl.RegisterName(xlanguage.Spanish, "Lilangeni")
}
