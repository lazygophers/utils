//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Korean, "포르투갈")
	dataPortugal.RegisterOfficialName(xlanguage.Korean, "포르투갈 공화국")
	dataPortugal.RegisterCapital(xlanguage.Korean, "리스본")
}
