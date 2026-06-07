//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Korean, "오스트리아")
	dataAustria.RegisterOfficialName(xlanguage.Korean, "오스트리아 공화국")
	dataAustria.RegisterCapital(xlanguage.Korean, "빈")
}
