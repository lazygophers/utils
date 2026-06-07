//go:build (lang_ko || lang_all) && (country_africa || country_all || country_gm || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Korean, "감비아")
	dataGambia.RegisterOfficialName(xlanguage.Korean, "감비아 공화국")
	dataGambia.RegisterCapital(xlanguage.Korean, "반줄")
}
