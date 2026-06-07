//go:build (lang_ko || lang_all) && (country_all || country_au || country_australia_and_new_zealand || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Korean, "오스트레일리아")
	dataAustralia.RegisterOfficialName(xlanguage.Korean, "오스트레일리아 연방")
	dataAustralia.RegisterCapital(xlanguage.Korean, "캔버라")
}
