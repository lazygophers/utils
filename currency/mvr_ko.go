//go:build (lang_ko || lang_all) && (country_all || country_asia || country_mv || country_southern_asia || currency_all || currency_mvr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MVR.RegisterName(xlanguage.Korean, "몰디브 루피야")
}
