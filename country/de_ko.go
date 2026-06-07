//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.Korean, "독일")
	dataGermany.RegisterOfficialName(xlanguage.Korean, "독일 연방 공화국")
	dataGermany.RegisterCapital(xlanguage.Korean, "베를린")
}
