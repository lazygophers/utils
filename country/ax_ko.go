//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Korean, "올란드 제도")
	dataAlandIslands.RegisterOfficialName(xlanguage.Korean, "올란드 제도")
	dataAlandIslands.RegisterCapital(xlanguage.Korean, "마리에함")
}
