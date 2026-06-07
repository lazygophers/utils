//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_er)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.Korean, "에리트레아")
	dataEritrea.RegisterOfficialName(xlanguage.Korean, "에리트레아")
	dataEritrea.RegisterCapital(xlanguage.Korean, "아스마라")
}
