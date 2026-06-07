//go:build (lang_ko || lang_all) && (country_all || country_europe || country_mk || country_southern_europe || currency_all || currency_mkd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MKD.RegisterName(xlanguage.Korean, "마케도니아 데나르")
}
