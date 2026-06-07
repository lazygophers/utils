//go:build (lang_ko || lang_all) && (country_all || country_asia || country_bn || country_south_eastern_asia || currency_all || currency_bnd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bnd.RegisterName(xlanguage.Korean, "브루나이 달러")
}
