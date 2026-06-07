//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_cw || country_sx || currency_all || currency_ang)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ANG.RegisterName(xlanguage.Korean, "네덜란드령 안틸레스 길더")
}
