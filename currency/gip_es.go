//go:build (lang_es || lang_all) && (country_all || country_europe || country_gi || country_southern_europe || currency_all || currency_gip)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gip.RegisterName(xlanguage.Spanish, "Libra gibraltareña")
}
