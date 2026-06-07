//go:build (lang_ko || lang_all) && (country_all || country_asia || country_la || country_south_eastern_asia || currency_all || currency_lak)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lak.RegisterName(xlanguage.Korean, "라오스 킵")
}
