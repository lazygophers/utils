//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.Korean, "인도")
	dataIndia.RegisterOfficialName(xlanguage.Korean, "인도 공화국")
	dataIndia.RegisterCapital(xlanguage.Korean, "뉴델리")
}
