//go:build (lang_es || lang_all) && (country_all || country_asia || country_central_asia || country_kz || currency_all || currency_kzt)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kzt.RegisterName(xlanguage.Spanish, "Tenge")
}
