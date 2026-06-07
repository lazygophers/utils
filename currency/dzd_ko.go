//go:build (lang_ko || lang_all) && (country_africa || country_all || country_dz || country_northern_africa || currency_all || currency_dzd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dzd.RegisterName(xlanguage.Korean, "알제리 디나르")
}
