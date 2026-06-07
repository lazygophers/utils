//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Korean, "아르메니아")
	dataArmenia.RegisterOfficialName(xlanguage.Korean, "아르메니아 공화국")
	dataArmenia.RegisterCapital(xlanguage.Korean, "예레반")
}
