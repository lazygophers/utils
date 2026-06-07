//go:build (lang_ko || lang_all) && (country_all || country_am || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Korean, "아르메니아")
	dataArmenia.RegisterOfficialName(xlanguage.Korean, "아르메니아 공화국")
	dataArmenia.RegisterCapital(xlanguage.Korean, "예레반")
}
