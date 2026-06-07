//go:build (lang_ko || lang_all) && (country_all || country_ch || country_europe || country_li || country_western_europe || currency_all || currency_chf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CHF.RegisterName(xlanguage.Korean, "스위스 프랑")
}
