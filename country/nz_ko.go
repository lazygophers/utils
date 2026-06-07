//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Korean, "뉴질랜드")
	dataNewZealand.RegisterOfficialName(xlanguage.Korean, "뉴질랜드")
	dataNewZealand.RegisterCapital(xlanguage.Korean, "웰링턴")
}
