//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Korean, "잠비아")
	dataZambia.RegisterOfficialName(xlanguage.Korean, "잠비아 공화국")
	dataZambia.RegisterCapital(xlanguage.Korean, "루사카")
}
