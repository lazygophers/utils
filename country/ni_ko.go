//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.Korean, "니카라과")
	dataNicaragua.RegisterOfficialName(xlanguage.Korean, "니카라과 공화국")
	dataNicaragua.RegisterCapital(xlanguage.Korean, "마나과")
}
