//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss || currency_all || currency_ssp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ssp.RegisterName(xlanguage.Korean, "남수단 파운드")
}
