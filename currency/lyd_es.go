//go:build (lang_es || lang_all) && (country_africa || country_all || country_ly || country_northern_africa || currency_all || currency_lyd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lyd.RegisterName(xlanguage.Spanish, "Dinar libio")
}
