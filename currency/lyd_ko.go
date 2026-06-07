//go:build (lang_ko || lang_all) && (country_africa || country_all || country_ly || country_northern_africa || currency_all || currency_lyd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LYD.RegisterName(xlanguage.Korean, "리비아 디나르")
}
