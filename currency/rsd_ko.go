//go:build (lang_ko || lang_all) && (country_all || country_europe || country_rs || country_southern_europe || currency_all || currency_rsd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Rsd.RegisterName(xlanguage.Korean, "세르비아 디나르")
}
