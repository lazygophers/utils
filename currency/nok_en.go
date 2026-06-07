//go:build country_all || country_antarctic || country_bv || country_europe || country_no || country_northern_europe || country_sj || currency_all || currency_nok

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nok.RegisterName(xlanguage.English, "Norwegian Krone")
}
