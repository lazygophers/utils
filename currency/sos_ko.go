//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_so || currency_all || currency_sos)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SOS.RegisterName(xlanguage.Korean, "소말리아 실링")
}
