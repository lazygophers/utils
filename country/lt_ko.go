//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Korean, "리투아니아")
	dataLithuania.RegisterOfficialName(xlanguage.Korean, "리투아니아 공화국")
	dataLithuania.RegisterCapital(xlanguage.Korean, "빌뉴스")
}
