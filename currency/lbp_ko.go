//go:build (lang_ko || lang_all) && (country_all || country_asia || country_lb || country_western_asia || currency_all || currency_lbp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lbp.RegisterName(xlanguage.Korean, "레바논 파운드")
}
