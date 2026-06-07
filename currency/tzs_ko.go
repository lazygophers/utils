//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz || currency_all || currency_tzs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tzs.RegisterName(xlanguage.Korean, "탄자니아 실링")
}
