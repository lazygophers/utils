//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.Korean, "아제르바이잔")
	dataAzerbaijan.RegisterOfficialName(xlanguage.Korean, "아제르바이잔 공화국")
	dataAzerbaijan.RegisterCapital(xlanguage.Korean, "바쿠")
}
