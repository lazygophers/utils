//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Korean, "탄자니아")
	dataTanzania.RegisterOfficialName(xlanguage.Korean, "탄자니아 연합 공화국")
	dataTanzania.RegisterCapital(xlanguage.Korean, "도도마")
}
