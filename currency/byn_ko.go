//go:build (lang_ko || lang_all) && (country_all || country_by || country_eastern_europe || country_europe || currency_all || currency_byn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Byn.RegisterName(xlanguage.Korean, "벨라루스 루블")
}
