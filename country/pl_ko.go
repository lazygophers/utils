//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Korean, "폴란드")
	dataPoland.RegisterOfficialName(xlanguage.Korean, "폴란드 공화국")
	dataPoland.RegisterCapital(xlanguage.Korean, "바르샤바")
}
