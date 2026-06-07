//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Korean, "시리아")
	dataSyria.RegisterOfficialName(xlanguage.Korean, "시리아 아랍 공화국")
	dataSyria.RegisterCapital(xlanguage.Korean, "다마스쿠스")
}
