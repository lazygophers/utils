//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Korean, "콜롬비아")
	dataColombia.RegisterOfficialName(xlanguage.Korean, "콜롬비아 공화국")
	dataColombia.RegisterCapital(xlanguage.Korean, "보고타")
}
