//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.Korean, "프랑스")
	dataFrance.RegisterOfficialName(xlanguage.Korean, "프랑스 공화국")
	dataFrance.RegisterCapital(xlanguage.Korean, "파리")
}
