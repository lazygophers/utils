//go:build (lang_ko || lang_all) && (country_all || country_asia || country_pk || country_southern_asia || currency_all || currency_pkr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pkr.RegisterName(xlanguage.Korean, "파키스탄 루피")
}
