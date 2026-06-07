//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.Korean, "쿡 제도")
	dataCookIslands.RegisterOfficialName(xlanguage.Korean, "쿡 제도")
	dataCookIslands.RegisterCapital(xlanguage.Korean, "아바루아")
}
