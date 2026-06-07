//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_ke || currency_all || currency_kes)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kes.RegisterName(xlanguage.Korean, "케냐 실링")
}
