//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Korean, "탄자니아")
	dataTanzania.RegisterOfficialName(xlanguage.Korean, "탄자니아 연합 공화국")
	dataTanzania.RegisterCapital(xlanguage.Korean, "도도마")
}
