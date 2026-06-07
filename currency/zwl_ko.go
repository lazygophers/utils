//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw || currency_all || currency_zwl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ZWL.RegisterName(xlanguage.Korean, "짐바브웨 달러")
}
