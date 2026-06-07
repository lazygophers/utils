//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Korean, "토고")
	dataTogo.RegisterOfficialName(xlanguage.Korean, "토고 공화국")
	dataTogo.RegisterCapital(xlanguage.Korean, "로메")
}
